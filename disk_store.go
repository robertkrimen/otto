package otto

import (
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// ErrDiskStoreNotConfigured is returned by SaveToDisk and LoadFromDisk when
// the Otto instance's DiskStore field has not been set. Disk-backed storage
// is opt-in: an *Otto returned by New() has DiskStore == nil, so these
// methods are no-ops (return this error) until a caller explicitly enables
// the feature.
var ErrDiskStoreNotConfigured = errors.New("otto: DiskStore not configured")

// DiskStoreConfig enables optional disk-backed storage of large JavaScript
// values for a single Otto instance. Set Otto.DiskStore after New() to opt
// in; leave it nil (the default) to leave the runtime's behavior unchanged.
type DiskStoreConfig struct {
	// Dir is the directory SaveToDisk writes files into. Defaults to
	// os.TempDir() when empty.
	Dir string

	// Compress gzip-compresses the JSON-encoded value before writing it to
	// disk, and LoadFromDisk decompresses it transparently on read.
	Compress bool
}

func (c *DiskStoreConfig) dir() string {
	if c == nil || c.Dir == "" {
		return os.TempDir()
	}
	return c.Dir
}

// SaveToDisk serializes v to JSON (via Value's existing MarshalJSON, i.e. the
// same encoding JSON.stringify would produce) and writes it to a file named
// name inside o.DiskStore.Dir, gzip-compressing it first when
// o.DiskStore.Compress is true. It returns the full path written.
//
// SaveToDisk returns ErrDiskStoreNotConfigured if o.DiskStore is nil.
func (o Otto) SaveToDisk(v Value, name string) (string, error) {
	if o.DiskStore == nil {
		return "", ErrDiskStoreNotConfigured
	}

	data, err := v.MarshalJSON()
	if err != nil {
		return "", fmt.Errorf("otto: marshal value for disk store: %w", err)
	}

	path := filepath.Join(o.DiskStore.dir(), name)

	f, err := os.Create(path)
	if err != nil {
		return "", fmt.Errorf("otto: create disk store file: %w", err)
	}
	defer f.Close()

	if !o.DiskStore.Compress {
		if _, err := f.Write(data); err != nil {
			return "", fmt.Errorf("otto: write disk store file: %w", err)
		}
		return path, nil
	}

	gw := gzip.NewWriter(f)
	if _, err := gw.Write(data); err != nil {
		gw.Close()
		return "", fmt.Errorf("otto: write compressed disk store file: %w", err)
	}
	if err := gw.Close(); err != nil {
		return "", fmt.Errorf("otto: flush compressed disk store file: %w", err)
	}

	return path, nil
}

// LoadFromDisk reads path back (transparently gunzipping it when
// o.DiskStore.Compress is true) and parses it into a Value using this
// runtime's own JSON.parse, so the result behaves like any other JavaScript
// value created inside o.
//
// LoadFromDisk returns ErrDiskStoreNotConfigured if o.DiskStore is nil.
func (o Otto) LoadFromDisk(path string) (Value, error) {
	if o.DiskStore == nil {
		return Value{}, ErrDiskStoreNotConfigured
	}

	f, err := os.Open(path)
	if err != nil {
		return Value{}, fmt.Errorf("otto: open disk store file: %w", err)
	}
	defer f.Close()

	var r io.Reader = f
	if o.DiskStore.Compress {
		gr, err := gzip.NewReader(f)
		if err != nil {
			return Value{}, fmt.Errorf("otto: open compressed disk store file: %w", err)
		}
		defer gr.Close()
		r = gr
	}

	data, err := io.ReadAll(r)
	if err != nil {
		return Value{}, fmt.Errorf("otto: read disk store file: %w", err)
	}

	return o.Call("JSON.parse", nil, string(data))
}
