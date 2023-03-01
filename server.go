package main
import "github.com/gin-gonic/gin"




func main(){
	server := gin.Default()

	server.GET("/getHello", getHello)

	server.Run(":8080")
}

func getHello(ctx *gin.Context){

	ctx.JSON(200, gin.H{
		"message": "Hello World",
	})
}