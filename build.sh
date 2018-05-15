#/usr/bin/env bash
CONFIGURATIONS=(
    darwin,386 \
    windows,386 \
    windows,amd64 \
    linux,386 \
    linux,amd64 \
    linux,arm \
    linux,arm64 \
    solaris,amd64
)

for config in ${CONFIGURATIONS[@]}; do 
IFS=","
set $config

os="$1"
arch="$2"
path="build/$os/$arch"
mkdir -p "build/$os/$arch"
bin_name="goiban-service"

if [ $os = "windows" ]; then
    bin_name="$bin_name.exe"
fi

GOOS="$os" GOARCH="$arch" go build -o "$path/$bin_name"
cp -r $GOPATH/src/github.com/fourcube/goiban-data-loader/data "$path/"
cp -r ./static "$path/"
tar -czf "build/goiban-service-$os-$arch.tar.gz" "$path"

unset IFS;
done