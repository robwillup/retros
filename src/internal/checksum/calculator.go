package checksum

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
)

func CalcChecksum(romPath string) (ROM, error) {
	rom := ROM{}
	f, err := os.Open(filepath.Clean(romPath))

	if err != nil {
		return rom, err
	}

	defer f.Close()

	sha256 := sha256.New()
	size, err := io.Copy(sha256, f)

	if err != nil {
		return rom, err
	}

	_, err = f.Seek(0, 0)
	if err != nil {
		return rom, err
	}

	md5 := md5.New()
	_, err = io.Copy(md5, f)

	if err != nil {
		return rom, err
	}

	_, err = f.Seek(0, 0)
	if err != nil {
		return rom, err
	}

	sha1 := sha1.New()
	_, err = io.Copy(sha1, f)

	if err != nil {
		return rom, err
	}

	rom.Name = filepath.Base(romPath)
	rom.MD5 = hex.EncodeToString(md5.Sum(nil))
	rom.SHA1 = hex.EncodeToString(sha1.Sum(nil))
	rom.SHA256 = hex.EncodeToString(sha256.Sum(nil))
	rom.Size = int(size)

	return rom, nil
}
