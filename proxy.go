package nodepmproxy

import (
	"errors"
	"os"
	"path"

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

func detectPM(o *Options) error {
	log.Debug().Msg("Mode AUTO. Trying to detect package manager.")
	if fileExists(path.Join(o.SitePath, "bun.lock")) {
		log.Debug().Msg("Detected Bun")
		WithBun(o)
	} else if fileExists(path.Join(o.SitePath, "pnpm-lock.yaml")) {
		log.Debug().Msg("Detected PNPM")
		WithPnpm(o)
	} else if fileExists(path.Join(o.SitePath, "yarn.lock")) {
		log.Debug().Msg("Detected Yarn")
		WithYarn(o)
	} else {
		return errors.New("no package manager lock file detected")
	}
	return nil
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
