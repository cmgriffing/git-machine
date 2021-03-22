#!/bin/bash

# determine platform
platform=""
case $(uname -s) in
    Darwin) platform="darwin";;
    Linux) platform="linux";;
esac
export platform

# determine arch
architecture=""
case $(uname -m) in
    i386)   architecture="386" ;;
    i686)   architecture="386" ;;
    x86_64) architecture="amd64" ;;
    arm)    dpkg --print-architecture | grep -q "arm64" && architecture="arm64" || architecture="arm" ;;
esac
export architecture

export SUPPORTED_PLATFORMS=(
  "darwin/amd64"
  # "darwin/386"
  # "darwin/arm"
  "linux/amd64"
  "linux/386"
  "linux/arm"
  "linux/arm64"
)
