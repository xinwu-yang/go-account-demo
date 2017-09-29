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
		stream = domain.Stream{StreamName: streamName, CreateTime: createTime}
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

func GetStream() base.Json {
	json := base.GetJson()
	stream := dao.GetNewStream()
	json.Ok = base.SUCCESS
	json.Content = stream
	return json
}

func CreateSubtitle(streamId int64, content string, createTime int64) base.Json {
	json := base.GetJson()
	subtitle := domain.Subtitle{StreamId: streamId, Content: content, CreateTime: createTime}
	orm.NewOrm().Insert(&subtitle)
	json.Ok = base.SUCCESS
	return json
}

func GetSubtitles(streamId int64) base.Json {
	json := base.GetJson()
	subtitles := dao.GetSubtitles(streamId)
	json.Array = subtitles
	json.Ok = base.SUCCESS
	return json
}
