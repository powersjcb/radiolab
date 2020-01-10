package lib

import (
	"github.com/mjibson/go-dsp/fft"
)

type SpectralPoint struct {
	Frequency float64
	Value     complex128
}

func NewSpectrum(data []complex128, sampleFreq uint64, sampleRate float64) []SpectralPoint{
	dft := bruteDFT(data)
	res := make([]SpectralPoint, len(data))
	for n, fBin:= range dft {
		// frequency detected relative to sampleFrequency
		var b float64
		if n < len(data) / 2 {
			b = float64(n)
		} else {
			b = - float64(n - len(data)/2.0)
		}
		freq := b * sampleRate / float64(len(data)) + float64(sampleFreq)
		// offset of detected frequency relative
		//offset := - sampleRate / 2
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
	return fft.FFT(data)
	//bins := make([]complex128, len(data))
	//for binIdx := 0; binIdx < len(data); binIdx++ {
	//	for sampleIdx := 0; sampleIdx < len(data); sampleIdx ++ {
	//		bins[binIdx] += data[sampleIdx] * cmplx.Exp(complex(0, -(float64(2) * math.Pi / float64(len(data))) * float64(binIdx * sampleIdx)))
	//	}
	//}
	//return bins
}

