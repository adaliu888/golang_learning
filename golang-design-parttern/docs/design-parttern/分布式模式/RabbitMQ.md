# RabbitMQ 安装步骤

## Windows 安装步骤

1. **安装 Erlang**
   - RabbitMQ 需要 Erlang 环境
   - 访问 [Erlang 官网](https://www.erlang.org/downloads) 下载适合你系统的安装包
   - 运行安装程序并按照向导完成安装
   - 将 Erlang 的 bin 目录添加到系统环境变量 PATH 中

2. **安装 RabbitMQ**
   - 访问 [RabbitMQ 官网](https://www.rabbitmq.com/download.html) 下载 Windows 安装包
   - 运行安装程序并按照向导完成安装
   - 默认安装路径通常为 `C:\Program Files\RabbitMQ Server`

3. **启用管理插件**
   - 打开命令提示符或 PowerShell（以管理员身份运行）
   - 导航到 RabbitMQ 的 sbin 目录：
     ```
     cd "C:\Program Files\RabbitMQ Server\rabbitmq_server-x.y.z\sbin"
     ```
   - 启用管理插件：
     ```
     rabbitmq-plugins enable rabbitmq_management
     ```

4. **启动 RabbitMQ 服务**
   - 使用命令：
     ```
     rabbitmq-service start
     ```
   - 或通过 Windows 服务管理器启动 "RabbitMQ" 服务

5. **验证安装**
   - 打开浏览器访问管理界面：http://localhost:15672/
   - 默认用户名和密码都是 `guest`

## Linux (Ubuntu/Debian) 安装步骤

1. **安装 Erlang**
   ```bash
   sudo apt update
   sudo apt install erlang
   ```

2. **添加 RabbitMQ 仓库**
   ```bash
   curl -s https://packagecloud.io/install/repositories/rabbitmq/rabbitmq-server/script.deb.sh | sudo bash
   ```

3. **安装 RabbitMQ**
   ```bash
   sudo apt update
   sudo apt install rabbitmq-server
   ```

4. **启动 RabbitMQ 服务**
   ```bash
   sudo systemctl start rabbitmq-server
   sudo systemctl enable rabbitmq-server
   ```

5. **启用管理插件**
   ```bash
   sudo rabbitmq-plugins enable rabbitmq_management
   ```

6. **验证安装**
   - 打开浏览器访问管理界面：http://localhost:15672/
   - 默认用户名和密码都是 `guest`

## macOS 安装步骤

1. **使用 Homebrew 安装**
   ```bash
   brew update
   brew install rabbitmq
   ```

2. **启动 RabbitMQ 服务**
   ```bash
   brew services start rabbitmq
   ```
   或
   ```bash
   rabbitmq-server
   ```

3. **启用管理插件**
   ```bash
   rabbitmq-plugins enable rabbitmq_management
   ```

4. **验证安装**
   - 打开浏览器访问管理界面：http://localhost:15672/
   - 默认用户名和密码都是 `guest`

## Docker 安装方式

1. **拉取 RabbitMQ 镜像**
   ```bash
   docker pull rabbitmq:3-management
   ```

2. **运行 RabbitMQ 容器**
   ```bash
   docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management
   ```

3. **验证安装**
   - 打开浏览器访问管理界面：http://localhost:15672/
   - 默认用户名和密码都是 `guest`

## 创建新用户和虚拟主机

1. **添加新用户**
   ```bash
   rabbitmqctl add_user myuser mypassword
   ```

2. **设置用户权限**
   ```bash
   rabbitmqctl set_user_tags myuser administrator
   ```

3. **创建虚拟主机**
   ```bash
   rabbitmqctl add_vhost myvhost
   ```

4. **设置虚拟主机权限**
   ```bash
   rabbitmqctl set_permissions -p myvhost myuser ".*" ".*" ".*"
   ```

安装完成后，您可以通过管理界面或客户端库开始使用 RabbitMQ 进行消息队列操作。