package v1

import (
	"github.com/labstack/echo"
	"strconv"
	"net/http"
	"git.kolaente.de/konrad/list/models"
)

func DeleteListByID(c echo.Context) error {
	// Check if we have our ID
	id := c.Param("id")
	// Make int
	itemID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Message{"Invalid ID."})
	}

	err = models.DeleteListByID(itemID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Message{"An error occured."})
	}

	return c.JSON(http.StatusOK, models.Message{"The list was deleted with success."})
}