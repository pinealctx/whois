package simple

import (
	"github.com/pinealctx/neptune/cache"
	"github.com/pinealctx/neptune/syncx/keylock"
	"github.com/pinealctx/whois/pkg/store"
	"github.com/pinealctx/whois/pkg/uid"
)

type WhoIsSimple struct {
	c *Config
	//id gen
	uidGen uid.UserIDGen

	//cache
	userCa      *cache.LRUCache
	mobCa       *cache.LRUCache
	weOpenCa    *cache.LRUCache
	weMiniAppCa *cache.LRUCache

	//store
	weStore  store.WeStore
	uStore   store.UserStore
	mobStore store.MobileStore

	//load data locker map
	loadMobDocker       *keylock.KeyLocker
	loadUserDocker      *keylock.KeyLocker
	loadWeOpenDocker    *keylock.KeyLocker
	loadWeMiniAppDocker *keylock.KeyLocker
	//update data locker map
	updMobDocker       *keylock.KeyLocker
	updUserDocker      *keylock.KeyLocker
	updWeOpenDocker    *keylock.KeyLocker
	updWeMiniAppDocker *keylock.KeyLocker
}

func NewWhoIsSimple(c *Config,
	ug uid.UserIDGen, ws store.WeStore, us store.UserStore, mbs store.MobileStore) *WhoIsSimple {
	var w = &WhoIsSimple{}
	w.c = c
	w.userCa = cache.NewLRUCache(int64(w.c.UserLRUSize))
	w.mobCa = cache.NewLRUCache(int64(w.c.MobileLRUSize))
	w.weOpenCa = cache.NewLRUCache(int64(w.c.WeChatOpenLRUSize))
	w.weMiniAppCa = cache.NewLRUCache(int64(w.c.WechatMiniAppLRUSize))

	w.uidGen = ug
	w.weStore, w.uStore, w.mobStore = ws, us, mbs

	w.loadMobDocker = keylock.NewKeyLocker()
	w.loadUserDocker = keylock.NewKeyLocker()
	w.loadWeOpenDocker = keylock.NewKeyLocker()
	w.loadWeMiniAppDocker = keylock.NewKeyLocker()

	w.updMobDocker = keylock.NewKeyLocker()
	w.updUserDocker = keylock.NewKeyLocker()
	w.updWeOpenDocker = keylock.NewKeyLocker()
	w.updWeMiniAppDocker = keylock.NewKeyLocker()
	return w
}
