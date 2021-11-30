package simple

import (
	"github.com/pinealctx/neptune/jsonx"
	"github.com/pinealctx/neptune/ulog"
	"go.uber.org/zap"
)

var (
	lessLRUSize = 4 * 1024

	defaultUserLRUSize          = 1024 * 8
	defaultMobileLRUSize        = 1024 * 8
	defaultWechatOpenLRUSize    = 1024 * 16
	defaultWechatMiniAppLRUSize = 1024 * 16

	defaultCnf = &Config{
		UserLRUSize:          defaultUserLRUSize,
		MobileLRUSize:        defaultMobileLRUSize,
		WeChatOpenLRUSize:    defaultWechatOpenLRUSize,
		WechatMiniAppLRUSize: defaultWechatMiniAppLRUSize,
	}
)

type Config struct {
	//UserLRUSize : user lru size in each slot
	UserLRUSize int `json:"user_lru_size" toml:"user_lru_size"`
	//MobileLRUSize : mobile lru size in each slot
	MobileLRUSize int `json:"mobile_lru_size" toml:"mobile_lru_size"`
	//WeChatOpenLRUSize : we chat open lru size in each slot
	WeChatOpenLRUSize int `json:"wechat_open_lru_size" toml:"wechat_open_lru_size"`
	//WechatMiniAppLRUSize : wechat mini app lru size in each slot
	WechatMiniAppLRUSize int `json:"wechat_mini_app_lru_size" toml:"wechat_mini_app_lru_size"`
}

func (c *Config) Normalize() {
	if c.UserLRUSize <= lessLRUSize {
		c.UserLRUSize = defaultUserLRUSize
	}
	if c.MobileLRUSize <= lessLRUSize {
		c.MobileLRUSize = defaultMobileLRUSize
	}
	if c.WeChatOpenLRUSize <= lessLRUSize {
		c.WeChatOpenLRUSize = defaultWechatOpenLRUSize
	}
	if c.WechatMiniAppLRUSize <= lessLRUSize {
		c.WechatMiniAppLRUSize = defaultWechatMiniAppLRUSize
	}
}

func LoadCnf(fName string) *Config {
	if fName == "" {
		return defaultCnf
	}
	var c Config
	var err = jsonx.LoadJSONFile2Obj(fName, &c)
	if err != nil {
		ulog.Error("load.whois.simple.config",
			zap.String("fileName", fName),
			zap.Error(err))
		return defaultCnf
	}
	c.Normalize()
	return &c
}
