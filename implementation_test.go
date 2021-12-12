package lab2

import (
	"fmt"

	"testing"
	. "gopkg.in/check.v1"
)

func TestImplementation(t *testing.T) { TestingT(t) }

type ImplementationSuite struct{}

var _ = Suite(&ImplementationSuite{})

func (s *ImplementationSuite) TestPostfixExpressionCalculator(c *C) {
	str := "4 2 - 3 * 5 +"
	res, err := PostfixExpressionCalculator(str)
	c.Assert(err, IsNil)
	c.Assert(res, Equals, 11.0)
}

func ExamplePostfixExpressionCalculator() {
	res, _ := PostfixExpressionCalculator("2 2 +")
	fmt.Println(res)

	// Output:
	// 4
}
