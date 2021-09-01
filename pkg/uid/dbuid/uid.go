package dbuid

import (
	"github.com/pinealctx/whois/api/model"
	"github.com/pinealctx/whois/pkg/uid"
	"go.uber.org/atomic"
	"gorm.io/gorm"
)

type UIDGenOfDB struct {
	db  *gorm.DB
	min int32
	inc *atomic.Int32
}

func NewDBUIDGen(db *gorm.DB, min int32) uid.UserIDGen {
	return &UIDGenOfDB{
		db:  db,
		min: min,
	}
}

func (u *UIDGenOfDB) Load() error {
	var maxID, err = model.GetLastUID(u.db)
	if err != nil {
		return err
	}
	if maxID < u.min {
		maxID = u.min
	}
	u.inc = atomic.NewInt32(maxID)
	return nil
}

func (u *UIDGenOfDB) Next() int32 {
	return u.inc.Inc()
}
