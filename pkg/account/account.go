package account

import (
	"bytes"
	"fmt"
)

// Number represents a bank account number
type Number [9]int

func (an Number) String() string {
	buf := bytes.NewBuffer(nil)
	var missing bool
	for i := 0; i < 8; i++ {
		if an[i] == -1 {
			buf.WriteString("?")
			missing = true
		} else {
			buf.WriteString(fmt.Sprintf("%d", an[i]))
		}
	}
	if missing {
		buf.WriteString(" ILL")
	} else if !IsValid(an) {
		buf.WriteString(" ERR")
	}
	return buf.String()
}

// IsValid determines if the bank account number is valid
func IsValid(an Number) bool {
	d1, d2, d3, d4, d5, d6, d7, d8, d9 :=
		an[8], an[7], an[6], an[5], an[4], an[3], an[2], an[1], an[0]
	return (d1+(2*d2)+(3*d3)+(4*d4)+(5*d5)+(6*d6)+(7*d7)+(8*d8)+(9*d9))%11 == 0
}
