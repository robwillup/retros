package checksum

import "testing"

//goland:noinspection SpellCheckingInspection
var checksum = ROM{
	Name:   "main.go",
	MD5:    "28766063f362c5d97e2288f3a3be3e78",
	SHA1:   "6c3c98f00928520e80efcb2aa1cb8a11ba0f5d6f",
	SHA256: "14ae8e67feedb5c9f2bae2d21d9b87260ec6922c6717210da39fe7b4e401d043",
	Size:   1168,
}

func TestCalcChecksum(t *testing.T) {
	actual, err := CalcChecksum("../../main.go")

	if err != nil {
		t.Fatalf("Failed to TestCalcChecksum().\nError: %v\n", err)
	}

	if actual != checksum {
		t.Fatalf("TestCalcChecksum failed.\nExpected: %v\nActual: %v\n", checksum, actual)
	}
}
