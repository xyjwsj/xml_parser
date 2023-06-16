package xml

import (
	"log"
	"testing"
)

func TestParseXml(t *testing.T) {
	//xml := ParseXml("/Users/wushaojie/Downloads/test/AndroidManifest_1.xml")
	////xml := ParseXml("/Users/wushaojie/Downloads/values/colors.xml")
	//json := util.Struct2EscapeJson(xml, true)
	//log.Println(json)
	//Serializer(xml, true, "/Users/wushaojie/Downloads/test.xml")
	xml := ParseXml("/Users/wushaojie/Downloads/aaa1.xml")
	log.Println(xml)
	Serializer(xml, XmlHeaderType, "/Users/wushaojie/Downloads/aaa2.xml")
}

func TestMergeValues(t *testing.T) {
	ParseXml("/Users/wushaojie/Documents/project/golang/package-core/apkBuild/channelBuildApkDir/res/values/strings.xml")
}

func TestMergeValueNew(t *testing.T) {
	dirFileXml := "/Users/wushaojie/Library/Caches/access_package/sdk/expand/hoolaichannelDir/res/values/strings.xml"
	ParseXml(dirFileXml)
}
