zerorpc是一个在服务端进程上提供分布式通信的轻量级的、可靠的跨语言的库。它基于[ZeroMQ](http://www.zeromq.org/)和[MessagePack](http://msgpack.org/)。提供流式相应(streamed response)——就像[Python](http://lib.csdn.net/base/python)的生成器(generators)——这让zerorpc不仅仅是个典型的RPC引擎。内置心跳包、超时监测以及从失败请求中恢复。自我修复能力、第一类异常以及命令行工具让debug也变得极其简单。

## 安装

```
pip install zerorpc
```

## HelloWorld

Server.py

```
import zerorpc

class HelloRPC(object):
    def hello(self, name):
        return "Hello, %s" % name

s = zerorpc.Server(HelloRPC())
s.bind("tcp://0.0.0.0:4242")
s.run()123456789123456789
```

Client.py

```
import zerorpc

c = zerorpc.Client()
c.connect("tcp://127.0.0.1:4242")
print c.hello("RPC")1234512345
```

## 流式相应(Streaming Responses)

Server.py

```
import zerorpc

class StreamingRPC(object):
    @zerorpc.stream
    def streaming_range(self, fr, to, step):
        return xrange(fr, to, step)

s = zerorpc.Server(StreamingRPC())
s.bind("tcp://0.0.0.0:4242")
s.run()1234567891012345678910
```

Client.py

```
import zerorpc

c = zerorpc.Client()
c.connect("tcp://127.0.0.1:4242")

for item in c.streaming_range(10, 20, 2):
    print item12345671234567
```

## 第一类异常(First-class exception)

Server.py

```
import zerorpc

class ExceptionalRPC(object):
    def bad(self):
        raise Exception(":P")

s = zerorpc.Server(ExceptionalRPC())
s.bind("tcp://0.0.0.0:4242")
s.run()123456789123456789
```

Client.py

```
import zerorpc

c = zerorpc.Client()
c.connect("tcp://127.0.0.1:4242")

try:
    c.bad()
except Exception, e:
    print "An error occurred: %s" % e123456789123456789
```