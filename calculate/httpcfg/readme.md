根据你提供的 `HttpClientConfig` 结构体定义，以下是对应的 JSON 格式示例：

```json
{
    "enable": true,
    "enable_autotest": false,
    "enable_debug": true,
    "name": "MyHttpClient",
    "tag": "v1.0",
    "url": "http://example.com",
    "connect-read-timeout": 30,
    "do-timeinterval": 10
}
```

### JSON 字段说明：
- **enable**: 布尔值，表示是否启用 HTTP 客户端。
- **enable_autotest**: 布尔值，表示是否启用自动测试。
- **enable_debug**: 布尔值，表示是否启用调试模式。
- **name**: 字符串，表示 HTTP 客户端的名称。
- **tag**: 字符串，表示版本标签。
- **url**: 字符串，表示请求的 URL。
- **connect-read-timeout**: 整数，表示连接读取超时时间（单位：秒）。
- **do-timeinterval**: 整数，表示执行时间间隔（单位：秒）。

### 使用示例
确保将上述 JSON 内容保存为 `config.json` 文件，并放置在你的 Go 程序可以访问的目录中。运行程序后，它将读取该配置文件并将其解码为 `HttpClientConfig` 结构体。
