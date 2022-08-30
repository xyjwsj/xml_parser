package xml

import (
	"github.com/xyjwsj/xml_parser/util"
	"io"
	"strings"
)

var (
	endSymbol = [...]string{"/>", "</"}
	tagSymbol = [...]string{"/>", "</", "<", ">"}
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
		filter := strings.TrimSpace(headerFilter(line))
		if filter == "" {
			return
		}
		descriptor := parseLine(filter)
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

		if descriptor.NoSymbol {
			operateTag.Value = descriptor.Text
			return
		}

		if descriptor.Text != "" {
			operateTag.Value = descriptor.Text
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
					operateTag.Attribute[key] = fixVal(val)
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
		NoSymbol:   false,
		Start:      false,
		End:        false,
		Pop:        false,
		TagName:    "",
		Attribute:  make(map[string]string, 0),
		StartChild: false,
		Text:       "",
	}
	for _, item := range endSymbol {
		if strings.Contains(line, item) {
			descriptor.End = true
			break
		}
	}

	if !existSymbol(line) {
		descriptor.NoSymbol = true
		descriptor.Text = line
	}

	if val := valueSymbol(line); val != "" {
		descriptor.Text = val
		line = strings.ReplaceAll(line, val, "")
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
		if item == "" { //
			continue
		}
		if !strings.Contains(item, "=") {
			descriptor.TagName = item
			continue
		}
		attr := strings.Split(item, "=")
		val := strings.ReplaceAll(attr[1], "/"+descriptor.TagName, "")
		descriptor.Attribute[attr[0]] = fixVal(val)
	}
	return descriptor
}

func originContent(line string) string {
	match := util.Match(line, "<(.*?)>")
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

func existSymbol(line string) bool {
	return strings.Contains(line, "<") ||
		strings.Contains(line, ">") ||
		strings.Contains(line, "/>") ||
		strings.Contains(line, "</")
}

func valueSymbol(line string) string {
	match := util.Match(line, ">(.*)<")
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

func fixVal(val string) string {
	return strings.ReplaceAll(val, "\"", "")
}
