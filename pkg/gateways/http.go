package gateways

import (
	"fmt"
	"github.com/powersjcb/radiolab/pkg"
	"github.com/powersjcb/radiolab/pkg/lib"
	"github.com/powersjcb/radiolab/pkg/views"
	"log"
	"net/http"
)

type HTTPServer struct {
	Application *pkg.Application
}

func NewHTTPServer(app *pkg.Application) HTTPServer {
	return HTTPServer{
		Application: app,
	}
}

func (s *HTTPServer)  Start() {
	http.HandleFunc("/", s.RootHandler)
	http.HandleFunc("/fft", s.FFTHandler)
	fmt.Println("starting http server")
	err := http.ListenAndServe("127.0.0.1:9999", nil)
	if err != nil {
		log.Fatalf("http server stopped: %s", err.Error())
	}
}

func (*HTTPServer) RootHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "ok")
	if err != nil {
		log.Fatalf("error: %s", err.Error())
	}
}

func (s *HTTPServer) FFTHandler(w http.ResponseWriter, r *http.Request) {
	samples := s.Application.Store.GetIQ()
	if len(samples) == 0 {
		return
	}
	data := samples[0].Data
	config := s.Application.Store.GetConfig()
	fft := lib.NewSpectrum(lib.DecodeIQ(data), config.SampleFrequency, config.SampleRate)
	views.FFTPlot(w, fft)
	//err := views.FFTPlot()
}
