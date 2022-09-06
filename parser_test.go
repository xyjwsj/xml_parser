package xml

import (
	"testing"
)

func TestParseXml(t *testing.T) {
	//xml := ParseXml("/Users/wushaojie/Downloads/test/AndroidManifest_1.xml")
	////xml := ParseXml("/Users/wushaojie/Downloads/values/colors.xml")
	//json := util.Struct2EscapeJson(xml, true)
	//log.Println(json)
	//Serializer(xml, true, "/Users/wushaojie/Downloads/test.xml")
	ParseXml("/Users/wushaojie/Documents/project/golang/package-core/apkBuild/srcSdkApkDir/res/values/strings.xml")
}

func TestMergeValues(t *testing.T) {
	ParseXml("/Users/wushaojie/Documents/project/golang/package-core/apkBuild/channelBuildApkDir/res/values/strings.xml")
}
