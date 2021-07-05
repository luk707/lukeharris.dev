#!/bin/bash

declare -a Packages=("linkedList" "r30" "testUtils")

for Package in ${Packages[@]}; do
   ../../bin/godocdown go.lukeharris.dev/$GoPackage > web/src/content/packages/$Package.md
done
