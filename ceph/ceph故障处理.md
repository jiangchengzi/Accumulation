- ceph osd down时，重启osd

  ```
   /etc/init.d/ceph -a start osd
  ```

- ceph重启

  SELECT tenx_project.name as project_name,tenx_project.description as description,tenx_project.creation_time as creation_time,tenx_project.update_time as update_time,tenx_project_account.balance as balance,count(distinct tenx_project_user.user_id) as user,count(distinct tenx_project_resource_ref.resource_id) as cluster FROM tenx_project LEFT JOIN tenx_project_account ON tenx_project.id=tenx_project_account.project_id INNER JOIN tenx_project_user ON tenx_project.id=tenx_project_user.project_id LEFT JOIN tenx_project_resource_ref ON tenx_project_resource_ref.project_id=tenx_project.id GROUP BY project_name;



SELECT tenx_project.name as project_name,tenx_project.id as project_id,tenx_project.description as description,tenx_project.creation_time as creation_time,tenx_project.update_time as update_time,tenx_project_account.balance as balance,count(distinct tenx_project_user.user_id) as user,count(distinct tenx_project_resource_ref.resource_id) as cluster FROM tenx_project LEFT JOIN tenx_project_account ON tenx_project.id=tenx_project_account.project_id LEFT JOIN tenx_project_user ON tenx_project.id=tenx_project_user.project_id LEFT JOIN tenx_project_resource_ref ON tenx_project_resource_ref.project_id=tenx_project.id WHERE tenx_project.id IN ('PID-dL39eKNi2dgz','PID-izVY2rmiiqSR','PID-nvDTw6mKAn7D') and tenx_project_resource_ref.resource_type = 0 GROUP BY project_name]





SELECT tenx_project.name as project_name,tenx_project.id as project_id,tenx_project.description as description,tenx_project.creation_time as creation_time,tenx_project.update_time as update_time,tenx_project_account.balance as balance,count(distinct tenx_project_user.user_id) as user,count(distinct tenx_project_resource_ref.resource_id) as cluster FROM tenx_project LEFT JOIN tenx_project_account ON tenx_project.id=tenx_project_account.project_id LEFT JOIN tenx_project_user ON tenx_project.id=tenx_project_user.project_id LEFT JOIN tenx_project_resource_ref ON tenx_project_resource_ref.project_id=tenx_project.id and tenx_project_resource_ref.resource_type = 0 WHERE tenx_project.id IN ('PID-dL39eKNi2dgz','PID-izVY2rmiiqSR','PID-nvDTw6mKAn7D') GROUP BY project_name]





/var/run/ceph/rbd-clients/ceph-client.admin.117926.140020765692032.asok

# AdminSocketConfigObs::init: failed: rbd-clients

解决方法：

mkdir /var/run/ceph/rbd-clients/



目前会出现如下bug：

Ceph servers are connected to Calamari, but no Ceph cluster has been created 





升级python至2.7.6

```
wget http://www.python.org/ftp/python/2.7.6/Python-2.7.6.tgz
tar -zxvf Python-2.7.6.tgz
cd Python-2.7.6
mkdir /usr/local/python27
./configure --prefix=/usr/local/python27
make
make install
mv /usr/bin/python /usr/bin/python_old
ln -s /usr/local/python27/bin/python2.7 /usr/bin/python
```





 Salt: 0.17.5

​         Python: 2.7.5 (default, Nov  6 2016, 00:28:07)

​         Jinja2: 2.7.2

​       M2Crypto: 0.21.1

 msgpack-python: 0.4.8

   msgpack-pure: Not Installed

​       pycrypto: 2.6.1

​         PyYAML: 3.10

​          PyZMQ: 14.0.1

​            ZMQ: 4.0.3



​         Salt: 0.17.5

​         Python: 2.7.6 (default, Jun 22 2015, 17:58:13)

​         Jinja2: 2.7.2

​       M2Crypto: 0.21.1

 msgpack-python: 0.3.0

   msgpack-pure: Not Installed

​       pycrypto: 2.6.1

​         PyYAML: 3.10

​          PyZMQ: 14.0.1

​            ZMQ: 4.0.4





需要安装setuptools

const transport = _.cloneDeep(globalConfig.mail_server)

const _ = require('lodash')