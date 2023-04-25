package kio

import (
	"bufio"
	"io"
	"os"
	"testing"
)

func TestPipeline(t *testing.T) {
	o := &RunOptions{
		InputPath: "./testdata/kcl-run.yaml",
	}
	if err := o.Run(); err != nil {
		t.Fatal(err)
	}
}

// RunOptions is the options for the run command
type RunOptions struct {
	// InputPath is the -f flag
	InputPath string
	// OutputPath is the -o flag
	OutputPath string
}

// RunOptions creates a new options for the run command.
func NewRunOptions() *RunOptions {
	return &RunOptions{}
}

// Run the with the run command options.
func (o *RunOptions) Run() error {
	reader, err := o.reader()
	if err != nil {
		return err
	}
	writer, err := o.writer()
	if err != nil {
		return err
	}
	pipeline := NewPipeline(reader, writer, true)
	return pipeline.Execute()
}

func (o *RunOptions) reader() (io.Reader, error) {
	if o.InputPath == "-" {
		return os.Stdin, nil
	} else {
		file, err := os.Open(o.InputPath)
		if err != nil {
			return nil, err
		}
		return bufio.NewReader(file), nil
	}
}

func (o *RunOptions) writer() (io.Writer, error) {
	if o.OutputPath == "" {
		return os.Stdout, nil
	} else {
		file, err := os.OpenFile(o.OutputPath, os.O_CREATE|os.O_RDWR, 0744)
		if err != nil {
			return nil, err
		}
		return bufio.NewWriter(file), nil
	}
}
