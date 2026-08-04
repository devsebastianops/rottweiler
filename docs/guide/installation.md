# Installation

We offer multiple ways to install Rottweiler, depending on your preferences and environment.

## Installer

```bash
curl -sSL https://raw.githubusercontent.com/devsebastianops/rottweiler/main/install.sh | bash
```

## Via Go

```bash
go install github.com/devsebastianops/rottweiler@latest
```

## From Source

```bash
git clone https://github.com/devsebastianops/rottweiler.git
cd rottweiler
go build -o rottweiler ./cmd/rottweiler/main.go
```

## Docker

```bash
docker run --rm -v $(pwd):/data devsebastianops/rottweiler check --input /data/input.json --policy /data/policy.yaml
```

## Pre-built Binaries

You can download the latest release from the [GitHub releases page](https://github.com/devsebastianops/rottweiler/releases).

## Verify your installation

Simply run the following command to verify that Rottweiler is installed correctly:

```bash
rottweiler --help
```

