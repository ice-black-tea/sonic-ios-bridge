/*
 *   sonic-ios-bridge  Connect to your iOS Devices.
 *   Copyright (C) 2022 SonicCloudOrg
 *
 *   This program is free software: you can redistribute it and/or modify
 *   it under the terms of the GNU Affero General Public License as published
 *   by the Free Software Foundation, either version 3 of the License, or
 *   (at your option) any later version.
 *
 *   This program is distributed in the hope that it will be useful,
 *   but WITHOUT ANY WARRANTY; without even the implied warranty of
 *   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *   GNU Affero General Public License for more details.
 *
 *   You should have received a copy of the GNU Affero General Public License
 *   along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */
package cmd

import (
	"os"

	"github.com/SonicCloudOrg/sonic-ios-bridge/src/util"
	"github.com/spf13/cobra"
)

var isJson, isDetail, isFormat bool
var udid string

var strExeccutable string = (func() string {
	exe, _ := os.Executable()
	return exe
})()
var rootCmd = &cobra.Command{
	Use:   strExeccutable,
	Short: "Bridge of iOS Devices",
	Long: `
	▄▄▄▄      ▄▄▄▄    ▄▄▄   ▄▄   ▄▄▄▄▄▄      ▄▄▄▄
  ▄█▀▀▀▀█    ██▀▀██   ███   ██   ▀▀██▀▀    ██▀▀▀▀█
  ██▄       ██    ██  ██▀█  ██     ██     ██▀
   ▀████▄   ██    ██  ██ ██ ██     ██     ██
	   ▀██  ██    ██  ██  █▄██     ██     ██▄
  █▄▄▄▄▄█▀   ██▄▄██   ██   ███   ▄▄██▄▄    ██▄▄▄▄█
   ▀▀▀▀▀      ▀▀▀▀    ▀▀   ▀▀▀   ▀▀▀▀▀▀      ▀▀▀▀
 
	 Copyright (C) 2022 SonicCloudOrg AGPLv3
 https://github.com/SonicCloudOrg/sonic-ios-bridge
 `,
}

func init() {
	rootCmd.Flags().StringVarP(&udid, "udid", "u", "", "device's serialNumber ( default first device )")
}

// Execute error
func Execute() {
	rootCmd.PersistentFlags().String("log-level", "info", "Valid values: [panic, fatal, error, warn, info, debug, trace]")
	rootCmd.ParseFlags(os.Args) // parse flags and set log level
	util.InitLogger(rootCmd.PersistentFlags().Lookup("log-level").Value.String())
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
	//err1 := doc.GenMarkdownTree(rootCmd, "doc")
	//if err1 != nil {
	//	log.Fatal(err1)
	//}
}
