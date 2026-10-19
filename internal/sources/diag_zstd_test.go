package sources

import "testing"

func TestDiagRestoreNoZstd(t *testing.T) {
	DiagSnapshot()
	diagNoZstd("a")
	got := DiagNoZstdPaths()
	DiagSnapshot()
	if len(DiagNoZstdPaths()) != 0 {
		t.Fatal("snapshot kept the paths")
	}
	DiagRestoreNoZstd(got)
	if !DiagNoZstdPaths()["a"] {
		t.Fatal("path not restored")
	}
	DiagSnapshot()
}
