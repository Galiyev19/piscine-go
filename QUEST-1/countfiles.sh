#1/bin/bash

countFiles=$(find . -type f -o -type d | wc -l)

printf $countFiles "\n"