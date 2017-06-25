tensorflow于时速云上的实践







map[manager:[29854] admin:[29854]] map[manager:[434 433] advisor:[29854]]



&{{94} [{TID-JtBXFhxDDTse aaatest  1232 2017-05-22 12:25:42 +0800 CST 1 4
 2 false}

 {TID-crUhqV535dSz baremetal  1232 2017-03-29 16:01:06 +0800 CST 2 0 8 false}

 {TID-XEQybXKLnVVi hhhdd  1232 2017-05-09 17:06:15 +0800 CST 0 1 0 false}

 {TID-gSxmoERNw2tA new-team  1232 2017-05-08 11:34:27 +0800 CST 1 3 2 false}

 {TID-fqveSuMPoDJB newtest  1232 2017-04-18 15:32:48 +0800 CST 0 2 2 false}

 {TID-373w85wCCitK oolfyyy  1232 2017-03-03 14:31:08 +0800 CST 1 0 2 false} 

{TID-ZZxt5Erg6SQT prod-demo  1232 2017-02-27 16:14:00 +0800 CST 1 3 4 false}

 {TID-f2gMhmtF9nEK sit-demo  1232 2017-02-27 16:13:52 +0800 CST 1 0 3 false} 

{TID-5TnA2QagUh6x test-eva  1232 2017-03-27 15:40:39 +0800 CST 2 1 7 false}

 {TID-U29j7jphSa7n test-syy-1  1232 2017-05-11 17:08:59 +0800 CST 2 2 1 false}]}





I0614 10:57:40.166394   75702 base.go:287] {0 0   false} api-server/controllers/volume.(*VolumeController).CreateSnapsh
ot
E0614 10:57:40.166431   75702 base.go:308] controllers/BaseController.GetAuditInfo Failed to get audit info for api api
-server/controllers/volume.(*VolumeController).CreateSnapshot
E0614 10:57:40.166444   75702 base.go:309] controllers/BaseController.GetAuditInfo admin POST /api/v2/clusters/CID-fe23
111d77cb/volumes/ssss/snapshot







Running transaction

  Erasing    : calamari-clients-1.2.2-32_g931ee58.el7.centos.x86_64                                                            1/2 

rm: cannot remove ‘/etc/httpd/conf.d/calamari.conf’: No such file or directory

rm: cannot remove ‘/etc/httpd/conf.d/wsgi.conf’: No such file or directory

mv: cannot stat ‘/etc/httpd/conf.d/welcome.conf.orig’: No such file or directory

Redirecting to /bin/systemctl stop  supervisord.service

Redirecting to /bin/systemctl start  supervisord.service

  Erasing    : calamari-server-1.5.2-13_g768f37d.el7.centos.x86_64                                                             2/2 

setsebool:  SELinux is disabled.

setsebool:  SELinux is disabled.

  Verifying  : calamari-server-1.5.2-13_g768f37d.el7.centos.x86_64                                                             1/2 

  Verifying  : calamari-clients-1.2.2-32_g931ee58.el7.centos.x86_64   



