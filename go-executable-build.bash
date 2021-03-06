#!/usr/bin/env bash

platforms=("windows/amd64" "windows/386" "linux/amd64")

package=$1
if [[ -z "$package" ]]; then
  echo "usage: $0 <package-name>"
  echo "example: env GOPATH=\$PWD $0 github.com/educhastenier/defer"
  exit 1
fi
package_split=(${package//\// })
package_name=${package_split[-1]}

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name=bin/$package_name'-'$GOOS'-'$GOARCH
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi

    env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name $package
    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
done
