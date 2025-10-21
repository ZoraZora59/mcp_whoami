package handler

import (
	"business-service/model"
	"business-service/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// InternalHandler 内部接口处理器（供MCP服务调用）
type InternalHandler struct {
	personService *service.PersonService
}

// NewInternalHandler 创建内部接口处理器
func NewInternalHandler(personService *service.PersonService) *InternalHandler {
	return &InternalHandler{
		personService: personService,
	}
}

// RegisterRoutes 注册内部路由
func (h *InternalHandler) RegisterRoutes(router *gin.Engine) {
	internal := router.Group("/internal")
	{
		internal.POST("/person", h.CreatePerson)
		internal.GET("/person/:workId", h.GetPerson)
		internal.GET("/persons", h.ListPersons)
		internal.PUT("/person/:workId", h.UpdatePerson)
		internal.DELETE("/person/:workId", h.DeletePerson)
	}
}

// CreatePerson 创建人员
func (h *InternalHandler) CreatePerson(c *gin.Context) {
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
func (h *InternalHandler) GetPerson(c *gin.Context) {
	workId := c.Param("workId")

	person, err := h.personService.GetPerson(workId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": person})
}

// ListPersons 列出所有人员
func (h *InternalHandler) ListPersons(c *gin.Context) {
	persons, err := h.personService.ListPersons()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": persons})
}

// UpdatePerson 更新人员信息
func (h *InternalHandler) UpdatePerson(c *gin.Context) {
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
func (h *InternalHandler) DeletePerson(c *gin.Context) {
	workId := c.Param("workId")

	if err := h.personService.DeletePerson(workId); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}
