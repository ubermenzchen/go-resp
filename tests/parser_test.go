package goresp_test

import (
	"bytes"
	"fmt"
	"testing"

	parser "github.com/ubermenzchen/go-resp/parser"
)

func TestParseInt(t *testing.T) {
	b := 1674
	buffer := bytes.NewBufferString(fmt.Sprintf(":%d\r\n", b))
	p := parser.NewParser[int](&parser.IntegerParser{})
	v, err := p.Parse(buffer)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if v != b {
		t.Fatalf("Error: value parsed: %d value expected: %d\n", v, b)
	}
}

func TestParseSimpleString(t *testing.T) {
	b := "kappa kappa kappa"
	p := parser.NewParser[string](&parser.SimpleStringParser{})
	buffer := bytes.NewBufferString(fmt.Sprintf("+%s\r\n", b))
	v, err := p.Parse(buffer)
	if err != nil {
		t.Fatal(err.Error())
	}

	if v != b {
		t.Fatalf("Error: value parsed: %s value expected: %s\n", v, b)
	}
}
