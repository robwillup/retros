package platform

import "path/filepath"

func FindPlatformFromExtension(romFile string) string {
	switch filepath.Ext(romFile) {
		case ".sfc":
			return "snes"
		case ".md":
			return "genesis"
		case ".sms":
			return "mastersystem"
		case ".gba":
			return "gba"
		case ".a26":
			return "atari2600"
		default:
			return ""
	}
}