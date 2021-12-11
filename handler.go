package lab2

import (
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

	_, writeErr := ch.Output.Write([]byte(res))
	if writeErr != nil {
		return writeErr
	}

	res, computeErr := PostfixExpressionCalculator(e)
	if computeErr != nil {
		return computeErr
	}

	return nil
}
