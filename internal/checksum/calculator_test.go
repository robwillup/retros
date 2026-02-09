package checksum

import "testing"

//goland:noinspection SpellCheckingInspection
var checksum = ROM{
	Name:   "main.go",
	MD5:    "b7beefafe8c7bfc37e2f127465fdc60d",
	SHA1:   "c82b6fd625357a04fde423f6580b677c5fab5e6d",
	SHA256: "a8cb0bb73c28ba4b79a76f161315eaad734af2bffc6000a91cbb93925d63d9a7",
	Size:   1164,
}

func TestCalcChecksum(t *testing.T) {
	actual, err := CalcChecksum("../../cmd/retros/main.go")

	if err != nil {
		t.Fatalf("Failed to TestCalcChecksum().\nError: %v\n", err)
	}

	if actual != checksum {
		t.Fatalf("TestCalcChecksum failed.\nExpected: %v\nActual: %v\n", checksum, actual)
	}
}
