package simple

import "github.com/pinealctx/whois/api/pb"

type WhoIsRunner struct {
	inner *WhoIsSimple
	pb.UnimplementedUserModServer
	pb.UnimplementedUserAdminServer
}
