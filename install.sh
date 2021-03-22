#!/usr/bin/env bash

echo "Installing git-machine"

echo "SOME ASCII JAMES BROWN"

# shellcheck source=PLATFORMS.sh
source PLATFORMS.sh

# check if platform supported
if [[ ! "${SUPPORTED_PLATFORMS[@]}" =~ "${platform}/${architecture}" ]]; then
    echo "Sorry, your platform is not yet ready for the glory of git-machine."
    exit 1
fi

filename="gitm-${platform}-${architecture}"

# download executable for platform
curl -o "$filename" -s "https://raw.githubusercontent.com/cmgriffing/git-machine/main/release/${filename}"

if [ ! -f "$filename" ]; then
  printf "\\nExecutable %s not downloaded properly." "$filename"
  exit 1
fi

default_prefix_dir="/usr/local"

# use getopt if the number of params grows
prefix_dir=$default_prefix_dir
prefix_param=${1:-}
prefix_value=${2:-}
if [[ "$prefix_param" = "-p" || "$prefix_param" = "--prefix" ]]; then
  if [[ -z "$prefix_value" ]]; then
    echo "Failed finding bin dir"
    exit 1
  else
    prefix_dir="$prefix_value"
  fi
fi

bin_dir="$prefix_dir/bin"

install -m755 "$filename" "$bin_dir/gitm"
rm "$filename"

# scaffold $HOME/.git-machine
mkdir -p ~/.git-machine

echo "git-machine is installed as gitm. You can configure it by creating a ~/.git-machine/config file and setting options such as Aliases."

# echo instructions for aliases
echo "To use Aliases, run 'gitm config aliases add'. You can update the Aliases prop in the config file to allow for more than the defaults: lets just need want."
