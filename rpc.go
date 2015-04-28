package types

// 云代码Session
type CloudSession struct {

	// Session 用户Id
	UserId string `json:"userId"`

	// 是否以Master Key权限调用
	Master bool `json:"master"`

	// 用户角色
	Roles []string `json:"roles"`

	// 用户状态
	Disabled bool `json:"disabled"`
}

type CloudError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type CloudLog struct {
	CreatedAt string `json:"createdAt"`
	Content   string `json:"content"`
	Flag      string `json:"flag"`
}

// 云代码传入参数
type CloudRequest struct {

	// 资源Id.
	// 在对单个资源操作时才会有
	ObjectId string `json:"objectId"`

	// Post 或 Put 等操作时前台传入的值
	Data map[string]interface{} `json:"data"`

	// 额外数据, 如微信服务器推过来的XMLstring
	ExtraData string `json:"extraData"`

	// 用户操作时的Session对象
	Session CloudSession `json:"session"`
}

// 云代码条用后返回结构
type CloudeResponse struct {

	// 是否成功返回
	Successed bool `json:"successed"`

	// 有修改字段的值
	Data map[string]interface{} `json:"data"`

	// 调用函数时返回的数据
	Result interface{} `json:"result"`

	// 额外字段, 如微信被动返回值, 直接编码友xml返回
	ExtraData string `json:"extraData"`

	// 需要隐藏字段
	Hide []string `json:"hide"`

	// 保护字段， 此字段设置后， 前台API无法修改
	Protect []string `json:"protect"`

	// 返回的错误
	Errors CloudError `json:"error"`

	// 调试用的log
	Logs []CloudLog `json:"logs"`
}
