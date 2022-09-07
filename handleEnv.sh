#! /bin/bash

echo ${{secrets.TEST}}
if [  -z ${{secrets.TEST}} ]; then
echo hello
export TEST=${{secrets.TEST}}
fi