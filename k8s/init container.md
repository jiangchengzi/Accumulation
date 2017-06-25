##### 什么是init container

Init container 存在于pod之中，在其他常规容器启动之前启动，一个pod可以有一个或多个init container，除了比常规容器启动顺序靠前外，它还必须遵循:	

- init container总是需要成功完成
- 如果有多个init container的话，它们并不是并发执行，而是按照一定顺序

```
root@ceph-node2:~# rbd create --size 512 rbd/ima1
root@ceph-node2:~# rbd map rbd/ima1
/dev/rbd5
root@ceph-node2:~# mkfs.ext4 /dev/rbd5
root@ceph-node2:~# mount /dev/rbd5 ~/testclone
root@ceph-node2:~# echo 'edit1' >> ~/testclone/clone
root@ceph-node2:~# rbd snap create rbd/ima1@snap1
root@ceph-node2:~# rbd snap protect rbd/ima1@snap1
root@ceph-node2:~# rbd clone rbd/ima1@snap1 rbd/ima2
root@ceph-node2:~# rbd map rbd/ima2
root@ceph-node2:~# mount /dev/rbd6 ~/testclone2
root@ceph-node2:~# cat ~/testclone2/clone ##为空
```

RUN systemctl restart  salt-master

RUN systemctl restart  diamond

RUN systemctl enable postgresql-9.5.service

RUN systemctl start postgresql-9.5.service





RUN systemctl start postgresql-9.5

RUN /usr/pgsql-9.5/bin/postgresql95-setup initdb

RUN systemctl restart  salt-master

RUN systemctl restart  diamond



rbd create --size 512 rbd/rbd0001

rbd map rbd/rbd0001

mkfs.ext4 /dev/rbd3 &> /dev/null

mkdir /mnt/test0001

mount /dev/rbd3 /mnt/test0001

echo 'edit1' >> /mnt/test0001/edit

cat /mnt/test0001/edit

fsfreeze -f /mnt/test0001

rbd snap create rbd/rbd0001@snap1

rbd snap protect rbd/rbd0001@snap1

rbd clone rbd/rbd0001@snap1 rbd/rbd0002

mkdir /mnt/test0002

rbd map rbd/rbd0002

mount /dev/rbd4 /mnt/test0002

cat /mnt/test0002/edit



'598f9711-ecb4-43bb-9982-6f696b397d13', 'rbd', '{\"name\":\"\",\"url\":\"http://192.168.1.87:8011\",\"agent\":{\"user\":\"system\",\"password\":\"31e120b3-512a-4e3b-910c-85c747fb1ec2\"},\"config\":{\"monitors\":[\"192.168.1.86:6789\",\"192.168.1.87:6789\",\"192.168.1.88:6789\"],\"pool\":\"tenx-k8s\",\"user\":\"admin\",\"keyring\":\"/etc/ceph/ceph.client.admin.keyring\",\"fsType\":\"\"}}', '2017-05-26 15:34:24', ''





{"name":"","url":"http://192.168.1.87:8011","config":{"monitors":["192.168.1.86:6789","192.168.1.87:6789","192.168.1.88:6789"],"pool":"tenx-pool","user":"admin","keyring":"/etc/ceph/ceph.client.admin.keyring","fsType":""},"agent":{"user":"system","password":"31e120b3-512a-4e3b-910c-85c747fb1ec2"}}



{"name":"","url":"http://192.168.1.87:8011","agent":{"string":"","password":"31e120b3-512a-4e3b-910c-85c747fb1ec2"},"config":{"monitors":["192.168.1.86:6789","192.168.1.87:6789","192.168.1.88:6789"],"pool":"tenx-pool","user":"admin","keyring":"/etc/ceph/ceph.client.admin.keyring"},"calamariUrl":"120"}

{"name":"","url":"http://192.168.1.87:8011","agent":{"user":"system","password":"31e120b3-512a-4e3b-910c-85c747fb1ec2"},"config":{"monitors":["192.168.1.86:6789","192.168.1.87:6789","192.168.1.88:6789"],"pool":"tenx-k8s","user":"admin","keyring":"/etc/ceph/ceph.client.admin.keyring","fsType":""}}

{"name":"","url":"http://192.168.1.87:8011","agent":{"user":"","password":"31e120b3-512a-4e3b-910c-85c747fb1ec2"},"config":{"monitors":["192.168.1.86:6789","192.168.1.87:6789","192.168.1.88:6789"],"pool":"tenx-pool","user":"admin","keyring":"/etc/ceph/ceph.client.admin.keyring"},"calamariUrl":"120"}





