package platform

import "testing"

func TestFindPlatformFromExtension(t *testing.T) {
	expected := "snes"
	actual := FindPlatformFromExtension("C:\\Users\\Gamer\\ReallyCoolGame.sfc")

	if expected != actual {
		t.Fatalf("Failed to FindPlatformFromExtension().\nexpected: %s\nactual: %s", expected, actual)
		return
	}

	expected = "genesis"
	actual = FindPlatformFromExtension("gamer/games/ReallyHardGame.md")

	if expected != actual {
		t.Fatalf("Failed to FindPlatformFromExtension().\nexpected: %s\nactual: %s", expected, actual)
		return
	}

	expected = "mastersystem"
	actual = FindPlatformFromExtension("/home/gamer/ReallyNostalgicGame.sms")

	if expected != actual {
		t.Fatalf("Failed to FindPlatformFromExtension().\nexpected: %s\nactual: %s", expected, actual)
		return
	}

	expected = "gba"
	actual = FindPlatformFromExtension("C:\\Users\\Gamer\\ReallyShortGame.gba")

	if expected != actual {
		t.Fatalf("Failed to FindPlatformFromExtension().\nexpected: %s\nactual: %s", expected, actual)
		return
	}

	expected = "atari2600"
	actual = FindPlatformFromExtension("C:\\Users\\Gamer\\ReallyOld.a26")

	if expected != actual {
		t.Fatalf("Failed to FindPlatformFromExtension().\nexpected: %s\nactual: %s", expected, actual)
		return
	}

	expected = "sega32x"
	actual = FindPlatformFromExtension("/home/gamer/ReallyRareGame.32x")

	if expected != actual {
		t.Fatalf("Failed to FindPlatformFromExtension().\nexpected: %s\nactual: %s", expected, actual)
		return
	}
}