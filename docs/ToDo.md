# To Do

## Remove

Removes ROM files from RetroPie:

```bash
rosy rm Game.sms
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