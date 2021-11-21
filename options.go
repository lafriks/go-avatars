package avatars

// Options for the avatar service.
type Options struct {
	// Size of the rendered avatar.
	Size int
	// Padding between the rendered avatar and the border.
	Padding int
}

// Option is a function that applies an option to the avatar service.
type Option interface {
	apply(a *Options)
}

// RenderSize sets the size of the rendered avatar in pixels.
type RenderSize int

func (s RenderSize) apply(a *Options) {
	a.Size = int(s)
}

// Padding sets the padding between the rendered avatar and the border.
type Padding int

func (s Padding) apply(a *Options) {
	a.Padding = int(s)
}

func options(opts ...Option) *Options {
	opt := &Options{
		Size: 256,
	}
	for _, o := range opts {
		o.apply(opt)
	}
	return opt
}
