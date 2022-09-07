#! /bin/bash

if [ ${{secrets.TEST}} ]; then
echo hello
export TEST=${{secrets.TEST}}
fi