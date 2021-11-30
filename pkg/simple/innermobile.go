package simple

import (
	"github.com/pinealctx/whois/api/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

var (
	ErrMobileExist = status.Errorf(codes.AlreadyExists, "mobile.already.exist")
)

//get mobile info
func (w *WhoIsSimple) getMobileInfo(m *model.MobileInfo) (*model.UMobile, error) {
	var (
		r  interface{}
		ok bool
	)
	var key = m.Key()
	r, ok = w.mobCa.Get(key)
	if ok {
		return r.(*model.UMobile), nil
	}

	var loadLock = w.loadMobDocker
	loadLock.Lock(key)
	defer loadLock.Unlock(key)

	//retry -- cause other go routine can figure out
	r, ok = w.mobCa.Get(key)
	if ok {
		return r.(*model.UMobile), nil
	}

	var mo, err = w.mobStore.LoadMobile(m)
	if err != nil {
		return nil, err
	}
	w.mobCa.Set(key, mo)
	return mo, nil
}

//add mobile info
func (w *WhoIsSimple) addMobileInfo(m *model.MobileInfo, uid int32, now time.Time) (*model.UMobile, error) {
	var key = m.Key()
	var _, exist = w.mobCa.Peek(key)
	if exist {
		return nil, ErrMobileExist
	}

	var loadLock = w.loadMobDocker
	loadLock.Lock(key)
	defer loadLock.Unlock(key)

	var err = w.mobStore.AddMobile(m, uid, now)
	if err != nil {
		if w.mobStore.IsDupErr(err) {
			return nil, ErrMobileExist
		}
		return nil, err
	}
	var mo = newUMobModel(m, uid, now)
	w.mobCa.Set(key, mo)
	return nil, err
}

//upsert mobile user id
func (w *WhoIsSimple) upsertMobileUID(m *model.MobileInfo, uid int32, now time.Time) (*model.UMobile, error) {
	var key = m.Key()
	var loadLock = w.loadMobDocker
	loadLock.Lock(key)
	defer loadLock.Unlock(key)

	var err = w.mobStore.UpsertMobileUID(m, uid, now)
	if err != nil {
		return nil, err
	}
	var r, ok = w.mobCa.Get(key)
	var nw *model.UMobile
	if !ok {
		//set new in cache
		nw = newUMobModel(m, uid, now)
	} else {
		//refresh cache
		var pre = r.(*model.UMobile)
		nw = &model.UMobile{}
		*nw = *pre
		nw.UserID, nw.UpdatedAt = uid, now
	}

	w.mobCa.Set(uid, nw)
	return nw, nil
}

func newUMobModel(m *model.MobileInfo, uid int32, now time.Time) *model.UMobile {
	return &model.UMobile{
		Mobile:    m.Mobile,
		Area:      m.Area,
		UserID:    uid,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
