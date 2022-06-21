package math

import "errors"

// Sum 값들을 모두 더한다.
func Sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// Div a를 b로 나눈다.
func Div(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0.0, errors.New("can't divide by zero")
	}
	return a / b, nil
}

// StrRepeat 문자열을 count 만큼 반복하고 결과를 반환한다.
func StrRepeat(s string, count int) string {
	b := make([]byte, len(s)*count)
	bp := copy(b, s)
	for bp < len(b) {
		copy(b[bp:], b[bp:])
		bp *= 2
	}
	return string(b)
}
