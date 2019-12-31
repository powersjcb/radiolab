package main

import (
	"errors"
	"fmt"
	"github.com/powersjcb/radiolab/core"
	"github.com/tochlab/go-hackrf/hackrf"
	"log"
	"time"
)

func main() {
	fmt.Println("starting device")
	dev, err := initDevice()
	if err != nil {
		log.Printf("%v+", err)
		return
	}
	defer teardown(dev)

	if dev.SetFreq(104_500_000) != nil {
		log.Printf("%v+", err)
		return
	}
	if dev.SetSampleRate(20_000_000) != nil {
		log.Printf("%v+", err)
		return
	}
	fmt.Println("starting RX...")
	app, _ := core.NewApplication()
	if dev.StartRX(app.NoopCallback) != nil {
		log.Printf("%v+", err)
		return
	}

	ticker := time.NewTicker(250 * time.Millisecond)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case _ = <-ticker.C:
				fmt.Println("checking the buffer:")
				for _, sample := range app.Store.Get() {
					fmt.Printf("%s: samples: %d\n", sample.OccurredAt, len(sample.Data))
				}
			}
		}
	}()

	// todo enter event loop for processing input changes, until then we wait forever
	time.Sleep(10 * time.Second)
	fmt.Printf("%x+", app.Store.Get()[0].Data[:20])
}

func initDevice() (*hackrf.Device, error) {
	err := hackrf.Init()
	if err != nil {
		return nil, err
	}
	fmt.Println("getting device list...")
	devices, err := hackrf.DeviceList()
	if err != nil {
		return nil, err
	}
	if len(devices) == 0 {
		return nil, errors.New("no devices found")
	}

	fmt.Println("initializing first available Device...")
	// open first available hackrf device
	fmt.Printf("%v+\n", devices[0])
	dev, err := hackrf.OpenBySerial(devices[0].SerialNumber)
	if err != nil {
		return nil, err
	}
	return dev, nil
}

func teardown(dev *hackrf.Device) {
	err := dev.Close()
	if err != nil {
		log.Println(err.Error())
	}
	err = hackrf.Exit()
	if err != nil {
		log.Fatal(err.Error())
	}
}
