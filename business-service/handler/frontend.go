package handler

import (
	"business-service/model"
	"business-service/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FrontendHandler 前端接口处理器
type FrontendHandler struct {
	personService *service.PersonService
}

// NewFrontendHandler 创建前端接口处理器
func NewFrontendHandler(personService *service.PersonService) *FrontendHandler {
	return &FrontendHandler{
		personService: personService,
	}
}

// RegisterRoutes 注册前端路由
func (h *FrontendHandler) RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/person", h.CreatePerson)
		api.GET("/person/:workId", h.GetPerson)
		api.GET("/persons", h.ListPersons)
		api.PUT("/person/:workId", h.UpdatePerson)
		api.DELETE("/person/:workId", h.DeletePerson)
	}
}

// CreatePerson 创建人员
func (h *FrontendHandler) CreatePerson(c *gin.Context) {
	var person model.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误: " + err.Error()})
		return
	}

	if person.WorkId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "工号不能为空"})
		return
	}

	if err := h.personService.CreatePerson(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "创建成功", "data": person})
}

// GetPerson 获取人员信息
func (h *FrontendHandler) GetPerson(c *gin.Context) {
	workId := c.Param("workId")

	person, err := h.personService.GetPerson(workId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": person})
}

// ListPersons 列出所有人员
func (h *FrontendHandler) ListPersons(c *gin.Context) {
	persons, err := h.personService.ListPersons()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": persons})
}

// UpdatePerson 更新人员信息
func (h *FrontendHandler) UpdatePerson(c *gin.Context) {
	workId := c.Param("workId")

	var person model.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误: " + err.Error()})
		return
	}

	person.WorkId = workId

	if err := h.personService.UpdatePerson(workId, &person); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "更新成功", "data": person})
}

// DeletePerson 删除人员
func (h *FrontendHandler) DeletePerson(c *gin.Context) {
	workId := c.Param("workId")

	if err := h.personService.DeletePerson(workId); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}
