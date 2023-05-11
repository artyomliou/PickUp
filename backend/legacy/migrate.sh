#!/bin/sh

cd "$(dirname "$0")"
. ./.env

if ! which migrate >/dev/null; then
    echo "Install golang-migrate binary..."

    version="v4.15.2"

    unameOut="$(uname -s)"
    case "${unameOut}" in
    Linux*) os=linux ;;
    Darwin*) os=darwin ;;
    CYGWIN*) os=windows ;;
    MINGW*) os=windows ;;
    *) echo "UNKNOWN OS:${unameOut}" return 1 ;;
    esac

    unameOut="$(uname -m)"
    case "${unameOut}" in
    x86_64*) arch=amd64 ;;
    amd64*) arch=amd64 ;;
    arm64*) arch=arm64 ;;
    *) echo "UNKNOWN ARCH:${unameOut}" return 1 ;;
    esac

    dir="$(mktemp -d)"
    cd $dir
    curl -s -L https://github.com/golang-migrate/migrate/releases/download/$version/migrate.$os-$arch.tar.gz | tar xvz
    sudo mv ./migrate /usr/bin/migrate
    sudo rm -rf $dir
fi

echo "Run migration..."
migrate -source file://migrations -database mysql://$MYSQL_USER:$MYSQL_PASSWORD@tcp\($MYSQL_HOST:$MYSQL_PORT\)/$MYSQL_DATABASE up
# https://github.com/golang-migrate/migrate/tree/master/database/mysql
