package converter

import (
	"io/fs"
	"testing"
)

type mockDirEntry struct {
	name string
}

func (m mockDirEntry) Name() string               { return m.name }
func (m mockDirEntry) IsDir() bool                { return false }
func (m mockDirEntry) Type() fs.FileMode          { return 0 }
func (m mockDirEntry) Info() (fs.FileInfo, error) { return nil, nil }

func TestSortPagesAndChapters(t *testing.T) {
	files := []fs.DirEntry{
		mockDirEntry{name: "2-10-uuid.jpg"},
		mockDirEntry{name: "2-2-uuid.jpg"},
		mockDirEntry{name: "1-5-uuid.jpg"},
		mockDirEntry{name: "1-3-uuid.jpg"},
		mockDirEntry{name: "2-1-uuid.jpg"},
	}

	sortPagesAndChapters(files)

	expectedOrder := []string{
		"1-3-uuid.jpg",
		"1-5-uuid.jpg",
		"2-1-uuid.jpg",
		"2-2-uuid.jpg",
		"2-10-uuid.jpg",
	}

	for i, file := range files {
		if file.Name() != expectedOrder[i] {
			t.Errorf("File order mismatch at index %d: got %s, want %s", i, file.Name(), expectedOrder[i])
		}
	}
}
