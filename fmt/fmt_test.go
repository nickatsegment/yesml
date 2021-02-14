package yesmlfmt

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFiles(t *testing.T) {
	fns, err := filepath.Glob("testdata/*.in.yaml")
	require.NoError(t, err)
	for _, fn := range fns {
		tname := strings.Split(fn, ".")[0]
		t.Run(tname, func(t *testing.T) {
			bin, err := ioutil.ReadFile(fn)
			require.NoError(t, err)
			// TODO test error cases?

			expectedBout, err := ioutil.ReadFile(tname + ".out.yaml")
			require.NoError(t, err)

			bout, err := Format(bin)
			require.NoError(t, err)

			assert.Equal(t, string(expectedBout), string(bout))
		})
	}
}
