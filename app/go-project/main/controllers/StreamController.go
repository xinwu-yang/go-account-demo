package controllers

import (
	"github.com/astaxie/beego"
	"com.cxria/modules/stream/service"
	"time"
)

type StreamController struct {
	beego.Controller
}

func (s *StreamController) URLMapping() {
	s.Mapping("createStream", s.CreateStream)
	s.Mapping("createSubtitle", s.CreateSubtitle)
	s.Mapping("getStream", s.GetStream)
	s.Mapping("getSubtitles", s.GetSubtitles)
}

// @Title 创建直播流
// @Description 创建一条相对时间的直播流
// @Param   url  path     true "播放地址，请urlEncode"
// @Success 200  {string} {"b" : 1}
// @router /createStream/:url [get]
func (s *StreamController) CreateStream() {
	url := s.GetString(":url")
	j := service.CreateStream(url, time.Now().Unix())
	s.Ctx.WriteString(j.String())
}

// @Title 创建直播流的字幕
// @Description 创建一条相对时间的直播流字幕
// @Param   streamId    path     true "流id"
// @Param   content     path     true "字幕内容，请urlEncode"
// @Param   createTime  path     true "字幕出现时间，第几秒"
// @Success 200  {string} {"b" : 1}
// @router /createSubtitle/:streamId/:content/:createTime [get]
func (s *StreamController) CreateSubtitle() {
	streamId, _ := s.GetInt64(":streamId")
	createTime, _ := s.GetInt64(":createTime")
	content := s.GetString(":content")
	j := service.CreateSubtitle(streamId, content, createTime)
	s.Ctx.WriteString(j.String())
}

// @Title 获取当前正在直播的流地址
// @Description 获取当前正在直播的流地址
// @Success 200  {string} {"b" : 1}
// @router /getStream [get]
func (s *StreamController) GetStream() {
	j := service.GetStream()
	s.Ctx.WriteString(j.String())
}

// @Title 获取当前正在直播的流的所有字幕
// @Description 获取当前正在直播的流的所有字幕
// @Param   streamId    path     true "流id"
// @Success 200  {string} {"b" : 1}
// @router /getSubtitles/:streamId [get]
func (s *StreamController) GetSubtitles() {
	streamId, _ := s.GetInt64(":streamId")
	j := service.GetSubtitles(streamId)
	s.Ctx.WriteString(j.String())
}
