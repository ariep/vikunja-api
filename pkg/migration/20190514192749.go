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

package migration

import (
	"src.techknowlogick.com/xormigrate"
	"xorm.io/xorm"
)

type listTask20190514192749 struct {
	DoneAtUnix int64 `xorm:"INDEX null" json:"done_at"`
}

func (listTask20190514192749) TableName() string {
	return "tasks"
}

func init() {
	migrations = append(migrations, &xormigrate.Migration{
		ID:          "20190514192749",
		Description: "Add task done at",
		Migrate: func(tx *xorm.Engine) error {
			return tx.Sync2(listTask20190514192749{})
		},
		Rollback: func(tx *xorm.Engine) error {
			return dropTableColum(tx, "tasks", "done_at_unix")
		},
	})
}
