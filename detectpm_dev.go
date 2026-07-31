//go:build !prod

package nodepmproxy

import (
	"errors"
	"os"
	"path"

	"github.com/rs/zerolog/log"
)

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
