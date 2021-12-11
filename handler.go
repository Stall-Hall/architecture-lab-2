package lab2

import (
	"io"
	"io/ioutil"
)

type ComputeHandler struct {
		Input  io.Reader
		Output io.Writer
}

func (ch *ComputeHandler) Compute() error 
{
	bytes, readErr := ioutil.ReadAll(ch.Input)
	if readErr != nil 
	{
		return readErr
	}

	e := string(bytes)

	_, writeErr := ch.Output.Write([]byte(res))
	if writeErr != nil 
	{
		return writeErr
	}

	res, computeErr := PostfixExpressionCalculator(e)
	if computeErr != nil 
	{
		return computeErr
	}

	return nil
}
