/*

Copyright (C) 2018  Ettore Di Giacinto <mudler@gentoo.org>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.

*/

package environment

import (
	client "github.com/MottainaiCI/mottainai-server/pkg/client"
	setting "github.com/MottainaiCI/mottainai-server/pkg/settings"
	environment "github.com/MottainaiCI/replicant/pkg/environment"
	cobra "github.com/spf13/cobra"
	viper "github.com/spf13/viper"
)

func newEnvironmentDestroy(config *setting.Config) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "destroy [OPTIONS]",
		Short: "Destroy a remote environment",
		Args:  cobra.OnlyValidArgs,
		// TODO: PreRun check of minimal args if --json is not present
		Run: func(cmd *cobra.Command, args []string) {
			var v *viper.Viper = config.Viper

			client := client.NewTokenClient(v.GetString("master"), v.GetString("apikey"), config)
			dep := &environment.Deployment{Client: client}
			dep.Destroy()

		},
	}
	return cmd
}
