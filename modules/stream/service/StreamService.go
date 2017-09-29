package service

import (
	"com.cxria/base"
	"com.cxria/modules/stream/dao"
	"com.cxria/modules/stream/domain"
	"github.com/astaxie/beego/orm"
)

func CreateStream(streamName string, createTime int64) base.Json {
	json := base.GetJson()
	stream := dao.GetStream(1)
	if stream.StreamId == 0 {
		stream = domain.Stream{StreamId: 1, StreamName: streamName, CreateTime: createTime}
		orm.NewOrm().Insert(&stream)
	} else {
		stream.StreamName = streamName
		stream.CreateTime = createTime
		orm.NewOrm().Update(&stream)
	}
	json.Ok = base.SUCCESS
	json.Content = stream
	return json
}
