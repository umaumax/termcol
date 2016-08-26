package termcol

import (
	"fmt"
)

func (c Color) Wrap(a ...interface{}) string {
	return string(c) + fmt.Sprint(a...) + string(Reset)
}

func (c Color) WrapSpace(a ...interface{}) string {
	s := string(c)
	for i, v := range a {
		if _, ok := v.(Color); !ok &&
			i > 0 {
			s += " "
		}
		s += fmt.Sprint(v)
	}
	s += string(Reset)
	return s
}

func (c Color) WrapF(format string, a ...interface{}) string {
	return fmt.Sprintf(string(c)+format+string(Reset), a...)
}
