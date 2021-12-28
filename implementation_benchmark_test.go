package lab2

import (
	"fmt"
	"testing"
)

var calcResult float64

func BenchmarkPostfixExpressionCalculator(b *testing.B) {
	const baseExp = "2 3 ^ 10 5 / - 2 + "
	currentExpression := baseExp
	for i := 0; i < 18; i++ {
		currentExpression += currentExpression + "+ "
		b.Run(fmt.Sprintf("ExpressionLen=%d", len(currentExpression)), func(b *testing.B) {
			calcResult, _ = PostfixExpressionCalculator(currentExpression)
		})
	}
}
