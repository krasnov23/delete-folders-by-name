package main

import "github.com/gin-gonic/gin"
import "github.com/krasnov23/delete-folders-by-name/service"

func main() {
	server := gin.Default()

	server.GET("/clean-folder", func(context *gin.Context) {
		folderName := context.Query("name")

		service.CleanAllFilesFromFolder(folderName)

		context.JSON(200, gin.H{
			"message": "success",
		})
	})

	server.Run(":8080")
}
