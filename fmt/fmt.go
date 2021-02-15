package yesmlfmt

import (
	"fmt"

	"github.com/nickatsegment/yesml"
	yaml "github.com/nickatsegment/yesml/yamlv3"
)

type Formatter struct {
}

const head = `%YAML 1.2
%yesml 0.1
---
`

const indent = 2

type ErrPos struct {
	Line int
	Col  int
	Err  error
}

func (e *ErrPos) Error() string {
	return fmt.Sprintf("(%d, %d) %s", e.Line, e.Col, e.Err)
}

func (e *ErrPos) Unwrap() error {
	return e.Err
}

// Formats n and its Content recursively
func (f *Formatter) walk(n *yaml.Node) error {
	if n.Tag == "!!yaml" {
		fmtv, err := f.Format([]byte(n.Value))
		if err != nil {
			return &ErrPos{Line: n.Line, Col: n.Column, Err: err}
		}
		n.Value = string(fmtv)
	}
	for _, c := range n.Content {
		err := f.walk(c)
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *Formatter) Format(bin []byte) ([]byte, error) {
	var out yaml.Node
	err := yesml.Unmarshal(bin, &out)
	if err != nil {
		return nil, fmt.Errorf("unmarshal: %w", err)
	}

	err = f.walk(&out)
	if err != nil {
		return nil, fmt.Errorf("walking node: %w", err)
	}

	bout, err := yesml.Marshal(&out)
	if err != nil {
		return nil, fmt.Errorf("remarshal: %w", err)
	}
	// TODO: when to have the head? it's annoying
	// bout = append([]byte(head), bout...)
	return bout, nil
}

func Format(bin []byte) ([]byte, error) {
	return (&Formatter{}).Format(bin)
}
