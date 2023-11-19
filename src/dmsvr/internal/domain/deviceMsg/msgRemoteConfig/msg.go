package msgRemoteConfig

type RemoteConfigMsg struct {
	Method  string `json:"method"`           //操作方法
	Code    int64  `json:"code,omitempty"`   //状态码
	Status  string `json:"status,omitempty"` //返回信息
	Payload string `json:"payload,optional"` //配置信息
}

const (
	TypeGet  = "get"  //远程配置请求
	TypePush = "push" //远程配置推送
)
