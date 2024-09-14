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
	"fmt"
	"os"

	"github.com/SonicCloudOrg/sonic-ios-bridge/src/util"

	"github.com/spf13/cobra"
)

var timeout int = 30

var mountCmd = &cobra.Command{
	Use:   "mount",
	Short: "Mount device development disk",
	Long:  "Mount device development disk",
	RunE: func(cmd *cobra.Command, args []string) error {
		device := util.GetDeviceByUdId(udid)
		if device == nil {
			os.Exit(0)
		}
		util.CheckMount(device, timeout)
		fmt.Println("mount successful!")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(mountCmd)
	mountCmd.Flags().IntVarP(&timeout, "timeout", "t", 30, "download timeout seconds")
	mountCmd.Flags().StringVarP(&udid, "udid", "u", "", "device's serialNumber ( default first device )")
}
