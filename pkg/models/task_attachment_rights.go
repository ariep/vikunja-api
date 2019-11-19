//   Vikunja is a todo-list application to facilitate your life.
//   Copyright 2019 Vikunja and contributors. All rights reserved.
//
//   This program is free software: you can redistribute it and/or modify
//   it under the terms of the GNU General Public License as published by
//   the Free Software Foundation, either version 3 of the License, or
//   (at your option) any later version.
//
//   This program is distributed in the hope that it will be useful,
//   but WITHOUT ANY WARRANTY; without even the implied warranty of
//   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//   GNU General Public License for more details.
//
//   You should have received a copy of the GNU General Public License
//   along with this program.  If not, see <https://www.gnu.org/licenses/>.

package models

import "code.vikunja.io/web"

// CanRead checks if the user can see an attachment
func (ta *TaskAttachment) CanRead(a web.Auth) (bool, error) {
	t := &Task{ID: ta.TaskID}
	return t.CanRead(a)
}

// CanDelete checks if the user can delete an attachment
func (ta *TaskAttachment) CanDelete(a web.Auth) (bool, error) {
	t := &Task{ID: ta.TaskID}
	return t.CanWrite(a)
}

// CanCreate checks if the user can create an attachment
func (ta *TaskAttachment) CanCreate(a web.Auth) (bool, error) {
	t, err := GetTaskByIDSimple(ta.TaskID)
	if err != nil {
		return false, err
	}
	return t.CanCreate(a)
}