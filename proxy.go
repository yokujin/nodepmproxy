package nodepmproxy

const (
	YARN = iota
	BUN
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

	return &NodePMProxy{
		Options: o,
	}
}
