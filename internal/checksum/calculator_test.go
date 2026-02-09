package checksum

import "testing"

//goland:noinspection SpellCheckingInspection
var checksum = ROM{
	Name:   "main.go",
	MD5:    "7fbd5e9ea3b5f622ba3a3e47badee8fe",
	SHA1:   "58598c1f53f58075ea7ec7e35114b4ddaf8b671b",
	SHA256: "e74d549319164be4084c0d22f3e1b5b6cf3356c6ad8ea843521a39565d0f33c0",
	Size:   1174,
}

func TestCalcChecksum(t *testing.T) {
	actual, err := CalcChecksum("../../cmd/main.go")

	if err != nil {
		t.Fatalf("Failed to TestCalcChecksum().\nError: %v\n", err)
	}

	if actual != checksum {
		t.Fatalf("TestCalcChecksum failed.\nExpected: %v\nActual: %v\n", checksum, actual)
	}
}
