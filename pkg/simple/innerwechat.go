package simple

import (
	"github.com/pinealctx/neptune/cache"
	"github.com/pinealctx/neptune/syncx/keylock"
	"github.com/pinealctx/whois/api/model"
	"time"
)

//setup we chat user id
func (w *WhoIsSimple) setWeUID(wk *model.WeChatKey, uid int32, now time.Time, isOpen bool) (*model.UWechat, error) {
	var (
		updLock *keylock.KeyLocker
		key     = wk.Key()
	)

	if isOpen {
		updLock = w.updWeOpenDocker
	} else {
		updLock = w.updWeMiniAppDocker
	}

	updLock.Lock(key)
	defer updLock.Unlock(key)

	var pre, err = w.getWeInfo(wk, isOpen)
	if err != nil {
		return nil, err
	}
	return w.updateWeUID(key, wk, pre, uid, now, isOpen)
}

//upsert we chat info
func (w *WhoIsSimple) upsertWeInfo(wf *model.WechatInfo, now time.Time, isOpen bool) (*model.UWechat, error) {

	var (
		updLock *keylock.KeyLocker
		key     = wf.Key()
	)

	if isOpen {
		updLock = w.updWeOpenDocker
	} else {
		updLock = w.updWeMiniAppDocker
	}

	updLock.Lock(key)
	defer updLock.Unlock(key)

	var wk = wf.WeKey()
	var pre, err = w.getWeInfo(wk, isOpen)
	if err != nil {
		if !w.weStore.IsNotFoundErr(err) {
			//other error
			return nil, err
		}
		//insert
		return w.addWeInfo(key, wf, now, isOpen)
	} else {
		//pre exist, update
		return w.updateWeInfo(key, wf, pre, now, isOpen)
	}
}

//get wechat info
func (w *WhoIsSimple) getWeInfo(wk *model.WeChatKey, isOpen bool) (*model.UWechat, error) {
	var (
		r        interface{}
		ok       bool
		loadLock *keylock.KeyLocker
		ca       *cache.LRUCache
		fn       func(weKey *model.WeChatKey) (*model.UWechat, error)
	)

	if isOpen {
		loadLock = w.loadWeOpenDocker
		ca = w.weOpenCa
		fn = w.weStore.LoadOpenInfo
	} else {
		loadLock = w.loadWeMiniAppDocker
		ca = w.weMiniAppCa
		fn = w.weStore.LoadMiniAppInfo
	}

	var key = wk.Key()
	r, ok = ca.Get(key)
	if ok {
		return r.(*model.UWechat), nil
	}

	loadLock.Lock(key)
	defer loadLock.Unlock(key)

	//retry -- cause other go routine can figure out
	r, ok = ca.Get(key)
	if ok {
		return r.(*model.UWechat), nil
	}

	var mo, err = fn(wk)
	if err != nil {
		return nil, err
	}
	ca.Set(key, mo)
	return mo, nil
}

//add we chat info
func (w *WhoIsSimple) addWeInfo(key string, wf *model.WechatInfo, now time.Time, isOpen bool) (*model.UWechat, error) {
	var (
		loadLock *keylock.KeyLocker
		addFn    func(*model.WechatInfo, time.Time) error
		ca       *cache.LRUCache
	)
	if isOpen {
		loadLock = w.loadWeOpenDocker
		addFn = w.weStore.AddOpenInfo
		ca = w.weOpenCa
	} else {
		loadLock = w.loadWeMiniAppDocker
		addFn = w.weStore.AddMiniAppInfo
		ca = w.weMiniAppCa
	}

	loadLock.Lock(key)
	defer loadLock.Unlock(key)

	var err = addFn(wf, now)
	if err != nil {
		return nil, err
	}
	var wc = &model.UWechat{}
	wc.UpdFrom(wf, now)
	wc.CreatedAt = now
	ca.Set(key, wc)
	return wc, nil
}

//update we chat info
func (w *WhoIsSimple) updateWeInfo(key string, wf *model.WechatInfo, pre *model.UWechat,
	now time.Time, isOpen bool) (*model.UWechat, error) {
	if pre.HasNoChange(wf) {
		return pre, nil
	}

	var (
		updFn func(*model.WechatInfo, time.Time) error
		ca    *cache.LRUCache
	)

	if isOpen {
		updFn = w.weStore.UpdateOpenInfo
		ca = w.weOpenCa
	} else {
		updFn = w.weStore.UpdateMiniAppInfo
		ca = w.weMiniAppCa
	}
	var err = updFn(wf, now)
	if err != nil {
		return nil, err
	}
	var nw = &model.UWechat{}
	*nw = *pre
	nw.UpdFrom(wf, now)
	ca.Set(key, nw)
	return nw, nil
}

//update we chat user id
func (w *WhoIsSimple) updateWeUID(key string, wk *model.WeChatKey, pre *model.UWechat,
	uid int32, now time.Time, isOpen bool) (*model.UWechat, error) {

	if pre.UserID == uid {
		return pre, nil
	}

	var (
		uidFn func(*model.WeChatKey, int32, time.Time) error
		ca    *cache.LRUCache
	)

	if isOpen {
		uidFn = w.weStore.BindOpenUserID
		ca = w.weOpenCa
	} else {
		uidFn = w.weStore.BindMiniAppUserID
		ca = w.weMiniAppCa
	}
	var err = uidFn(wk, uid, now)
	if err != nil {
		return nil, err
	}
	var nw = &model.UWechat{}
	*nw = *pre
	nw.UserID = uid
	ca.Set(key, nw)
	return nw, nil
}
