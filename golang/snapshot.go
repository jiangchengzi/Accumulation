package models


import (
	"time"
	"github.com/astaxie/beego/orm"
	"github.com/golang/glog"
)

type Snapshot struct {
	ID           string    `orm:"pk;column(id)"`
	Cluster      string    `orm:"size(45);column(cluster)"`
	Namespace    string    `orm:"size(255);column(namespace)"`
	Name         string    `orm:"size(100);column(name)"`
	UserID       uint32    `orm:"column(user_id)"`
	Driver       string    `orm:"size(255);column(driver)"`
	DriverDetail string    `orm:"column(driver_detail)"`
	Volume       string    `orm:"size(45);column(volume)"`
	CreateTime   time.Time `orm:"type(datetime);column(creation_time)"`
}

func (s *Snapshot) TableName() string{
	return "tenx_snapshot"
}

func (s *Snapshot) Insert(o orm.Ormer) error{
	_, err := o.Insert(s)
	if err != nil {
		glog.Errorln("Insert snapshot table failed.", err)
	}
	return err
}


func (s *Snapshot) GetSnapInfoByName(clusterID, namespace, SnapshotName string) (snapshot Snapshot,err error){
	o:=orm.NewOrm()
	snapshot=Snapshot{}
	if err=o.QueryTable(s.TableName()).
		Filter("cluster",clusterID).
		Filter("namespace",namespace).
		Filter("name",SnapshotName).
		One(&snapshot);err!=nil {
		return snapshot,err
	}
	return snapshot,nil

}
func (t *Snapshot) DeleteMulti(o orm.Ormer,clusterID, namespace string, snapshots []string) error {
	method := "DeleteMulti"
	num, err := o.QueryTable(t.TableName()).
		Filter("namespace", namespace).
		Filter("cluster", clusterID).
		Filter("name__in", snapshots).
		Delete()
	if num != int64(len(snapshots)) {
		glog.Errorf("%s volume count:%v delete count %v\n", method, len(snapshots), num)
	}
	return err
}




func (t Snapshot) List(clusterID, namespace string, fields ...string) ([]Snapshot, error) {
	o := orm.NewOrm()
	snapshot := make([]Snapshot, 0, 1)
	_, err := o.QueryTable(t.TableName()).
		Filter("cluster", clusterID).
		Filter("namespace", namespace).
		All(&snapshot, fields...)
	return snapshot, err
}

func (t Snapshot) CheckSnapExist(clusterID,namespace,SnapshotName string) bool{
	o:=orm.NewOrm()
	return o.QueryTable(t.TableName()).
		Filter("cluster",clusterID).
		Filter("namespace",namespace).
		Filter("name",SnapshotName).
		Exist()
}



