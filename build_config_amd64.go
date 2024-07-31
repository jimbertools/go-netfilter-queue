//go:build linux && amd64

package netfilter

/*
#cgo CFLAGS: -Wall -I/usr/include
#cgo LDFLAGS: -L/usr/lib/mips-linux-gnu -lnetfilter_queue
*/

import "C"
