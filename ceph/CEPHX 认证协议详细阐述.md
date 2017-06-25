##### CEPHX 认证协议详细阐述

本文档提供了更详细的Cephx授权协议，其高级流程由Yehuda的备忘录（12/19/09）描述。因为这个备忘录讨论了调用的例程和使用的变量的细节，它代表一个快照。在创建此文档之后，代码可能会更改，并且文档不太可能在锁定步骤中更新。幸运的是，代码注释将显示协议实现方式的重大变化。





配置/etc/hosts



No zlib  

1. install zlib-devel
2. make and install python again.



安装python-virtualenv

​	yum install python-virtualenv

安装python2.7.6

​	yum install zlib-devel

​	yum install openssl openssl

​	./configure --prefix=/usr/local

​	make && make install

激活隔离环境

​	virtualenv --system-site-packages -p /usr/local/bin/python calamari-agent

​	cd  ~/calamari-agent

​	source bin/activate





pip install Jinja2==2.7.2

 pip install M2Crypto==0.21.1的过程中会报错command 'swig' failed with exit status 1

​	安装swig-3.0.4注意版本

​	swig-3.0.4.tar.gz

​	yum install pcre-devel

​	./configure

​	make && make install

​	env SWIG_FEATURES="-cpperraswarn -includeall -D\_\_\`uname -m\`\_\_ -I/usr/include/openssl" pip 

pip install  PyYAML==3.10

pip install PyZMQ==14.0.1

pip install pycrypto==2.6.1

pip install msgpack-python==0.3.0

pip install salt.0.17.5.tar.gz



配置salt

​	mkdir /etc/salt

​	vi /etc/salt/minion_id

​	echo 'master: calamari-server' > /etc/salt/minion

​	cat << EOF >> /etc/salt/minion  

​	providers:

​	  pkg: yumpkg5

​	 EOF

​	编辑salt-minion 服务脚本：

​		vi /etc/init.d/salt-minion

​		。。。

​		chmod 755 /etc/init.d/salt-minion

​		chkconfig salt-minion	

安装diamond：

​	mv /etc/diamond/diamond.conf.example /etc/diamond/diamond.conf

​	service diamond restart

复制所有的site-package下的文件到当前环境下

​	\cp -Rf  /usr/lib/python2.7/site-packages/* ~/calamari-agent/lib/python2.7/site-packages/







\- name: restart calamari-server salt-master

  sudo: true

  shell: docker exec calamari-server bash -c 'service salt-master restart'

\- name: restart calamari-server apache2

  sudo: true

  shell: docker exec calamari-server bash -c 'service apache2 restart'

\- name: chmod cthulhu.log

  sudo: true

  shell: docker exec calamari-server bash -c 'chmod 777 /var/log/calamari/cthulhu.log'

\- name: chmod calamari.log

  sudo: true

  shell: docker exec calamari-server bash -c 'chmod 777 /var/log/calamari/calamari.log'

\- name: restart PostgreSQL

  sudo: true

  shell: docker exec calamari-server bash -c 'service postgresql restart'

\- name: initialize calamari-server user 

  sudo: true

  shell: docker exec calamari-server bash -c 'calamari-ctl add_user admin --password admin --email admin@ceph.example.com'

\- name: assign calamari-server role

  sudo: true

  shell: docker exec calamari-server bash -c 'calamari-ctl assign_role admin --role superuser'









yum install -y --downloadonly  zlib-devel openssl openssl python-virtualenv swig pcre-devel --downloaddir=/tmp/rpm/