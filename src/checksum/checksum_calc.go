package checksum

import (
	"crypto/sha256"
	"io"
	"os"
	"path/filepath"
)

func CalcSha256(romPath string) ([]byte, error) {
	f, err := os.Open(filepath.Clean(romPath))

	if err != nil {
		return []byte{}, err
	}

	defer f.Close()

	h := sha256.New()

	if _, err := io.Copy(h, f); err != nil {
		return []byte{}, err
	}

	return h.Sum(nil), nil
}