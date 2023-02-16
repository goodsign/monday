//go:build debug
// +build debug

package monday

import "log"

const (
	debugLayoutDef = on
)

func init() {
	log.SetFlags(log.Flags() | log.Lshortfile)
}
