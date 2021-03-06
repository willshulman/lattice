package app

import (
	goflag "flag"

	"github.com/mlab-lattice/lattice/pkg/api/server/authentication/authenticator/token/tokenfile"
	"github.com/mlab-lattice/lattice/pkg/api/server/rest"
	mockbackend "github.com/mlab-lattice/lattice/pkg/backend/mock/api/server/backend"
	mockresolver "github.com/mlab-lattice/lattice/pkg/backend/mock/definition/component/resolver"
	"github.com/mlab-lattice/lattice/pkg/definition/resolver"
	"github.com/mlab-lattice/lattice/pkg/util/cli"
	"github.com/mlab-lattice/lattice/pkg/util/cli/flags"
	"github.com/mlab-lattice/lattice/pkg/util/git"

	"github.com/spf13/pflag"
)

func Command() *cli.RootCommand {
	// https://flowerinthenight.com/blog/2017/12/01/golang-cobra-glog
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)

	var (
		port          int32
		tokenAuthFile string
		workDirectory string
	)

	command := &cli.RootCommand{
		Name: "api-server",
		Command: &cli.Command{
			Flags: cli.Flags{
				"port": &flags.Int32{
					Usage:   "port to bind to",
					Default: 8080,
					Target:  &port,
				},
				"token-auth-file": &flags.String{
					Usage:   "path for token file for bearer token authenticator",
					Default: "",
					Target:  &tokenAuthFile,
				},
				"work-directory": &flags.String{
					Usage:   "directory used to download git repositories",
					Default: "/tmp/lattice/mock/api-server",
					Target:  &workDirectory,
				},
			},
			Run: func(args []string, flags cli.Flags) error {
				templateStore := mockresolver.NewMemoryTemplateStore()
				secretStore := mockresolver.NewMemorySecretStore()
				gitResolver, err := git.NewResolver(workDirectory, false)
				if err != nil {
					return err
				}

				r := resolver.NewComponentResolver(gitResolver, templateStore, secretStore)
				backend := mockbackend.NewMockBackend(r)
				// construct server options
				options := createServerOptions(tokenAuthFile)
				rest.RunNewRestServer(backend, r, port, options)
				return nil
			},
		},
	}

	return command
}

func createServerOptions(tokenAuthFile string) *rest.ServerOptions {
	options := rest.NewServerOptions()

	// enable api authentication key as needed
	if tokenAuthFile != "" {
		tokenAuthenticator, err := tokenfile.NewFromCSV(tokenAuthFile)
		if err != nil {
			panic(err)
		}
		options.AuthOptions.Token = tokenAuthenticator
	}
	return options
}
