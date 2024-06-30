#!/bin/bash

# Array of target directories
directories=("ex00" "ex01" "ex02" "ex03" "ex04" "ex05" "ex06" "ex07")

# Iterate over each directory and initialize a Go module
for dir in "${directories[@]}"; do
    # Check if the directory exists
    if [ -d "$dir" ]; then
        cd "$dir" || exit
        # Initialize the Go module
        go mod init "$dir"
        cd ..
    else
        echo "Directory $dir does not exist."
    fi
done
