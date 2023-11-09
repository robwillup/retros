# RoSy

RoSy is a CLI tool to help you sync your ROM files with your RetroPie setup
on a remote machine, e.g.: Raspberry Pi.
Instead of having to manually copy files into the correct folder, you can simply
call `rosy cp MyGameROM.a26` and RoSy will do the rest for you.

![build workflow](https://github.com/robwillup/rosy/actions/workflows/build.yml/badge.svg)
![Latest Release](https://img.shields.io/github/v/release/robwillup/rosy?label=Rosy%20(Linux%20binary)&sort=semver)
![Latest Release](https://img.shields.io/github/v/release/robwillup/rosy?label=Rosy%20(Windows%20binary)&sort=semver)

<div align="center">
    <img
        src="https://repository-images.githubusercontent.com/709978523/5298bcc2-e191-46a5-a5f7-1bda2c7f8075"
        alt="RoSy" style="width: 600px;"/>
        <p>Image by @robwillup<p>
</div>

> **Status**
>
> Early development. There are many features to be added and there
may be bugs in current versions.

## Download

You can download the latest binary from [releases](https://github.com/robwillup/rosy/releases)
or use one of the commands below:

### Linux

```bash
# wget
wget https://github.com/robwillup/rosy/releases/download/v1.0.0/rosy \
&& chmod +x rosy
```

```bash
# cURL
curl -L https://github.com/robwillup/rosy/releases/download/v1.0.0/rosy -o rosy \
&& chmod +x rosy
```

### Windows

```powershell
iwr "https://github.com/robwillup/rosy/releases/download/v1.0.0/rosy.exe" -o "rosy.exe"
```

## Setup

RoSy will need the `host (IP address)` and the `username` of the remote machine
where RetroPie is running. You can configure that by running:

```bash
rosy cf
```

RoSy assumes your `SSH private key` is in the default path, i.e.: `$HOME/.ssh/id_rsa`.

## Operations

These are the commands currently being implemented.

### Copy

Adding a single ROM file to RetroPie:

```bash
rosy cp Game.md
```

The ROM will be added to the corresponding folder in RetroPie based on
the ROM file extension, in the case above `/genesis`.

Adding all ROM files in a directory:

```bash
rosy cp /home/gamer/roms
```

All ROM files in the provided directory will be copied to their respective
platforms in RetroPie based on the file extensions.

#### Specifying the platform

If a ROM file has a different extension, you can copy it by
specifying the platform it should go into:

```bash
rosy cp --platform=atari2600 Game.bin
```

## List

Lists the ROM files in RetroPie:

```bash
rosy ls
```

To list ROM files for specific platforms:

```bash
rosy ls -p=mastersystem
```
