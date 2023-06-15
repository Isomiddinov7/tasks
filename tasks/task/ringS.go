package task

func RingS(r1, r2 float64) float64 {
	var rings float64
	S1 := Surface(r1)
	S2 := Surface(r2)
	if S1 > S2 {
		rings = S1 - S2
	} else {
		rings = S2 - S1
	}
	return rings
}
