package v1

import (
	"github.com/labstack/echo"
	"net/http"
)

// GetAllNamespacesByCurrentUser ...
func GetAllNamespacesByCurrentUser(c echo.Context) error {
	// swagger:operation GET /namespaces namespaces getNamespaces
	// ---
	// summary: Get all namespaces the currently logged in user has at least read access
	// consumes:
	// - application/json
	// produces:
	// - application/json
	// responses:
	//   "200":
	//     "$ref": "#/responses/Namespace"
	//   "500":
	//     "$ref": "#/responses/Message"

	return echo.NewHTTPError(http.StatusNotImplemented)
	/*

		user, err := models.GetCurrentUser(c)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, models.Message{"Could not get the current user."})
		}

		namespaces, err := models.GetAllNamespacesByUserID(user.ID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, models.Message{"Could not get namespaces."})
		}

		return c.JSON(http.StatusOK, namespaces)*/
}