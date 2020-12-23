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

package cmd

import (
	"code.vikunja.io/api/pkg/initialize"
	"code.vikunja.io/api/pkg/log"
	"code.vikunja.io/api/pkg/mail"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(testmailCmd)
}

var testmailCmd = &cobra.Command{
	Use:   "testmail [email]",
	Short: "Send a test mail using the configured smtp connection",
	Args:  cobra.ExactArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		initialize.LightInit()

		// Start the mail daemon
		mail.StartMailDaemon()
	},
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Sending testmail...")
		email := args[0]
		if err := mail.SendTestMail(email); err != nil {
			log.Errorf("Error sending test mail: %s", err.Error())
			return
		}
		log.Info("Testmail successfully sent.")
	},
}
