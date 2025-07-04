package dockercore

import (
	_ "embed"
	"fmt"
	"net/url"
	"os"
	"path"
	"path/filepath"

	"github.com/epos-eu/epos-opensource/common"

	"github.com/joho/godotenv"
)

const platform = "docker"

//go:embed static/docker-compose.yaml
var ComposeFile string

//go:embed static/.env
var EnvFile string

// NewEnvDir creates a new environment directory with .env and docker-compose.yaml files.
// If customEnvFilePath or customComposeFilePath are provided, it reads the content from those files.
// If they are empty strings, it uses default content for the respective files.
// Returns the path to the created environment directory.
// If any error occurs after directory creation, the directory and its contents are automatically cleaned up.
func NewEnvDir(customEnvFilePath, customComposeFilePath, customPath, name string) (string, error) {
	envPath, err := common.BuildEnvPath(customPath, name, platform)
	if err != nil {
		return "", err
	}

	// Check if directory already exists
	if _, err := os.Stat(envPath); err == nil {
		return "", fmt.Errorf("directory %s already exists", envPath)
	} else if !os.IsNotExist(err) {
		return "", fmt.Errorf("failed to check directory %s: %w", envPath, err)
	}

	// Create the directory
	if err := os.MkdirAll(envPath, 0777); err != nil {
		return "", fmt.Errorf("failed to create env directory %s: %w", envPath, err)
	}

	var success bool
	// Ensure cleanup of directory if any error occurs after creation
	defer func() {
		if !success {
			if removeErr := os.RemoveAll(envPath); removeErr != nil {
				common.PrintError("Failed to cleanup directory '%s' after error: %v. You may need to remove it manually.", envPath, removeErr)
			}
		}
	}()

	// Get .env file content (from file path or use default)
	envContent, err := common.GetContentFromPathOrDefault(customEnvFilePath, EnvFile)
	if err != nil {
		return "", fmt.Errorf("failed to get .env file content: %w", err)
	}

	// Create .env file
	if err := common.CreateFileWithContent(path.Join(envPath, ".env"), envContent); err != nil {
		return "", fmt.Errorf("failed to create .env file: %w", err)
	}

	// Get docker-compose.yaml file content (from file path or use default)
	composeContent, err := common.GetContentFromPathOrDefault(customComposeFilePath, ComposeFile)
	if err != nil {
		return "", fmt.Errorf("failed to get docker-compose.yaml file content: %w", err)
	}

	// Create docker-compose.yaml file
	if err := common.CreateFileWithContent(path.Join(envPath, "docker-compose.yaml"), composeContent); err != nil {
		return "", fmt.Errorf("failed to create docker-compose.yaml file: %w", err)
	}

	success = true
	return envPath, nil
}

func buildEnvURLs(dir string) (portalURL, gatewayURL, backofficeURL string, err error) {
	env, err := godotenv.Read(filepath.Join(dir, ".env"))
	if err != nil {
		return "", "", "", fmt.Errorf("failed to read .env file at %s: %w", filepath.Join(dir, ".env"), err)
	}

	dataportalPort, ok := env["DATAPORTAL_PORT"]
	if !ok {
		return "", "", "", fmt.Errorf("environment variable DATAPORTAL_PORT is not set")
	}

	gatewayPort, ok := env["GATEWAY_PORT"]
	if !ok {
		return "", "", "", fmt.Errorf("environment variable GATEWAY_PORT is not set")
	}

	backofficePort, ok := env["BACKOFFICE_PORT"]
	if !ok {
		return "", "", "", fmt.Errorf("environment variable BACKOFFICE_PORT is not set")
	}

	apiPath, ok := env["API_PATH"]
	if !ok {
		return "", "", "", fmt.Errorf("environment variable API_PATH is not set")
	}

	portalURL = fmt.Sprintf("http://localhost:%s", dataportalPort)

	gatewayURL, err = url.JoinPath(fmt.Sprintf("http://localhost:%s", gatewayPort), apiPath)
	if err != nil {
		return "", "", "", fmt.Errorf("error building gateway URL: %w", err)
	}

	backofficeURL = fmt.Sprintf("http://localhost:%s", backofficePort)

	return portalURL, gatewayURL, backofficeURL, nil
}
