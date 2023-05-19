# Allwinner A133 libcedarc(VE) for GNU Linux

## 编译

```bash
./bootstrap
CC=aarch64-linux-gnu-gcc CXX=aarch64-linux-gnu-g++ ./configure --prefix=/home/calvin/helper/libcedarc/arm64 --host=aarch64-linux-gnu CFLAGS="-Wno-error=switch-unreachable -DCONF_KERNEL_VERSION_4_9 -DCONF_IMG_GPU_USE_COMMON_STRUCT -DCONF_USE_IOMMU -DCONF_KERN_BITWIDE=64 -DCONFIG_VE_IPC_ENABLE -DGPU_ALIGN_STRIDE=32 -DCONF_VE_FREQ_ENABLE_SETUP -DCONF_CERES_VE_FREQ_ENABLE_SETUP -DCONF_PIE_AND_NEWER -DCONF_ARMV7_A_NEON" LDFLAGS="-L/home/calvin/helper/libcedarc/library/aarch64-linux-gnu"
make
make install
```

-D宏定义参考config/cedarc_config.go编译脚本中的a100a100_Q_cflags，这是a133的编译宏定义，不添加这些宏编译出的so库无法操作VE

## 魔改支持GNU Linux GLIBC

**library增加了aarch64-linux-gnu动态链接库，该库由library/androidq_64下的动态链接库魔改elf文件后得到，可在GNU Linux下使用glibc编译连接使用。**

1. 经测试发现，将安卓的全部so进行魔改时解码会有问题，因此最新commit中**仅**魔改了**libVE.so**，其余so文件由toolchain-sunxi-aarch64-glibc中so文件拷贝而来
2. libawavs2.so与libawvp9HwAL.so只有androidq_64有而toolchain-sunxi-aarch64-glibc中没有，因此这两个so也由安卓so库魔改得到，但**未进行测试**
3. toolchain-sunxi-aarch64-glibc不能使用的根本原因是因为其中的libVE.so未适配a133的VE，全志仅放出了安卓平台的含a133 VE的libVE.so。这点应该是全志故意为之，因为Linux内核包含了完整的VE驱动代码，并且安卓端的libVE.so也仅仅用到了libc而未使用安卓其他so库的任何功能，本质上编译出一份Linux glibc版本的so库是没有任何额外工作的
4. 使用魔改的libVE.so + toolchain-sunxi-aarch64-glibc下的其余so库，经测试**jpeg编码、H264编码、mjpeg解码**正常，其余功能有待进一步测试

魔改大致步骤如下：

1. 去除bionic libc依赖，替换成glibc，去除不必要的安卓so库依赖

2. glibc不存在的bionic函数使用libcwrapper实现

3. liblog库使用liblogwrapper替换

4. 修改安卓使用的.rela.android重定位表，解码成Linux下的.rela重定位表

5. 去除.fini_array避免程序结束后glibc执行清理工作时产生段错误
