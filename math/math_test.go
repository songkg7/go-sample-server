package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Sum(t *testing.T) {
	v0 := Sum(1, 2, 3)
	assert.Equal(t, v0, 6, "1,2,3 == 6")

	v1 := Sum(6, 5)
	assert.Equal(t, v1, 11, "6+5 == 11")
}

func Test_Div(t *testing.T) {
	v2, _ := Div(0, 2)
	t.Log("0/2 =", v2)
}

func Test_StrRepeat(t *testing.T) {
	str := StrRepeat("a", 3)
	if len(str) != 3 {
		t.Fatal("Repeat fail")
	}
}
