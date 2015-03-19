package mp

import "encoding/xml"

const (
	// 微信服务器推送过来的消息类型
	MsgTypeText                    = "text"                      // 文本消息
	MsgTypeImage                   = "image"                     // 图片消息
	MsgTypeVoice                   = "voice"                     // 语音消息
	MsgTypeVideo                   = "video"                     // 视频消息
	MsgTypeLocation                = "location"                  // 地理位置消息
	MsgTypeLink                    = "link"                      // 链接消息
	MsgTypeEvent                   = "event"                     // 事件推送
	MsgTypeMusic                   = "music"                     // 音乐消息
	MsgTypeNews                    = "news"                      // 图文消息
	MsgTypeTransferCustomerService = "transfer_customer_service" // 将消息转发到多客服
)

// 微信服务器推送过来的消息(事件)通用的消息头
type CommonMessageHeader struct {
	ToUserName   string `mapstructure:"ToUserName"   json:"ToUserName"`
	FromUserName string `mapstructure:"FromUserName" json:"FromUserName"`
	CreateTime   int64  `mapstructure:"CreateTime"   json:"CreateTime"`
	MsgType      string `mapstructure:"MsgType"      json:"MsgType"`
}

// 文本消息
type ReqText struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonMessageHeader

	MsgId   int64  `xml:"MsgId"   json:"MsgId"`   // 消息id, 64位整型
	Content string `xml:"Content" json:"Content"` // 文本消息内容
}

func GetText(data string) (*ReqText, error) {
	var result ReqText
	err := UnmarshalXML([]byte(data), &result)
	return &result, err
}

// 图片消息
type ReqImage struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonMessageHeader

	MsgId   int64  `xml:"MsgId"   json:"MsgId"`   // 消息id, 64位整型
	MediaId string `xml:"MediaId" json:"MediaId"` // 图片消息媒体id，可以调用多媒体文件下载接口拉取数据。
	PicURL  string `xml:"PicUrl"  json:"PicUrl"`  // 图片链接
}

func GetImage(data string) (*ReqImage, error) {
	var result ReqImage
	err := UnmarshalXML([]byte(data), &result)
	return &result, err
}

// 语音消息
type ReqVoice struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonMessageHeader

	MsgId   int64  `xml:"MsgId"   json:"MsgId"`   // 消息id, 64位整型
	MediaId string `xml:"MediaId" json:"MediaId"` // 语音消息媒体id，可以调用多媒体文件下载接口拉取该媒体
	Format  string `xml:"Format"  json:"Format"`  // 语音格式，如amr，speex等

	// 语音识别结果，UTF8编码，
	// NOTE: 需要开通语音识别功能，否则该字段为空，即使开通了语音识别该字段还是有可能为空
	Recognition string `xml:"Recognition,omitempty" json:"Recognition,omitempty"`
}

func GetVoice(data string) (*ReqVoice, error) {
	var result ReqVoice
	err := UnmarshalXML([]byte(data), &result)
	return &result, err
}

// 视频消息
type ReqVideo struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonMessageHeader

	MsgId        int64  `xml:"MsgId"        json:"MsgId"`        // 消息id, 64位整型
	MediaId      string `xml:"MediaId"      json:"MediaId"`      // 视频消息媒体id，可以调用多媒体文件下载接口拉取数据。
	ThumbMediaId string `xml:"ThumbMediaId" json:"ThumbMediaId"` // 视频消息缩略图的媒体id，可以调用多媒体文件下载接口拉取数据。
}

func GetVideo(data string) (*ReqVideo, error) {
	var result ReqVideo
	err := UnmarshalXML([]byte(data), &result)
	return &result, err
}

// 地理位置消息
type ReqLocation struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonMessageHeader

	MsgId     int64   `xml:"MsgId"      json:"MsgId"`      // 消息id, 64位整型
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

// 链接消息
type ReqLink struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonMessageHeader

	MsgId       int64  `xml:"MsgId"       json:"MsgId"`       // 消息id, 64位整型
	Title       string `xml:"Title"       json:"Title"`       // 消息标题
	Description string `xml:"Description" json:"Description"` // 消息描述
	URL         string `xml:"Url"         json:"Url"`         // 消息链接
}

func GetLink(data string) (*ReqLink, error) {
	var result ReqLink
	err := UnmarshalXML([]byte(data), &result)
	return &result, err
}

func UnmarshalXML(data []byte, v interface{}) error {
	err := xml.Unmarshal(data, v)
	return err
}
