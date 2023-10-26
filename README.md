# RoSy

ROM Sync for RetroPie through SSH

## Features

* Copies ROM files
* Lists ROM files
* Removes ROM files
* Compares ROM files
* Verifies Checksum of ROM files
* Syncs ROM files

## Commands

These are the commands currently being implemented.

### cp

Adding ROM files to RetroPie:

```bash
rosy cp Game.md

# or multiple ROM files
rosy cp Game.md OtherGame.sfc
```

The ROM will be added to the corresponding folder in RetroPie based on
the ROM file extension, in the cases above Mega Drive/Genesis and SNES.

To force overwrite:

```bash
rosy cp -f Game.md
```

## ls

Lists the ROM files in RetroPie:

```bash
rosy ls
```

To list ROM files for specific platforms:

```bash
rosy ls -p mastersystem
```

## rm

Removes ROM files from RetroPie:

```bash
rosy rm Game.ms
```

Removes a ROM file only if a copy exists locally:

```bash
rosy rm -s Game.ms
```

## diff

Compares two ROM files:

```bash
rosy diff Game.gba Game.gba
```

The first ROM file is in the local system and the second in RetroPie.

## cs

Verifies the integrity of a ROM file using the checksum:

```bash
rosy cs Game.ms
```

## Sync

Copies missing verified ROM files from the local system to RetroPie:

```bash
rosy sync
```
