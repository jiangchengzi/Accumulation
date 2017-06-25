calamari 

在详述calamari安装过程之前，我先简要地介绍一下calamari这个开源项目的大概架构。

​	calamari主要是用来管理ceph集群，ceph本身也提供多种管理方式，如librados、radgw，calamari在此基础上重新封装了一层，对外提供更加丰富、简洁的restapi，以及界面。

​	calamari从大的框架上可以分为两条线：

​	ceph—(admin socket)——>Diamod——(tcp)——>carton_cache———>graphite_web

​	ceph ———(salt-minion)——>salt-master——>cthultu———>calamari_rest

​	Diamond负责收集监控数据，它支持非常多的数据类型和metrics,它除了收集Ceph本身的状态信息，它还可以收集关键的资源使用情况和性能数据，包括CPU，内存，网络，I / O负载和磁盘指标,Diamond将收集上来的数据通过carbon-cache交由Graphite来显示，Graphite不仅是一个企业级的监控工具, 还可以实时绘图。carbon-cache是Python实现的高度可扩展的事件驱动的I/O架构的后端进程，它可以有效地跟大量的客户端通信并且以较低的开销处理大量的业务量。

​	calamari的主要组件如下：

​		(master节点上)

​			calamari-client:包含四个模块，分别是manage，admin ，login，dashboard,为用户提供web界面展示

​			cthulhu：类似于一个中间层，对下通过salt-master向ceph节点上的salt-minion传达指令，完成ceph集群的信息收集、管理工作，

​			salt-master：接受cthulhu的指令，完成与salt-minion的交互

​			apache：为calamari提供		

​	通过以上介绍，大家可能也只是简要地了解了calamari的各个组件，至于其究竟ceph的一条查询指令是如何传达至ceph集群，并且通过何种方式与ceph进行交互，并实时收集ceph集群的有用信息，如osd status、mon status，以及各个节点上的iops、cpu等。下面我抛砖引玉，简要介绍下一个calamari请求的心路历程。

​	

​	calamari









service salt-master restart

service apache2 restart



​								