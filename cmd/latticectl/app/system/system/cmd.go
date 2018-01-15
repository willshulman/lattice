package service

import (
	"log"
	"os"

	"github.com/mlab-lattice/system/pkg/cli"
	"github.com/mlab-lattice/system/pkg/constants"
	"github.com/mlab-lattice/system/pkg/managerapi/client"
	"github.com/mlab-lattice/system/pkg/managerapi/client/rest"
	"github.com/mlab-lattice/system/pkg/types"

	"github.com/spf13/cobra"
)

var (
	output string

	systemIDString string
	url            string
	systemID       types.SystemID
	userClient     client.Interface
	systemClient   client.SystemClient
)

var Cmd = &cobra.Command{
	Use:  "service",
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(1)
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list services",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		services, err := systemClient.Services(systemID).List()
		if err != nil {
			log.Panic(err)
		}

		if err := cli.ShowServices(services, cli.OutputFormat(output)); err != nil {
			log.Panic(err)
		}
	},
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get service",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := types.ServiceID(args[0])
		build, err := systemClient.Services(systemID).Get(id)
		if err != nil {
			log.Panic(err)
		}

		if err := cli.ShowService(build, cli.OutputFormat(output)); err != nil {
			log.Panic(err)
		}
	},
}

func init() {
	cobra.OnInitialize(initCmd)

	Cmd.PersistentFlags().StringVar(&url, "url", "", "URL of the manager-api for the system")
	Cmd.PersistentFlags().StringVar(&systemIDString, "system", string(constants.SystemIDDefault), "system to use")

	Cmd.AddCommand(listCmd)
	Cmd.AddCommand(getCmd)
	getCmd.Flags().StringVarP(&output, "output", "o", "table", "whether or not to display output as JSON")
	listCmd.Flags().StringVarP(&output, "output", "o", "table", "whether or not to display output as JSON")
}

func initCmd() {
	systemID = types.SystemID(systemIDString)

	userClient = rest.NewClient(url)
	systemClient = userClient.Systems()
}