/*获取结构体的地址：
使用 & 操作符来获取结构体实例的地址，这将返回一个指向该结构体的指针。

访问字段：
通过指针访问字段时，使用 pointerName.FieldName 语法。

修改字段：
要通过指针修改字段，可以直接使用 pointerName.FieldName = newValue 语法。

调用方法：
如果结构体定义了方法，可以通过指针或实例来调用这些方法。

检查空指针：
在尝试访问或修改字段之前，应该检查指针是否为 nil，以避免空指针异常。

以下是一个示例，演示了如何通过指针访问和修改结构体实例的字段：
*/

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// 创建一个 Person 实例
	person := Person{Name: "Alice", Age: 30}

	// 获取 Person 实例的地址，得到一个指针
	personPtr := &person

	// 通过指针访问字段
	fmt.Println("Name before change:", personPtr.Name) // 输出原始名称

	// 通过指针修改字段
	personPtr.Name = "Bob"

	// 再次通过指针访问字段，查看修改结果
	fmt.Println("Name after change:", personPtr.Name) // 输出修改后的名称

	// 直接修改实例的字段也是可能的，因为实例本身也隐式地有自己的地址
	person.Age = 40

	// 验证字段修改
	fmt.Printf("%+v\n", person) // 输出修改后的实例
}
