package cmd

import (
	"fmt"
	"os"

	"github.com/civo/cli/config"
	"github.com/civo/cli/utility"
	"github.com/spf13/cobra"
)

var instancePasswordCmd = &cobra.Command{
	Use:     "password",
	Example: "civo instance public-ip ID/HOSTNAME",
	Short:   "Show instance's default password",
	Aliases: []string{"pw"},
	Long: `Show the specified instance's default SSH password by part of the instance's ID or name.
If you wish to use a custom format, the available fields are:

	* ID
	* Hostname
	* Password
	* User`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := config.CivoAPIClient()
		if err != nil {
			utility.Error("Creating the connection to Civo's API failed with %s", err)
			os.Exit(1)
		}

		instance, err := client.FindInstance(args[0])
		if err != nil {
			utility.Error("Finding instance failed with %s", err)
			os.Exit(1)
		}

		if outputFormat == "human" {
			fmt.Printf("The instance %s (%s) has the password %s (and user %s)\n", utility.Green(instance.Hostname), instance.ID, utility.Green(instance.InitialPassword), utility.Green(instance.InitialUser))
		} else {
			ow := utility.NewOutputWriter()
			ow.StartLine()
			ow.AppendData("ID", instance.ID)
			ow.AppendData("Hostname", instance.Hostname)
			ow.AppendData("Password", instance.InitialPassword)
			ow.AppendData("User", instance.InitialUser)
			if outputFormat == "json" {
				ow.WriteSingleObjectJSON()
			} else {
				ow.WriteCustomOutput(outputFields)
			}
		}
	},
}
