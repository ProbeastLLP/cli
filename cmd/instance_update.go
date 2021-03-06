package cmd

import (
	"fmt"
	"os"

	"github.com/civo/cli/config"
	"github.com/civo/cli/utility"

	"github.com/spf13/cobra"
)

var notes, reverseDNS, hostname string

var instanceUpdateCmd = &cobra.Command{
	Use:     "update",
	Example: "civo instance update ID/HOSTNAME --reverse-dns=foo.example.com",
	Aliases: []string{"set"},
	Short:   "Change the instance",
	Long: `Change the notes, hostname or reverse DNS for an instance with partial ID/name provided.
If you wish to use a custom format, the available fields are:

	* ID
	* Hostname
	* ReverseDNS
	* Notes`,
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

		if notes != "" {
			instance.Notes = notes
		}
		if reverseDNS != "" {
			instance.ReverseDNS = reverseDNS
		}
		if hostname != "" {
			instance.Hostname = hostname
		}

		_, err = client.UpdateInstance(instance)
		if err != nil {
			utility.Error("Updating instance failed with %s", err)
			os.Exit(1)
		}

		if outputFormat == "human" {
			fmt.Printf("The instance %s (%s) has been updated\n", utility.Green(instance.Hostname), instance.ID)
		} else {
			ow := utility.NewOutputWriter()
			ow.StartLine()
			ow.AppendData("ID", instance.ID)
			ow.AppendData("Hostname", instance.Hostname)
			ow.AppendDataWithLabel("ReverseDNS", instance.ReverseDNS, "Reverse DNS")
			ow.AppendData("Notes", instance.Notes)
			if outputFormat == "json" {
				ow.WriteSingleObjectJSON()
			} else {
				ow.WriteCustomOutput(outputFields)
			}
		}
	},
}
