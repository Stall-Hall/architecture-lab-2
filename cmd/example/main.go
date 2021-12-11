package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	lab2 "github.com/Stall-Hall/architecture-lab-2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile       = flag.String("f", "", "Input file with expression")
	outputFile      = flag.String("o", "", "Output file for calculation result")
)

func main() {
	flag.Parse()

	var reader io.Reader
	var writer io.Writer
	var err error

	if *inputFile != "" {
		reader, err = os.Open(*inputFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			return
		}
	} else if *inputExpression != "" {
		reader = strings.NewReader(*inputExpression)
	} else {
		fmt.Fprintln(os.Stderr, "missing expression")
		return
	}

	if *outputFile != "" {
		writer, err = os.Create(*outputFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			return
		}
	} else {
		writer = os.Stdout
	}

	handler := &lab2.ComputeHandler{
		//Input:  reader,
		//Output: writer,
	}
	err = handler.Compute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}
