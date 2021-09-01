package simple

import (
	"github.com/pinealctx/neptune/cache"
	"github.com/pinealctx/neptune/lock"
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
	loadMobDocker       *lock.SimpleLockDocker
	loadUserDocker      *lock.SimpleLockDocker
	loadWeOpenDocker    *lock.SimpleLockDocker
	loadWeMiniAppDocker *lock.SimpleLockDocker
	//update data locker map
	updMobDocker       *lock.SimpleLockDocker
	updUserDocker      *lock.SimpleLockDocker
	updWeOpenDocker    *lock.SimpleLockDocker
	updWeMiniAppDocker *lock.SimpleLockDocker
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

	w.loadMobDocker = lock.NewSimpleLockDocker()
	w.loadUserDocker = lock.NewSimpleLockDocker()
	w.loadWeOpenDocker = lock.NewSimpleLockDocker()
	w.loadWeMiniAppDocker = lock.NewSimpleLockDocker()

	w.updMobDocker = lock.NewSimpleLockDocker()
	w.updUserDocker = lock.NewSimpleLockDocker()
	w.updWeOpenDocker = lock.NewSimpleLockDocker()
	w.updWeMiniAppDocker = lock.NewSimpleLockDocker()
	return w
}
