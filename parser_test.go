package xml

import (
	"github.com/xyjwsj/xml_parser/util"
	"log"
	"testing"
)

func TestParseXml(t *testing.T) {
	xml := ParseXml("/Users/wushaojie/Downloads/AndroidManifest.xml")
	//xml := ParseXml("/Users/wushaojie/Downloads/values/colors.xml")
	json := util.Struct2EscapeJson(xml, true)
	log.Println(json)
	Serializer(xml, true, "/Users/wushaojie/Downloads/test.xml")
}
