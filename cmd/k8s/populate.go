package k8s

import (
	"fmt"

	"github.com/epos-eu/epos-opensource/cmd/k8s/k8score"
	"github.com/epos-eu/epos-opensource/display"

	"github.com/spf13/cobra"
)

var PopulateCmd = &cobra.Command{
	Use:   "populate [env-name] [ttl-directory...]",
	Short: "Ingest TTL files from one or more directories into an environment",
	Long: `Populate an existing environment with all *.ttl files found in the specified directories (recursively). 
Multiple directories can be provided and will be processed in order.

NOTE: to execute the population it will try to use port-forwarding to the cluster. If that fails it will retry using the external api`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		ttlDirs := args[1:]

		k, err := k8score.Populate(name, ttlDirs)
		if err != nil {
			display.Error("%v", err)
			return
		}

		display.Urls(k.GuiUrl, k.ApiUrl, k.BackofficeUrl, fmt.Sprintf("epos-opensource kubernetes deploy %s", name))
	},
}
