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

	var l = len(uc.Mobiles)
	if l > 0 {
		rsp.Mobiles = make([]*pb.RspMobile, l)
		for i := 0; i < l; i++ {
			rsp.Mobiles[i] = &pb.RspMobile{}
			rsp.Mobiles[i].Mobile, rsp.Mobiles[i].Area = uc.Mobiles[i].Mobile, uc.Mobiles[i].Area
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
