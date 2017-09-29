package domain

type Stream struct {
	StreamId   int64 `orm:"pk"`
	StreamName string
	CreateTime int64
}

type Subtitle struct {
	SubtitleId int64 `orm:"pk"`
	StreamId   int64
	Content    string
	CreateTime int64
}
