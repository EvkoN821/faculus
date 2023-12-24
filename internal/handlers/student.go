package handlers

import (
	"github.com/IlyaZayats/faculus/internal/requests"
	"github.com/IlyaZayats/faculus/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type StudentHandlers struct {
	svc    *services.StudentService
	engine *gin.Engine
}

func NewStudentHandlers(engine *gin.Engine, svc *services.StudentService) (*StudentHandlers, error) {
	h := &StudentHandlers{
		svc:    svc,
		engine: engine,
	}
	h.initRoute()
	return h, nil
}

func (h *StudentHandlers) initRoute() {
	h.engine.GET("/student", h.GetStudents)      //
	h.engine.DELETE("/student", h.DeleteStudent) //
	h.engine.PUT("/student", h.InsertStudent)    //
	h.engine.POST("/student", h.UpdateStudent)   //
}

func (h *StudentHandlers) GetStudents(c *gin.Context) {
	//req, ok := GetRequest[requests.GetStudentsRequest](c)
	//if !ok {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "get students request error", "text": ok})
	//	return
	//}
	students, err := h.svc.GetStudents()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "get students error", "text": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "data": students})
}

func (h *StudentHandlers) DeleteStudent(c *gin.Context) {

	req, ok := GetRequest[requests.DeleteStudentRequest](c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "delete student request error", "text": ok})
		return
	}

	if err := h.svc.DeleteStudent(req.Id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "delete student error", "text": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *StudentHandlers) InsertStudent(c *gin.Context) {

	req, ok := GetRequest[requests.InsertStudentRequest](c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "insert student request error", "text": ok})
		return
	}

	if err := h.svc.InsertStudent(req.FirstName, req.LastName, req.MiddleName, req.PhoneNumber, req.BirthDate, req.GroupId, req.Gender); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "insert student error", "text": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *StudentHandlers) UpdateStudent(c *gin.Context) {

	req, ok := GetRequest[requests.UpdateStudentRequest](c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "update student request error", "text": ok})
		return
	}

	if err := h.svc.UpdateStudent(req.FirstName, req.LastName, req.MiddleName, req.PhoneNumber, req.BirthDate, req.Id, req.Gender); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "update student error", "text": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
