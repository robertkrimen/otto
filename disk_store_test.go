package otto

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDiskStoreNotConfigured(t *testing.T) {
	vm := New()

	if _, err := vm.SaveToDisk(UndefinedValue(), "x.json"); err != ErrDiskStoreNotConfigured {
		t.Fatalf("SaveToDisk: expected ErrDiskStoreNotConfigured, got %v", err)
	}
	if _, err := vm.LoadFromDisk("x.json"); err != ErrDiskStoreNotConfigured {
		t.Fatalf("LoadFromDisk: expected ErrDiskStoreNotConfigured, got %v", err)
	}
}

func testDiskStoreRoundTrip(t *testing.T, compress bool) {
	dir := t.TempDir()
	vm := New()
	vm.DiskStore = &DiskStoreConfig{Dir: dir, Compress: compress}

	v, err := vm.Run(`
		(function() {
			var big = [];
			for (var i = 0; i < 1000; i++) {
				big.push({ id: i, name: "item-" + i, active: i % 2 === 0 });
			}
			return big;
		})()
	`)
	if err != nil {
		t.Fatalf("Run: %v", err)
	}

	path, err := vm.SaveToDisk(v, "big.json")
	if err != nil {
		t.Fatalf("SaveToDisk: %v", err)
	}
	if filepath.Dir(path) != dir {
		t.Fatalf("SaveToDisk: expected file under %q, got %q", dir, path)
	}
	if _, err := os.Stat(path); err != nil {
		t.Fatalf("expected file to exist: %v", err)
	}

	loaded, err := vm.LoadFromDisk(path)
	if err != nil {
		t.Fatalf("LoadFromDisk: %v", err)
	}

	// Compare via JSON.stringify inside the VM so we're checking JS-level
	// equality of the round-tripped structure, not Go representation details.
	orig, err := vm.Call("JSON.stringify", nil, v)
	if err != nil {
		t.Fatalf("JSON.stringify(orig): %v", err)
	}
	got, err := vm.Call("JSON.stringify", nil, loaded)
	if err != nil {
		t.Fatalf("JSON.stringify(loaded): %v", err)
	}
	if orig.String() != got.String() {
		t.Fatalf("round-trip mismatch:\nwant %s\ngot  %s", orig.String(), got.String())
	}
}

func TestDiskStoreRoundTripPlain(t *testing.T) {
	testDiskStoreRoundTrip(t, false)
}

func TestDiskStoreRoundTripCompressed(t *testing.T) {
	testDiskStoreRoundTrip(t, true)
}

func TestDiskStoreCompressedIsSmaller(t *testing.T) {
	dir := t.TempDir()
	vm := New()

	v, err := vm.Run(`
		(function() {
			var s = [];
			for (var i = 0; i < 5000; i++) { s.push("repeat-me-" + (i % 5)); }
			return s;
		})()
	`)
	if err != nil {
		t.Fatalf("Run: %v", err)
	}

	vm.DiskStore = &DiskStoreConfig{Dir: dir, Compress: false}
	plainPath, err := vm.SaveToDisk(v, "plain.json")
	if err != nil {
		t.Fatalf("SaveToDisk(plain): %v", err)
	}

	vm.DiskStore = &DiskStoreConfig{Dir: dir, Compress: true}
	gzPath, err := vm.SaveToDisk(v, "compressed.json.gz")
	if err != nil {
		t.Fatalf("SaveToDisk(compressed): %v", err)
	}

	plainInfo, err := os.Stat(plainPath)
	if err != nil {
		t.Fatal(err)
	}
	gzInfo, err := os.Stat(gzPath)
	if err != nil {
		t.Fatal(err)
	}

	if gzInfo.Size() >= plainInfo.Size() {
		t.Fatalf("expected compressed size (%d) < plain size (%d)", gzInfo.Size(), plainInfo.Size())
	}
}
