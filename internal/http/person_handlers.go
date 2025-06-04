package http

import (
	"github.com/artengin/rest-api-go/internal/app"
	"github.com/artengin/rest-api-go/internal/logic"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type PersonHandler struct {
	logic    *logic.PersonLogic
	validate *validator.Validate
}

func NewPersonHandler(logic *logic.PersonLogic) *PersonHandler {
	return &PersonHandler{
		logic:    logic,
		validate: validator.New(),
	}
}

func (h *PersonHandler) CreatePerson(c echo.Context) error {
	p := new(app.Person)

	if err := c.Bind(p); err != nil {
		log.Errorf("Failed to bind person: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	if err := h.validate.Struct(p); err != nil {
		log.Errorf("Validation error: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	ctx := c.Request().Context()

	err := h.logic.CreatePerson(ctx, p)
	if err != nil {
		log.Errorf("Failed to create person: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create person"})
	}

	return c.JSON(http.StatusCreated, p)
}

func (h *PersonHandler) GetAllPerson(c echo.Context) error {
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	offset, _ := strconv.Atoi(c.QueryParam("offset"))
	search := c.QueryParam("search")

	ctx := c.Request().Context()
	persons, err := h.logic.GetAllPerson(ctx, limit, offset, search)
	if err != nil {
		log.Errorf("Failed to get persons: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get persons"})
	}
	if persons == nil {
		persons = make([]*app.Person, 0)
	}
	return c.JSON(http.StatusOK, persons)
}

func (h *PersonHandler) GetPerson(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Invalid ID parameter"})
	}
	id := int64(idP)

	ctx := c.Request().Context()
	person, err := h.logic.GetPerson(ctx, id)
	if err != nil {
		log.Errorf("Failed to get person: %v", err)
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Person not found"})
	}
	return c.JSON(http.StatusOK, person)
}

func (h *PersonHandler) UpdatePerson(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Invalid ID parameter"})
	}
	id := int64(idP)

	p := new(app.Person)
	if err := c.Bind(p); err != nil {
		log.Errorf("Failed to bind person: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	if err := h.validate.Struct(p); err != nil {
		log.Errorf("Validation error: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	ctx := c.Request().Context()
	err = h.logic.UpdatePerson(ctx, id, p)
	if err != nil {
		log.Errorf("Failed to update person: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update person"})
	}
	return c.JSON(http.StatusOK, p)
}

func (h *PersonHandler) DeletePerson(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Invalid ID parameter"})
	}
	id := int64(idP)

	ctx := c.Request().Context()
	err = h.logic.DeletePerson(ctx, id)
	if err != nil {
		log.Errorf("Failed to delete person: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete person"})
	}
	return c.NoContent(http.StatusNoContent)
}
