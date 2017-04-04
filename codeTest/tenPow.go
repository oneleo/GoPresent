// +build OMIT
package myMath

func tenPow(in int) (out float64) { // OMIT START
	out = 1
	for i := 1; i <= in; i++ {
		out = out * 10 // OMIT END
	}
	return
}
