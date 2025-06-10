# Go 语言创建型设计模式

创建型设计模式关注对象的创建机制，帮助创建对象的同时隐藏创建逻辑，使系统更加灵活。以下是 Go 语言中常用的创建型设计模式实现：

## 1. 单例模式 (Singleton)

确保一个类只有一个实例，并提供全局访问点。

```go
package singleton

import (
    "sync"
)

type singleton struct {
    // 单例的属性
    data string
}

var (
    instance *singleton
    once     sync.Once
)

// GetInstance 获取单例实例
func GetInstance() *singleton {
    once.Do(func() {
        instance = &singleton{data: "初始化数据"}
    })
    return instance
}

// SetData 设置数据
func (s *singleton) SetData(data string) {
    s.data = data
}

// GetData 获取数据
func (s *singleton) GetData() string {
    return s.data
}
```

使用示例：

```go
s1 := singleton.GetInstance()
s2 := singleton.GetInstance()
// s1 和 s2 是同一个实例
```

## 2. 工厂方法模式 (Factory Method)

定义一个用于创建对象的接口，让子类决定实例化哪个类。

```go
package factory

// Product 产品接口
type Product interface {
    Use() string
}

// ConcreteProductA 具体产品A
type ConcreteProductA struct{}

func (p *ConcreteProductA) Use() string {
    return "使用产品A"
}

// ConcreteProductB 具体产品B
type ConcreteProductB struct{}

func (p *ConcreteProductB) Use() string {
    return "使用产品B"
}

// Factory 工厂接口
type Factory interface {
    CreateProduct() Product
}

// ConcreteFactoryA 具体工厂A
type ConcreteFactoryA struct{}

func (f *ConcreteFactoryA) CreateProduct() Product {
    return &ConcreteProductA{}
}

// ConcreteFactoryB 具体工厂B
type ConcreteFactoryB struct{}

func (f *ConcreteFactoryB) CreateProduct() Product {
    return &ConcreteProductB{}
}
```

使用示例：

```go
factoryA := &factory.ConcreteFactoryA{}
productA := factoryA.CreateProduct()
fmt.Println(productA.Use()) // 输出: 使用产品A
```

## 3. 抽象工厂模式 (Abstract Factory)

提供一个创建一系列相关或相互依赖对象的接口，而无需指定它们具体的类。

```go
package abstractfactory

// Button 按钮接口
type Button interface {
    Paint() string
}

// WinButton Windows按钮
type WinButton struct{}

func (b *WinButton) Paint() string {
    return "渲染Windows按钮"
}

// MacButton Mac按钮
type MacButton struct{}

func (b *MacButton) Paint() string {
    return "渲染Mac按钮"
}

// Checkbox 复选框接口
type Checkbox interface {
    Paint() string
}

// WinCheckbox Windows复选框
type WinCheckbox struct{}

func (c *WinCheckbox) Paint() string {
    return "渲染Windows复选框"
}

// MacCheckbox Mac复选框
type MacCheckbox struct{}

func (c *MacCheckbox) Paint() string {
    return "渲染Mac复选框"
}

// GUIFactory GUI工厂接口
type GUIFactory interface {
    CreateButton() Button
    CreateCheckbox() Checkbox
}

// WinFactory Windows工厂
type WinFactory struct{}

func (f *WinFactory) CreateButton() Button {
    return &WinButton{}
}

func (f *WinFactory) CreateCheckbox() Checkbox {
    return &WinCheckbox{}
}

// MacFactory Mac工厂
type MacFactory struct{}

func (f *MacFactory) CreateButton() Button {
    return &MacButton{}
}

func (f *MacFactory) CreateCheckbox() Checkbox {
    return &MacCheckbox{}
}
```

使用示例：

```go
var factory abstractfactory.GUIFactory

// 根据操作系统选择工厂
if os == "Windows" {
    factory = &abstractfactory.WinFactory{}
} else {
    factory = &abstractfactory.MacFactory{}
}

button := factory.CreateButton()
checkbox := factory.CreateCheckbox()
fmt.Println(button.Paint())
fmt.Println(checkbox.Paint())
```

## 4. 建造者模式 (Builder)

将一个复杂对象的构建与它的表示分离，使得同样的构建过程可以创建不同的表示。

