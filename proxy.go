package nodepmproxy

const (
	YARN = iota
	BUN
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
