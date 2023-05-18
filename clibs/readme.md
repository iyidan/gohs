requires:
 - system: centos7.2+
 - tools: 
    - gcc/gcc-c++ 4.8.5
    - GNU Make 3.82
    - cmake-3.26.3-linux-x86_64
 - packages:
    - boost: 1_80_0
    - pcre 8.45

compile:
```
#!/bin/bash
set -uex

TARGET=hs_centos72_install
WORKSPACE=$(cd `dirname $0` && pwd -P)
cd $WORKSPACE

yum update make gcc gcc-c++
chmod +x ./cmake-3.26.3-linux-x86_64.sh
./cmake-3.26.3-linux-x86_64.sh --skip-license --prefix=./ --include-subdir

tar zxvf boost_1_80_0.tar.gz
tar zxvf pcre-8.45.tar.gz

rm -rf ./hyperscan*
wget https://github.com/intel/hyperscan/archive/refs/tags/v5.4.2.tar.gz -O hyperscan-5.4.2.tar.gz
tar zxvf hyperscan-5.4.2.tar.gz
mv hyperscan-5.4.2 hyperscan
cd ./hyperscan/

ln -s $WORKSPACE/boost_1_80_0/boost include/boost

$WORKSPACE/cmake-3.26.3-linux-x86_64/bin/cmake -DCMAKE_BUILD_TYPE=Release -DBUILD_CHIMERA=on -DCMAKE_INSTALL_PREFIX=$WORKSPACE/hyperscan/$TARGET -DPCRE_SOURCE=$WORKSPACE/pcre-8.45/
$WORKSPACE/cmake-3.26.3-linux-x86_64/bin/cmake --build .
make install

cp ./lib/libpcre.a ./$TARGET/lib64/
tar zcvf "${TARGET}.tgz" $TARGET
echo "build end"
```
