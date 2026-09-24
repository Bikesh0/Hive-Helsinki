#!/bin/bash

count=$(find . | wc -l)
count=$((count * 5))

printf "\t\vTotal files * 5: %d\v\n" "$count"
