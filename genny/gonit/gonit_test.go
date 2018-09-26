package gonit

import (
	"testing"

	"github.com/gobuffalo/genny/gentest"
	"github.com/stretchr/testify/require"
)

func Test_New(t *testing.T) {
	r := require.New(t)

	g, err := New(&Options{})
	r.NoError(err)

	run := gentest.NewRunner()
	run.WithGroup(g)

	r.NoError(run.Run())

	res := run.Results()

	r.Len(res.Commands, 5)
	r.Len(res.Files, 4)

	f := res.Files[0]
	r.Equal(".gitignore", f.Name())

}
