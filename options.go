package avatars

// Options for the avatar service.
type Options struct{}

// Option is a function that applies an option to the avatar service.
type Option interface {
	apply(a *Options)
}

func options(opts ...Option) *Options {
	opt := &Options{}
	for _, o := range opts {
		o.apply(opt)
	}
	return opt
}

// RenderOptions is a set of options for rendering an avatar.
type RenderOptions struct {
	// Size of the rendered avatar.
	Size int
	// Padding between the rendered avatar and the border.
	Padding int
}

// Render is a function that applies a render option to the avatar service.
type RenderOption interface {
	apply(a *RenderOptions)
}

// RenderSize sets the size of the rendered avatar in pixels.
type RenderSize int

func (s RenderSize) apply(a *RenderOptions) {
	a.Size = int(s)
}

// RenderPadding sets the padding between the rendered avatar and the border.
type RenderPadding int

func (s RenderPadding) apply(a *RenderOptions) {
	a.Padding = int(s)
}

func renderOptions(opts ...RenderOption) *RenderOptions {
	opt := &RenderOptions{
		Size: 256,
	}
	for _, o := range opts {
		o.apply(opt)
	}
	return opt
}
