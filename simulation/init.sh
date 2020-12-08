#!/bin/bash

root="`pwd`"
begin="$root/begin"
examples="$root/examples"

echo "开始创建模拟路径以及文件 $begin"

#复制examples下的所有目录到 $fullPath
rm -rf "$begin/examples"
cp -rf "$examples" "$begin/"
