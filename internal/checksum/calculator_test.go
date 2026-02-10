package checksum

import "testing"

//goland:noinspection SpellCheckingInspection
var checksum = ROM{
	Name:   "calculator.go",
	MD5:    "83e321154c56b253443bb031759e4b23",
	SHA1:   "1b6ff79552dc91d638b867064012f028f4e7e096",
	SHA256: "3dacea3020cbdd866d2c1156bedaf5f065cfd37f5516aed64258fb6bf247b9e0",
	Size:   895,
}

func TestCalcChecksum(t *testing.T) {
	actual, err := CalcChecksum("calculator.go")

	if err != nil {
		t.Fatalf("Failed to TestCalcChecksum().\nError: %v\n", err)
	}

	if actual != checksum {
		t.Fatalf("TestCalcChecksum failed.\nExpected: %v\nActual: %v\n", checksum, actual)
	}
}
