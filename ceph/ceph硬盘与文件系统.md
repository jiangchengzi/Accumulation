### ceph 硬盘与文件系统

- osd最好与文件系统放在不同硬盘上，原因如下：

  - 内核版本小于2.6的系统会将数据

    一直读取不到，并不会堵塞job的进程，因为只是在watch，



