/*
* Copyright (c) 2008-2016 Allwinner Technology Co. Ltd.
* All rights reserved.
*
* File : log.h
* Description :
* History :
*   Author  : xyliu <xyliu@allwinnertech.com>
*   Date    : 2015/04/13
*   Comment :
*
*
*/

#ifndef LOG_H
#define LOG_H

#include <cdc_config.h>

#ifndef LOG_TAG
#define LOG_TAG "cedarc"
#endif

enum CDX_LOG_LEVEL_TYPE {
    LOG_LEVEL_VERBOSE = 2,
    LOG_LEVEL_DEBUG = 3,
    LOG_LEVEL_INFO = 4,
    LOG_LEVEL_WARNING = 5,
    LOG_LEVEL_ERROR = 6,
};

extern enum CDX_LOG_LEVEL_TYPE GLOBAL_LOG_LEVEL;

#ifdef __ANDROID__
    #ifdef CONF_OREO_AND_NEWER
    #include <log/log.h>
    #else
    #include <cutils/log.h>
    #endif

    #define LOG_LEVEL_ERROR     ANDROID_LOG_ERROR
    #define LOG_LEVEL_WARNING   ANDROID_LOG_WARN
    #define LOG_LEVEL_INFO      ANDROID_LOG_INFO
    #define LOG_LEVEL_VERBOSE   ANDROID_LOG_VERBOSE
    #define LOG_LEVEL_DEBUG     ANDROID_LOG_DEBUG

    #define AWLOG(level, fmt, arg...)  \
        LOG_PRI(level, LOG_TAG, "<%s:%u>: " fmt, __FUNCTION__, __LINE__, ##arg)

    #define CDCLOG(level, fmt, arg...)  \
        do { \
            if (level >= (android_LogPriority)GLOBAL_LOG_LEVEL) \
                LOG_PRI((android_LogPriority)level, LOG_TAG, "<%s:%u>: " fmt, __FUNCTION__, __LINE__, ##arg); \
        } while (0)



#define CC_LOG_ASSERT(e, fmt, arg...)                               \
        LOG_ALWAYS_FATAL_IF(                                        \
                !(e),                                               \
                "<%s:%d>check (%s) failed:" fmt,                    \
                __FUNCTION__, __LINE__, #e, ##arg)                  \

#else

#include <stdio.h>
#include <string.h>

/*
#define LOG_LEVEL_ERROR     "error  "
#define LOG_LEVEL_WARNING   "warning"
#define LOG_LEVEL_INFO      "info   "
#define LOG_LEVEL_VERBOSE   "verbose"
#define LOG_LEVEL_DEBUG     "debug  "
*/

extern const char *CDX_LOG_LEVEL_NAME[];
#define CDCLOG(level, fmt, arg...)  \
    do { \
        if ((level >= GLOBAL_LOG_LEVEL || level >= CONFIG_LOG_LEVEL) && \
                level <= LOG_LEVEL_ERROR) \
            printf("%s: %s <%s:%u>: " fmt "\n", \
                    CDX_LOG_LEVEL_NAME[level], LOG_TAG, __FUNCTION__, __LINE__, ##arg); \
    } while (0)


#define AWLOG(level, fmt, arg...)  \
    printf("%d: %s <%s:%u>: "fmt"\n", level, LOG_TAG, __FUNCTION__, __LINE__, ##arg)

#define CC_LOG_ASSERT(e, fmt, arg...)                                       \
                do {                                                        \
                    if (!(e))                                               \
                    {                                                       \
                        loge("check (%s) failed:"fmt, #e, ##arg);           \
                        assert(0);                                          \
                    }                                                       \
                } while (0)

#endif

#define logd(fmt, arg...) AWLOG(LOG_LEVEL_DEBUG, fmt, ##arg)

#if CONFIG_LOG_LEVEL == OPTION_LOG_LEVEL_CLOSE

#define loge(fmt, arg...)
#define logw(fmt, arg...)
#define logi(fmt, arg...)
#define logv(fmt, arg...)

#elif CONFIG_LOG_LEVEL == OPTION_LOG_LEVEL_ERROR

#define loge(fmt, arg...) AWLOG(LOG_LEVEL_ERROR, "\033[40;31m" fmt "\033[0m", ##arg)
#define logw(fmt, arg...)
#define logi(fmt, arg...)
#define logv(fmt, arg...)

#elif CONFIG_LOG_LEVEL == OPTION_LOG_LEVEL_WARNING

#define loge(fmt, arg...) AWLOG(LOG_LEVEL_ERROR, "\033[40;31m" fmt "\033[0m", ##arg)
#define logw(fmt, arg...) AWLOG(LOG_LEVEL_WARNING, fmt, ##arg)
#define logi(fmt, arg...)
#define logv(fmt, arg...)

#elif CONFIG_LOG_LEVEL == OPTION_LOG_LEVEL_DEFAULT

#define loge(fmt, arg...) AWLOG(LOG_LEVEL_ERROR, "\033[40;31m" fmt "\033[0m", ##arg)
#define logw(fmt, arg...) AWLOG(LOG_LEVEL_WARNING, fmt, ##arg)
#define logi(fmt, arg...) AWLOG(LOG_LEVEL_INFO, fmt, ##arg)
#define logv(fmt, arg...)

#elif CONFIG_LOG_LEVEL == OPTION_LOG_LEVEL_DETAIL

#define loge(fmt, arg...) AWLOG(LOG_LEVEL_ERROR, "\033[40;31m" fmt "\033[0m", ##arg)
#define logw(fmt, arg...) AWLOG(LOG_LEVEL_WARNING, fmt, ##arg)
#define logi(fmt, arg...) AWLOG(LOG_LEVEL_INFO, fmt, ##arg)
#define logv(fmt, arg...) AWLOG(LOG_LEVEL_VERBOSE, fmt, ##arg)

#else
    #error "invalid configuration of debug level."
#endif

#define omx_logd(fmt, arg...) CDCLOG(LOG_LEVEL_DEBUG , fmt, ##arg)
#define omx_loge(fmt, arg...) CDCLOG(LOG_LEVEL_ERROR, "\033[40;31m" fmt "\033[0m", ##arg)
#define omx_logw(fmt, arg...) CDCLOG(LOG_LEVEL_WARNING, fmt, ##arg)
#define omx_logi(fmt, arg...) CDCLOG(LOG_LEVEL_INFO, fmt, ##arg)
#define omx_logv(fmt, arg...) CDCLOG(LOG_LEVEL_VERBOSE, fmt, ##arg)


//If you want to use omx_log you should call CdcIniParserInit when the module inits.
void log_set_level(unsigned level);


#define CEDARC_UNUSE(param) (void)param  //just for remove compile warning

#define CEDARC_DEBUG (0)

#endif

