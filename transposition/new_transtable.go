// +build !linux

package transposition

import "unsafe"
import . "github.com/David-Orson/casperin/utils"

func NewTransTable(megabytes int) TranspositionTable {
	size := NearestPowerOfTwo(1024 * 1024 * megabytes / int(unsafe.Sizeof(transEntry{})))
	return TranspositionTable{make([]transEntry, size), size - 1}
}
