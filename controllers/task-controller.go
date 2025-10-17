package controllers

import "github.com/gin-gonic/gin"

func ListTasksController(c *gin.Context) {
	c.JSON(201, gin.H{"message": "ListTasksController called"})
}

func GetTaskController(c *gin.Context) {
	c.JSON(201, gin.H{"message": "GetTaskController called"})
}

func CreateTaskController(c *gin.Context) {
	c.JSON(201, gin.H{"message": "CreateTaskController called"})
}

func UpdateTaskController(c *gin.Context) {
	c.JSON(201, gin.H{"message": "UpdateTaskController called"})
}

func DeleteTaskController(c *gin.Context) {
	c.JSON(201, gin.H{"message": "DeleteTaskController called"})
}
