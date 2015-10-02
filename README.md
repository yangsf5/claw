Run
====

game server framework & engine (golang)

### Run template
#### 1. Run master
```
./run.py -t master
```
#### 2. Run a harbor
```
./run.py -t harbor1
```

### Run examples
please see the README.md in examples


Architecture
====
TODO


Code Flow
====
![Code Flow](http://7xlbbk.dl1.z0.glb.clouddn.com/claw/arch/code-flow.png)

Services 
====
### master 集群中心控制节点
监听一个端口，等待各harbor链接上来并处理或转发harbor传过来的消息。

目前消息只有两种类型service.master.LOGIN和service.master.BROADCAST。

### harbor 工作节点
各具体业务逻辑服务均工作在harbor上。harbor启动时去master注册自己，然后监听master广播过来的消息，并转发给自己内部相应服务。

### error 集中处理错误信息
接受各服务传过来的错误消息，然后做集中处理，例如打印在控制台或者记录日志，或者入库。

### gate 门卫
所有TCP链接接入gate服务，gate服务充当服务器上各服务于客户端的沟通桥梁。

### web 网页服务
web服务实现一个http服务器

### test 测试服务
可以往这里发送测试数据或者调试信息

### agent 代理服务
TODO