{"name":"","url":"http://192.168.1.87:8011","agent":{"user":"","password":"31e120b3-512a-4e3b-910c-85c747fb1ec2"},"config":{"monitors":["192.168.1.86:6789","192.168.1.87:6789","192.168.1.88:6789"],"pool":"tenx-pool","user":"admin","keyring":"/etc/ceph/ceph.client.admin.keyring"},"calamariUrl":"120"}







{"name":"","url":"http://192.168.1.87:8011","agent":{"user":"","password":"31e120b3-512a-4e3b-910c-85c747fb1ec2"},"config":{"monitors":["192.168.1.86:6789","192.168.1.87:6789","192.168.1.88:6789"],"pool":"tenx-pool","user":"admin","keyring":"/etc/ceph/ceph.client.admin.keyring"},"calamariUrl":"120"}
I0526 16:14:15.169644   88974 storage.go:413] { http://192.168.1.87:8011 { 31e120b3-512a-4e3b-910c-85c747fb1ec2} {[192.168.1.86:6789 192.168.1.87:6789 192.168.1.88:6789] tenx-pool admin /etc/ceph/ceph.client.admin.keyring } 120}
I0526 16:14:15.169682   88974 storage.go:422] {"name":"","url":"http://192.168.1.87:8011","agent":{"user":"","password":"31e120b3-512a-4e3b-910c-85c747fb1ec2"},"config":{"monitors":["192.168.1.86:6789","192.168.1.87:6789","192.168.1.88:6789"],"pool":"tenx-pool","user":"admin","keyring":"/etc/ceph/ceph.client.admin.keyring"},"calamariUrl":"120"}
[ORM] - 2017-05-26 16:14:15 - [Queries/default] - [  OK /     db.Exec /     5.2ms] - [UPDATE `tenx_configs` T0 SET T0.`config_detail` = ?, T0.`create_time` = ? WHERE T0.`config_type` = ? ] - `{"name":"","url":"http://192.168.1.87:8011","agent":{"user":"","password":"31e120b3-512a-4e3b-910c-85c747fb1ec2"},"config":{"monitors":["192.168.1.86:6789","192.168.1.87:6789","192.168.1.88:6789"],"pool":"tenx-pool","user":"admin","keyring":"/etc/ceph/ceph.client.admin.keyring"},"calamariUrl":"120"}`, `2017-05-26 16:14:15.169688109 +0800 CST`, `rbd`











 {{"name":"","url":"http://192.168.1.87:8011","agent":{"user":"","password":"31e120b3-512a-4e3b-910c-85c747fb1ec2"},"config":{"monitors":["192.168.1.86:6789","192.168.1.87:6789","192.168.1.88:6789"],"pool":"tenx-pool","user":"admin","keyring":"/etc/ceph/ceph.client.admin.keyring"},"calamariUrl":"120"}}
I0526 16:18:02.834818   89041 storage.go:406] {"name":"","url":"http://192.168.1.87:8011","agent":{"user":"","password":"31e120b3-512a-4e3b-910c-85c747fb1ec2"},"config":{"monitors":["192.168.1.86:6789","192.168.1.87:6789","192.168.1.88:6789"],"pool":"tenx-pool","user":"admin","keyring":"/etc/ceph/ceph.client.admin.keyring"},"calamariUrl":"120"}
I0526 16:18:02.834869   89041 storage.go:413] { http://192.168.1.87:8011 { 31e120b3-512a-4e3b-910c-85c747fb1ec2} {[192.168.1.86:6789 192.168.1.87:6789 192.168.1.88:6789] tenx-pool admin /etc/ceph/ceph.client.admin.keyring } 120}
I0526 16:18:02.834901   89041 storage.go:422] {"name":"","url":"http://192.168.1.87:8011","agent":{"user":"","password":"31e120b3-512a-4e3b-910c-85c747fb1ec2"},"config":{"monitors":["192.168.1.86:6789","192.168.1.87:6789","192.168.1.88:6789"],"pool":"tenx-pool","user":"admin","keyring":"/etc/ceph/ceph.client.admin.keyring"},"calamariUrl":"120"}
[ORM] - 2017-05-26 16:18:02 - [Queries/default] - [  OK /     db.Exec /     6.0ms] - [UPDATE `tenx_configs` T0 SET T0.`config_detail` = ?, T0.`create_time` = ? WHERE T0.`config_type` = ? ] - `{"name":"","url":"http://192.168.1.87:8011","agent":{"user":"","password":"31e120b3-512a-4e3b-910c-85c747fb1ec2"},"config":{"monitors":["192.168.1.86:6789","192.168.1.87:6789","192.168.1.88:6789"],"pool":"tenx-pool","user":"admin","keyring":"/etc/ceph/ceph.client.admin.keyring"},"calamariUrl":"120"}`, `2017-05-26 16:18:02.834908574 +0800 CST`, `rbd`















