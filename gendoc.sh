#!/bin/bash

declare -a GoPackages=("linkedList" "testUtils" )

for GoPackage in ${GoPackages[@]}; do
   ../../bin/godocdown lukeharris.dev/$GoPackage > web/src/content/go-packages/$GoPackage.md
done
