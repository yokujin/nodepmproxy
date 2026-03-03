package nodepmproxy

import (
	"io/fs"

	"github.com/rs/zerolog/log"
)

type OptionFn func(*Options)

type Options struct {
	Port       int
	PM         int
	Framework  int
	SitePath   string
	Embedded   bool
	EmbeddedFS fs.FS
}

func defaultOptions() Options {
	port, err := getFreePort()
	if err != nil {
		log.Fatal().Err(err).Msg("get free port")
	}

	return Options{
		Port:       port,
		PM:         YARN,
		Framework:  NUXT,
		SitePath:   "",
		Embedded:   false,
		EmbeddedFS: nil,
	}
}

func WithYarn(o *Options) {
	o.PM = YARN
}

func WithBun(o *Options) {
	o.PM = BUN
}

func WithNuxt(o *Options) {
	o.Framework = NUXT
}

func WithSvelte(o *Options) {
	o.Framework = SVELTE
}

func WithSitePath(pth string) OptionFn {
	return func(o *Options) {
		o.SitePath = pth
	}
}

func WithEmbedded(emb fs.FS) OptionFn {
	return func(o *Options) {
		o.Embedded = true
		o.EmbeddedFS = emb
	}
}
