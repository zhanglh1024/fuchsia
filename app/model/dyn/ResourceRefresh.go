package dyn

import (
	"time"
)

type ResourceRefresh struct {
	Id          int64
	RoleId		string
	TypeId	 	int16
	First       int
	RefreshTime		time.Time
}
