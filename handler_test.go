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

func (s *HandlerSuite) TestingUseResultForOutput(c *C) {
	output := bytes.NewBufferString("")
	handler := ComputeHandler{
		Input:  strings.NewReader("2 5 -"),
		Output: output,
	}
	err := handler.Compute()
	c.Assert(err, IsNil)
	c.Assert(output.String(), Equals, "-3\n")
}

func (s *HandlerSuite) TestingBackError(c *C) {
	handler := ComputeHandler{
		Input:  strings.NewReader("New text is read"),
		Output: bytes.NewBufferString(""),
	}
	err := handler.Compute()
	c.Assert(err.Error(), Equals, "invalid input expression")
}
