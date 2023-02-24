package cedarc_config

import (
    "android/soong/android"
    "android/soong/cc"
    "fmt"
)
// Add other board config related with Android sdkVersion alike below.
var a50_Q_cflags = []string {
    "-DCONF_KERNEL_VERSION_4_9",
    "-DCONF_MALI_GPU",
    "-DCONF_USE_IOMMU",
    "-DCONF_KERN_BITWIDE=32",
    "-DCONFIG_VE_IPC_ENABLE",
    "-DGPU_ALIGN_STRIDE=32",
}

var a100_Q_cflags = []string {
    "-DCONF_KERNEL_VERSION_4_9",
    "-DCONF_IMG_GPU_USE_COMMON_STRUCT",//After ceres, MALI GPU and IMG GPU use the same structure for users
    "-DCONF_USE_IOMMU",
    "-DCONF_KERN_BITWIDE=64",
    "-DCONFIG_VE_IPC_ENABLE",
    "-DGPU_ALIGN_STRIDE=32",
    "-DCONF_VE_FREQ_ENABLE_SETUP",
    "-DCONF_CERES_VE_FREQ_ENABLE_SETUP",
}
var t3_Q_cflags = []string {
    "-DCONF_KERNEL_VERSION_4_9",
    "-DCONF_MALI_GPU",
    "-DCONF_KERN_BITWIDE=32",
    "-DCONFIG_VE_IPC_ENABLE",
    "-DGPU_ALIGN_STRIDE=32",
}

var dolphin_Q_cflags = []string {
    "-DCONF_KERNEL_VERSION_4_9",
    "-DCONF_MALI_GPU",
    "-DCONF_ENABLE_OPENMAX_DI_FUNCTION",
    "-DCONF_KERN_BITWIDE=32",
    "-DCONFIG_VE_IPC_ENABLE",
    "-DOMX_DROP_B_FRAME_4K",
    "-DGPU_ALIGN_STRIDE=32",
}
//CONF_VP9_DEC: 0-libawvp9soft; 1-libawvp9Hw; 2-libawvp9HwAL
var cupid_Q_cflags = []string {
    "-DCONF_KERNEL_VERSION_4_9",
    "-DCONF_MALI_GPU",
    "-DCONF_USE_IOMMU",
    "-DCONF_ENABLE_OPENMAX_DI_FUNCTION",
    "-DCONF_KERN_BITWIDE=64",
    "-DCONF_AFBC_ENABLE",
    "-DCONF_HIGH_DYNAMIC_RANGE_ENABLE",
    "-DCONFIG_VE_IPC_ENABLE",
    "-DCONF_VP9_DEC=2",
    "-DGPU_ALIGN_STRIDE=32",
    "-DCONF_DI_300_SUPPORT",
}

var mercury_Q_cflags = []string {
    "-DCONF_KERNEL_VERSION_4_9",
    "-DCONF_MALI_GPU",
    "-DCONF_USE_IOMMU",
    "-DCONF_ENABLE_OPENMAX_DI_FUNCTION",
    "-DCONF_KERN_BITWIDE=64",
    "-DCONF_AFBC_ENABLE",
    "-DCONF_HIGH_DYNAMIC_RANGE_ENABLE",
    "-DCONFIG_VE_IPC_ENABLE",
    "-DCONF_VP9_DEC=2",
    "-DGPU_ALIGN_STRIDE=32",
    "-DCONF_DI_300_SUPPORT",
}

var default_cflags = []string {
    "-DCONF_KERNEL_VERSION_4_9",
    "-DCONF_MALI_GPU",
    "-DCONF_USE_IOMMU",
    "-DCONF_KERN_BITWIDE=32",
    "-DCONFIG_VE_IPC_ENABLE",
    "-DGPU_ALIGN_STRIDE=32",
    //"-DCONF_ENABLE_OPENMAX_DI_FUNCTION", //if open di support
    //"-DCONF_DI_V2X_SUPPORT", //if di processes 3 input pictures
    //"-DCONF_DI_300_SUPPORT", //if di300 supported
}
func globalDefaults(ctx android.BaseContext) ([]string) {
    var cppflags []string

    sdkVersion := ctx.AConfig().PlatformSdkVersionInt()
    fmt.Println("sdkVersion:", sdkVersion)

    if sdkVersion >= 29 { //after Q
        cppflags = append(cppflags,"-DCONF_Q_AND_NEWER")
        board := ctx.AConfig().VendorConfig("vendor").String("board")
        fmt.Println("board:",board)
        cppflags = append(cppflags,"-DTARGET_BOARD_PLATFORM="+board)
        switch board{
            case "venus":
                cppflags = append(cppflags,a50_Q_cflags...)
            case "ceres":
                cppflags = append(cppflags,a100_Q_cflags...)
            case "t3":
                cppflags = append(cppflags,t3_Q_cflags...)
            case "dolphin":
                cppflags = append(cppflags,dolphin_Q_cflags...)
            case "cupid":
                cppflags = append(cppflags,cupid_Q_cflags...)
            case "mercury":
                cppflags = append(cppflags,mercury_Q_cflags...)
            default:
                cppflags = append(cppflags,default_cflags...)
                fmt.Println("Platform config has NOT been set. Please check it.")
        }
    }
    if sdkVersion >= 28 { //after P
        cppflags = append(cppflags,"-DCONF_PIE_AND_NEWER")
    }
    if sdkVersion >= 26 { //after O
        cppflags = append(cppflags,"-DCONF_OREO_AND_NEWER")
    }
    if sdkVersion >= 24 { //after N
        cppflags = append(cppflags,"-DCONF_NOUGAT_AND_NEWER")
    }
    if sdkVersion >= 23 { //after M
        cppflags = append(cppflags,"-DCONF_MARSHMALLOW_AND_NEWER")
    }
    if sdkVersion >= 21 { //after L
        cppflags = append(cppflags,"-DCONF_LOLLIPOP_AND_NEWER")
    }
    if sdkVersion >= 19 { //after KK
        cppflags = append(cppflags,"-DCONF_KITKAT_AND_NEWER")
    }
    if sdkVersion >= 17 { //after JB42
        cppflags = append(cppflags,"-DCONF_JB42_AND_NEWER")
    }

    cryptolevel := ctx.AConfig().VendorConfig("widevine").String("cryptolevel")
    playreadytype := ctx.AConfig().VendorConfig("playready").String("playreadytype")
    fmt.Println("cryptolevel:",cryptolevel)
    fmt.Println("playreadytype:",playreadytype)
    if cryptolevel == "1" || playreadytype == "2"{
         cppflags = append(cppflags,"-DPLATFORM_SURPPORT_SECURE_OS=1",
                                    "-DSECURE_OS_OPTEE=1",
                                    "-DADJUST_ADDRESS_FOR_SECURE_OS_OPTEE=1")
    }

    cppflags = append(cppflags,"-DCONF_ARMV7_A_NEON")

    config  := ctx.Config().VendorConfig("gpu").String("public_include_file")
    cppflags = append(cppflags, "-DGPU_PUBLIC_INCLUDE=\"" + config + "\"")

    return cppflags
}

func configDefaults(ctx android.LoadHookContext) {
    type props struct {
        Cflags []string
    }
    p := &props{}
    p.Cflags = globalDefaults(ctx)
    ctx.AppendProperties(p)
}

func configDefaultsFactory() android.Module {
    module := cc.DefaultsFactory()
    android.AddLoadHook(module, configDefaults)
    return module
}

func init() {
    fmt.Println("cedarc config go file start")
    android.RegisterModuleType("config_defaults", configDefaultsFactory)
}
