#!/bin/bash
cd $1
rm -rf "$3"

echo "开始拉取代码 cd $1<br>"
git clone https://$4:$5%@$2.git $3
#git clone https://$2.git $3
#git clone http://$2:$3@$4.git $5  //有用户名密码使用
echo "拉取代码完成<br>"
git config --global user.email "$2@163.com"
git config --global user.name "$2"