```go
package builder

// Product 产品
type Product struct {
    PartA string
    PartB string
    PartC string
}

// Builder 建造者接口
type Builder interface {
    BuildPartA()
    BuildPartB()
    BuildPartC()
    GetResult() *Product
}

// ConcreteBuilder 具体建造者
type ConcreteBuilder struct {
    product *Product
}

func NewConcreteBuilder() *ConcreteBuilder {
    return &ConcreteBuilder{product: &Product{}}
}

func (b *ConcreteBuilder) BuildPartA() {
    b.product.PartA = "部件A"
}

func (b *ConcreteBuilder) BuildPartB() {
    b.product.PartB = "部件B"
}

func (b *ConcreteBuilder) BuildPartC() {
    b.product.PartC = "部件C"
}

func (b *ConcreteBuilder) GetResult() *Product {
    return b.product
}

// Director 指挥者
type Director struct {
    builder Builder
}

func NewDirector(builder Builder) *Director {
    return &Director{builder: builder}
}

func (d *Director) Construct() *Product {
    d.builder.BuildPartA()
    d.builder.BuildPartB()
    d.builder.BuildPartC()
    return d.builder.GetResult()
}
```

使用示例：

```go
builder := builder.NewConcreteBuilder()
director := builder.NewDirector(builder)
product := director.Construct()
fmt.Printf("产品部件: %s, %s, %s\n", product.PartA, product.PartB, product.PartC)
```

## 5. 原型模式 (Prototype)

通过复制现有的实例来创建新的实例，而不是通过实例化类。

```go
package prototype

import "fmt"

// Cloneable 可克隆接口
type Cloneable interface {
    Clone() Cloneable
    GetName() string
}

// ConcretePrototype 具体原型
type ConcretePrototype struct {
    name string
    data map[string]string
}

func NewConcretePrototype(name string) *ConcretePrototype {
    return &ConcretePrototype{
        name: name,
        data: make(map[string]string),
    }
}

func (p *ConcretePrototype) Clone() Cloneable {
    clone := &ConcretePrototype{
        name: p.name + "_clone",
        data: make(map[string]string),
    }
    
    // 深拷贝数据
    for k, v := range p.data {
        clone.data[k] = v
    }
    
    return clone
}

func (p *ConcretePrototype) GetName() string {
    return p.name
}

func (p *ConcretePrototype) SetData(key, value string) {
    p.data[key] = value
}

func (p *ConcretePrototype) GetData(key string) string {
    return p.data[key]
}
```

使用示例：

```go
original := prototype.NewConcretePrototype("原型1")
original.SetData("key1", "value1")

// 克隆原型
clone := original.Clone().(*prototype.ConcretePrototype)
fmt.Println(clone.GetName())       // 输出: 原型1_clone
fmt.Println(clone.GetData("key1")) // 输出: value1
```

## 6. 对象池模式 (Object Pool)

通过重用对象来减少创建和销毁对象的开销。

```go
package objectpool

import (
    "errors"
    "sync"
)

// PooledObject 池化对象
type PooledObject struct {
    ID int
}

// ObjectPool 对象池
type ObjectPool struct {
    idle   []*PooledObject
    active []*PooledObject
    capacity int
    mutex sync.Mutex
}

// NewObjectPool 创建对象池
func NewObjectPool(capacity int) *ObjectPool {
    pool := &ObjectPool{
        idle:     make([]*PooledObject, 0),
        active:   make([]*PooledObject, 0),
        capacity: capacity,
    }
    
    // 预创建对象
    for i := 0; i < capacity; i++ {
        pool.idle = append(pool.idle, &PooledObject{ID: i})
    }
    
    return pool
}

// Acquire 获取对象
func (p *ObjectPool) Acquire() (*PooledObject, error) {
    p.mutex.Lock()
    defer p.mutex.Unlock()
    
    if len(p.idle) == 0 {
        return nil, errors.New("对象池已空")
    }
    
    obj := p.idle[0]
    p.idle = p.idle[1:]
    p.active = append(p.active, obj)
    return obj, nil
}

// Release 释放对象
func (p *ObjectPool) Release(obj *PooledObject) error {
    p.mutex.Lock()
    defer p.mutex.Unlock()
    
    for i, activeObj := range p.active {
        if activeObj == obj {
            p.active = append(p.active[:i], p.active[i+1:]...)
            p.idle = append(p.idle, obj)
            return nil
        }
    }
    
    return errors.New("对象不在活动列表中")
}
```

使用示例：

```go
pool := objectpool.NewObjectPool(10)

obj1, _ := pool.Acquire()
fmt.Printf("获取对象: %d\n", obj1.ID)

// 使用对象...

pool.Release(obj1)
fmt.Println("释放对象")
```

## 实际应用场景

1. **单例模式**：数据库连接池、配置管理器
2. **工厂方法**：根据用户输入创建不同的服务实例
3. **抽象工厂**：创建跨平台UI组件
4. **建造者模式**：构建复杂的HTTP请求或配置对象
5. **原型模式**：复制预配置的对象模板
6. **对象池**：连接池、内存池等资源管理

Go语言的创建型模式实现通常比其他面向对象语言更简洁，得益于Go的接口和组合特性。这些模式可以帮助你写出更加灵活、可维护和可扩展的代码。