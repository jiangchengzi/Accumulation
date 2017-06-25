ceph 日常维护指令

##### monitor

```
$ ceph mon dump #查看monitor映射信息
dumped monmap epoch 1                      
epoch 1
fsid 04bad187-3633-401a-a600-6c52b2a6b098    #集群ID
last_changed 0.000000						 #最近更新时间
created 0.000000						
0: 192.168.0.83:6789/0 mon.node1			 #监控主机名以及IP地址端口号
1: 192.168.0.84:6789/0 mon.node2
2: 192.168.0.85:6789/0 mon.node3
```



ceph 存储池

​	获取存储池相关参数

​		ceph osd pool get rbd {参数}

​	      example:

​			ceph osd pool get rbd pg-num

​	创建存储池：

```
ceph osd pool create {pool-name} {pg-num} [{pgp-num}] [replicated] \
        [crush-ruleset-name] [expected-num-objects]
ceph osd pool create {pool-name} {pg-num}  {pgp-num}   erasure \
        [erasure-code-profile] [crush-ruleset-name] [expected_num_objects]
```

