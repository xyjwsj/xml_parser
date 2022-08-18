package xml

import (
	"github.com/xyjwsj/xml_parser/util"
	"log"
)

func ParseXml(xmlPath string) Tag {
	tag := Tag{}
	util.ReadFileLine(xmlPath, func(err error, line string) {
		log.Println(line)
	})
	return tag
}
