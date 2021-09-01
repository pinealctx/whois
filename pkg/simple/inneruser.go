package simple

import (
	"github.com/pinealctx/whois/api/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

var (
	ErrUserExist = status.Errorf(codes.AlreadyExists, "user.already.exist")
)

type UCaItem struct {
	UserMod *model.User
	Mobiles []model.MobileInfo
}

func (u *UCaItem) Size() int {
	return 1
}

//get user info
func (w *WhoIsSimple) getUserInfo(uid int32) (*UCaItem, error) {
	var (
		r   interface{}
		ok  bool
		uc  *UCaItem
		err error
	)

	r, ok = w.userCa.Get(uid)
	if ok {
		return r.(*UCaItem), nil
	}

	var loadLock = w.loadUserDocker.GetLock(uid)
	loadLock.Lock()
	defer loadLock.Unlock()

	//retry -- cause other go routine can figure out
	r, ok = w.userCa.Get(uid)
	if ok {
		return r.(*UCaItem), nil
	}

	uc, err = w.loadUserAndMobiles(uid)
	if err != nil {
		return nil, err
	}

	w.userCa.Set(uid, uc)
	return uc, nil
}

//add user info
func (w *WhoIsSimple) addUserInfo(uid int32, nick, avatar string,
	mob model.MobileInfo, now time.Time) (*UCaItem, error) {
	var _, exist = w.userCa.Peek(uid)
	if exist {
		return nil, ErrUserExist
	}

	var loadLock = w.loadUserDocker.GetLock(uid)
	loadLock.Lock()
	defer loadLock.Unlock()

	var err = w.uStore.AddUser(uid, nick, avatar, now)
	if err != nil {
		if w.uStore.IsDupErr(err) {
			return nil, ErrUserExist
		}
		return nil, err
	}
	var uc = &UCaItem{
		UserMod: &model.User{
			ID:        uid,
			NickName:  nick,
			Avatar:    avatar,
			CreatedAt: now,
			UpdatedAt: now,
		},
		Mobiles: []model.MobileInfo{
			mob,
		},
	}
	w.userCa.Set(uid, uc)
	return uc, nil
}

//update nick
func (w *WhoIsSimple) updateNick(uid int32, nick string, now time.Time) (*UCaItem, error) {
	var updLock = w.updUserDocker.GetLock(uid)
	updLock.Lock()
	defer updLock.Unlock()

	var uc, err = w.getUserInfo(uid)
	if err != nil {
		return nil, err
	}

	if nick == uc.UserMod.NickName {
		return uc, nil
	}

	err = w.uStore.UpdateNick(uid, nick, now)
	if err != nil {
		return nil, err
	}
	var un = cloneUCaItem(uc)
	un.UserMod.NickName = nick
	un.UserMod.UpdatedAt = now
	w.userCa.Set(uid, un)
	return un, nil
}

//update nick
func (w *WhoIsSimple) updateAvatar(uid int32, avatar string, now time.Time) (*UCaItem, error) {
	var updLock = w.updUserDocker.GetLock(uid)
	updLock.Lock()
	defer updLock.Unlock()

	var uc, err = w.getUserInfo(uid)
	if err != nil {
		return nil, err
	}

	if avatar == uc.UserMod.Avatar {
		return uc, nil
	}

	err = w.uStore.UpdateAvatar(uid, avatar, now)
	if err != nil {
		return nil, err
	}
	var un = cloneUCaItem(uc)
	un.UserMod.NickName = avatar
	un.UserMod.UpdatedAt = now
	w.userCa.Set(uid, un)
	return un, nil
}

//add user info
func (w *WhoIsSimple) updateMobileInUser(uid int32, mob model.MobileInfo) {
	var loadLock = w.loadUserDocker.GetLock(uid)
	loadLock.Lock()
	defer loadLock.Unlock()

	var pre, ok = w.userCa.Peek(uid)
	if !ok {
		//do nothing
		return
	}
	var nw = cloneUCaItem(pre.(*UCaItem))
	nw.Mobiles = []model.MobileInfo{mob}
	w.userCa.Set(uid, nw)
}

//load user and mobile
func (w *WhoIsSimple) loadUserAndMobiles(uid int32) (*UCaItem, error) {
	var (
		u   *model.User
		mbs []model.MobileInfo
		err error
	)
	u, err = w.uStore.LoadUser(uid)
	if err != nil {
		return nil, err
	}

	mbs, err = w.uStore.LoadUserMobiles(uid)
	if err != nil {
		return nil, err
	}

	return &UCaItem{
		UserMod: u,
		Mobiles: mbs,
	}, nil
}

//clone user cache item
func cloneUCaItem(o *UCaItem) *UCaItem {
	//use new memory
	var un = &UCaItem{
		UserMod: &model.User{},
		Mobiles: make([]model.MobileInfo, len(o.Mobiles)),
	}
	*un.UserMod = *o.UserMod
	copy(un.Mobiles, o.Mobiles)
	return un
}
