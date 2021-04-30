package tateru

import (
	"github.com/chunni/fiptoml"
	"unsafe"
)

// https://stackoverflow.com/a/17982725
//go:nosplit
func ExposeTomlDict(t **fiptoml.Toml) *map[string]interface{} {
	return *(**map[string]interface{}) (unsafe.Pointer(t))
}
