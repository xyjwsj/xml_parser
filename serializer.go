package xml

import (
	"fmt"
	"log"
)

const (
	xmlHeader = "<?xml version=\"1.0\" encoding=\"utf-8\" standalone=\"no\"?>"
)

func Serializer(tag Tag, supportHeader bool) {
	if supportHeader {
		log.Println(xmlHeader)
	}
	writeSingleDom(&tag, 0)
}

func writeSingleDom(tag *Tag, hierarchy int) {
	dom := createTagDom(tag.Name, tag.Attribute, tag.Value, len(tag.ChildTags) > 0)
	//log.Println(dom)
	fmt.Println(createTab(hierarchy) + dom)
	for idx, item := range tag.ChildTags {
		writeSingleDom(item, hierarchy+1)
		if idx == len(tag.ChildTags)-1 {
			fmt.Println(createTab(hierarchy) + "</" + tag.Name + ">")
		}
	}
}

func createTagDom(tagName string, attributes map[string]string, val string, existChild bool) string {
	dom := "<" + tagName
	for key, val := range attributes {
		dom += " " + key + "=\"" + val + "\""
	}
	if existChild {
		dom += ">"
	} else {
		if val != "" {
			dom += ">" + val + "</" + tagName + ">"
		} else {
			dom += "/>"
		}
	}
	return dom
}

func createTab(num int) string {
	space := ""
	for i := 0; i < num; i++ {
		space += "    "
	}
	return space
}
