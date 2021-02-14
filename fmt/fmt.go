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

func (f *Formatter) Format(bin []byte) ([]byte, error) {
	var out yaml.Node
	err := yesml.Unmarshal(bin, &out)
	if err != nil {
		return nil, fmt.Errorf("unmarshal: %w", err)
	}
	bout, err := yesml.Marshal(&out)
	if err != nil {
		return nil, fmt.Errorf("remarshal: %w", err)
	}
	bout = append([]byte(head), bout...)
	return bout, nil
}

func Format(bin []byte) ([]byte, error) {
	return (&Formatter{}).Format(bin)
}
