# RomSync

Syncs ROM files through SSH

## Features

* Adds ROM files
* Lists ROM files
* Deletes ROM files
* Syncs ROM files
* Compares ROM files
* Gets Checksum of ROM files

## Commands

These are the commands currently being implemented.

### Add

Adding ROM files to RetroPie:

```bash
rosy add Game.md

# or multiple ROM files
rosy add Game.md OtherGame.sfc
```

The ROM will be added to the corresponding folder in RetroPie based on
the ROM file extension, in the cases above Mega Drive/Genesis and SNES.

To force overwrite:

```bash
rosy add -f Game.md
```
