# RetroS

RetroS is a tool to help manage files for retro gaming emulation on remote or
local machines.

![build workflow](https://github.com/robwillup/retros/actions/workflows/build.yml/badge.svg)
![Latest Release](https://img.shields.io/github/v/release/robwillup/retros?label=RetroS%20(Linux%20binary)&sort=semver)
![Latest Release](https://img.shields.io/github/v/release/robwillup/retros?label=RetroS%20(Windows%20binary)&sort=semver)

<div  style="text-align: center;">
    <img
        src="https://repository-images.githubusercontent.com/709978523/1ebe6c81-8dfd-499a-a194-4bbfefe65243"
        alt="RetroS" style="width: 550px;"/>
        <p>Image by @robwillup<p>
</div>

## Status

Early development. There are many features to be added and there may be bugs in
current versions.

## Download

You can download the latest binary from [releases](https://github.com/robwillup/retros/releases)
or use one of the commands below:

> For Windows, you may need to download using the command below. When downloading
> from the browser, Windows Defender may remove it as a virus.

### Linux

```bash
# wget
wget https://github.com/robwillup/retros/releases/download/v1.0.6/retros \
&& chmod +x retros
```

```bash
# cURL
curl -L https://github.com/robwillup/retros/releases/download/v1.0.6/retros \
-o retros && chmod +x retros
```

### Windows

```powershell
iwr "https://github.com/robwillup/retros/releases/download/v1.0.6/retros.exe" -o "retros.exe"
```

### macOS

```bash
curl -L https://github.com/robwillup/retros/releases/download/v1.0.6/osx-retros \
-o retros && chmod +x retros
```

## Setup

RetroS will need the `host (IP address)` and the `username` of the remote machine
where RetroPie is running. You can configure that by running:

```bash
retros cf
```

### SSH keys

To be able to connect to the remote machine securely, you will need to create an SSH key pair and place the private key
on your machine and the public key in the remote machine. You can follow these steps:

1. Create the key pair

```shell
ssh-keygen -t rsa -b 4096
```

2. Copy the public key to the remote machine

```shell
ssh-copy-id pi@RETROPIE_IP_ADDRESS
```

This installs your public key in the remote machine's `~/.ssh/authorized_keys` file.

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

### Remove

Removes specific ROM files:

```bash
retros rm "Game with awesome sprites.gba"
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

## Disclaimer

This project does NOT provide any game ROM files.
