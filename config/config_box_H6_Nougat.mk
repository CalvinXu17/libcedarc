############################################
## configurations: BOX-H6-7.1
############################################
include $(TOP)/frameworks/av/media/libcedarc/config/config_common.mk

## 1. linux kernel version
#LOCAL_CFLAGS += -DCONF_KERNEL_VERSION_3_4
LOCAL_CFLAGS += -DCONF_KERNEL_VERSION_3_10

## 2. gpu
LOCAL_CFLAGS += -DCONF_MALI_GPU
#LOCAL_CFLAGS += -DCONF_IMG_GPU

## 3. new display framework
NEW_DISPLAY = yes

## 4. suporrt 4K 30fps recoder
LOCAL_CFLAGS += -DCONF_SUPPORT_4K_30FPS_RECORDER

## 5. audio decoder soundtouch
#LOCAL_CFLAGS += -CONFIG_ADECODER_SUPPORT_SOUNDTOUCH

## 6. ve ipc
#LOCAL_CFLAGS += -DCONFIG_VE_IPC_ENABLE
## 7. iommu
LOCAL_CFLAGS += -DCONF_USE_IOMMU

## 8. for enable afbc
LOCAL_CFLAGS += -DCONF_AFBC_ENABLE

## 9. for the function of VE frequency setup
LOCAL_CFLAGS += -DCONF_VE_FREQ_ENABLE_SETUP

## 10. now, the function of openmax-di just work well on h6 platform.
LOCAL_CFLAGS += -DCONF_ENABLE_OPENMAX_DI_FUNCTION
## 11. config kernel bitwide value explicitly as 64 FOR DI process.
LOCAL_CFLAGS += -DCONF_KERN_BITWIDE=64
## 12. DI process 3 input pictures. Otherwise, 2 input pictures would be processed anyway.
LOCAL_CFLAGS += -DCONF_DI_PROCESS_3_PICTURE

##13. for 10bit and hdr
LOCAL_CFLAGS += -DCONF_HIGH_DYNAMIC_RANGE_ENABLE
