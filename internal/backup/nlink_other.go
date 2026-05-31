//go:build !unix

package backup

// nlink returns a sentinel link count of 2 on platforms where st_nlink is not
// readily available (e.g. Windows). Returning >1 makes gcOrphanedCAS
// conservative: it never deletes a CAS object it cannot prove is orphaned, so
// the store may grow but never loses content a batch still needs.
func nlink(_ string) uint64 {
	return 2
}
