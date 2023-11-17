package checksum

import "testing"

var keystone = ROM{
	Name: "Keystone Kapers (USA).a26",
	MD5: "be929419902e21bd7830a7a7d746195d",
	SHA1: "3eefc193dec3b242bcfd43f5a4d9f023e55378a4",
	SHA256: "",
	Size: 4096,
}

func TestGetChecksums(t *testing.T) {
	checksums, err := GetChecksums("atari2600", "data/")

	if err != nil {
		t.Fatalf("Failed to TestGetChecksums().\nError: %v\n", err)
	}

	c, ok := checksums[keystone.MD5]

	if !ok {
		t.Fatalf("Failed to TestGetChecksums().\nExpected to find MD5: %s\n", keystone.MD5)
	}

	if c != keystone {
		t.Fatalf("TestGetChecksums failed.\nExpected: %v\nActual: %v\n", keystone, c)
	}
}