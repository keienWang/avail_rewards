## Avail 项目奖励验证

这个程序用于验证 Avail 项目的奖励资格。它使用了 Ethereum 区块链上的智能合约来确认用户的资格，并通过向 Avail 项目的 Web 服务发送签名消息来进行验证。

### 用法

1. **安装 Go 编程语言**

   确保你的计算机上已经安装了 Go 编程语言。你可以在 [Go 官方网站](https://golang.org/) 上找到安装说明。

2. **获取代码**

   通过 `git clone` 命令或者下载 ZIP 压缩包的方式获取代码到本地环境。

3. **修改配置**

   在代码中的 `GetAvail` 函数中，你需要修改以下参数：

   - `YOUR_INFURA_PROJECT_ID`：将其替换为你的 Infura 项目 ID，用于连接到以太坊网络。
   - `const urlLogin = "https://claim-api.availproject.org/check-rewards"`：此处定义了 Avail 项目的 Web 服务 URL。如果需要更改，请替换为正确的 URL。

4. **运行程序**

   使用命令行进入到代码所在的目录，然后运行以下命令：

   ```
   go run main.go
   ```

   程序将会开始生成私钥并发送验证请求。请注意，这可能需要一些时间。

### 注意事项

- **网络连接**：确保你的计算机能够连接到互联网，以便访问 Infura 和 Avail 项目的服务器。
- **Gas 费用**：此代码示例中的验证请求不会产生任何 Gas
