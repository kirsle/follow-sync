#!/bin/bash

# Make a distributable EXE for this program. Usage:
#
# ./dist.sh <target> <arch>
#   target: one of `linux`, `darwin` or `windows`
#   arch: one of `32` or `64`

target="$1"
arch="$2"
binname="follow-sync"
VERSION=`go run main.go -version`

usage() {
	echo "Usage: dist.sh <target> <arch>"
	echo -e "\tTarget must be one of: linux, darwin, windows"
	echo -e "\tArch must be one of: 32, 64"
	exit 1
}

# Validate the arch param.
if [[ $arch == "32" ]]; then
	arch="386"
elif [[ $arch == "64" ]]; then
	arch="amd64"
else
	usage
fi

# Validate the target param.
if [[ $target =~ ^(linux|darwin|windows)$ ]]; then
	[[ $target == "windows" ]] && binname="follow-sync.exe"
	echo "Target platform: $target $arch"
else
	usage
fi

# Prepare the dist root.
root="dist/${target}-${arch}"
mkdir -p "$root"

# Build it.
export GOOS="$target"
export GOARCH="$arch"
go build -ldflags '-s' -o "$root/$binname"

# Copy supplemental files.
cp README.md screenshot.png LICENSE "$root/"

# Zip it up.
cd "$root"
zip "follow-sync-${VERSION}-${target}-${arch}.zip" *
cd -
