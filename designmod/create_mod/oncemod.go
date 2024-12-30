package create_mod

//单例模式（Singleton Pattern）确保一个类只有一个实例，并提供一个全局访问点。单例模式通常用于管理共享资源，如数据库连接、配置文件、日志记录等
/*
Singleton 是单例模式的结构体。
instance 是一个全局变量，用于存储单例的实例。
once 是一个sync.Once类型的变量，用于确保GetInstance函数只被执行一次。
GetInstance 函数是全局访问点，用于获取单例的实例。如果实例不存在，它将创建一个实例并返回。
使用sync.Once确保即使在并发环境下，GetInstance函数也只会被执行一次，从而保证单例的实例只被创建一次

*/
import (
	"fmt"
	"sync"
)

// Singleton 结构体定义
type Singleton struct {
}

// instance 存储单例实例
var instance *Singleton
var once sync.Once

// GetInstance 是一个全局访问点，确保只创建一个实例
func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}

func OnceMod() {
	s1 := GetInstance()
	s2 := GetInstance()

	fmt.Println(s1 == s2) // 输出 true，说明s1和s2是同一个实例

	//test Signal module
	s1.doSomething()
	s2.doSomething() // 输出 Singleton is working...，Singleton可以执行方法

}

// doSomething 是单例可以执行的方法
func (s *Singleton) doSomething() {
	fmt.Println("Singleton is working...")
	fmt.Println("This is a temp")
}
