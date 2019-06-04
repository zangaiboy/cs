package model

type UserInfo struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`  //姓名
	State int8   `json:"state"` //状态
}
