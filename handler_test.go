package lab2

import (
	"bytes"
	"testing"
	. "gopkg.in/check.v1"
	"strings"
)

func TestHandler(t *testing.T) { TestingT(t) }

type HandlerSuite struct{}

var _ = Suite(&HandlerSuite{})

func (s *HandlerSuite) TestWrongTypeOfInput(c *C) {
	handler := ComputeHandler {
		Input:  strings.NewReader("Letters are not numbers"),
		Output: bytes.NewBufferString(""),
	}
	err := handler.Compute()
	c.Assert(err.Error(), Equals, "invalid input expression")
}

func (s *HandlerSuite) TestWrongNumberOfNumbers(c *C) {
	handler := ComputeHandler {
		Input: strings.NewReader("2 3 7 +"),
		Output: bytes.NewBufferString(""),
	}
	err := handler.Compute()
	c.Assert(err.Error(), Equals, 
	"invalid input expression(missed operator(s))")
}

func (s *HandlerSuite) TestHandlerComputeInput(c *C) {
	output := bytes.NewBufferString("")
	handler := ComputeHandler {
		Input: strings.NewReader("1 7 +"),
		Output: output,
	}
	err := handler.Compute()
	c.Assert(err,IsNil)
	c.Assert(output.String(), Equals, "8\n")
}
