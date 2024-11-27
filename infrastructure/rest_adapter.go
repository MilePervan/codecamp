package infrastructure

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"todo-cc/config"
	"todo-cc/service"
)

type CreateTaskRequestDTO struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
	Completed   bool      `json:"completed"`
}

type Controller struct {
	router  *gin.Engine
	service service.Todo
}

func NewRestController(service service.Todo) Controller {
	c := Controller{
		router:  gin.Default(),
		service: service,
	}

	c.setupRoutes()

	return c
}

func (c *Controller) setupRoutes() {
	c.router.Use(gin.Logger())

	c.router.GET("/", c.health)

	v1 := c.router.Group("v1")
	{
		v1.POST("/tasks", c.createNewTask)
	}
}

func (c *Controller) createNewTask(ctx *gin.Context) {
	var taskDTO CreateTaskRequestDTO

	if err := ctx.ShouldBindJSON(&taskDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Errorf("error while binding json: %v", err.Error()),
		})
		return
	}

	err := c.service.CreateNewTask(taskDTO.Title, taskDTO.Description, taskDTO.Deadline, taskDTO.Completed)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Errorf("error while saving task: %v", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "success",
	})
}

func (c *Controller) health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "todo-cc is running...",
	})
}

func (c *Controller) Run() {
	err := c.router.Run(config.SERVER_PORT)
	if err != nil {
		panic(err)
	}
}