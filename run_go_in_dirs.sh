#!/bin/bash

# Array of target directories
directories=("ex00" "ex01" "ex02" "ex03" "ex04" "ex05" "ex06" "ex07")

# Iterate over each directory and run `go run .`
for dir in "${directories[@]}"; do
    # Check if the directory exists
    if [ -d "$dir" ]; then
        cd "$dir" || exit
        # Run `go run .`
        echo "Running go run . in $dir"
        go run .
        cd ..
    else
        echo "Directory $dir does not exist."
    fi
done
