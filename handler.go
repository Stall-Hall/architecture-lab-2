package lab2

import (
	"fmt"
	"io"
	"io/ioutil"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.
type ComputeHandler struct {
		Input  io.Reader
		Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	bytes, readErr := ioutil.ReadAll(ch.Input)
	if readErr != nil {
		return readErr
	}

	e := string(bytes)

	res, computeErr := PostfixExpressionCalculator(e)
	if computeErr != nil {
		return computeErr
	}

	_, writeErr := fmt.Fprintln(ch.Output, res)
	if writeErr != nil {
		return writeErr
	}

	return nil
}
