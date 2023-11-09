package emulators

import "path/filepath"

func FindEmulatorFromExtension(romFile string) string {
	switch filepath.Ext(romFile) {
		case ".a26":
			return "atari2600"
		case ".a78":
			return "atari7800"
		case ".a52":
			return "atari5200"
		case ".j64":
			fallthrough
		case ".jag":
			return "atarijaguar"
		case ".lnx":
			return "atarilynx"
		case ".st":
			fallthrough
		case ".stx":
			return "atarist"
		case ".col":
			return "coleco"
		case ".gba":
			return "gba"
		case ".gbc":
			return "gbc"
		case ".gb":
			return "gb"
		case ".gcm":
			fallthrough
		case ".gcz":
			return "gc"
		case ".gg":
			return "gamegear"
		case ".gen":
			fallthrough
		case ".md":
			return "megadrive"
		case ".sms":
			return "mastersystem"
		case ".z64":
			fallthrough
		case ".v64":
			fallthrough
		case ".n64":
			return "n64"
		case ".nds":
			return "nds"
		case ".fds":
			fallthrough
		case ".nes":
			return "nes"
		case ".32x":
			return "sega32x"
		case ".cue":
			fallthrough
		case ".chd":
			return "segacd"
		case ".sfc":
			fallthrough
		case ".smc":
			return "snes"
		default:
			return ""
	}
}
