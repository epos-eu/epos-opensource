package k8s

import (
	"fmt"

	"github.com/epos-eu/epos-opensource/cmd/k8s/k8score"
	"github.com/epos-eu/epos-opensource/common"

	"github.com/spf13/cobra"
)

var PopulateCmd = &cobra.Command{
	Use:   "populate [env-name] [ttl-directory]",
	Short: "Ingest TTL files into an existing Kubernetes environment",
	Long:  "Uploads and registers all .ttl files from the given directory (recursively) into the specified environment, using a temporary metadata server and port-forwarding.",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		ttlDir := args[1]

		k, err := k8score.Populate(name, ttlDir)
		if err != nil {
			common.PrintError("%v", err)
			return
		}

		common.PrintUrls(k.GuiUrl, k.ApiUrl, k.BackofficeUrl, fmt.Sprintf("epos-opensource kubernetes deploy %s", name))
	},
}
