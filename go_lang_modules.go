// это коммент к этому модулю

package go_lang_modules

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a T, b T) T {
	return a + b
}
