//go:build unix

package backup

import (
	"os"
	"syscall"
)

// nlink returns the hardlink count of the file at path, or 0 if it cannot be
// stat'd. On unix this reads st_nlink via syscall.Stat_t.
func nlink(path string) uint64 {
	info, err := os.Lstat(path)
	if err != nil {
		return 0
	}
	st, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		return 0
	}
	return uint64(st.Nlink)
}
