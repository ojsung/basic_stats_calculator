package calculator

import op "github.com/ojsung/basic_stats_calculator/pkg/operand"

func generateTaylorTerms[T op.FloatNumber](terms int) (termsSlice []op.Operand[T]) {
	termsSlice = make([]op.Operand[T], terms)
	for index := range terms {
		termIndex := index + 1
		var sign op.Operand[T]
		if termIndex % 2 == 0 {
			println(sign)
		}
		
	}
	return
}