 [English](./README.md) | 简体中文

# ms-agent

ms-agent 是一个使用 go 语言编写接收zabbix 的告警消息并发送到 [ZbxTable](https://github.com/canghai908/zbxtable) 平台的工具，需配合 ZbxTable 平台使用。

## 编译

``` bash
mkdir -p $GOPATH/src/github.com/canghai908
cd $GOPATH/src/github.com/canghai908
git clone https://github.com/canghai908/ms-agent.git
cd ms-agent
./control build
./control pack
```

## 更新记录
V1.0.4
2022.07.11 
1. 默认配置文件修改为程序目录下
2. 配置文件host地址不需要增加/v1/receive


V1.0.1
2020.07.24 修复 log 日志权限问题

会编译生成二进制文件，并打包到压缩包

## 配置

ms-agent 部署需部署在 Zabbix Server，ms-agent 接收 zabbix 的告警消息，通过 http 协议发送到 ZbxTable 平台，使用 zbxtable 完成 ms-agent 在 zabbix server 平台配置

``` 
cd /usr/local/zbxtable
./zbxtable install
```

显示如下日志

``` 
2022/07/04 16:27:48.252 [I] [command.go:163]  Create media type successfully!
2022/07/04 16:27:48.320 [I] [command.go:163]  Create user group successfully!
2022/07/04 16:27:48.575 [I] [command.go:163]  Create alarm user successfully!
2022/07/04 16:27:48.575 [I] [command.go:163]  Username : ms-agent
2022/07/04 16:27:48.575 [I] [command.go:163]  Password : qynNlKzMBx
2022/07/04 16:27:48.668 [I] [command.go:163]  Create alarm action successfully!
2022/07/04 16:27:48.668 [I] [command.go:163]  MS-Agent plugin configured successfully!
2022/07/04 16:27:48.668 [I] [command.go:163]  MS-Agent token is de0c0d234f054c74b3d87d715f69afb6
```

此步骤会在 Zabbix Server 创建 ms-agent，密码为随机，并配置相关 action 和 media，并关联到用户

## 安装
下载二进制文件，并解压

```
cd /opt/
wget https://dl.cactifans.com/zbxtable/ms-agent-1.0.4.tar.gz
tar zxvf ms-agent-1.0.4.tar.gz
mv ms-agent-1.0.4 ms-agent
```

解压之后生成一个 ms-agent 二进制文件,一个 app.ini 配置文件。

| 程序       | 作用                                             |
| :-------  | :----------------------------------------------- |
| ms-agent  | 接收 Zabbix 平台产生的告警并发送到 ZbxTable 平台 |
| app.ini   | ms-agent 配置文件                                |

拷贝 ms-agent 到你的 zabbix server 的 Alertscripts 目录下，默认路径为/usr/lib/zabbix/alertscripts/，也可通过修改 Zabbix Server 的配置文件指定alertscripts 目录。
修改zabbix server的Alertscripts目录
vi zabbix_server.conf

```
AlertScriptsPath=/usr/lib/zabbix/alertscripts
```

重启 Zabbix Server 生效.
拷贝ms-agent二进制及app.ini配置文件到zabbix server配置的告警脚本目录

```
cp ms-agent/* /usr/lib/zabbix/alertscripts/
```

赋予ms-agent脚本可执行权限

```
chmod a+x /usr/lib/zabbix/alertscripts/ms-agent
```

至此完成基本安装

### 配置文件
zabbix server会调用ms-agent进行告警的发送，同时会读取ms-agent程序目录下的app.ini配置文件,默认内容如下

```
[app]
Debug = 0
TenantID = zabbix01
LogSavePath = /tmp
Host = http://192.168.10.10:8088
Token = 2d7a7ab0b0be493ab0bb9a925e4a30d2
```

Debug 为程序日志级别 0 是 debug，1 为 info

LogSavePath 为日志目录，默认为/tmp 目录

TenantID 租户id，默认即可，如有多套ms-agent发送到同一个zbxtable，建议补重复即可

Host 为 ZbxTable 系统的访问地址，默认为 http:+ 服务器 IP:8088

Token 与 ZbxTable 通信的 Token,可自行修改,需要与 ZbxTable 平台配置保持一致即可，否则无法接收告警。

#### Debug

可修改配置文件打开 Debug 模式，查看日志文件名格式如下/tmp/ms-agent_yyyymmdd.log


## License

<img alt="Apache-2.0 license" src="https://s3-gz01.didistatic.com/n9e-pub/image/apache.jpeg" width="128">

Nightingale is available under the Apache-2.0 license. See the [LICENSE](LICENSE) file for more info.
