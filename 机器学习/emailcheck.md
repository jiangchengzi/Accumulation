```
// verify email
// @router /email/verification [post]
func (c *SPIController) VerifyEmail() {
       method:="VerifyEmail"
       body:= struct {
              Email string `json:"email"`
       }{}
       if err := json.Unmarshal(c.Ctx.Input.RequestBody, &body); err != nil {
              c.ErrorBadRequest("invalid request", nil)
              return
       }
       if !c.IsUserSuperAdmin(){
              c.ErrorForbidden("email","operation forbidden","you have not pemission")
              return
       }
       cf:=&models.Configs{}
       mcd,err:=cf.GetMailConfig()
       if err!=nil{
              glog.Errorf("error in Getting information from config model")
              c.ErrorInternalServerError(err)
              return
       }
       if mcd.SenderMail!=body.Email{
              glog.Errorf("the mail server which will be verified doesn't exist in cluster")
              c.ErrorInternalServerError(errors.New("the mail server which will be verified doesn't exist in cluster"))
              return
       }
       mcd.Verification=false
       mcd.Code=misc.GenSelfCheckCode()
       configDetail,err:=json.Marshal(mcd)
       if err!=nil{
              glog.Errorln(method,err)
              c.ErrorInternalServerError(err)
              return
       }
       cf.ConfigDetail=string(configDetail)
       glog.Infoln(cf.ConfigDetail,mcd)
       err=cf.UpdateMailConfig()
       if err!=nil{
              glog.Errorln(method,err)
              c.ErrorInternalServerError(err)
              return
       }
       result:=struct {
              Addr string `json:"addr"`
              Code string `json:"code"`
       }{
              Addr:mcd.SenderMail,
              Code:mcd.Code,
       }
       c.ResponseSuccess(map[string]interface{}{
              "email": result,
       })
}

// update the verification status when user chick on the link we provided
//  @router /email/join [post]
func (c *SPIController) JoinMail(){
       method := "JoinMail"
       body := struct {
              Code string `json:"code"`
       }{}
       if err := json.Unmarshal(c.Ctx.Input.RequestBody, &body); err != nil {
              glog.Infof("%s unmarshal failed. error: %s\n", method, err)
              c.ErrorBadRequest("invalid request", nil)
              return
       }
       cf:=models.Configs{}
       mcd,err:=cf.GetMailConfig()
       if err!=nil{
              glog.Errorf("error in Getting information from config model")
              c.ErrorInternalServerError(err)
              return
       }
       if mcd.Code !=body.Code{
              glog.Errorln(method,"code not found")
              c.ErrorNotFound("mail code not found","mail")
              return
       }
       mcd.Verification=true
       configDetail,err:=json.Marshal(mcd)
       if err!=nil{
              glog.Errorln(method,err)
              c.ErrorInternalServerError(err)
              return
       }
       cf.ConfigDetail=string(configDetail)
       err=cf.UpdateMailConfig()
       if err!=nil{
              glog.Errorln(method,err)
              c.ErrorInternalServerError(err)
              return
       }
       c.ResponseSuccess("")
}

// check the mail verification status
// @router /email/status [get]
func (c *SPIController) CheckEmailStatus(){
       cf:=models.Configs{}
       status:=cf.CheckMailVerification()
       result:= struct {
              Status bool `json:"emailStatus"`
       }{
              Status:status,
       }
       c.ResponseSuccess(result)
}
```
```
type MailConfigDetail struct {
       Secure       bool `json:"secure"`
       SenderMail   string `json:"senderMail"`
       SenderPassword string `json:"senderPassword"`
       MailServer     string `json:"mailServer"`
       Service_mail    string `json:"service_mail"`
       Verification    bool `json:"verification"`
       Code            string `json:"code"`

}
```
```
func (t *Configs) GetMailConfig()(mcd *MailConfigDetail,err error){
       o:=orm.NewOrm()
       if err=o.QueryTable(t.TableName()).Filter("config_type","mail").One(t,"config_detail");err!=nil{
              return nil,err
       }
       mcd=&MailConfigDetail{}
       if err:=json.Unmarshal([]byte(t.ConfigDetail),mcd);err!=nil{
              glog.Infoln(t.ConfigDetail,err)
              return nil,err
       }
       return mcd,err
}
func (t *Configs) UpdateMailConfig() error{
       o:=orm.NewOrm()
       t.CreateTime=time.Now()
       _,err:=o.QueryTable(t.TableName()).Filter("config_type","mail").Update(orm.Params{
              "config_detail": t.ConfigDetail,
              "create_time":t.CreateTime,
       })
       return err
}
func (t *Configs) CheckMailVerification() bool{
       mcd,err:=t.GetMailConfig()
       if err!=nil{
              glog.Fatalf("error in Getting information from config model")
              return false
       }
       return mcd.Verification
}
```