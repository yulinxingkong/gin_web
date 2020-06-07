package main

import "github.com/gin-gonic/gin"

func main(){
    // 定义路由，使用默认
    r:=gin.Default()
    // gin.H 本质上是map
    r.GET("/ping",func(c *gin.Context){
        c.JSON(200,gin.H{
            "msg":"pong",
        }
    }
    // 运行
    r.Run()
}
