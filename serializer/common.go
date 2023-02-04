package serializer

// Response 基础序列化器
type Response struct {
	StatusCode int         `json:"status"`
	Data       interface{} `json:"data"`
	StatusMsg  string      `json:"msg"`
	Error      string      `json:"error"`
	Token      string      `json:"token"`
}

// DataList 带有总数的Data结构
type DataList struct {
	Item  interface{} `json:"item"`
	Total uint        `json:"total"`
}

// TokenData 带有token的Data结构
type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

// TrackedErrorResponse 有追踪信息的错误反应
type TrackedErrorResponse struct {
	Response
	TrackID string `json:"track_id"`
}

// BulidListResponse 带有总数的列表构建器
func BuildListResponse(items interface{}, total uint) Response {
	return Response{
		StatusCode: 200,
		Data: DataList{
			Item:  items,
			Total: total,
		},
		StatusMsg: "ok",
	}
}
