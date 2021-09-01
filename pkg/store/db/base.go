package db

import (
	"github.com/pinealctx/neptune/store/gormx"
)

type Base struct {
}

func (b Base) IsDupErr(err error) bool {
	return gormx.IsDupError(err)
}

func (b Base) IsNotFoundErr(err error) bool {
	return gormx.IsNotFoundErr(err)
}
