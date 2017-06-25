- ```
  sudo yum install -y yum-utils && sudo yum-config-manager --add-repo https://dl.fedoraproject.org/pub/epel/7/x86_64/ && sudo yum install --nogpgcheck -y epel-release && sudo rpm --import /etc/pki/rpm-gpg/RPM-GPG-KEY-EPEL-7 && sudo rm /etc/yum.repos.d/dl.fedoraproject.org*
  ```

  过程中出现pub/epel/7/x86_64/找不到，原因是/etc/sysconfig/network-scripts/ifcfg-enp0s3种onBoot为No，应该设置为yes

- 安装到ceph-deploy install admin node1 node2 node3...的时候，出现了no section 'ceph'的错误，原因是yum源没有设置优先级，在admin这台机子上ceph.repo源设置priority=2,并且：编辑/etc/yum/pluginconf.d/priorities.conf确保配置文件中包含下面的行

  ```
  [main]
  enabled=1
  ```

  保存并关闭该文件

- 最好能用国内的源进行ceph的安装

  ```
  CentOS:
  export CEPH_DEPLOY_REPO_URL=http://mirrors.163.com/ceph/rpm-jewel/el7
  export CEPH_DEPLOY_GPG_URL=http://mirrors.163.com/ceph/keys/release.asc

  Ubuntu:
  export CEPH_DEPLOY_REPO_URL=http://mirrors.163.com/ceph/debian-jewel
  export CEPH_DEPLOY_GPG_URL=http://mirrors.163.com/ceph/keys/release.asc
  ```

- 安装过程中需要activate osd ，如果出现ceph permission denied 则：

  ```i
  chown ceph:ceph /dev/sdd1
  ```

- 如果出现时间偏移，clock skew则需要同步一下ntp

  #### 记录在centos7上三节点安装ceph的完整过程：

  安装架构图：

  ​	![ditaa-cffd08dd3e192a5f1d724ad7930cb04200b9b425](/Users/dengqiaoling/博客/docker/ditaa-cffd08dd3e192a5f1d724ad7930cb04200b9b425.png)

- 安装虚拟机，保证每台主机都能够通过主机名ping通：

  - vim /etc/hosts

    ```shell
    192.168.1.20 admin
    192.168.1.21 node1
    192.168.1.22 node2
    192.168.1.23 node3
    ```

  - 最好采用国内的源进行安装

    ```shell
    $vim ~/.bashrc
    CentOS:
    export CEPH_DEPLOY_REPO_URL=http://mirrors.163.com/ceph/rpm-jewel/el7
    export CEPH_DEPLOY_GPG_URL=http://mirrors.163.com/ceph/keys/release.asc
    $source ~/.bashrc 
    ```

  - 添加ceph安装源

    ```shell
    [ceph]
    name=Ceph packages for $basearch
    baseurl=http://download.ceph.com/rpm-{ceph-release}/{distro}/$basearch
    enabled=1
    priority=2
    gpgcheck=1
    type=rpm-md
    gpgkey=https://download.ceph.com/keys/release.asc

    [ceph-noarch]
    name=Ceph noarch packages
    baseurl=http://download.ceph.com/rpm-{ceph-release}/{distro}/noarch
    enabled=1
    priority=2
    gpgcheck=1
    type=rpm-md
    gpgkey=https://download.ceph.com/keys/release.asc

    [ceph-source]
    name=Ceph source packages
    baseurl=http://download.ceph.com/rpm-{ceph-release}/{distro}/SRPMS
    enabled=0
    priority=2
    gpgcheck=1
    type=rpm-md
    gpgkey=https://download.ceph.com/keys/release.asc
    ```

    可以在 /etc/yum.repos.d/ 目录下新增一个 Ceph 库：创建 ceph.repo 。在下例中，需要用 Ceph 主要发布名（如 dumpling 、 emperor）替换 {ceph-release} 、用 Linux 发行版名（ el6、 rhel6 等）替换 {distro} 。你可以到 [http://download.ceph.com/rpm](http://download.ceph.com/rpm)-{ceph-release}/ 看看 Ceph 支持哪些发行版。有些 Ceph 包（如 EPEL ）必须优先于标准包，所以你必须确保设置了 priority=2 。

  - 更新软件库并安装 ceph-deploy

    ```shell
    $sudo yum update && sudo yum install ceph-deploy
    ```

- 安装前准备

  - 安装ntp

    ```shell
    $sudo yum install ntp ntpdate ntp-doc
    ```

  - 安装ssh服务器

    ```Shell
    $sudo yum install openssh-server
    ```

  - 创建部署ceph的用户

    - 在各 Ceph 节点创建新用户

    ```Shell
    $ssh houge@node1
    $sudo useradd -d /home/houge -m houge
    $sudo passwd houge
    ```

    - 确保各 Ceph 节点上新创建的用户都有 `sudo` 权限

    ```shell
    $echo "houge ALL = (root) NOPASSWD:ALL" | sudo tee /etc/sudoers.d/houge
    $sudo chmod 0440 /etc/sudoers.d/houge
    ```

  - 允许无密码ssh登录

    生成密码对的时候不要用root用户或者sudo

    ```shell
    $ssh-keygen     #直接enter
    ```

    将密码拷贝到各ceph节点

    ```Shell
    $ssh-copy-id houge@node1
    $ssh-copy-id houge@node2
    $ssh-copy-id houge@node3
    ```

  - 引导时联网

    进入 `/etc/sysconfig/network-scripts` 目录并确保 `ifcfg-{iface}` 文件中的 `ONBOOT` 设置成了 `yes`，主要是因为如果网络默认为 `off` ，那么 Ceph 集群在启动时就不能上线，直到你打开网络

  - 关闭防火墙

    ```shell
    $systemctl stop firewalld
    ```

  - SELINUX

    ```Shell
    $sudo setenforce 0
    ```

  - 禁用tty

    如果你的 Ceph 节点默认设置了 `requiretty` ，执行 `sudo visudo` 禁用它，并找到 `Defaults requiretty` 选项，把它改为 `Defaults:ceph !requiretty` 或者直接注释掉，这样 `ceph-deploy` 就可以用之前创建的用户连接了

  - 启用优先级/首选项包

    ```shell
    $sudo yum install yum-plugin-priorities
    ```

    编辑/etc/yum/pluginconf.d/priorities.conf确保配置文件中包含下面的行

    ```shell
    [main]
    enabled=1
    ```

    保存并关闭该文件

- 存储集群安装

  集群包含了三个 Monitor 和三个 OSD 守护进程

  - 创建目录

    ```shell
    $mkdir my-cluster
    $cd my-cluster
    ```

    管理节点上创建一个目录，用于保存 `ceph-deploy` 生成的配置文件和密钥对

  - 创建monitors

    ```shell
    $ceph-deploy new node1 node2 node3 
    ```

  - ​

