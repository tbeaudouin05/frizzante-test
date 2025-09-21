package runtime

import (
	"github.com/dop251/goja"
	"main/lib/core/js"
)

// WithFunction sets a function.
func WithFunction(run *goja.Runtime, name string, call js.Function) error {
	return run.Set(name, call)
}
