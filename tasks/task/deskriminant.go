package task

func Descriminant(a, b, c int) int {
	D := b*b - 4*a*c
	if D > 0 {
		return 2
	} else if D == 0 {
		return 1
	} else {
		return 0
	}
}
