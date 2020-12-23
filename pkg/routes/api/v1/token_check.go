// Vikunja is a to-do list application to facilitate your life.
// Copyright 2018-2020 Vikunja and contributors. All rights reserved.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public Licensee as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public Licensee for more details.
//
// You should have received a copy of the GNU Affero General Public Licensee
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package v1

import (
	"fmt"

	"code.vikunja.io/api/pkg/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// CheckToken checks prints a message if the token is valid or not. Currently only used for testing pourposes.
func CheckToken(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)

	fmt.Println(user.Valid)

	return c.JSON(418, models.Message{Message: "🍵"})
}
