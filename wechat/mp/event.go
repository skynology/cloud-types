package mp

import (
	"fmt"
	"strings"

	"github.com/skynology/wechat/mp"
)

const (
	EventTypeClick = "CLICK" // 点击菜单拉取消息时的事件推送
	EventTypeView  = "VIEW"  // 点击菜单跳转链接时的事件推送

	EventTypeSubscribe   = "subscribe"   // 订阅, 包括点击订阅和扫描二维码
	EventTypeUnsubscribe = "unsubscribe" // 取消订阅
	EventTypeScan        = "SCAN"        // 已经订阅的用户扫描二维码事件
	EventTypeLocation    = "LOCATION"    // 上报地理位置事件

	// 请注意, 下面的事件仅支持微信iPhone5.4.1以上版本, 和Android5.4以上版本的微信用户,
	// 旧版本微信用户点击后将没有回应, 开发者也不能正常接收到事件推送.
	EventTypeScanCodePush    = "scancode_push"      // scancode_push：扫码推事件的事件推送
	EventTypeScanCodeWaitMsg = "scancode_waitmsg"   // scancode_waitmsg：扫码推事件且弹出“消息接收中”提示框的事件推送
	EventTypePicSysPhoto     = "pic_sysphoto"       // pic_sysphoto：弹出系统拍照发图的事件推送
	EventTypePicPhotoOrAlbum = "pic_photo_or_album" // pic_photo_or_album：弹出拍照或者相册发图的事件推送
	EventTypePicWeixin       = "pic_weixin"         // pic_weixin：弹出微信相册发图器的事件推送
	EventTypeLocationSelect  = "location_select"    // location_select：弹出地理位置选择器的事件推送
)

// 关注事件(普通关注)
type ReqSubscribeEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	mp.CommonMessageHeader

	Event string `xml:"Event" json:"Event"` // 事件类型，subscribe(订阅)
}

func GetSubscribeEvent(data string) (*ReqSubscribeEvent, error) {
	var result ReqSubscribeEvent
	err := UnmarshalXML([]byte(data), &result)
	return &result, err
}

// 取消关注
type ReqUnsubscribeEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	mp.CommonMessageHeader

	Event string `xml:"Event" json:"Event"` // 事件类型，unsubscribe(取消订阅)
}

func GetUnsubscribeEvent(data string) (*ReqUnsubscribeEvent, error) {
	var result ReqUnsubscribeEvent
	err := UnmarshalXML([]byte(data), &result)
	return &result, err
}

// 用户未关注时，扫描带参数二维码进行关注后的事件推送
type ReqSubscribeByScanEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	mp.CommonMessageHeader

	Event    string `xml:"Event"    json:"Event"`    // 事件类型，subscribe
	EventKey string `xml:"EventKey" json:"EventKey"` // 事件KEY值，qrscene_为前缀，后面为二维码的参数值
	Ticket   string `xml:"Ticket"   json:"Ticket"`   // 二维码的ticket，可用来换取二维码图片
}

// 获取二维码参数
func (event *ReqSubscribeByScanEvent) Scene() (scene string, err error) {
	const prefix = "qrscene_"
	if !strings.HasPrefix(event.EventKey, prefix) {
		err = fmt.Errorf("EventKey 应该以 %q 为前缀: %q", prefix, event.EventKey)
		return
	}
	scene = event.EventKey[len(prefix):]
	return
}

func GetSubscribeByScanEvent(data string) (*ReqSubscribeByScanEvent, error) {
	var result ReqSubscribeByScanEvent
	err := UnmarshalXML([]byte(data), &result)
	return &result, err
}

// 用户已关注时，扫描带参数二维码的事件推送
type ReqScanEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	mp.CommonMessageHeader

	Event    string `xml:"Event"    json:"Event"`    // 事件类型，SCAN
	EventKey string `xml:"EventKey" json:"EventKey"` // 事件KEY值，是一个32位无符号整数，即创建二维码时的二维码scene_id
	Ticket   string `xml:"Ticket"   json:"Ticket"`   // 二维码的ticket，可用来换取二维码图片
}

func GetScanEvent(data string) (*ReqScanEvent, error) {
	var result ReqScanEvent
	err := UnmarshalXML([]byte(data), &result)
	return &result, err
}

// 上报地理位置事件
type ReqLocationEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	mp.CommonMessageHeader

	Event     string  `xml:"Event"     json:"Event"`     // 事件类型，LOCATION
	Latitude  float64 `xml:"Latitude"  json:"Latitude"`  // 地理位置纬度
	Longitude float64 `xml:"Longitude" json:"Longitude"` // 地理位置经度
	Precision float64 `xml:"Precision" json:"Precision"` // 地理位置精度
}

func GetLocationEvent(data string) (*ReqLocationEvent, error) {
	var result ReqLocationEvent
	err := UnmarshalXML([]byte(data), &result)
	return &result, err
}
