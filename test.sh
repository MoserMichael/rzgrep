#!/usr/bin/env bash

set -ex

make_zip_in_zip() {
  zip zip.jar src/rzgrep.go go.mod
  zip zip.ear zip.jar
  tar cvfz zip.tgz zip.jar
}

test_it() {
    ./rzgrep -e 'Cl.se' -in zip.jar
    ./rzgrep -e Cl.se -in zip.ear
    ./rzgrep -e Cl.se -in zip.tgz

    ./rzgrep -C 3 -e 'Cl.se' -in zip.jar
    ./rzgrep -C 3 -e Cl.se -in zip.ear
    ./rzgrep -C 3 -e Cl.se -in zip.tgz
}

make_zip_in_zip
test_it
echo "*** eof test ***"
