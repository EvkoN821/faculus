package handlers

import (
	"github.com/IlyaZayats/faculus/internal/requests"
	"github.com/IlyaZayats/faculus/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GroupHandlers struct {
	svc    *services.GroupService
	engine *gin.Engine
}

func NewGroupHandlers(engine *gin.Engine, svc *services.GroupService) (*GroupHandlers, error) {
	h := &GroupHandlers{
		svc:    svc,
		engine: engine,
	}
	h.initRoute()
	return h, nil
}

func (h *GroupHandlers) initRoute() {
	h.engine.GET("/group", h.GetGroups)      //
	h.engine.DELETE("/group", h.DeleteGroup) //
	h.engine.PUT("/group", h.InsertGroup)    //
	h.engine.POST("/group", h.UpdateGroup)   //
}

func (h *GroupHandlers) GetGroups(c *gin.Context) {
	//req, ok := GetRequest[requests.GetGroupsRequest](c)
	//if !ok {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "get groups request error", "text": ok})
	//	return
	//}
	groups, err := h.svc.GetGroups()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "get groups error", "text": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "data": groups})
}

func (h *GroupHandlers) DeleteGroup(c *gin.Context) {

	req, ok := GetRequest[requests.DeleteGroupRequest](c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "delete group request error", "text": ok})
		return
	}

	if err := h.svc.DeleteGroup(req.Id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "delete group error", "text": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *GroupHandlers) InsertGroup(c *gin.Context) {

	req, ok := GetRequest[requests.InsertGroupRequest](c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "insert group request error", "text": ok})
		return
	}

	if err := h.svc.InsertGroup(req.FacultyId, req.Name); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "insert group error", "text": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *GroupHandlers) UpdateGroup(c *gin.Context) {

	req, ok := GetRequest[requests.UpdateGroupRequest](c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "update group request error", "text": ok})
		return
	}

	if err := h.svc.UpdateGroup(req.Id, req.Name); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "update group error", "text": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
