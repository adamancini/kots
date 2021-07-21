package cli

import (
	"os"
	"strings"

	"github.com/replicatedhq/kots/pkg/apiserver"
	"github.com/replicatedhq/kots/pkg/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func APICmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "api",
		Short: "Starts the API server",
		Long:  ``,
		PreRun: func(cmd *cobra.Command, args []string) {
			viper.BindPFlags(cmd.Flags())
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			v := viper.GetViper()

			if v.GetString("log-level") == "debug" {
				logger.SetDebug()
			}

			params := apiserver.APIServerParams{
				Version:                os.Getenv("VERSION"),
				PostgresURI:            os.Getenv("POSTGRES_URI"),
				AutocreateClusterToken: os.Getenv("AUTOCREATE_CLUSTER_TOKEN"),
				EnableIdentity:         true,
			}

			apiserver.Start(&params)
			return nil
		},
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	return cmd
}
