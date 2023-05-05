package omx_core

import (
    "android/soong/android"
    "android/soong/cc"
    "fmt"
)

func init() {
    android.RegisterModuleType("omx_core_defaults", omxCoreDefaultsFactory)
}

func omxCoreDefaultsFactory() android.Module {
    module := cc.DefaultsFactory()
    android.AddLoadHook(module, omxCoreDefaults)

    return module
}

func omxCoreDefaults(ctx android.LoadHookContext) {
    type props struct {
        Srcs       []string
    }

    p := &props{}

    board := ctx.AConfig().VendorConfig("vendor").String("board")
    switch board{
        case "ceres":
            p.Srcs = append(p.Srcs,"tablelist/ceres/aw_registry_table.c")
        case "venus":
            p.Srcs = append(p.Srcs,"tablelist/venus/aw_registry_table.c")
        case "petrel":
            p.Srcs = append(p.Srcs,"tablelist/petrel/aw_registry_table.c")
        case "dolphin":
            p.Srcs = append(p.Srcs,"tablelist/dolphin/aw_registry_table.c")
        case "tulip":
            p.Srcs = append(p.Srcs,"tablelist/tulip/aw_registry_table.c")
        case "cupid":
            p.Srcs = append(p.Srcs,"tablelist/cupid/aw_registry_table.c")
        case "mercury":
            p.Srcs = append(p.Srcs,"tablelist/mercury/aw_registry_table.c")
        default:
            p.Srcs = append(p.Srcs,"tablelist/default/aw_registry_table.c")
            fmt.Println("registry_table has NOT been set. Please check it.")
    }

    ctx.AppendProperties(p)
}
