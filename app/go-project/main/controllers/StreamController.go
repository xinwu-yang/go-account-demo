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
	s.Mapping("getSubtitleList", s.GetSubtitleList)
}

// @Title 创建直播流
// @Description 创建一条相对时间的直播流
// @Param   url path     CreateStreamParam true "播放地址，请urlEncode"
// @Success 200  {string} {"b" : 1}
// @router /createStream/:url [get]
func (s *StreamController) CreateStream() {
	url := s.GetString(":url")
	j := service.CreateStream(url, time.Now().Unix())
	s.Ctx.WriteString(j.String())
}

func (s *StreamController) CreateSubtitle() {

}

func (s *StreamController) GetStream() {

}

func (s *StreamController) GetSubtitleList() {

}
