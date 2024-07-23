package main

import (
	"encoding/xml" //用于XML的编解码
	"fmt"
	"io/ioutil" //用于读取文件的全部内容到内存
	"os"        //用于文件操作
)

type Recurlyservers struct { //用于匹配XML文档的顶层结构，包含服务器列表、版本号和描述
	XMLName     xml.Name `xml:"server"`
	Version     string   `xml:"version,attr"`
	Svs         []server `xml:"server"`
	Description string   `xml:",innerxml"`
}

type server struct { //用于匹配XML中的<server>元素，包含服务器名称和IP地址
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

func main() {
	file, err := os.Open("server.xml") // .
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file) //使用ioutil.ReadAll读取文件的全部内容
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	v := Recurlyservers{}
	err = xml.Unmarshal(data, &v) //使用xml.Unmarshal函数将读取的XML数据解析到Recurlyservers类型的变量v中
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println(v)
}

/*将xml文件解析成对应的struct对象是通过xml.Unmarshal来完成的，这个过程是如何实现的？可以看到我们的struct定义后面多了一些类似于xml:"serverName"这样的内容,这个是struct的一个特性，它们被称为 struct tag，它们是用来辅助反射的。我们来看一下Unmarshal的定义：


func Unmarshal(data []byte, v interface{}) error
我们看到函数定义了两个参数，第一个是XML数据流，第二个是存储的对应类型，目前支持struct、slice和string，XML包内部采用了反射来进行数据的映射，所以v里面的字段必须是导出的。Unmarshal解析的时候XML元素和字段怎么对应起来的呢？这是有一个优先级读取流程的，首先会读取struct tag，如果没有，那么就会对应字段名。必须注意一点的是解析的时候tag、字段名、XML元素都是大小写敏感的，所以必须一一对应字段。

*/
