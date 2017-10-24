package controllers

import (
	_ "onlineshop/models"

	"fmt"
	"onlineshop/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

// 获取用户信息
func (c *MainController) UserInfo() {
	defer c.ServeJSON()
	user_id, _ := c.GetInt(":user_id")
	user := models.User{Id: user_id}
	o := orm.NewOrm()
	o.Read(&user)

	c.Data["json"] = map[string]interface{}{"ok": true,
		"data": map[string]interface{}{"UserId": user_id, "username": user.Username, "passowrd": user.Password}}
}

//用户注册
func (this *MainController) UserRegister() {
	defer this.ServeJSON()
	username := this.GetString("Username")
	password := this.GetString("Password")
	o := orm.NewOrm()
	user := models.User{Username: username, Password: password}
	//判断该用户是否存在
	if o.QueryTable("user").Filter("Username", username).Exist() {
		this.Data["json"] = map[string]interface{}{"ok": false, "message": "Username: " + username + ", 该用户已经存在"}
		return
	} else {
		id, err := o.Insert(&user)
		if err != nil {
			//debug使用
			fmt.Println("err: ", err.Error())
			return
		} else {
			//debug使用
			fmt.Println("id: ", id)
		}
		this.Data["json"] = map[string]interface{}{
			"ok": true,
			"data": map[string]interface{}{
				"data":     "恭喜你注册成功",
				"username": "用户名为:" + username,
				"password": "账号密码为:" + password,
				"用户Id":     id,
			}}
	}
}
