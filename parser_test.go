package xml

import (
	"github.com/xyjwsj/xml_parser/util"
	"log"
	"testing"
)

func TestParseXml(t *testing.T) {
	xml := ParseXml("/Users/wushaojie/Downloads/AndroidManifest.xml")
	json := util.Struct2EscapeJson(xml, true)
	log.Println(json)
}
