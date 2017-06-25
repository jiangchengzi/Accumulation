K近邻算法

**思想：**如果一个样本在特征空间中的k个最相似(即特征空间中最邻近)的样本中的大多数属于某一个类别，则该样本也属于这个类别

**图例：**

![b151f8198618367ad0414a912d738bd4b21ce5b5](/Users/dengqiaoling/博客/机器学习/b151f8198618367ad0414a912d738bd4b21ce5b5.jpg)

俗语说，物以类聚，人以群分，如果K=3绿色的这个待分类点属于红色的三角形一类,如果K=5,判定绿色的这个待分类点属于蓝色的正方形一类

**特点：**

- KNN算法中，所选择的邻居都是已经正确分类的对象
- 是一种 lazy-learning 算法,分类器不需要使用训练集进行训练，训练时间复杂度为0

**三个基本要素：**

- K 值的选择
  -  K 值较大，优点是可以减少学习的估计误差，但缺点是学习的近似误差增大，这时与输入实例较远的训练实例也会对预测起作用，使预测发生错误
  - K 值一般选择一个较小的数值，通常采用交叉验证的方法来选择最优的 K 值。随着训练实例数目趋向于无穷和 K=1 时，误差率不会超过贝叶斯误差率的2倍，如果K也趋向于无穷，则误差率趋向于贝叶斯误差率
- 距离度量
- 分类决策规则

default            192.168.1.1        UGSc           24        5     en0

127                127.0.0.1          UCS             0        0     lo0

127.0.0.1          127.0.0.1          UH             41   384629     lo0

169.254            link#4             UCS             0        0     en0

192.168.1          link#4             UCS             1        0     en0

192.168.1.1/32     link#4             UCS             1        0     en0

192.168.1.1        b8:f8:83:e4:14:10  UHLWIir        27       81     en0   1180

192.168.1.103      link#4             UHLWI           0       37     en0

192.168.1.124/32   link#4             UCS             0        0     en0

224.0.0/4          link#4             UmCS            2        0     en0

224.0.0.251        1:0:5e:0:0:fb      UHmLWI          0        0     en0

239.255.255.250    1:0:5e:7f:ff:fa    UHmLWI          0       12     en0

255.255.255.255/32 link#4             UCS             0        0     en0



default            192.168.1.1        UGSc           23        5     en0

127                127.0.0.1          UCS             0        0     lo0

127.0.0.1          127.0.0.1          UH             40   385141     lo0

169.254            link#4             UCS             0        0     en0

**192.168.0/23       192.168.217.1      UGSc            0        0   utun2**

192.168.1          link#4             UCS             1        0     en0

192.168.1.1/32     link#4             UCS             1        0     en0

192.168.1.1        b8:f8:83:e4:14:10  UHLWIir        24      103     en0   1194

192.168.1.103      link#4             UHLWI           0       37     en0

192.168.1.124/32   link#4             UCS             0        0     en0

**192.168.217        192.168.217.13     UGSc            1        0   utun2**

**192.168.217.13     192.168.217.13     UH              1        0   utun2**

224.0.0/4          link#4             UmCS            2        0     en0

224.0.0.251        1:0:5e:0:0:fb      UHmLWI          0        0     en0

239.255.255.250    1:0:5e:7f:ff:fa    UHmLWI          0       16     en0

255.255.255.255/32 link#4             UCS             0        0     en0



 {"secure":true,"senderMail":"13240063391@163.com","senderPassword":"vb1990522","mailServer":"smtp.exmail.qq.com:465","service_mail":"smtp.exmail.qq.com:465","code":"DtbnuPxU9ucs6dPjDRBP"} &{true 13240063391@163.com vb1990522 smtp.exmail.qq.com:465 smtp.exmail.qq.com:465 false DtbnuPxU9ucs6dPjDRBP}
`{"secure":true,"senderMail":"13240063391@163.com","senderPassword":"vb1990522","mailServer":"smtp.exmail.qq.com:465","service_mail":"smtp.exmail.qq.com:465","code":"DtbnuPxU9ucs6dPjDRBP"}`, `2017-05-31 00:50:02.15258253 +0800 CST`, `mail`



- `{"secure":true,"senderMail":"13240063391@163.com","senderPassword":"vb1990522","mailServer":"smtp.exmail.qq.com:465","service_mail":"smtp.exmail.qq.com:465","code":"ndKUbmKkVCQ8cAHbMzbc"}`, `2017-05-31 01:00:12.487381169 +0800 CST`, `mail`
  2017/05



