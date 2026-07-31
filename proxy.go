package nodepmproxy

import (
	"github.com/rs/zerolog/log"
)

const (
	MANUAL = iota
	YARN
	BUN
	PNPM
	AUTO
)

const (
	NUXT = iota
	SVELTE
)

type NodePMProxy struct {
	Options
}

func New(opts ...OptionFn) *NodePMProxy {
	o := defaultOptions()

	for _, fn := range opts {
		fn(&o)
	}

	if o.PM == AUTO {
		err := detectPM(&o)
		if err != nil {
			log.Fatal().Err(err).Msg("cannot auto-detect package manager")
		}
	}

	return &NodePMProxy{
		Options: o,
	}
}