- [0f6642d](https://github.com/saltstack/salt/commit/0f6642d533703822f9ebb45abf49234adf837356)





ABCabc123







 python setup.py install

  203  cp /root/salt/conf/minion /etc/salt/

  204  vi /etc/salt/minion

  205  salt-minion --version

  206  /etc/init.d/salt-minion restart

  207  chkconfig salt-minion on

  208  service salt-minion start

  209  vi /etc/salt/minion

  210  service salt-minion start

  211  service salt-minion restart

  212  rm -rf /etc/salt/pki

  213  rm -f  /etc/salt/minion.d/calamari.conf 

  214  rm /etc/salt/minion

  215  service salt-minion restart

  216  vi /etc/salt/minion

  217  service salt-minion restart

  218  /etc/init.d/ceph -a  restart osd

  219  mkdir /var/run/ceph/rbd-clients/

  220  vi /etc/ceph/ceph.conf

  221  /etc/init.d/ceph -a restart

  222  service diamond restart

  223  salt --versions-report

  224  pip upgrade python==2.7.6

  225  easy_install pip

  226  python -v

  227  salt --versions-report

  228  history

  229  cd  salt-0.17.3-py2.7.egg-info

  230  cd /usr/lib/python2.7/site-packages/

  231  ls

  232  rm  salt-0.17.3-py2.7.egg-info

  233  yum install python-devel

  234  salt --versions-report

  235  pip install pyzmq==14.0.1

  236  ls

  237  cd /root/

  238  ls

  239  mkdir diamond

  240  cd diamond/

  241  git clone https://github.com/ceph/Diamond.git --branch=calamari

  242  ls

  243  cd Diamond/

  244  ls

  245  ./version.sh 

  246  make rpm

  247  yum -y install rpm-build

  248  pip

  249  ls

  250  make rpm

  251  yum remove diamond

  252  yum install -y build/bdist.linux-x86_64/rpm/RPMS/noarch/diamond-3.4.67-0.noarch.rpm

  253  cd build/bdist.linux-x86_64/rpm/RPMS/noarch/

  254  ls

  255  cd -

  256  cd dist/

  257  ls

  258  yum install diamond-3.4.67-0.noarch.rpm 

  259  mv /etc/diamond/diamond.conf.example /etc/diamond/diamond.conf

  260  service diamond restart

  261  scp diamond-3.4.67-0.noarch.rpm root@node2:/root/

  262  ssh root@node2

  263  service salt-minion restart

  264  vi /etc/salt/minion

  265  salt -l debug

  266  salt-minion -l debug

  267  rm /etc/salt/pki/minion/minion_master.pub

  268  salt-minion -l debug

  269  ceph -s

  270  cd ..

  271  ls

  272  cd rpm/

  273  ls

  274  cd ..

  275  ls

  276  salt

  277  ps -ef|grep salt

  278  ps -ef|grep salt*

  279  service salt-minion start

  280  ps -ef|grep salt*

  281  history

  282  tail -f /var/log/salt/minion 

  283  service salt-minion restart

  284  tail -f /var/log/salt/minion 

  285  rm -rf /etc/salt/pki/minion/minion_master.pub

  286  service salt-minion restart

  287  tail -f /var/log/salt/minion 

  288  ps -ef|grep salt*

  289  vi /etc/init.d/salt-minion 

  290  tail -f /var/log/salt/minion 

  291  diamond --version

  292  vi /etc/diamond/diamond.conf

  293  tail -100 /var/log/diamond/archive.log

  294  tail -100 /var/log/diamond/diamond.log

  295  service diamond restart

  296  systemctl daemon-reload

  297  tail -100 /var/log/diamond/diamond.log

  298  tail -100 /var/log/diamond/archive.log

  299  tail -f /var/log/salt/minion 

  300  history

  301  tail -100 /var/log/diamond/archive.log

  302  tail -100 /var/log/diamond/diamond.log

  303  service diamond restart

  304  tail -100 /var/log/diamond/diamond.log

  305  tail -100 /var/log/diamond/archive.log

  306  tail -100 /var/log/diamond/diamond.log

  307  tail -f /var/log/salt/minion 

  308  vi /etc/salt/minion

  309  history







Jinja2

M2Crypto

msgpack-python

pycrypto

PyYAML

pyzmq >= 2.1.9

markupsafe



yum install -y python-jinja2.noarch --downloadonly --downloaddir=/root/





python-devel

Openssl-devel



 201  ceph -s

  202  ls

  203  yum install -y diamond-3.4.67-0.noarch.rpm

  204  mv /etc/diamond/diamond.conf.example /etc/diamond/diamond.conf

  205  service diamond restart

  206  vi /etc/hosts

  207  tar -xzf rpm.tar.gz 

  208  tar -xzf salt-0.17.5.tar.gz 

  209  yum localinstall -y rpm-packages/*.rpm

  210  python salt-0.17.5/setup.py install 

  211  cd salt-0.17.5/

  212  python setup.py install

  213  ls

  214  find ./ -name salt-minion

  215  vi ./scripts/salt-minion 

  216  vi ./build/scripts-2.7/salt-minion 

  217  ls

  218  mkdir /etc/salt

  219  chmod 755 /etc/salt/

  220  echo 'master: ceph-server' > /etc/salt/minion

  221  echo 'node1' > /etc/salt/minion_id

  222  salt-minion -d -l

  223  salt-minion -d 

  224  tail -f /var/log/salt/minion 

  225  tail -100 /var/log/diamond/diamond.log

  226  vi /etc/diamond/diamond.conf

  227  getenforce

  228  iptables -L

  229  systemctl stop firewalld

  230  tail -f /var/log/salt/minion 

  231  vi /etc/ceph/ceph.conf

  232  ceph -s

  233  tail -f /var/log/salt/minion 

  234   ceph osd lspools

  235  vi /etc/ceph/ceph.conf

  236  /etc/init.d/ceph restart

  237  ceph -s

  238  /etc/init.d/ceph -a start osd

  239  vi /etc/ceph/ceph.conf

  240  /etc/init.d/ceph -a start osd

  241  ceph -s

  242  vi /etc/ceph/ceph.conf

  243  ceph-deploy

  244  tail -f /var/log/salt/minion 

  245  ls

  246  cd ..

  247  ls

  248  history