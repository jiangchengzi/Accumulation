如何打包rpm

使用工具frm

#### 安装frm	

FPM功能简单说就是将一种类型的包转换成另一种类型。

FPM的github：https://github.com/jordansissel/fpm 

1.支持的源类型包：

 dir: 将目录打包成所需要的类型，可以用于源码编译安装的软件包

 rpm: 对rpm进行转换

 gem: 对rubygem包进行转换

 python: 将Python模块打包成相应的类型

2.安装ruby环境和gem命令： 这里我会提前把开发包装好

fpm 是 ruby写的，因此系统环境需要ruby，且版本必须大于1.8.5

```shell
yum -y install ruby rubygems ruby-devel
```

3.查看当前ruby源：

添加国内源：

```shell
gem sources -a http://mirrors.aliyun.com/rubygems/
```

移除国外源：

```shell
gem sources --remove http://rubygems.org/
```

4.安装FPM工具：

​	gem install frm

5.打包：

​	fpm -s dir -f -t rpm -n python  --epoch 0 -v '2.7.8' --verbose  --description 'python2.7.8 build' --url 'www.easemob.com' --license 'BSD'  -C /tmp/installdir-python27 .