package lib_test

import (
	"fmt"
	"github.com/powersjcb/radiolab/pkg/lib"
	"math"
	"math/cmplx"
	"testing"
)
func generateSinData(carrierFreq, signalFreq float64, numSamples int, sampleRate float64) []complex128 {
	samples := make([]complex128, numSamples)
	for i := 0; i < numSamples; i++ {
		t := float64(i) / sampleRate
		samples[i] = complex(math.Sin(), math.Cos())
	}
	return samples
}

const float64EqualityThreshold = 1e-9

func almostEqual(a, b complex128) bool {
	return math.Abs(real(a) - real(b)) <= float64EqualityThreshold &&
		math.Abs(imag(a) - imag(b)) <= float64EqualityThreshold
}

func TestNewSpectrum(t *testing.T) {
	sampleRate := float64(4)
	samples := 8
	frequency := float64(1)

	expectedSpectra := []lib.SpectralPoint{
		{
			Frequency: -3,
			Value: 0,
		},
		{
			Frequency: -2,
			Value:     0.5,
		},
		{
			Frequency: -1,
			Value:     0,
		},
		{
			Frequency: 0,
			Value:     0,
		},
		{
			Frequency: 1,
			Value:     0.5,
		},
		{
			Frequency: 2,
			Value:     0,
		},
		{
			Frequency: 3,
			Value: 0,
		},
	}

	data := generateSinData(frequency, samples, sampleRate)
	spectra := lib.NewSpectrum(data, sampleRate)

	if len(spectra) != len(expectedSpectra) {
		fmt.Printf("%v+", spectra)
		t.Fatalf("expected %d points, got %d", len(expectedSpectra), len(spectra))
	}

	for i := 0; i < len(expectedSpectra); i++ {
		if !almostEqual(spectra[i].Value, expectedSpectra[i].Value) {
			fmt.Printf("%v+\n%v+\n", expectedSpectra, spectra)
			t.Fatalf("got: %v+, expected: %v+)", spectra[i], expectedSpectra[i])
		}
	}
}
