行情发布接口消息格式定义
===
*Protocol Buffers v3*

目录结构说明
---
    目录 models 是 protobuf 模式文件目录
    目录 models 下的 *.proto 是各接口使用的消息以及数据定义


目标语言协议构建说明
---

你可以手工使用`protoc`编译器生成自己需要的协议语言版本，例如 javanano
```bash
cd models
mkdir output
protoc --javanano_out=output *.proto
```

也可以使用`make`利用预先写好的`Makefile`规则在当前目录自动生成对应语言版本的协议源码，例如构建 c++, java 版的协议：
```bash
make cpp java
```

下面是目前`make`支持的语言列表及示例

```help
Usage: make <target language>
Support target language list:
    go        Generate Golang source file.
    cpp       Generate C++ header and source.
    java      Generate Java source file.
    objc      Generate Objective C header and source.
    javanano  Generate Java Nano source file.
For example:
    make go
    make go [target]
    make all
```

