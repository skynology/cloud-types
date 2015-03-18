package wechat

func getInt(v interface{}) int {
	t, _ := v.(int)
	return t
}

func getInt64(v interface{}) int64 {
	t := getInt(v)
	return int64(t)
}

func getString(v interface{}) string {
	t, _ := v.(string)
	return t
}
