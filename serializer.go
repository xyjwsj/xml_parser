package xml

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type HeaderType int

const (
	No = iota
	XmlHeaderType
	StandaloneYesXmlHeaderType
	StandaloneNoXmlHeaderType
)
const (
	standaloneNoXmlHeader  = "<?xml version=\"1.0\" encoding=\"utf-8\" standalone=\"no\"?>"
	standaloneYesXmlHeader = "<?xml version=\"1.0\" encoding=\"utf-8\" standalone=\"yes\"?>"
	xmlHeader              = "<?xml version=\"1.0\" encoding=\"utf-8\"?>"
)

var headerTypeContent = map[HeaderType]string{
	XmlHeaderType:              xmlHeader,
	StandaloneYesXmlHeaderType: standaloneYesXmlHeader,
	StandaloneNoXmlHeaderType:  standaloneNoXmlHeader,
}

type LineContent struct {
	Content string
	End     bool
}

func Serializer(tag Tag, headerType HeaderType, filePath string) {
	contentChan := make(chan LineContent, 1000)
	stopChan := make(chan int, 2)
	go writeFile(filePath, contentChan, stopChan)
	if headerType > No {
		contentChan <- LineContent{
			Content: headerTypeContent[headerType],
			End:     false,
		}
	}
	writeSingleDom(&tag, 0, contentChan)
	contentChan <- LineContent{
		Content: "",
		End:     true,
	}
	<-stopChan
}

func writeSingleDom(tag *Tag, hierarchy int, contentChan chan<- LineContent) {
	dom := createTagDom(tag.Name, tag.Attribute, tag.Value, len(tag.ChildTags) > 0)
	content := createTab(hierarchy) + dom
	contentChan <- LineContent{
		Content: content,
		End:     false,
	}
	for idx, item := range tag.ChildTags {
		writeSingleDom(item, hierarchy+1, contentChan)
		if idx == len(tag.ChildTags)-1 {
			cStr := createTab(hierarchy) + "</" + tag.Name + ">"
			contentChan <- LineContent{
				Content: cStr,
				End:     false,
			}
		}
	}
}

func createTagDom(tagName string, attributes map[string]string, val string, existChild bool) string {
	dom := "<" + tagName
	for key, va := range attributes {
		dom += " " + key + "=\"" + va + "\""
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

func writeFile(filePath string, contentChan <-chan LineContent, stop chan<- int) {
	f, err := os.Create(filePath)
	if err != nil {
		log.Println("Create File Error:" + err.Error())
		os.Exit(1)
	}
	defer f.Close()

	w := bufio.NewWriter(f)

	for {
		select {
		case content := <-contentChan:
			if content.End {
				goto Loop
			}
			fmt.Fprintln(w, content.Content)
		default:
		}
	}
Loop:
	w.Flush()
	stop <- 1
}
