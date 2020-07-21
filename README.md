# ms-agent

ms-agent 是一个使用 go 语言编写接受 zabbix 的告警消息并发送到 [ZbxTable](https://github.com/canghai908/zbxtable) 平台的工具，需配合 ZbxTable 平台使用。

## 编译

```bash
mkdir -p $GOPATH/src/github.com/canghai908
cd $GOPATH/src/github.com/canghai908
git clone https://github.com/canghai908/ms-agent.git
cd ms-agent
./control build
./control pack
```

会编译生成二进制文件，并打包到压缩包

## 配置

ms-agent 部署需部署在 Zabbix Server，ms-agent 接收 zabbix 的告警消息，通过 http 协议发送到 ZbxTable 平台，使用 zbxtable 完成 ms-agent 在 zabbix server 平台配置

```
cd /usr/local/zbxtable
./zbxtable install
```

显示如下日志

```
2020/07/18 23:22:16.881 [I] [install.go:43]  Zabbix API Address: http://zabbix-server/api_jsonrpc.php
2020/07/18 23:22:16.881 [I] [install.go:44]  Zabbix Admin User: Admin
2020/07/18 23:22:16.881 [I] [install.go:45]  Zabbix Admin Password: xxxxx
2020/07/18 23:22:17.716 [I] [install.go:52]  登录zabbix平台成功!
2020/07/18 23:22:17.879 [I] [install.go:69]  创建告警媒介成功!
2020/07/18 23:22:18.027 [I] [install.go:82]  创建告警用户组成功!
2020/07/18 23:22:18.198 [I] [install.go:113]  创建告警用户成功!
2020/07/18 23:22:18.198 [I] [install.go:114]  用户名:ms-agent
2020/07/18 23:22:18.198 [I] [install.go:115]  密码:xxxx
2020/07/18 23:22:18.366 [I] [install.go:167]  创建告警动作成功!
2020/07/18 23:22:18.366 [I] [install.go:168]  插件安装完成!
```

此步骤会在 Zabbix Server 创建 ms-agent，密码为随机，并配置相关 action 和 media，并关联到用户

## 安装

此程序必须部署在 Zabbix Server

```
yum install https://dl.cactifans.com/zabbix/ms-agent-1.0.0-1.el7.x86_64.rpm -y
```

环境信息

| 程序     | 路径                                  | 作用                                             |
| :------- | :------------------------------------ | :----------------------------------------------- |
| ms-agent | /usr/lib/zabbix/alertscripts/ms-agent | 接收 Zabbix 平台产生的告警并发送到 ZbxTable 平台 |
| app.ini  | /etc/ms-agent/app.ini                 | ms-agent 配置文件                                |

如果你的 Zabbix Server 的 alertscripts 目录不为/usr/lib/zabbix/alertscripts/ 需要移动 ms-agen 到你的 zabbix server 的 alertscripts 目录下即可,否则会在 Zabbix 告警页面出现找不到 ms-agent 的错误提示，也无法收到告警消息。
也可以修改 Zabbix Server 的配置文件，将 alertscripts 目录指向/usr/lib/zabbix/alertscripts/

vi zabbix_server.conf

```
AlertScriptsPath=/usr/lib/zabbix/alertscripts
```

修改后重启 Zabbix Server 生效

## Debug

可修改配置文件打开 Debug 模式，查看日志/tmp/ms-agent_yyyymmdd.log

## License

<img alt="Apache-2.0 license" src="https://s3-gz01.didistatic.com/n9e-pub/image/apache.jpeg" width="128">

Nightingale is available under the Apache-2.0 license. See the [LICENSE](LICENSE) file for more info.
