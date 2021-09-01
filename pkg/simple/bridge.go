package simple

import (
	"github.com/pinealctx/whois/api/model"
	"github.com/pinealctx/whois/api/pb"
)

func UserCaToRsp(uc *UCaItem) *pb.RspUserInfo {
	var u = &pb.RspUser{
		ID:        uc.UserMod.ID,
		State:     uc.UserMod.State,
		NickName:  uc.UserMod.NickName,
		Avatar:    uc.UserMod.Avatar,
		CreatedAt: uc.UserMod.CreatedAt.Unix(),
		UpdatedAt: uc.UserMod.UpdatedAt.Unix(),
	}
	var rsp = &pb.RspUserInfo{
		User: u,
	}
	if uc.Mob == nil {
		rsp.Mob = nil
	} else {
		rsp.Mob = &pb.RspMobile{
			Mobile: uc.Mob.Mobile,
			Area:   uc.Mob.Area,
		}
	}

	return rsp
}

func WeInfoReqToModel(wq *pb.ReqWechat) *model.WechatInfo {
	return &model.WechatInfo{
		OpenID:     wq.OpenID,
		AppID:      wq.AppID,
		NickName:   wq.NickName,
		Country:    wq.Country,
		Province:   wq.Province,
		City:       wq.City,
		HeadImgUrl: wq.HeadImgUrl,
		Sex:        wq.Sex,
	}
}

func WeInfoNoUIDToRsp(m *model.UWechat, ts int64) *pb.RspUserInfo {
	var u = &pb.RspUser{
		NickName:  m.NickName,
		Avatar:    m.HeadImgUrl,
		CreatedAt: ts,
		UpdatedAt: ts,
	}
	return &pb.RspUserInfo{User: u}
}
