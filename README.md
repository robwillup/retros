# RetroS

RetroS is a tool to help you manage files with your retro gaming setup
on remote or local machines.

![build workflow](https://github.com/robwillup/retros/actions/workflows/build.yml/badge.svg)
![Latest Release](https://img.shields.io/github/v/release/robwillup/retros?label=RetroS%20(Linux%20binary)&sort=semver)
![Latest Release](https://img.shields.io/github/v/release/robwillup/retros?label=RetroS%20(Windows%20binary)&sort=semver)

<div align="center">
    <img
        src="https://repository-images.githubusercontent.com/709978523/1ebe6c81-8dfd-499a-a194-4bbfefe65243"
        alt="RetroS" style="width: 550px;"/>
        <p>Image by @robwillup<p>
</div>

## Status

Early development. There are many features to be added and there may be bugs in
current versions.

## Piracy

This project does NOT provide any game ROM files and does not encourage obtaining
such files illegally.

## Download

You can download the latest binary from [releases](https://github.com/robwillup/retros/releases)
or use one of the commands below:

> For Windows, you may need to download using the command below. When downloading
> from the browser, Windows Defender may remove it as a virus.

### Linux

```bash
# wget
wget https://github.com/robwillup/retros/releases/download/v1.0.5/retros \
&& chmod +x retros
```

```bash
# cURL
curl -L https://github.com/robwillup/retros/releases/download/v1.0.5/retros \
-o retros && chmod +x retros
```

### Windows

```powershell
iwr "https://github.com/robwillup/retros/releases/download/v1.0.5/retros.exe" -o "retros.exe"
```

### macOS

```bash
curl -L https://github.com/robwillup/retros/releases/download/v1.0.5/osx-retros \
-o retros && chmod +x retros
```

## Setup

RetroS will need the `host (IP address)` and the `username` of the remote machine
where RetroPie is running. You can configure that by running:

```bash
retros cf
```

RetroS assumes your `SSH private key` is in the default path, i.e.: `$HOME/.ssh/id_rsa`.

## Commands

These are the commands currently available.

### Copy

Adding a single ROM file to RetroPie:

```bash
retros cp Game.md
```

The ROM file will be added to the corresponding folder in RetroPie based on
the ROM file extension, in the case above `megadrive`.

Adding all ROM files in a directory:

```bash
retros cp /home/gamer/roms
```

All ROM files in the provided directory will be copied to their respective
emulators in RetroPie based on the file extensions.

#### Specifying the emulator

If a ROM file has a different extension, you can copy it by
specifying the emulator it should go into:

```bash
retros cp --emulator=atari2600 Game.bin
```

### List

Lists the ROM files in RetroPie:

```bash
retros ls
```

To list ROM files for specific emulators:

```bash
retros ls -e=mastersystem
```

### Check

The `check` command verifies the integrity of ROM files.
The example below shows how to check all ROM files in the `snes` directory:

```bash
retros check ~/games/snes
```

## Roadmap

The intent for RetroS is that it will be used to sync other files and assets with
RetroPie besides ROMs.

Much is still under consideration, but some ideas are being added to the
[project here](https://github.com/users/robwillup/projects/1).
