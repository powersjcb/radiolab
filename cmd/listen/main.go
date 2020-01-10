package main

import (
	"errors"
	"fmt"
	"github.com/powersjcb/radiolab/pkg"
	"github.com/powersjcb/radiolab/pkg/gateways"
	"github.com/tochlab/go-hackrf/hackrf"
	"log"
)

const (
	sampleFreqHz uint64 = 104_500_000
	sampleRateHz float64 = 20_000_000
)

func main() {
	fmt.Println("starting device")
	dev, err := initDevice()
	if err != nil {
		log.Printf("%v+", err)
		return
	}
	defer teardown(dev)

	if dev.SetFreq(sampleFreqHz) != nil {
		log.Printf("%v+", err)
		return
	}
	if dev.SetSampleRate(sampleRateHz) != nil {
		log.Printf("%v+", err)
		return
	}

	fmt.Println("starting RX...")
	app, _ := pkg.NewApplication()
	if dev.StartRX(app.NoopCallback) != nil {
		log.Printf("%v+", err)
		return
	}
	err = app.Store.SetConfig(
		pkg.Config{
			SampleFrequency: sampleFreqHz,
			SampleRate: sampleRateHz,
		},
	)
	if err != nil {
		log.Fatalln(err.Error())
	}

	server := gateways.NewHTTPServer(&app)
	server.Start()

	// todo enter event loop for processing input changes, until then we wait forever
	//data  app.Store.Get()[0].Data
	//iqData := lib.DecodeIQ(
	//	,
	//)
	//err = views.IQPlot(iqData)
	//if err != nil {
	//	log.Println(err)
	//	log.Fatal("failed")
	//}

	//err = views.FFTPlot(lib.NewSpectrum(iqData, sampleFreqHz, sampleRateHz))
	//if err != nil {
	//	log.Println(err)
	//	log.Fatal("failed")
	//}

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

