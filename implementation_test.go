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

func (s *ImplementationSuite) TestTablePostfixExpressionCalculator(c *C){
	var tests = []struct{
		inputString string	
		expectedResult float64
	}{
		{"2 2 +",			4.0},
		{"12 1 * 4 +",		16.0},
		{"2 2 ^ 4 * 10 -",	6.0},
		{"4 5 + 5 2 - / ",	3.0},
		{"9 2 * 12 16 + -",	-10.0},
	}
	for _, test := range tests{
		res, err := PostfixExpressionCalculator(test.inputString)
		c.Assert(err, IsNil)
		c.Assert(res, Equals, test.expectedResult)
	}
}

func (s *ImplementationSuite) TestErrorPostfixExpressionCalculator(c *C){
	var tests = []struct{
		inputString string	
		expectedError string
	}{
		{"", `invalid input\(short input expression\)`},
		{"1 + +", `invalid input\(too many operators in the expression\)`},
		{"1 1 + 1", `invalid input expression\(missed operator\(s\)\)`},
		{"q w e r t y", `invalid input expression`},
	}
	for _, test := range tests{
		_, err := PostfixExpressionCalculator(test.inputString)
		c.Assert(err, ErrorMatches, test.expectedError)
	}
}
func ExamplePostfixExpressionCalculator() {
	res, _ := PostfixExpressionCalculator("2 2 +")
	fmt.Println(res)

	// Output:
	// 4
}
