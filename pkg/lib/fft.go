package lib

import (
	"math"
	"math/cmplx"
)

func Transform(data []complex128) []complex128 {
	//return fft.FFT(data)
	return bruteDFT(data)
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

