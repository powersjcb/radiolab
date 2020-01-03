package lib

func DecodeIQ(buf []int8) []complex128 {
	d := make([]complex128, len(buf) / 2)
	for i := 0; i < len(buf); i += 2 {
		d[i/2] = complex(float64(buf[i+1]), float64(buf[i]))
	}
	return d
}