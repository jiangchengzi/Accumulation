### calamari Centos7.2安装详细步骤

centos上没有现成的calamari安装包，需要手动编译安装，下面记录详细的安装过程

##### 获取calamari 相关源码

```shell
$ git clone https://github.com/ceph/calamari.git
$ git clone https://github.com/ceph/calamari-clients.git
$ git clone https://github.com/ceph/Diamond.git --branch=calamari
```

calamari-clients现在已经更新了名字：romana,有兴趣的朋友可以去下载最新的版本

##### 编译calamari server的rpm包

```Shell
$ yum install gcc gcc-c++ postgresql-libs python-virtualenv
$ yum install postgresql-devel httpd checkpolicy
$ yum install selinux-policy-devel selinux-policy-doc selinux-policy-mls
$ yum -y install rpm-build
$ cd calamari 
$ yum remove prelink //避免安装时出现cpio Dismatch 错误
$ ./build-rpm.sh
```

##### 编译diamond的rpm包

```Shell
$ cd Diamond/
$ make rpm 
```

##### 编译calamari-client的rpm包

```shell
$ yum install npm ruby rubygems
$ npm install -g grunt grunt-cli bower grunt-contrib-compass
$ gem update --sytem && gem install compass
$ cd calamari-clients
$ make build-real
$ make dist 
```

​	至此，这一步会在上层文件夹下生成tar包，将其进行解压，做以下操作

```Shell
$ cd ..
$ mkdir -p opt/calamari/webapp
$ cd calamari-clients
$ for dir in manage admin login dashboard
>do
>mkdir -p ../opt/calamari/webapp/content/"$dir"
>cp -pr "$dir"/dist/* ../opt/calamari/webapp/content/"$dir"/
>done
```

​	*重新制作压缩包*



Salt-minion:

vi /etc/salt/minion

master: ceph-server







### calamari centos tenxcloud

```shell
$ docker pull kairen/docker-calamari-server:1.3.1
$ docker run -d -p 80:80 -p 4505:4505 -p 4506:4506 -p 2003:2003 -p 2004:2004 -ti --name calamari-server kairen/docker-calamari-server:1.3.1 -d
```







```
docker run -d -p 8089:80 -p 4505:4505 -p 4506:4506 -ti --name calamari-server kairen/docker-calamari-server:1.3.1 -d
```

```
sudo supervisorctl restart cthulhu
sudo service apache2 restart
```

calamari-ctl add_user admin --password admin --email admin@ceph.example.com
calamari-ctl assign_role admin --role superuser

setuptools手动安装

msgpack-python手动安装,可以考虑用pip

pip install pyzmq-16.0.2-cp27-cp27mu-manylinux1_x86_64.whl

vi /etc/init.d/salt-minion

chkconfig salt-minion on



————salt-minion启动脚本:

```
#!/bin/sh
#
# Salt minion
###################################
# LSB header
### BEGIN INIT INFO
# Provides:          salt-minion
# Required-Start:    $all
# Required-Stop:
# Default-Start:     2 3 4 5
# Default-Stop:      0 1 6
# Short-Description: Salt minion daemon
# Description:       This is the Salt minion daemon that can be controlled by the
#                    Salt master.
### END INIT INFO

# chkconfig header
# chkconfig: 345 97 04
# description:  This is the Salt minion daemon that can be controlled by the Salt master.
#
# processname: /usr/bin/salt-minion

DEBIAN_VERSION=/etc/debian_version
SUSE_RELEASE=/etc/SuSE-release
# Source function library.
if [ -f $DEBIAN_VERSION ]; then
   break
elif [ -f $SUSE_RELEASE -a -r /etc/rc.status ]; then
    . /etc/rc.status
else
    . /etc/rc.d/init.d/functions
fi
# Default values (can be overridden below)
SALTMINION=/root/calamari-agent/bin/salt-minion
PYTHON=/root/calamari-agent/bin/python
MINION_ARGS=""
if [ -f /etc/default/salt ]; then
    . /etc/default/salt
fi
SERVICE=salt-minion
PROCESS=salt-minion
RETVAL=0
start() {
    echo -n $"Starting salt-minion daemon: "
    if [ -f $SUSE_RELEASE ]; then
        startproc -f -p /var/run/$SERVICE.pid $SALTMINION -d $MINION_ARGS
        rc_status -v
    elif [ -e $DEBIAN_VERSION ]; then
        if [ -f $LOCKFILE ]; then
            echo -n "already started, lock file found"
            RETVAL=1
        elif $PYTHON $SALTMINION -d $MINION_ARGS >& /dev/null; then
            echo -n "OK"
            RETVAL=0
        fi
    else
        if [[ ! -z "$(pidofproc -p /var/run/$SERVICE.pid $PROCESS)" ]]; then
            RETVAL=$?
            echo -n "already running"
        else
            daemon --check $SERVICE $SALTMINION -d $MINION_ARGS
        fi
    fi
    RETVAL=$?
    echo
    return $RETVAL
}
stop() {
    echo -n $"Stopping salt-minion daemon: "
    if [ -f $SUSE_RELEASE ]; then
        killproc -TERM $SALTMINION
        rc_status -v
        RETVAL=$?
    elif [ -f $DEBIAN_VERSION ]; then
        # Added this since Debian's start-stop-daemon doesn't support spawned processes
        if ps -ef | grep "$PYTHON $SALTMINION" | grep -v grep | awk '{print $2}' | xargs kill &> /dev/null; then
            echo -n "OK"
            RETVAL=0
        else
            echo -n "Daemon is not started"
            RETVAL=1
        fi
    else
        killproc $PROCESS
        RETVAL=$?
    fi
    echo
}
restart() {
   stop
   start
}
# See how we were called.
case "$1" in
    start|stop|restart)
        $1
        ;;
    status)
        if [ -f $SUSE_RELEASE ]; then
            echo -n "Checking for service salt-minion "
            checkproc $SALTMINION
            rc_status -v
        elif [ -f $DEBIAN_VERSION ]; then
            if [ -f $LOCKFILE ]; then
                RETVAL=0
                echo "salt-minion is running."
            else
                RETVAL=1
                echo "salt-minion is stopped."
            fi
        else
            status $PROCESS
            RETVAL=$?
        fi
        ;;
    condrestart)
        [ -f $LOCKFILE ] && restart || :
        ;;
    reload)
        echo "can't reload configuration, you have to restart it"
        RETVAL=$?
        ;;
    *)
        echo $"Usage: $0 {start|stop|status|restart|condrestart|reload}"
        exit 1
        ;;
esac
exit $RETVAL
```







首先安装pip

再利用pip 安装

pip install pyzmq-16.0.2-cp27-cp27mu-manylinux1_x86_64.whl