{{"name":"","url":"http://192.168.1.87:8011","agent":{"user":"","password":"31e120b3-512a-4e3b-910c-85c747fb1ec2"},"config":{"monitors":["192.168.1.86:6789","192.168.1.87:6789","192.168.1.88:6789"],"pool":"tenx-pool","user":"admin","keyring":"/etc/ceph/ceph.client.admin.keyring","fsType":""},"calamariUrl":"120"}}
I0526 16:19:49.502822   89041 storage.go:406] {"name":"","url":"http://192.168.1.87:8011","agent":{"user":"","password":"31e120b3-512a-4e3b-910c-85c747fb1ec2"},"config":{"monitors":["192.168.1.86:6789","192.168.1.87:6789","192.168.1.88:6789"],"pool":"tenx-pool","user":"admin","keyring":"/etc/ceph/ceph.client.admin.keyring","fsType":""},"calamariUrl":"120"}
I0526 16:19:49.502855   89041 storage.go:413] { http://192.168.1.87:8011 { 31e120b3-512a-4e3b-910c-85c747fb1ec2} {[192.168.1.86:6789 192.168.1.87:6789 192.168.1.88:6789] tenx-pool admin /etc/ceph/ceph.client.admin.keyring } 120}
I0526 16:19:49.502873   89041 storage.go:422] {"name":"","url":"http://192.168.1.87:8011","agent":{"user":"","password":"31e120b3-512a-4e3b-910c-85c747fb1ec2"},"config":{"monitors":["192.168.1.86:6789","192.168.1.87:6789","192.168.1.88:6789"],"pool":"tenx-pool","user":"admin","keyring":"/etc/ceph/ceph.client.admin.keyring"},"calamariUrl":"120"}
[ORM] - 2017-05-26 16:19:49 - [Queries/default] - [  OK /     db.Exec /     7.5ms] - [UPDATE `tenx_configs` T0 SET T0.`create_time` = ?, T0.`config_detail` = ? WHERE T0.`config_type` = ? ] - `2017-05-26 16:19:49.502943951 +0800 CST`, `{"name":"","url":"http://192.168.1.87:8011","agent":{"user":"","password":"31e120b3-512a-4e3b-910c-85c747fb1ec2"},"config":{"monitors":["192.168.1.86:6789","192.168.1.87:6789","192.168.1.88:6789"],"pool":"tenx-pool","user":"admin","keyring":"/etc/ceph/ceph.client.admin.keyring"},"calamariUrl":"120"}`, `rbd`





{"name":"","url":"http://192.168.1.87:8011","agent":{"user":"system","password":"31e120b3-512a-4e3b-910c-85c747fb1ec2"},"config":{"monitors":["192.168.1.86:6789","192.168.1.87:6789","192.168.1.88:6789"],"pool":"tenx-pool","user":"admin","keyring":"/etc/ceph/ceph.client.admin.keyring"},"calamariUrl":"192.168.1.68:8087"}



{"name":"","url":"http://192.168.1.87:8011","agent":{"user":"system","password":"31e120b3-512a-4e3b-910c-85c747fb1ec2"},"config":{"monitors":["192.168.1.86:6789","192.168.1.87:6789","192.168.1.88:6789"],"pool":"tenx-k8s","user":"admin","keyring":"/etc/ceph/ceph.client.admin.keyring","fsType":""}}

rpm-build







| /api/v2/configs/mail put



/spi/v2/configs 







[https://portal.tenxcloud.com/alerts/invitations/join?code=4pZ6tQwCa4S6Lw6FOyHK](https://portal.tenxcloud.com/alerts/invitations/join?code=4pZ6tQwCa4S6Lw6FOyHK)





告警组无admin



