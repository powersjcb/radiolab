package lib

import (
	"math"
	"math/cmplx"
)

type SpectralPoint struct {
	Frequency float64
	Value     complex128
}

func NewSpectrum(data []complex128, sampleFreq uint64, sampleRate float64) []SpectralPoint{
	dft := bruteDFT(data)
	res := make([]SpectralPoint, len(data))
	for n, fBin:= range dft {
		freq := float64(n) * float64(sampleFreq) / float64(len(data))
		res[n] = SpectralPoint{
			Frequency: freq,
			Value:     fBin,
		}
	}

	return res
}

// bruteDFT Discrete Fourier transform - per https://en.wikipedia.org/wiki/Discrete_Fourier_transform
// note: this is not optimized and runs at O(n^2)
func bruteDFT(data []complex128) []complex128 {
	bins := make([]complex128, len(data))
	for binIdx := 0; binIdx < len(data); binIdx++ {
		for sampleIdx := 0; sampleIdx < len(data); sampleIdx ++ {
			bins[binIdx] += data[sampleIdx] * cmplx.Exp(complex(0, -(float64(2) * math.Pi / float64(len(data))) * float64(binIdx * sampleIdx)))
		}
	}
	return bins
}

