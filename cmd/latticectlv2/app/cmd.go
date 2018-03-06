package app

import (
	"github.com/mlab-lattice/system/pkg/cli/latticectl"
	"github.com/mlab-lattice/system/pkg/cli/latticectl/commands/context"
	"github.com/mlab-lattice/system/pkg/cli/latticectl/commands/systems"
	"github.com/mlab-lattice/system/pkg/cli/latticectl/commands/systems/deploys"
)

var lctl = latticectl.Latticectl{
	Client:  latticectl.DefaultLatticeClient,
	Context: &latticectl.DefaultFileContext{},
	Root: &latticectl.BaseCommand{
		Name:  "latticectl",
		Short: "command line utility for interacting with lattice clusters and systems",
		Subcommands: []latticectl.Command{
			&context.Command{
				Subcommands: []latticectl.Command{
					&context.GetCommand{},
					&context.SetCommand{},
				},
			},
			&systems.Command{
				Subcommands: []latticectl.Command{
					&systems.CreateCommand{},
					&systems.GetCommand{},
					&systems.DeleteCommand{},
					&systems.BuildCommand{},
					&systems.DeployCommand{},
					&deploys.Command{
						Subcommands: []latticectl.Command{
							&deploys.GetCommand{},
						},
					},
				},
			},
		},
	},
}

func Execute() {
	//lctl.Execute()
	lctl.ExecuteColon()
}
