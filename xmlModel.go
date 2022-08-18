package xml

// Tag 标签描述
type Tag struct {
	Name      string            // 标签名
	Attribute map[string]string // 标签属性
	Value     string            // 标签值
	ChildTags []Tag             //子标签
}
