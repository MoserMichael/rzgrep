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

    echo "*** Highlight search results***"

    ./rzgrep -color -C 3 -e 'Cl.se' -in zip.jar
    ./rzgrep -color -C 3 -e Cl.se -in zip.ear
    ./rzgrep -color -C 3 -e Cl.se -in zip.tgz

    echo "*** Java decompiler: search in compiled classes ***"
    ./rzgrep -color -C 3 -e for -in rzgrep.jar -j

    echo "*** Java decompiler: show first twenty imports by popularity ***"
    ./rzgrep  -e ^import -in rzgrep.jar -j  | awk '{ print $3 }' | sort | uniq -c  | sort -k 1 -r | head -20
}

make_zip_in_zip

cp README.template  README.md
./rzgrep --help 2>>README.md

cat >>README.md <<EOF
</pre>
The test output
<pre>
EOF

exec 1>>README.md
exec 2>&1

test_it

echo "*** eof test ***"
