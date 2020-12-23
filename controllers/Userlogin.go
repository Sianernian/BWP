package controllers

import (
	"BWP/db_mysql"
	"BWP/models/user"
	"fmt"
	"github.com/astaxie/beego"
	"log"
)

type UserloginController struct {
	beego.Controller
}

func (u *UserloginController) Post() {
	//1.存储用户的信息
	user2 := user.User{}
	//2.解析前端的用户信息
	err := u.ParseForm(&user2)
	if  err != nil {
		log.Fatal("解析前端数据错误")
	}
	fmt.Println(user2)
	//3.将用户信息存入数据库中
	//3.1打开数据库连接
	if err := db_mysql.OpenDB(); err != nil {
		log.Fatal("数据库连接错误"+err.Error())
	}
	defer db_mysql.CloneDB()
	_, err = user2.SaveUserInfo()
	if err != nil {
		u.Ctx.WriteString("错误")
		log.Fatal("数据存储错误"+err.Error())
	}
	u.TplName="home.html"
}


////method :=models.Method{}
//JsonData := models.PraPareJson("getbestblockhash")
//DataString := models.ClienCount(JsonData)
//
//Code := DataString.StatusCode;
//if Code == 200 {
//bytes, err := ioutil.ReadAll(DataString.Body)
//if err != nil {
//c.Ctx.WriteString("解析错误")
//}
//var f models.RPCResult
//err = json.Unmarshal(bytes, &f);
//if err != nil {
//log.Fatal(err.Error())
//}
//fmt.Println(f.Result)
//
//c.Data["result"]= f.Result
////	c.ServeJSON()
//c.TplName = "home.html"
//} else {
//
//fmt.Println("请求失败")
//}
