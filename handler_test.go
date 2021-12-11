package lab2

import (
	"bytes"
	. "gopkg.in/check.v1"
	"strings"
)

type HandlerSuite struct{}

var _ = Suite(&HandlerSuite{})

func (s *HandlerSuite) TestingUseResultForOutput(c *C) {
	output := bytes.NewBufferString("")
	handler := ComputeHandler{
		Input:  strings.NewReader("2 5 -"),
		Output: output,
	}
	_ = handler.Compute()
	c.Assert(output.String(), Equals, "- 5 2")
}

func (s *HandlerSuite) TestingBackError(c *C) {
	handler := ComputeHandler{
		Input:  strings.NewReader("New text is read"),
		Output: bytes.NewBufferString(""),
	}
	err := handler.Compute()
	c.Assert(err.Error(), Equals, "invalid expression")
}