package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func _string(c *gin.Context){
	c.String(http.StatusOK,"hello!")
}

func _json(c *gin.Context){
	//json响应结构体
	type UserInfo struct{
		UserName 	string 		`json:"user_name"`
		Age 		int			`json:"age"`
		Password	string 		`json:"-"`  //忽略转换为json
	}
	// user := UserInfo{UserName:"ys",Age:23,Password:"123456"}
	// c.JSON(http.StatusOK,user)


	//json 响应map
	// userMap := map[string]string{
	// 	"user_name":"ys",
	// 	"age":		"23",
	// }
	// c.JSON(http.StatusOK,userMap)


	//直接响应jaon
	c.JSON(http.StatusOK,gin.H{"usaername":"ys","age":23})
}

func _yaml(c *gin.Context){
	c.YAML(http.StatusOK,gin.H{"username":"yy","age":23,"data":gin.H{"user":"ff"}})
}

func _xml(c *gin.Context){
	c.XML(http.StatusOK,gin.H{"username":"yy","age":23,"data":gin.H{"user":"ff"}})
}

func _html(c *gin.Context){
	type UserInfo struct{
		UserName 	string 		`json:"user_name"`
		Age 		int			`json:"age"`
		Password	string 		`json:"-"`  //忽略转换为json
	}
	user := UserInfo{UserName:"ys",Age:23,Password:"123456"}
	c.HTML(http.StatusOK,"index.html",user)
}

func main() {
	router := gin.Default()
	//加载模板目录下所有的模板文件
	router.LoadHTMLGlob("templates/*")
	//在golang中，没有相对文件的描述，他只有相对项目的路径
	//配置单个文件，网页请求的路由，文件的路径
	router.StaticFile("/img","static/img.png")
	//网页请求这个静态目录的前缀，第二个参数是一个目录，注意前缀不要重复
	router.StaticFS("/static",http.Dir("static/static"))
	router.GET("/",_string)
	router.GET("/json",_json)
	router.GET("/xml",_xml)
	router.GET("/yaml",_yaml)
	router.GET("/html",_html)
	router.Run(":8081")
}