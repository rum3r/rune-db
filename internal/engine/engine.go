package engine

import (
	"bufio"
	"errors"
	"io"
	"os"
	"sync"
)

// Storage is the behaviour every storage backend must provide.
type Storage interface {
	Set(key, value string) error
	Get(key string) (string, error)
	Delete(key string) error
}

// ErrKeyNotFound is returned by Get when the key is not in the store.
var ErrKeyNotFound = errors.New("key not found")

// location tells us where a key's most recent value lives in the log file.
type location struct {
	offset int64  // byte offset of the value within the file
	length uint32 // length of the value in bytes
}

// Engine is a log-structured key-value store. Every write is appended to a
// single file; an in-memory index maps each key to the location of its latest
// value so reads can seek straight to it instead of scanning the whole log.
type Engine struct {
	mu    sync.RWMutex
	file  *os.File
	index map[string]location
	end   int64 // size of the file, i.e. the offset of the next append
}

// Open opens (creating it if needed) a log-structured store backed by path.
func Open(path string) (*Engine, error) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	e := &Engine{
		file:  file,
		index: make(map[string]location),
	}

	if err := e.buildIndex(); err != nil {
		file.Close()
		return nil, err
	}
	return e, nil
}

// Close releases the underlying file.
func (e *Engine) Close() error {
	return e.file.Close()
}

func (e *Engine) buildIndex() error {
	if _, err := e.file.Seek(0, io.SeekStart); err != nil {
		return err
	}

	reader := bufio.NewReader(e.file)
	var offset int64
	for {
		rec, size, err := decodeRecord(reader)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		valueOffset := offset + headerSize + int64(len(rec.key))
		switch rec.op {
		case opPut:
			e.index[rec.key] = location{offset: valueOffset, length: uint32(len(rec.value))}
		case opDelete:
			delete(e.index, rec.key)
		}
		offset += int64(size)
	}

	e.end = offset
	return nil
}

// Set appends a put record to the log and points the index at its value.
func (e *Engine) Set(key, value string) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	if err := e.append(record{op: opPut, key: key, value: value}); err != nil {
		return err
	}

	valueOffset := e.end + headerSize + int64(len(key))
	e.index[key] = location{offset: valueOffset, length: uint32(len(value))}
	e.end += int64(headerSize + len(key) + len(value))
	return nil
}

// Get finds the key in the index and reads its value straight from the file.
func (e *Engine) Get(key string) (string, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()

	loc, ok := e.index[key]
	if !ok {
		return "", ErrKeyNotFound
	}

	buf := make([]byte, loc.length)
	if _, err := e.file.ReadAt(buf, loc.offset); err != nil {
		return "", err
	}
	return string(buf), nil
}

// Delete appends a delete record to the log and drops the key from the index.
func (e *Engine) Delete(key string) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	if err := e.append(record{op: opDelete, key: key}); err != nil {
		return err
	}

	delete(e.index, key)
	e.end += int64(headerSize + len(key))
	return nil
}

// append writes a single record to the end of the log file. The file is opened
// in append mode, so every write lands at the current end of the file.
func (e *Engine) append(rec record) error {
	_, err := e.file.Write(rec.encode())
	return err
}
