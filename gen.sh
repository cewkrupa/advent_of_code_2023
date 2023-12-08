#!/bin/bash
set -e

mkdir "$1"
cd "$1"
touch input.txt
touch test_input.txt
touch prompt.txt
cp ../main_template.txt .
mv ./main_template.txt ./main.go
echo
git add .
