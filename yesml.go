package yesml

import (
	"bytes"

	yaml "github.com/nickatsegment/yesml/yamlv3"
)

func Unmarshal(b []byte, v interface{}) error {
	// TODO
	return yaml.Unmarshal(b, v)
}

func Marshal(v interface{}) ([]byte, error) {
	bout := bytes.Buffer{}
	enc := yaml.NewEncoder(&bout)
	enc.SetIndent(2)
	err := enc.Encode(v)
	return bout.Bytes(), err
}
