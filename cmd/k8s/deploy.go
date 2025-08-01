// Package k8s contains the cobra cmd implementation for the k8s management
package k8s

import (
	"fmt"

	"github.com/epos-eu/epos-opensource/cmd/k8s/k8score"
	"github.com/epos-eu/epos-opensource/display"

	"github.com/spf13/cobra"
)

var DeployCmd = &cobra.Command{
	Use:   "deploy [env-name]",
	Short: "Create and deploy a new Kubernetes environment in a dedicated namespace",
	Long:  "Sets up a new Kubernetes environment in a fresh namespace, applying all required manifests and configuration. Fails if the namespace already exists.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		protocol := "http"
		if secure {
			protocol = "https"
		}

		k, err := k8score.Deploy(envFile, manifestsDir, path, name, context, protocol)
		if err != nil {
			display.Error("%v", err)
			return
		}
		display.Urls(k.GuiUrl, k.ApiUrl, k.BackofficeUrl, fmt.Sprintf("epos-opensource kubernetes deploy %s", name))
	},
}

func init() {
	DeployCmd.Flags().StringVarP(&envFile, "env-file", "e", "", "Path to the environment variables file (.env)")
	DeployCmd.Flags().StringVarP(&path, "path", "p", "", "Location for the environment files")
	DeployCmd.Flags().StringVarP(&manifestsDir, "manifests-dir", "m", "", "Path to the directory containing the manifests files")
	DeployCmd.Flags().StringVarP(&context, "context", "c", "", "kubectl context used for the environment deployment. Uses current if not set")
	DeployCmd.Flags().BoolVarP(&secure, "secure", "s", false, "Use https as the protocol. If not set uses http by default")
}
