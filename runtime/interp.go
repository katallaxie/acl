package runtime

type interp struct{}

// NewInterp ...
func NewInterp() *interp {
	return &interp{}
}

// HasAccess ...
func (i *interp) HasAccess(ctx Context, path string) bool {
	return true
}
