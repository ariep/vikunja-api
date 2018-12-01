//  Vikunja is a todo-list application to facilitate your life.
//  Copyright 2018 Vikunja and contributors. All rights reserved.
//
//  This program is free software: you can redistribute it and/or modify
//  it under the terms of the GNU General Public License as published by
//  the Free Software Foundation, either version 3 of the License, or
//  (at your option) any later version.
//
//  This program is distributed in the hope that it will be useful,
//  but WITHOUT ANY WARRANTY; without even the implied warranty of
//  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//  GNU General Public License for more details.
//
//  You should have received a copy of the GNU General Public License
//  along with this program.  If not, see <https://www.gnu.org/licenses/>.

package handler

import (
	"code.vikunja.io/web"
	"github.com/labstack/echo"
	"github.com/op/go-logging"
	"net/http"
)

// ReadOneWeb is the webhandler to get one object
func (c *WebHandler) ReadOneWeb(ctx echo.Context) error {
	// Get our model
	currentStruct := c.EmptyStruct()

	// Get the object & bind params to struct
	if err := ParamBinder(currentStruct, ctx); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "No or invalid model provided.")
	}

	// Get our object
	err := currentStruct.ReadOne()
	if err != nil {
		return HandleHTTPError(err, ctx)
	}

	// Check rights
	// We can only check the rights on a full object, which is why we need to check it afterwards
	authprovider := ctx.Get("AuthProvider").(*web.Auths)
	currentAuth, err := authprovider.AuthObject(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not determine the current user.")
	}
	if !currentStruct.CanRead(currentAuth) {
		ctx.Get("LoggingProvider").(*logging.Logger).Noticef("Tried to read one while not having the rights for it", currentAuth)
		return echo.NewHTTPError(http.StatusForbidden, "You don't have the right to see this")
	}

	return ctx.JSON(http.StatusOK, currentStruct)
}