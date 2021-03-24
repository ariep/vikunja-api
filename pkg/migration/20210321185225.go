// Vikunja is a to-do list application to facilitate your life.
// Copyright 2018-2021 Vikunja and contributors. All rights reserved.
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

type buckets20210321185225 struct {
	IsDoneBucket bool `xorm:"BOOL" json:"is_done_column"`
}

func (buckets20210321185225) TableName() string {
	return "buckets"
}

func init() {
	migrations = append(migrations, &xormigrate.Migration{
		ID:          "20210321185225",
		Description: "Add is done bucket to buckets",
		Migrate: func(tx *xorm.Engine) error {
			return tx.Sync2(buckets20210321185225{})
		},
		Rollback: func(tx *xorm.Engine) error {
			return nil
		},
	})
}
