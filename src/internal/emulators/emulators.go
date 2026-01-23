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

func GetEmulatorDisplayName(shortName string) string {
	switch shortName {
	case "amstradcpc":
		return "Amstrad CPC"
	case "arcade":
		return "Arcade"
	case "atari2600":
		return "Atari 2600"
	case "atari5200":
		return "Atari 5200"
	case "atari7800":
		return "Atari 7800"
	case "atari800":
		return "Atari 8-bit Family"
	case "atarilynx":
		return "Atari Lynx"
	case "channelf":
		return "Fairchild Channel F"
	case "coleco":
		return "ColecoVision"
	case "fba":
		return "FinalBurn Alpha"
	case "fds":
		return "Famicom Disk System"
	case "gamegear":
		return "Sega Game Gear"
	case "gb":
		return "Game Boy"
	case "gba":
		return "Game Boy Advance"
	case "gbc":
		return "Game Boy Color"
	case "genesis":
		return "Sega Genesis - Mega Drive"
	case "mame-libretro":
		return "MAME"
	case "mastersystem":
		return "Sega Master System"
	case "megadrive":
		return "Sega Genesis - Mega Drive"
	case "msx":
		return "MSX"
	case "n64":
		return "Nintendo 64"
	case "neogeo":
		return "Neo Geo"
	case "nes":
		return "NES - Nintendo Entertainment System"
	case "ngp":
		return "Neo Geo Pocket"
	case "ngpc":
		return "Neo Geo Pocket Color"
	case "pcengine":
		return "PC Engine / TurboGrafx-16"
	case "psx":
		return "PlayStation"
	case "sega32x":
		return "Sega 32X"
	case "segacd":
		return "Sega CD"
	case "sg-1000":
		return "SG-1000"
	case "snes":
		return "SNES - Super Nintendo Entertainment System"
	case "vectrex":
		return "Vectrex"
	case "virtualboy":
		return "Virtual Boy"
	case "zxspectrum":
		return "ZX Spectrum"
	default:
		return shortName
	}
}
