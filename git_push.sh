#!/bin/bash
cd $1
pwd
echo "开始push $1<br>"

git add -A
git commit  -m "$2"
echo $3
echo  $4

git push https://$4

git tag -a $3 -m "$2"
git push  https://$4 --tags

echo "push 完成"

