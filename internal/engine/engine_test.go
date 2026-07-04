package engine

import (
	"path/filepath"
	"testing"
)

// newTestEngine opens an engine backed by a fresh temp file that is cleaned up
// automatically when the test finishes.
func newTestEngine(t *testing.T) *Engine {
	t.Helper()
	path := filepath.Join(t.TempDir(), "test.log")
	db, err := Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	t.Cleanup(func() { db.Close() })
	return db
}

func TestSetAndGet(t *testing.T) {
	db := newTestEngine(t)

	if err := db.Set("name", "runedb"); err != nil {
		t.Fatalf("Set: %v", err)
	}

	got, err := db.Get("name")
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if got != "runedb" {
		t.Fatalf("Get = %q, want %q", got, "runedb")
	}
}

func TestGetMissingKey(t *testing.T) {
	db := newTestEngine(t)

	if _, err := db.Get("nope"); err != ErrKeyNotFound {
		t.Fatalf("Get err = %v, want ErrKeyNotFound", err)
	}
}

func TestSetOverwritesValue(t *testing.T) {
	db := newTestEngine(t)

	if err := db.Set("k", "old"); err != nil {
		t.Fatalf("Set old: %v", err)
	}
	if err := db.Set("k", "new"); err != nil {
		t.Fatalf("Set new: %v", err)
	}

	got, err := db.Get("k")
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if got != "new" {
		t.Fatalf("Get = %q, want %q", got, "new")
	}
}

func TestDelete(t *testing.T) {
	db := newTestEngine(t)

	if err := db.Set("k", "v"); err != nil {
		t.Fatalf("Set: %v", err)
	}
	if err := db.Delete("k"); err != nil {
		t.Fatalf("Delete: %v", err)
	}

	if _, err := db.Get("k"); err != ErrKeyNotFound {
		t.Fatalf("Get after delete err = %v, want ErrKeyNotFound", err)
	}
}

// TestReopenRebuildsIndex is the key persistence test: after closing and
// reopening the same file, the in-memory index must be rebuilt from the log
// alone and reflect the latest state (overwrites and deletes included).
func TestReopenRebuildsIndex(t *testing.T) {
	path := filepath.Join(t.TempDir(), "reopen.log")

	db, err := Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	mustSet(t, db, "a", "1")
	mustSet(t, db, "b", "2")
	mustSet(t, db, "a", "updated") // overwrite
	mustSet(t, db, "c", "3")
	if err := db.Delete("b"); err != nil {
		t.Fatalf("Delete: %v", err)
	}
	db.Close()

	reopened, err := Open(path)
	if err != nil {
		t.Fatalf("reopen: %v", err)
	}
	defer reopened.Close()

	if got, _ := reopened.Get("a"); got != "updated" {
		t.Fatalf("after reopen a = %q, want %q", got, "updated")
	}
	if got, _ := reopened.Get("c"); got != "3" {
		t.Fatalf("after reopen c = %q, want %q", got, "3")
	}
	if _, err := reopened.Get("b"); err != ErrKeyNotFound {
		t.Fatalf("after reopen b err = %v, want ErrKeyNotFound", err)
	}
}

func mustSet(t *testing.T, db *Engine, key, value string) {
	t.Helper()
	if err := db.Set(key, value); err != nil {
		t.Fatalf("Set(%q, %q): %v", key, value, err)
	}
}
