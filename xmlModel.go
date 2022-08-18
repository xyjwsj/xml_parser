package xml

// Tag 标签描述
type Tag struct {
	Start     bool              // 开始标签
	Name      string            // 标签名
	Attribute map[string]string // 标签属性
	ChildTags []*Tag            // 子标签
	Parent    *Tag              // 父标签引用
}

type LineDescriptor struct {
	Start      bool              // 是否开始标签
	End        bool              // 结束当前标签
	Pop        bool              // 返回上一层树
	StartChild bool              // 开始子标签
	TagName    string            // 标签名
	Attribute  map[string]string // 标签属性
}
