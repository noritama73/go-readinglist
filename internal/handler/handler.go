package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/noritama73/go-readinglist/internal/model"
)

type ItemHandler struct {
	itemRepository model.ItemRepository
}

func NewItemHandler(sqlService *SQLService) *ItemHandler {
	return &ItemHandler{
		itemRepository: sqlService,
	}
}

func (h *ItemHandler) GetItem(c echo.Context) error {
	id := c.QueryParam("id")

	if id == "" {
		return apiResponseErr(c, http.StatusBadRequest, clientErrMsg)
	}

	item, e := h.itemRepository.GetItem(model.ID(id))
	if e != nil {
		log.Println(e)
		return apiResponseErr(c, http.StatusInternalServerError, serverErrMsg)
	}

	return apiResponseOK(c, item)
}

func (h *ItemHandler) ListItems(c echo.Context) error {
	itemList, e := h.itemRepository.ListItems()
	if e != nil {
		return apiResponseErr(c, http.StatusInternalServerError, serverErrMsg)
	}
	return apiResponseOK(c, itemList)
}

func (h *ItemHandler) PutItemData(c echo.Context) error {
	var param model.PutItemData

	if e := c.Bind(&param); e != nil {
		return apiResponseErr(c, http.StatusBadRequest, clientErrMsg)
	}

	if e := h.itemRepository.PutItemData([]byte(param.Data)); e != nil {
		return apiResponseErr(c, http.StatusInternalServerError, serverErrMsg)
	}
	return apiResponseOK(c, "Successfully put data!")
}

func (h *ItemHandler) UpdateItemData(c echo.Context) error {
	var param model.UpdateItemData

	if e := c.Bind(&param); e != nil {
		return apiResponseErr(c, http.StatusBadRequest, clientErrMsg)
	}

	if e := h.itemRepository.UpdateItemData(param.ID, []byte(param.Data)); e != nil {
		return apiResponseErr(c, http.StatusInternalServerError, serverErrMsg)
	}
	return apiResponseOK(c, "Successfully update data!")
}

func (h *ItemHandler) DeleteItemData(c echo.Context) error {
	id := c.QueryParam("id")

	if id == "" {
		return apiResponseErr(c, http.StatusBadRequest, clientErrMsg)
	}

	if e := h.itemRepository.DeleteItemData(model.ID(id)); e != nil {
		return apiResponseErr(c, http.StatusInternalServerError, serverErrMsg)
	}

	return apiResponseOK(c, "Successfully delete data!")
}
