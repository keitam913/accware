#!/bin/bash

set -eu

build_ui() {
        (
                cd react
                npm install
                npm run build
        )
}

install_resources() {
        sudo install -d -o $(id -u) -g $(id -g) /usr/local/share/accware
        sudo install -d -o $(id -u) -g $(id -g) /var/lib/accware
        ln -sf $PWD/schema.sql /usr/local/share/accware/schema.sql
        ln -sf $PWD/react/build /usr/local/share/accware/assets
}

main() {
        build_ui
        install_resources
}

main
