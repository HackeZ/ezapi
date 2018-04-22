package model

// Category 类目信息
type Category struct {
	CID  int32  `json:"cid"`
	Name string `json:"name"`
}

// CategoryProp 类目属性
type CategoryProp struct {
	PID    int32               `json:"pid"`
	Name   string              `json:"name"`
	Values []CategoryPropValue `json:"value"`
	Type   int32               `json:"type"`
}

// CategoryPropValue 类目属性值
type CategoryPropValue struct {
	VID  int32  `json:"vid"`
	Name string `json:"name"`
}
