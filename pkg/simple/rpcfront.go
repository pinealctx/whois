package simple

import (
	"context"
	"github.com/pinealctx/whois/api/model"
	"github.com/pinealctx/whois/api/pb"
	"time"
)

//SignInByMobile : login or register by mobile
func (r *WhoIsRunner) SignInByMobile(_ context.Context, req *pb.ReqMobile) (*pb.RspUserInfo, error) {
	return r.getOrCreateUMob(req.Mobile, req.Area, "", "")
}

//SignInByWechatOpen : login or register by we chat open
func (r *WhoIsRunner) SignInByWechatOpen(_ context.Context, req *pb.ReqWechat) (*pb.RspUserInfo, error) {
	return r.signInByWechat(req, true)
}

//SignInByWechatMiniApp : login or register by we chat mini app
func (r *WhoIsRunner) SignInByWechatMiniApp(_ context.Context, req *pb.ReqWechat) (*pb.RspUserInfo, error) {
	return r.signInByWechat(req, false)
}

//BindMobileByWechatOpen : bind mobile by we chat open
func (r *WhoIsRunner) BindMobileByWechatOpen(_ context.Context, req *pb.ReqWechatBindMobile) (*pb.RspUserInfo, error) {
	return r.bindMobileByWechat(req, true)
}

//BindMobileByWechatMiniApp : bind mobile by we chat mini app
func (r *WhoIsRunner) BindMobileByWechatMiniApp(_ context.Context, req *pb.ReqWechatBindMobile) (*pb.RspUserInfo, error) {
	return r.bindMobileByWechat(req, false)
}

//UpdateNickname : update user nick name
func (r *WhoIsRunner) UpdateNickname(_ context.Context, req *pb.ReqNickname) (*pb.RspUserInfo, error) {
	return r.updateUserField(req.UserID, req.NickName, r.inner.updateNick)
}

//UpdateAvatar : update user avatar
func (r *WhoIsRunner) UpdateAvatar(_ context.Context, req *pb.ReqAvatar) (*pb.RspUserInfo, error) {
	return r.updateUserField(req.UserID, req.Avatar, r.inner.updateAvatar)
}

func (r *WhoIsRunner) signInByWechat(req *pb.ReqWechat, isOpen bool) (*pb.RspUserInfo, error) {
	var we = WeInfoReqToModel(req)
	var now = time.Now()
	var wf, err = r.inner.upsertWeInfo(we, now, isOpen)
	if err != nil {
		return nil, err
	}
	if wf.UserID == 0 {
		//empty
		return WeInfoNoUIDToRsp(wf, now.Unix()), nil
	}
	var rsp *pb.RspUserInfo
	rsp, err = r.getUserInfo(wf.UserID)
	if err != nil {
		return nil, err
	}
	if wf.NickName != "" {
		rsp.User.NickName = wf.NickName
	}
	if wf.HeadImgUrl != "" {
		rsp.User.Avatar = wf.HeadImgUrl
	}
	return rsp, nil
}

func (r *WhoIsRunner) bindMobileByWechat(req *pb.ReqWechatBindMobile, isOpen bool) (*pb.RspUserInfo, error) {
	var mk = &model.WeChatKey{OpenID: req.OpenID, AppID: req.AppID}
	var wf, err = r.inner.getWeInfo(mk, isOpen)
	if err != nil {
		return nil, err
	}

	var rsp *pb.RspUserInfo
	rsp, err = r.getOrCreateUMob(req.Mobile, req.Area, wf.NickName, wf.HeadImgUrl)
	if err != nil {
		return nil, err
	}
	_, err = r.inner.setWeUID(mk, rsp.User.ID, time.Now(), isOpen)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func (r *WhoIsRunner) updateUserField(uid int32, val string,
	fn func(int32, string, time.Time) (*UCaItem, error)) (*pb.RspUserInfo, error) {

	var uc, err = fn(uid, val, time.Now())
	if err != nil {
		return nil, err
	}
	return UserCaToRsp(uc), nil
}

func (r *WhoIsRunner) getOrCreateUMob(mobile string, area string, nick string, avatar string) (*pb.RspUserInfo, error) {
	var mk = &model.MobileInfo{Mobile: mobile, Area: area}
	var mob, err = r.inner.getMobileInfo(mk)

	if err == nil {
		//exist
		return r.getUserInfo(mob.UserID)
	}
	if !r.inner.mobStore.IsNotFoundErr(err) {
		//other error, exist
		return nil, err
	}

	//add uid, then bind it
	return r.addUserWithMobile(mk, nick, avatar)
}

func (r *WhoIsRunner) getUserInfo(uid int32) (*pb.RspUserInfo, error) {
	var existU, err = r.inner.getUserInfo(uid)
	if err != nil {
		return nil, err
	}
	return UserCaToRsp(existU), nil
}

func (r *WhoIsRunner) addUserWithMobile(mk *model.MobileInfo, nick, avatar string) (*pb.RspUserInfo, error) {
	//add uid, then bind it
	var now = time.Now()
	var uid = r.inner.uidGen.Next()
	var nu, err = r.inner.addUserInfo(uid, nick, avatar, *mk, now)
	if err != nil {
		return nil, err
	}
	_, err = r.inner.addMobileInfo(mk, uid, now)
	if err != nil {
		return nil, err
	}
	return UserCaToRsp(nu), nil
}
