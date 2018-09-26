package gonit

import (
	"os/exec"

	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/genny/movinglater/gotools/gomods"
	"github.com/gobuffalo/genny/movinglater/plushgen"
	"github.com/gobuffalo/packr"
	"github.com/gobuffalo/plush"
	"github.com/gobuffalo/release/genny/initgen"
	"github.com/pkg/errors"
)

func New(opts *Options) (*genny.Group, error) {
	gg := &genny.Group{}

	if err := opts.Validate(); err != nil {
		return gg, errors.WithStack(err)
	}

	g := genny.New()
	g.Command(exec.Command("git", "init"))

	if err := g.Box(packr.NewBox("../gonit/templates")); err != nil {
		return gg, errors.WithStack(err)
	}
	ctx := plush.NewContext()
	ctx.Set("opts", opts)
	g.Transformer(plushgen.Transformer(ctx))

	g.Transformer(genny.Dot())

	g.RunFn(func(r *genny.Runner) error {
		if !gomods.On() {
			return nil
		}
		return r.Exec(exec.Command("go", "mod", "init"))
	})

	gg.Add(g)

	g2, err := initgen.New(opts.Options)
	if err != nil {
		return gg, errors.WithStack(err)
	}
	gg.Merge(g2)

	g = genny.New()
	g.RunFn(func(r *genny.Runner) error {
		if !gomods.On() {
			return nil
		}
		return r.Exec(exec.Command("go", "mod", "tidy"))
	})
	gg.Add(g)

	return gg, nil
}
