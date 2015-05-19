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

// 原生支付通知
type NativePay struct {
	XMLName     struct{} `xml:"xml" json:"-"`
	AppId       string   `xml:"appid"   json:"appid"`
	MchId       string   `xml:"mch_id" json:"mch_id"`
	IsSubscribe string   `xml:"is_subscribe" json:"is_subscribe"`
	NonceStr    string   `xml:"nonce_str" json:"nonce_str"`
	ProductId   string   `xml:"product_id" json:"product_id"`
	Sign        string   `xml:"sign" json:"sign"`
}

func GetNativePay(data string) (*NativePay, error) {
	var result NativePay
	err := UnmarshalXML([]byte(data), &result)
	return &result, err
}

// 支付后通知
// 微信接口的变态... 谁特么在数据交互中返回动态字段? 返回个嵌套的array都是正常的吧
// 所以暂时只支持了4个代金券或立减优惠批次ID, 普通场景应该够用了.....
type PayNotify struct {
	XMLName        struct{} `xml:"xml" json:"-"`
	AppId          string   `xml:"appid"   json:"appid"`
	MchId          string   `xml:"mch_id" json:"mch_id"`
	DeviceInfo     string   `xml:"device_info" json:"device_info"`
	NonceStr       string   `xml:"nonce_str" json:"nonce_str"`
	Sign           string   `xml:"sign" json:"sign"`
	ResultCode     string   `xml:"result_code" json:"result_code"`
	ErrCode        string   `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrDescription string   `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	OpenId         string   `xml:"openid" json:"openid"`
	IsSubscribe    string   `xml:"is_subscribe" json:"is_subscribe"`
	TradeType      string   `xml:"trade_type" json:"trade_type"`
	BankType       string   `xml:"bank_type" json:"bank_type"`
	TotalFee       int      `xml:"total_fee" json:"total_fee"`
	FeeType        string   `xml:"fee_type,omitempty" json:"fee_type,omitempty"`
	CashFee        string   `xml:"cash_fee" json:"cash_fee"`
	CashFeeType    string   `xml:"cash_fee_type" json:"cash_fee_type"`
	CouponFee      int      `xml:"coupon_fee" json:"coupon_fee"`
	CouponCount    int      `xml:"coupon_count" json:"coupon_count"`
	TransactionId  string   `xml:"transaction_id" json:"transaction_id"`
	OutTradeNo     string   `xml:"out_trade_no" json:"out_trade_no"`
	Attach         string   `xml:"attach,omitempty" json:"attach,omitempty"`
	TimeEnd        string   `xml:"time_end" json:"time_end"`
	CouponBachId1  string   `xml:"coupon_batch_id_1,omitempty" json:"coupon_batch_id_1,omitempty"`
	CouponId1      string   `xml:"coupon_id_1,omitempty" json:"coupon_id_1,omitempty"`
	CouponFee1     int      `xml:"coupon_fee_1,omitempty" json:"coupon_fee_1,omitempty"`
	CouponBachId2  string   `xml:"coupon_batch_id_2,omitempty" json:"coupon_batch_id_2,omitempty"`
	CouponId2      string   `xml:"coupon_id_2,omitempty" json:"coupon_id_2,omitempty"`
	CouponFee2     int      `xml:"coupon_fee2" json:"coupon_fee2"`
	CouponBachId3  string   `xml:"coupon_batch_id_3,omitempty" json:"coupon_batch_id_3,omitempty"`
	CouponId3      string   `xml:"coupon_id_3,omitempty" json:"coupon_id_3,omitempty"`
	CouponFee3     int      `xml:"coupon_fee_3,omitempty" json:"coupon_fee_3,omitempty"`
	CouponBachId4  string   `xml:"coupon_batch_id_4,omitempty" json:"coupon_batch_id_4,omitempty"`
	CouponId4      string   `xml:"coupon_id_4,omitempty" json:"coupon_id_4,omitempty"`
	CouponFee4     int      `xml:"coupon_fee_4,omitempty" json:"coupon_fee_4,omitempty"`
}

func GetPayNotify(data string) (*PayNotify, error) {
	var result PayNotify
	err := UnmarshalXML([]byte(data), &result)
	return &result, err
}

func UnmarshalXML(data []byte, v interface{}) error {
	err := xml.Unmarshal(data, v)
	return err
}
