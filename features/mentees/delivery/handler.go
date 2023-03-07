package delivery

import (
	"immersiveApp/features/mentees"
	"immersiveApp/utils/helpers"
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
)

type MenteeHandler struct {
	Service mentees.MenteeServiceInterface
}

func New(srv mentees.MenteeServiceInterface) *MenteeHandler {
	return &MenteeHandler{
		Service: srv,
	}
}

func (m *MenteeHandler) GetAll(c echo.Context) error{
	menteeEntity, err := m.Service.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail("error read data"))
	}
	listMenteeResponse := ListMenteeToMenteeResponse(menteeEntity)
	return c.JSON(http.StatusOK, helpers.ResponseSuccess("all teams", listMenteeResponse))
}

func (m *MenteeHandler) GetById(c echo.Context) error{
	_id, _ := strconv.Atoi(c.Param("id"))
	id := uint(_id)
	menteeEntity, err := m.Service.GetById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.ResponseFail("data not found"))
	}
	menteeResponse := MenteeEntitytoMenteeResponse(menteeEntity)
	return c.JSON(http.StatusOK, helpers.ResponseSuccess("-", menteeResponse))
}

func (m *MenteeHandler) Create(c echo.Context) error{
	var formInput MenteeRequest
	if err := c.Bind(&formInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
	}

	mentee, err := m.Service.Create(menteeRequestToMenteeEntity(formInput))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusCreated, helpers.ResponseSuccess("Create Data Success", MenteeEntitytoMenteeResponse(mentee)))
}

func (m *MenteeHandler) Update(c echo.Context) error{
	var formInput MenteeRequest
	if err := c.Bind(&formInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
	}

	_id, _ := strconv.Atoi(c.Param("id"))
	id := uint(_id)

	mentee, err := m.Service.Update(menteeRequestToMenteeEntity(formInput), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("Update Data Success", MenteeEntitytoMenteeResponse(mentee)))
}

func (m *MenteeHandler) Delete(c echo.Context) error{
	_id, _ := strconv.Atoi(c.Param("id"))
	id := uint(_id)

	if err := m.Service.Delete(id); err != nil {
		return c.JSON(http.StatusNotFound, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("Delete Data Success", nil))
}