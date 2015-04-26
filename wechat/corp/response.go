package corp

import (
	"encoding/xml"
	"errors"
	"fmt"
)

type CommonResponseMessageHeader struct {
	ToUserName   string `mapstructure:"ToUserName"  xml:"ToUserName" json:"ToUserName"`
	FromUserName string `mapstructure:"FromUserName" xml:"FromUserName" json:"FromUserName"`
	CreateTime   int64  `mapstructure:"CreateTime"  xml:"CreateTime" json:"CreateTime"`
	MsgType      string `mapstructure:"MsgType"     xml:"MsgType" json:"MsgType"`
}

type ResText struct {
	XMLName                     xml.Name `xml:"xml" json:"-"`
	CommonResponseMessageHeader `mapstructure:",squash"`

	Content string `mapstructure:"Content" xml:"Content" json:"Content"` // 文本消息内容
}

func NewResText(to, from string, timestamp int64, content string) (text *ResText) {
	return &ResText{
		CommonResponseMessageHeader: CommonResponseMessageHeader{
			ToUserName:   to,
			FromUserName: from,
			CreateTime:   timestamp,
			MsgType:      MsgTypeText,
		},
		Content: content,
	}
}

type ResImage struct {
	XMLName                     xml.Name `xml:"xml" json:"-"`
	CommonResponseMessageHeader `mapstructure:",squash"`

	Image struct {
		MediaId string `mapstructure:"MediaId" xml:"MediaId" json:"MediaId"` // 图片文件id，可以调用上传媒体文件接口获取
	} `mapstructure:"Image" xml:"Image" json:"Image"`
}

func NewResImage(to, from string, timestamp int64, mediaId string) (image *ResImage) {
	image = &ResImage{
		CommonResponseMessageHeader: CommonResponseMessageHeader{
			ToUserName:   to,
			FromUserName: from,
			CreateTime:   timestamp,
			MsgType:      MsgTypeImage,
		},
	}
	image.Image.MediaId = mediaId
	return
}

type ResVoice struct {
	XMLName                     xml.Name `xml:"xml" json:"-"`
	CommonResponseMessageHeader `mapstructure:",squash"`

	Voice struct {
		MediaId string `mapstructure:"MediaId" xml:"MediaId" json:"MediaId"` // 语音文件id，可以调用上传媒体文件接口获取
	} `mapstructure:"Voice" xml:"Voice" json:"Voice"`
}

func NewResVoice(to, from string, timestamp int64, mediaId string) (voice *ResVoice) {
	voice = &ResVoice{
		CommonResponseMessageHeader: CommonResponseMessageHeader{
			ToUserName:   to,
			FromUserName: from,
			CreateTime:   timestamp,
			MsgType:      MsgTypeVoice,
		},
	}
	voice.Voice.MediaId = mediaId
	return
}

type ResVideo struct {
	XMLName                     xml.Name `xml:"xml" json:"-"`
	CommonResponseMessageHeader `mapstructure:",squash"`

	Video struct {
		MediaId     string `mapstructure:"MediaId"            xml:"MediaId"   json:"MediaId"`                // 视频文件id，可以调用上传媒体文件接口获取
		Title       string `mapstructure:"Title,omitempty"    xml:"MediaId"   json:"Title,omitempty"`        // 视频消息的标题
		Description string `mapstructure:"Description,omitempty" xml:"MediaId" json:"Description,omitempty"` // 视频消息的描述
	} `mapstructure:"Video" xml:"Video" json:"Video"`
}

func NewResVideo(to, from string, timestamp int64, mediaId, title, description string) (video *ResVideo) {
	video = &ResVideo{
		CommonResponseMessageHeader: CommonResponseMessageHeader{
			ToUserName:   to,
			FromUserName: from,
			CreateTime:   timestamp,
			MsgType:      MsgTypeVideo,
		},
	}
	video.Video.MediaId = mediaId
	video.Video.Title = title
	video.Video.Description = description
	return
}

type ResArticle struct {
	Title       string `mapstructure:"Title,omitempty" xml:"Title" json:"Title,omitempty"`                   // 图文消息标题
	Description string `mapstructure:"Description,omitempty" xml:"Description" json:"Description,omitempty"` // 图文消息描述
	PicURL      string `mapstructure:"PicUrl,omitempty"    xml:"PicUrl"  json:"PicUrl,omitempty"`            // 图片链接，支持JPG、PNG格式，较好的效果为大图360*200，小图200*200
	URL         string `mapstructure:"Url,omitempty"      xml:"Url"   json:"Url,omitempty"`                  // 点击图文消息跳转链接
}

type ResNews struct {
	XMLName                     xml.Name `xml:"xml" json:"-"`
	CommonResponseMessageHeader `mapstructure:",squash"`

	ArticleCount int          `mapstructure:"ArticleCount"        xml:"ArticleCount"    json:"ArticleCount"` // 图文条数，默认第一条为大图。图文数不能超过10，否则将会无响应
	Articles     []ResArticle `mapstructure:"Articles,omitempty"  xml:"Articles>item,omitempty" json:"Articles,omitempty"`
}

func NewResNews(to, from string, timestamp int64, articles []ResArticle) (news *ResNews) {
	news = &ResNews{
		CommonResponseMessageHeader: CommonResponseMessageHeader{
			ToUserName:   to,
			FromUserName: from,
			CreateTime:   timestamp,
			MsgType:      MsgTypeNews,
		},
	}
	news.Articles = articles
	news.ArticleCount = len(articles)
	return
}

// 检查 News 是否有效，有效返回 nil，否则返回错误信息
func (news *ResNews) CheckValid() (err error) {
	n := len(news.Articles)
	if n != news.ArticleCount {
		err = fmt.Errorf("图文消息的 ArticleCount == %d, 实际文章个数为 %d", news.ArticleCount, n)
		return
	}
	if n <= 0 {
		err = errors.New("图文消息里没有文章")
		return
	}
	if n > NewsArticleCountLimit {
		err = fmt.Errorf("图文消息的文章个数不能超过 %d, 现在为 %d", NewsArticleCountLimit, n)
		return
	}
	return
}
