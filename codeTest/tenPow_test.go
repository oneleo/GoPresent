// +build OMIT
package myMath

import (
	"math"
	"testing"
)

func TestAdd(t *testing.T) { // OMIT START
	for i := 0; i < 10; i++ {
		if m := tenPow(i); m != math.Pow10(i) {
			t.Error("10 的次方函數測試失敗！")
		} else {
			t.Log("10 的次方函數測試成功！") // OMIT END
		}
	}
}
