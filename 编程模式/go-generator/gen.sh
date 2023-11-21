#!/bin/bash
# 2. 函数生成脚本
# 需要 4 个参数：
# 模板源文件；
# 包名；
# 实际需要具体化的类型；
# 用于构造目标文件名的后缀。
set -e

SRC_FILE=${1}
PACKAGE=${2}
TYPE=${3}
DES=${4}

# uppcase the first char
PREFIX="$(tr '[:lower:]' '[:upper:]' <<< ${TYPE:0:1})${TYPE:1}"
DES_FILE=$(echo ${TYPE}| tr '[:upper:]' '[:lower:]')_${DES}.go
sed 's/PACKAGE_NAME/'"${PACKAGE}"'/g' ${SRC_FILE} | \
    sed 's/GENERIC_TYPE/'"${TYPE}"'/g' | \
    sed 's/GENERIC_NAME/'"${PREFIX}"'/g' > ${DES_FILE}