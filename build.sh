#!/usr/bin/env bash

# Directly stolen from https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04

package=$1
if [[ -z "$package" ]]; then
  echo "usage: $0 <package-name>"
  exit 1
fi
package_split=(${package//\// })
package_name=${package_split[-1]}

package_shortname=$(echo $package_name | sed --expression='s/.go//g')

platforms=("linux/amd64" "linux/386" "linux/arm" "linux/arm64" "netbsd/amd64" "netbsd/386" "netbsd/arm" "openbsd/amd64" "openbsd/386" "openbsd/arm" "windows/amd64" "windows/386" "darwin/amd64" "darwin/386" "darwin/arm" "darwin/arm64")

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name='dist/'$GOOS'/'$GOARCH'/'$package_shortname
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi  

    env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name $package
    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
done
