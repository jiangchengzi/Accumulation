### Node

### what is Node?

​	一个节点是Kubernetes中的一个工作节点，以前被称为minion。在不同集群里，节点可以是VM，也可以是物理机。每个节点必须具有运行pod的服务组件，并由主组件进行管理。节点上的服务包括Docker，kubelet和kube-proxy。

### Node拥有哪些属性 ？

- Address

  以下字段的使用取决于云提供商或裸机配置。 
  HostName：节点内核定义的主机名。可以通过kubelet --hostname-override参数重定义。 
  ExternalIP：通常是可外部路由的节点的IP地址（可从群集外部获得）。 
  InternalIP：通常只能在集群内进行路由的节点的IP地址。

- Condition

  条件字段描述所有运行节点的状态。 
  OutOfDisk：如果节点上没有足够的可用空间用于添加新的pod，则为True，否则False 

  Ready ：如果节点正常并准备好接受pod，则为True，如果节点不健康且不接受pod，则为False			

  ​		如果节点控制器在过去40秒内没有从节点中听到，则为未知
  MemoryPressure：如果节点没有内存压力，则为True，否则为False 
  DiskPressure：如果节点没有磁盘压力，则为True否则为False

  如果Ready状态的状态为“Unknown”或因为超过撤离超时时间而处于“False”状态，节点上的所有Pod都将被节点删除控制器所调度，撤离超时时间为传递给kube-controller-manager的一个参数。默认的撤离超时时间为五分钟。在某些情况下，当节点不可达时，apiserver无法与其上的kubelet进行通信。在重新建立与apiserver的通信之前，删除pod的决定不能传达到kubelet。同时，计划删除的pod可以继续在分区节点上运行。 
  在1.5之前的Kubernetes版本中，节点控制器将强制从apiserver中删除这些不可达的pod。但是，在1.5及更高版本中，节点控制器不会强制删除pod，直到确认它们已停止在群集中运行为止。可以看到这些可能在不可达节点上运行的pod处于“终止”或“未知”状态。如果节点永久离开集群，Kubernetes无法从底层基础架构推导出来，则集群管理员可能需要手动删除节点对象。从Kubernetes删除节点对象会导致运行在其上的所有Pod对象从apiserver中删除，释放其名称。

  # Tasks

  ## 监控、日志、调试

  ### 通过Elasticsearch and Kibana进行日志处理

  ​	在Google Compute Engine（GCE）平台上，默认的日志记录支持是Stackdriver Logging，这在Logging With Stackdriver Logging中有详细描述。 
  ​	本文介绍如何设置集群以将日志记录到Elasticsearch并使用Kibana进行查看，作为在GCE上运行时Stackdriver Logging的替代方法。请注意，Elasticsearch和Kibana无法在Google Container Engine上托管的Kubernetes群集中自动设置，因此必须手动进行部署。 
  ​	要使用Elasticsearch和Kibana进行集群日志记录，您应该使用kube-up.sh创建集群时，如下所示设置以下环境变量：

  ```
  KUBE_LOGGING_DESTINATION=elasticsearch
  ```

  ​	您还应确保KUBE_ENABLE_NODE_LOGGING = true（这是GCE平台的默认值）。 
  ​	现在，当您创建集群时，消息将指示在每个节点上运行的Fluentd日志收集守护程序将定位到Elasticsearch：		

  ```
  $ cluster/kube-up.sh
  ...
  Project: kubernetes-satnam
  Zone: us-central1-b
  ... calling kube-up
  Project: kubernetes-satnam
  Zone: us-central1-b
  +++ Staging server tars to Google Storage: gs://kubernetes-staging-e6d0e81793/devel
  +++ kubernetes-server-linux-amd64.tar.gz uploaded (sha1 = 6987c098277871b6d69623141276924ab687f89d)
  +++ kubernetes-salt.tar.gz uploaded (sha1 = bdfc83ed6b60fa9e3bff9004b542cfc643464cd0)
  Looking for already existing resources
  Starting master and configuring firewalls
  Created [https://www.googleapis.com/compute/v1/projects/kubernetes-satnam/zones/us-central1-b/disks/kubernetes-master-pd].
  NAME                 ZONE          SIZE_GB TYPE   STATUS
  kubernetes-master-pd us-central1-b 20      pd-ssd READY
  Created [https://www.googleapis.com/compute/v1/projects/kubernetes-satnam/regions/us-central1/addresses/kubernetes-master-ip].
  +++ Logging using Fluentd to elasticsearch
  ```

  ​	每个节点的Fluentd pod，Elasticsearch pod和Kibana pod都应该在集群生效后立即在kube-system命名空间中运行。

  ```
  $ kubectl get pods --namespace=kube-system
  NAME                                           READY     STATUS    RESTARTS   AGE
  elasticsearch-logging-v1-78nog                 1/1       Running   0          2h
  elasticsearch-logging-v1-nj2nb                 1/1       Running   0          2h
  fluentd-elasticsearch-kubernetes-node-5oq0     1/1       Running   0          2h
  fluentd-elasticsearch-kubernetes-node-6896     1/1       Running   0          2h
  fluentd-elasticsearch-kubernetes-node-l1ds     1/1       Running   0          2h
  fluentd-elasticsearch-kubernetes-node-lz9j     1/1       Running   0          2h
  kibana-logging-v1-bhpo8                        1/1       Running   0          2h
  kube-dns-v3-7r1l9                              3/3       Running   0          2h
  monitoring-heapster-v4-yl332                   1/1       Running   1          2h
  monitoring-influx-grafana-v1-o79xf             2/2       Running   0          2h
  ```

  ​	fluentd-elasticsearch从每个节点收集日志，并将它们发送到elasticsearch-logging的pods。这些Elasticsearch pod存储日志并通过REST API公开它们。 kibana-logging pod提供了一个用于读取存储在Elasticsearch中的日志的Web UI，并且是名为kibana-logging的服务的一部分。 
  Elasticsearch和Kibana服务都是在kube系统命名空间中，不会通过公开可访问的IP地址直接暴露。要访问它们，请按照访问集群中运行的服务的说明进行操作。 
  ​	如果您尝试访问浏览器中的elasticsearch-logging服务，您将看到一个状态页面，如下所示：![es-browser](/Users/dengqiaoling/博客/kubernetes中文文档/es-browser.png)

  Error: error connecting to the cluster;clustername: ceph;name: client.admin conf: None conffile: flags: 0None: errno EINVAL

  ​

  ​

  ​

  ​        self.cluster=cluster

  ​        self.name=name

  ​        self.flags=flags	



ceph --admin-daemon ceph-mon.node1.asok config show

