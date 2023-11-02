# RoSy

RoSy is a CLI tool to help you sync your ROM files with your RetroPie setup
on a remote machine, e.g.: Raspberry Pi.
Instead of having to manually copy files into the correct folder, you can simply call
`rosy cp MyGameROM.a26` and RoSy will do the rest for you.

![example workflow](https://github.com/robwillup/rosy/actions/workflows/build.yml/badge.svg)

## Download

You can download the latest binary from [releases](https://github.com/robwillup/rosy/releases) or use one of the commands below:

### Linux

```bash
# wget
wget https://github.com/robwillup/rosy/releases/download/v1.0.0/rosy && chmod +x rosy
```

```bash
# cURL
curl -L https://github.com/robwillup/rosy/releases/download/v1.0.0/rosy -o rosy && chmod +x rosy
```

### Windows

```powershell
iwr "https://github.com/robwillup/rosy/releases/download/v1.0.0/rosy.exe" -o "rosy.exe"
```

## Setup

RoSy will need the `host (IP address)` and the `username` of the remote machine where RetroPie is running. You can configure that by running:

```bash
./rosy cf
```

RoSy assumes your `SSH private key` is in the default path, i.e.: `$HOME/.ssh/id_rsa`.

## Operations

These are the commands currently being implemented.

### Copy

Adding ROM files to RetroPie:

```bash
./rosy cp Game.md
```

The ROM will be added to the corresponding folder in RetroPie based on
the ROM file extension, in the case above `/genesis`.

## List

Lists the ROM files in RetroPie:

```bash
./rosy ls
```

To list ROM files for specific platforms:

```bash
./rosy ls -p mastersystem
```
