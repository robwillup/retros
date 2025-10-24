package checksum

import "testing"

//goland:noinspection SpellCheckingInspection
var checksum = ROM{
	Name:   "main.go",
	MD5:    "3c6498d3c1a2ed05106815bffac3e235",
	SHA1:   "09d65f6d4a4dc451fd70c01028a06fc5d0adde7c",
	SHA256: "eb70de22ccf7b3ca3a56886d7e5adcb3e4f63ccbf1d7f6b5e2413935c38b8164",
	Size:   1154,
}

func TestCalcChecksum(t *testing.T) {
	actual, err := CalcChecksum("../main.go")

	if err != nil {
		t.Fatalf("Failed to TestCalcChecksum().\nError: %v\n", err)
	}

	if actual != checksum {
		t.Fatalf("TestCalcChecksum failed.\nExpected: %v\nActual: %v\n", checksum, actual)
	}
}
