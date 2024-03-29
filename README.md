English | [简体中文](./README.zh-CN.md)

# ms-agent

ms-agent is a tool that uses the Go language to write and receive zabbix alarm messages and send them to the [ZbxTable](https://github.com/canghai908/zbxtable) platform. It needs to be used with the ZbxTable platform.

## Compile

``` bash
mkdir -p $GOPATH/src/github.com/canghai908
cd $GOPATH/src/github.com/canghai908
git clone https://github.com/canghai908/ms-agent.git
cd ms-agent
./control build
./control pack
```

## Releases
V1.0.4
2022.07.11
1. Modify the default configuration file to the program directory
2. The host address of the configuration file does not need to be increased by /v1/receive


V1.0.1
2020.07.24 Fix log log permission issue

It will compile and generate binary files and package them into a compressed package

## Configuration

The ms-agent deployment needs to be deployed on Zabbix Server. The ms-agent receives zabbix alarm messages and sends them to the ZbxTable platform through the http protocol. Use zbxtable to complete the ms-agent configuration on the zabbix server platform

``` 
cd /usr/local/zbxtable
./zbxtable install
```

The display log is as follows

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

This step will create ms-agent on Zabbix Server with random password, configure related actions and media, and associate it with users

## Installation
````
cd /opt/
wget https://dl.cactifans.com/zbxtable/ms-agent-1.0.4.tar.gz
tar zxvf ms-agent-1.0.4.tar.gz
mv ms-agent-1.0.4 ms-agent
````

After decompression, a ms-agent binary file and an app.ini configuration file are generated.

| Program | Function |
| :------- | :-------------------------------------- -------- |
| ms-agent | Receive alerts generated by Zabbix platform and send to ZbxTable platform |
| app.ini | ms-agent configuration file |

Copy ms-agent to the Alertscripts directory of your zabbix server. The default path is /usr/lib/zabbix/alertscripts/. You can also specify the alertscripts directory by modifying the Zabbix Server configuration file.
Modify the Alertscripts directory of the zabbix server
vi zabbix_server.conf

````
AlertScriptsPath=/usr/lib/zabbix/alertscripts
````

Restart Zabbix Server to take effect.
Copy the ms-agent binary and app.ini configuration file to the alarm script directory configured by the zabbix server

````
cp ms-agent/* /usr/lib/zabbix/alertscripts/
````

Give the ms-agent script executable permission

````
chmod a+x /usr/lib/zabbix/alertscripts/ms-agent
````

So far the basic installation is complete

### config file
The zabbix server will call ms-agent to send alarms, and will also read the app.ini configuration file in the ms-agent program directory. The default content is as follows

````
[app]
Debug = 0
TenantID = zabbix01
LogSavePath = /tmp
Host = http://192.168.10.10:8088
Token = 2d7a7ab0b0be493ab0bb9a925e4a30d2
````

Debug is the program log level 0 is debug, 1 is info

LogSavePath is the log directory, the default is /tmp directory

TenantID Tenant id, the default is sufficient. If multiple sets of ms-agent are sent to the same zbxtable, it is recommended to repeat them.

Host is the access address of the ZbxTable system, the default is http:+ server IP:8088

Token The Token that communicates with ZbxTable can be modified by itself, and it needs to be consistent with the ZbxTable platform configuration, otherwise the alarm cannot be received.

## Debug

You can modify the configuration file to open the Debug mode, and view the log /tmp/ms-agent_yyyymmdd.log

## License

<img alt="Apache-2.0 license" src="https://s3-gz01.didistatic.com/n9e-pub/image/apache.jpeg" width="128">

Zbxtable is available under the Apache-2.0 license. See the [LICENSE](LICENSE) file for more info.
