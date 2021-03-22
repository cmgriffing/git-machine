#!/usr/bin/env bash}

# shellcheck source=PLATFORMS.sh
source PLATFORMS.sh

for platform in "${SUPPORTED_PLATFORMS[@]}"
do
    IFS='/'
    read -a platform_split <<< "$platform"

    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    echo "BUILDING $GOOS/$GOARCH"
    output_name='gitm-'$GOOS'-'$GOARCH
    if [ "$GOOS" = "windows" ]; then
        output_name+='.exe'
    fi

    env GOOS="$GOOS" GOARCH="$GOARCH" go build -o "release/$output_name" "gitm"
    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
done