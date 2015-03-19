package corp

import "encoding/xml"

const (
	MsgTypeText     = "text"
	MsgTypeImage    = "image"
	MsgTypeVoice    = "voice"
	MsgTypeVideo    = "video"
	MsgTypeFile     = "file"
	MsgTypeNews     = "news"
	MsgTypeMPNews   = "mpnews"
	MsgTypeLocation = "location" // 地理位置消息
	MsgTypeEvent    = "event"    // 事件推送
)

const NewsArticleCountLimit = 10

// 微信服务器推送过来的消息(事件)通用的消息头
type CommonMessageHeader struct {
	ToUserName   string `xml:"ToUserName"   json:"ToUserName"`
	FromUserName string `xml:"FromUserName" json:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"   json:"CreateTime"`
	MsgType      string `xml:"MsgType"      json:"MsgType"`
	AgentId      int64  `xml:"AgentID"      json:"AgentID"`
}

type ReqText struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonMessageHeader

	MsgId   int64  `xml:"MsgId"   json:"MsgId"`   // 消息id，64位整型
	Content string `xml:"Content" json:"Content"` // 文本消息内容
}

func UnmarshalXML(data []byte, v interface{}) error {
	err := xml.Unmarshal(data, v)
	return err
}
func GetText(data string) (*ReqText, error) {
	var result ReqText
	err := UnmarshalXML([]byte(data), &result)
	return &result, err
}

type ReqImage struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonMessageHeader

	MsgId   int64  `xml:"MsgId"   json:"MsgId"`   // 消息id，64位整型
	MediaId string `xml:"MediaId" json:"MediaId"` // 图片媒体文件id，可以调用获取媒体文件接口拉取数据
	PicURL  string `xml:"PicUrl"  json:"PicUrl"`  // 图片链接
}

func GetImage(data string) (*ReqImage, error) {
	var result ReqImage
	err := UnmarshalXML([]byte(data), &result)
	return &result, err
}

type ReqVoice struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonMessageHeader

	MsgId   int64  `xml:"MsgId"   json:"MsgId"`   // 消息id，64位整型
	MediaId string `xml:"MediaId" json:"MediaId"` // 语音媒体文件id，可以调用获取媒体文件接口拉取数据
	Format  string `xml:"Format"  json:"Format"`  // 语音格式，如amr，speex等
}

func GetVoice(data string) (*ReqVoice, error) {
	var result ReqVoice
	err := UnmarshalXML([]byte(data), &result)
	return &result, err
}

type ReqVideo struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonMessageHeader

	MsgId        int64  `xml:"MsgId"        json:"MsgId"`        // 消息id，64位整型
	MediaId      string `xml:"MediaId"      json:"MediaId"`      // 视频媒体文件id，可以调用获取媒体文件接口拉取数据
	ThumbMediaId string `xml:"ThumbMediaId" json:"ThumbMediaId"` // 视频消息缩略图的媒体id，可以调用获取媒体文件接口拉取数据
}

func GetVideo(data string) (*ReqVideo, error) {
	var result ReqVideo
	err := UnmarshalXML([]byte(data), &result)
	return &result, err
}

type ReqLocation struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonMessageHeader

	MsgId     int64   `xml:"MsgId"      json:"MsgId"`      // 消息id，64位整型
	LocationX float64 `xml:"Location_X" json:"Location_X"` // 地理位置纬度
	LocationY float64 `xml:"Location_Y" json:"Location_Y"` // 地理位置经度
	Scale     int     `xml:"Scale"      json:"Scale"`      // 地图缩放大小
	Label     string  `xml:"Label"      json:"Label"`      // 地理位置信息
}

func GetLocation(data string) (*ReqLocation, error) {
	var result ReqLocation
	err := UnmarshalXML([]byte(data), &result)
	return &result, err
}
