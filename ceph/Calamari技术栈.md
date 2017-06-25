![20937170_1453126140dWFl](/Users/dengqiaoling/博客/docker/20937170_1453126140dWFl.jpg)



```
//list information of  snapshot
// @router /volumes/snapshot/list [get]
func (c *Controller) ListSnapshot(){
       method:="ListSnapshot"
       snapshot:=&models.Snapshot{}
       clusterID:=c.GetString(":cluster")
       resp,err:=snapshot.List(clusterID,c.Namespace)
       glog.Infoln(resp)
       if err!=nil{
              msg:="list information of snapshot in DB error"
              glog.Errorln(method,msg,err)
              c.ErrorInternalServerError(err)
              return
       }
       //convert data to output
       respSnapshot:=make([]models.RespSnapshot,len(resp))
       for index,result:=range resp{
              detail,err:=parseRBDConfig(result.DriverDetail)
              if err!=nil{
                     glog.Errorf("%s:error happened when parse RBD config",method)
                     c.ErrorInternalServerError(err)
                     return
              }
              respSnapshot[index]=models.RespSnapshot{
                     SortID:index,
                     Snapshot:result,
                     FsType:detail.FsType,
                     Size:detail.Size,
              }
       }
       glog.Infoln(respSnapshot)
       c.ResponseSuccess(respSnapshot)
}
```