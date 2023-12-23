package handlers

import (
	"github.com/IlyaZayats/faculus/internal/requests"
	"github.com/IlyaZayats/faculus/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FacultyHandlers struct {
	svc       *services.FacultyService
	engine    *gin.Engine
	validator *Validate
}

func NewFacultyHandlers(engine *gin.Engine, svc *services.FacultyService) (*FacultyHandlers, error) {
	h := &FacultyHandlers{
		svc:    svc,
		engine: engine,
	}
	h.initRoute()
	return h, nil
}

func (h *FacultyHandlers) initRoute() {
	h.engine.GET("/faculty", h.GetFaculties)
	h.engine.DELETE("/faculty", h.DeleteFaculty)
	h.engine.PUT("/faculty", h.InsertFaculty)
	h.engine.POST("/faculty", h.UpdateFaculty)
}

func (h *FacultyHandlers) GetFaculties(c *gin.Context) {
	faculties, err := h.svc.GetFaculties()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "get faculties error", "text": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "data": faculties})
}

func (h *FacultyHandlers) DeleteFaculty(c *gin.Context) {

	req, ok := GetRequest[requests.DeleteFacultyRequest](c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "delete faculty request error", "text": ok})
		return
	}

	if err := h.svc.DeleteFaculty(req.Id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "delete faculty error", "text": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *FacultyHandlers) InsertFaculty(c *gin.Context) {

	req, ok := GetRequest[requests.InsertFacultyRequest](c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "insert faculty request error", "text": ok})
		return
	}

	if err := h.svc.InsertFaculty(req.Name); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "insert faculty error", "text": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *FacultyHandlers) UpdateFaculty(c *gin.Context) {

	req, ok := GetRequest[requests.UpdateFacultyRequest](c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "update faculty request error", "text": ok})
		return
	}

	if err := h.svc.UpdateFaculty(req.Id, req.Name); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "update faculty error", "text": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
