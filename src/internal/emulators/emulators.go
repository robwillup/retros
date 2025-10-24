package emulators

import "path/filepath"

func FindEmulatorFromExtension(romFile string) string {
	switch filepath.Ext(romFile) {
	case ".32x":
		return "sega32x"
	case ".a26":
		return "atari2600"
	case ".a52":
		return "atari5200"
	case ".a78":
		return "atari7800"
	case ".col":
		return "coleco"
	case ".gb":
		return "gb"
	case ".gba":
		return "gba"
	case ".gbc":
		return "gbc"
	case ".gcm":
		fallthrough
	case ".gcz":
		return "gc"
	case ".gen":
		return "genesis"
	case ".gg":
		return "gamegear"
	case ".j64":
		fallthrough
	case ".jag":
		return "atarijaguar"
	case ".lnx":
		return "atarilynx"
	case ".md":
		return "megadrive"
	case ".n64":
		return "n64"
	case ".nds":
		return "nds"
	case ".nes":
		return "nes"
	case ".pce":
		return "pcengine"
	case ".sfc":
		fallthrough
	case ".smc":
		return "snes"
	case ".sg":
		return "sg-1000"
	case ".sms":
		return "mastersystem"
	case ".st":
		fallthrough
	case ".stx":
		return "atarist"
	default:
		return ""
	}
}
