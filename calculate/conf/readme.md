要在 Go 中读取 YAML 配置文件，你可以使用 `gopkg.in/yaml.v2` 包。以下是一个示例代码，展示如何读取你提供的 `conf.yaml` 文件内容。

首先，确保你已经安装了 `yaml` 包：

```bash
go get gopkg.in/yaml.v2
```

然后，你可以使用以下代码读取配置文件：

```go
package main

import (
    "fmt"
    "io/ioutil"
    "log"

    "gopkg.in/yaml.v2"
)

// Config 结构体用于映射 YAML 配置
type Config struct {
    Database struct {
        Host     string `yaml:"host"`
        User     string `yaml:"user"`
        Dbname   string `yaml:"dbname"`
        Pwd      string `yaml:"pwd"`
    } `yaml:"database"`
}

func main() {
    // 读取 YAML 文件
    data, err := ioutil.ReadFile("calculate/conf/conf.yaml")
    if err != nil {
        log.Fatalf("error: %v", err)
    }

    // 解析 YAML 数据
    var config Config
    err = yaml.Unmarshal(data, &config)
    if err != nil {
        log.Fatalf("error: %v", err)
    }

    // 打印配置
    fmt.Printf("Database Host: %s\n", config.Database.Host)
    fmt.Printf("Database User: %s\n", config.Database.User)
    fmt.Printf("Database Name: %s\n", config.Database.Dbname)
    fmt.Printf("Database Password: %s\n", config.Database.Pwd)
}
```

### 代码说明：
1. **Config 结构体**：定义了一个结构体来映射 YAML 文件中的数据。
2. **读取文件**：使用 `ioutil.ReadFile` 读取 YAML 文件内容。
3. **解析 YAML**：使用 `yaml.Unmarshal` 将读取的数据解析到 `Config` 结构体中。
4. **打印配置**：最后打印出数据库的配置信息。

确保将代码中的文件路径 `calculate/conf/conf.yaml` 替换为你的实际文件路径。运行此程序后，它将输出配置文件中的数据库信息。



要将 `GetConfigDetail` 函数修改为返回配置数据而不是直接打印它，你可以将函数的返回类型更改为 `Config`，并返回解析后的配置对象。以下是更新后的代码：

```go
package conf

import (
	"log"

	"gopkg.in/yaml.v2"
)

// Config 结构体用于映射 YAML 配置
type Config struct {
	Database struct {
		Host   string `yaml:"host"`
		User   string `yaml:"user"`
		Dbname string `yaml:"dbname"`
		Pwd    string `yaml:"pwd"`
	} `yaml:"database"`
}

// GetConfigDetail 返回配置的详细信息
func GetConfigDetail() Config {
	// 读取 YAML 文件
	data, err := ioutil.ReadFile("calculate/conf/conf.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// 解析 YAML 数据
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// 返回配置
	return config
}
```

### 代码说明：
1. **返回类型**：将 `GetConfigDetail` 函数的返回类型更改为 `Config`。
2. **返回配置**：在函数的最后，返回解析后的 `config` 对象，而不是打印它。

### 使用示例：
你可以在其他地方调用 `GetConfigDetail` 函数并使用返回的配置对象，例如：

```go
func main() {
	config := GetConfigDetail()
	fmt.Printf("Database Host: %s\n", config.Database.Host)
	fmt.Printf("Database User: %s\n", config.Database.User)
	fmt.Printf("Database Name: %s\n", config.Database.Dbname)
	fmt.Printf("Database Password: %s\n", config.Database.Pwd)
}
```

这样，你就可以在需要的地方使用配置数据，而不是直接在 `GetConfigDetail` 函数中打印它。


要为 `GetConfigDetail` 函数添加测试，你可以创建一个新的测试文件，通常命名为 `config_test.go`。在这个文件中，你可以使用 Go 的内置测试框架来编写测试用例。以下是一个示例测试代码，假设你的配置文件 `conf.yaml` 是有效的。

首先，确保你的项目结构如下：

```
calculate/
├── conf/
│   ├── conf.yaml
│   ├── config.go
│   └── config_test.go
```

然后，在 `config_test.go` 文件中添加以下内容：

```go
package conf

import (
	"testing"
)

// TestGetConfigDetail 测试 GetConfigDetail 函数
func TestGetConfigDetail(t *testing.T) {
	expectedHost := "127.0.0.1"
	expectedUser := "root"
	expectedDbname := "test"
	expectedPwd := "123456"

	config := GetConfigDetail()

	if config.Database.Host != expectedHost {
		t.Errorf("expected host %s, got %s", expectedHost, config.Database.Host)
	}
	if config.Database.User != expectedUser {
		t.Errorf("expected user %s, got %s", expectedUser, config.Database.User)
	}
	if config.Database.Dbname != expectedDbname {
		t.Errorf("expected dbname %s, got %s", expectedDbname, config.Database.Dbname)
	}
	if config.Database.Pwd != expectedPwd {
		t.Errorf("expected password %s, got %s", expectedPwd, config.Database.Pwd)
	}
}
```

### 代码说明：
1. **测试包**：测试文件使用 `package conf`，与被测试的代码在同一包中。
2. **测试函数**：`TestGetConfigDetail` 是一个测试函数，Go 会自动识别以 `Test` 开头的函数。
3. **期望值**：定义了期望的数据库配置值。
4. **调用函数**：调用 `GetConfigDetail` 函数并将返回的配置与期望值进行比较。
5. **错误处理**：如果返回的值与期望值不匹配，使用 `t.Errorf` 输出错误信息。

### 运行测试：
在终端中，导航到项目目录并运行以下命令：

```bash
go test ./...
```

这将运行所有测试并显示结果。如果一切正常，你应该看到测试通过的消息。这样，你就为 `GetConfigDetail` 函数添加了测试。
