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

func GetNewStream() domain.Stream {
	var stream domain.Stream
	o := orm.NewOrm()
	o.QueryTable("stream").OrderBy("-stream_id").Limit(1).One(&stream)
	return stream
}

func GetSubtitles(streamId int64) []*domain.Subtitle {
	var subtitles []*domain.Subtitle
	o := orm.NewOrm()
	o.QueryTable("subtitle").Filter("stream_id", streamId).All(&subtitles)
	return subtitles
}
