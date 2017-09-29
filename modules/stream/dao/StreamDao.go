package dao

import (
	"com.cxria/modules/stream/domain"
	"github.com/astaxie/beego/orm"
)

func GetStream(streamId int64) domain.Stream {
	o := orm.NewOrm()
	stream := domain.Stream{StreamId: streamId}
	err := o.Read(&stream)
	if err == orm.ErrNoRows {
		return domain.Stream{}
	}
	return stream
}
