package gonit

import "github.com/gobuffalo/release/genny/initgen"

type Options struct {
	*initgen.Options
}

// Validate that options are usuable
func (opts *Options) Validate() error {
	if opts.Options == nil {
		opts.Options = &initgen.Options{}
	}
	return opts.Options.Validate()
}
