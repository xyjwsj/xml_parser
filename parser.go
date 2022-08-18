package xml

import (
	"github.com/xyjwsj/xml_parser/util"
	"io"
	"strings"
)

var (
	endSymbol = [...]string{"/>", "</"}
	tagSymbol = [...]string{"<", ">", "/>", "</"}
)

func ParseXml(xmlPath string) Tag {
	tag := Tag{
		Start:     false,
		Name:      "",
		Attribute: nil,
		ChildTags: nil,
		Parent:    nil,
	}
	operateTag := &tag
	util.ReadFileLine(xmlPath, func(err error, line string) {
		if err == io.EOF || line == "" {
			return
		}
		filter := headerFilter(line)
		descriptor := parseLine(strings.TrimSpace(filter))
		if descriptor.Pop {
			operateTag = operateTag.Parent
			return
		}
		if descriptor.Start && operateTag.Start {
			child := &Tag{
				Start:     true,
				Name:      "",
				Attribute: nil,
				ChildTags: nil,
				Parent:    operateTag,
			}
			operateTag.ChildTags = append(operateTag.ChildTags, child)
			operateTag = child
		}
		operateTag.Start = true
		if descriptor.TagName != "" {
			operateTag.Name = descriptor.TagName
		}
		if len(descriptor.Attribute) > 0 {
			if operateTag.Attribute == nil {
				operateTag.Attribute = descriptor.Attribute
			} else {
				for key, val := range descriptor.Attribute {
					operateTag.Attribute[key] = val
				}
			}
		}
		if descriptor.StartChild {
			if operateTag.ChildTags == nil {
				operateTag.ChildTags = make([]*Tag, 0)
			}
		}
		if descriptor.End {
			operateTag = operateTag.Parent
		}
	})
	return tag
}

func headerFilter(content string) string {
	flag := "?>"
	if !strings.Contains(content, flag) {
		return content
	}
	index := strings.Index(content, flag)
	return content[index+len(flag):]
}

func parseLine(line string) LineDescriptor {
	descriptor := LineDescriptor{
		Start:      false,
		End:        false,
		Pop:        false,
		TagName:    "",
		Attribute:  make(map[string]string, 0),
		StartChild: false,
	}
	for _, item := range endSymbol {
		if strings.Contains(line, item) {
			descriptor.End = true
			break
		}
	}

	if strings.HasPrefix(line, "<") {
		descriptor.Start = true
	}

	if strings.HasPrefix(line, "</") {
		descriptor.Pop = true
	}

	if strings.HasSuffix(line, ">") && !strings.HasSuffix(line, "/>") {
		descriptor.StartChild = true
	}

	content := originContent(line)

	attrs := strings.Split(content, " ")
	for _, item := range attrs {
		if !strings.Contains(item, "=") {
			descriptor.TagName = item
			continue
		}
		attr := strings.Split(item, "=")
		descriptor.Attribute[attr[0]] = attr[1]
	}
	return descriptor
}

func originContent(line string) string {
	for _, item := range tagSymbol {
		line = strings.ReplaceAll(line, item, "")
	}
	return line
}
