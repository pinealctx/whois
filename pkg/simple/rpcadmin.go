package simple

import (
	"context"
	"github.com/pinealctx/whois/api/model"
	"github.com/pinealctx/whois/api/pb"
	"time"
)

func (r *WhoIsRunner) SetMobileUID(_ context.Context, req *pb.ReqMobileUID) (*pb.RspUserInfo, error) {
	var mk = &model.MobileInfo{Mobile: req.Mobile, Area: req.Area}
	var _, err = r.inner.upsertMobileUID(mk, req.UserID, time.Now())
	if err != nil {
		return nil, err
	}
	r.inner.updateMobileInUser(req.UserID, *mk)
	return r.getUserInfo(req.UserID)
}

func (r *WhoIsRunner) GetUserInfoByID(_ context.Context, req *pb.ReqUserKey) (*pb.RspUserInfo, error) {
	return r.getUserInfo(req.UserID)
}

func (r *WhoIsRunner) GetUserInfoByMobile(_ context.Context, req *pb.ReqMobileKey) (*pb.RspUserInfo, error) {
	var mk = &model.MobileInfo{Mobile: req.Mobile, Area: req.Area}
	var mob, err = r.inner.getMobileInfo(mk)
	if err != nil {
		return nil, err
	}
	return r.getUserInfo(mob.UserID)
}
