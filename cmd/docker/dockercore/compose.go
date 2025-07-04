package dockercore

import (
	"fmt"
	"os/exec"

	"github.com/epos-eu/epos-opensource/common"
)

// composeCommand creates a docker compose command configured with the given directory and environment name
func composeCommand(dir, name string, args ...string) *exec.Cmd {
	cmd := exec.Command("docker", append([]string{"compose"}, args...)...)
	cmd.Dir = dir
	if name != "" {
		cmd.Env = append(cmd.Env, "ENV_NAME="+name)
	}
	return cmd
}

// pullEnvImages pulls docker images for the environment with custom messages
func pullEnvImages(dir, name string) error {
	common.PrintStep("Pulling images for environment: %s", name)
	if _, err := common.RunCommand(composeCommand(dir, "", "pull"), false); err != nil {
		return fmt.Errorf("pull images failed: %w", err)
	}
	common.PrintDone("Images pulled for environment: %s", name)
	return nil
}

// deployStack deploys the stack in the specified directory
func deployStack(dir, name string) error {
	common.PrintStep("Deploying stack")
	if _, err := common.RunCommand(composeCommand(dir, name, "up", "-d"), false); err != nil {
		return fmt.Errorf("deploy stack failed: %w", err)
	}
	common.PrintDone("Deployed environment: %s", name)
	return nil
}

// downStack stops the stack running in the given directory
func downStack(dir string, removeVolumes bool) error {
	if removeVolumes {
		_, err := common.RunCommand(composeCommand(dir, "", "down", "-v"), false)
		return err
	}
	_, err := common.RunCommand(composeCommand(dir, "", "down"), false)
	return err
}
