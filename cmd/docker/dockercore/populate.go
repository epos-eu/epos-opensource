package dockercore

import (
	"fmt"
	"net/url"
	"path/filepath"

	"github.com/epos-eu/epos-opensource/common"
)

func Populate(path, name, ttlDir string) (portalURL, gatewayURL string, err error) {
	common.PrintStep("Populating environment: %s", name)
	dir, err := common.GetEnvDir(path, name, pathPrefix)
	if err != nil {
		return "", "", fmt.Errorf("failed to get environment directory: %w", err)
	}
	common.PrintDone("Environment found in dir: %s", dir)

	ttlDir, err = filepath.Abs(ttlDir)
	if err != nil {
		return "", "", fmt.Errorf("error finding absolute path for given metadata path '%s': %w", ttlDir, err)
	}

	common.PrintStep("Starting metadata server")
	metadataServer, err := common.NewMetadataServer(ttlDir)
	if err != nil {
		return "", "", fmt.Errorf("creating metadata server for dir %q: %w", ttlDir, err)
	}

	if err = metadataServer.Start(); err != nil {
		return "", "", fmt.Errorf("starting metadata server: %w", err)
	}

	// Make sure the metadata server is stopped and URLs are printed last on success.
	defer func(env string) {
		common.PrintStep("Stopping metadata server")
		if err := metadataServer.Stop(); err != nil {
			common.PrintError("Error while removing metadata server deployment: %v. You might have to remove it manually.", err)
		} else {
			common.PrintDone("Metadata server stopped successfully")
		}
	}(name)

	portalURL, gatewayURL, err = buildEnvURLs(dir)
	if err != nil {
		return "", "", fmt.Errorf("error building env urls for environment '%s': %w", dir, err)
	}

	err = metadataServer.PostFiles(gatewayURL)
	if err != nil {
		return "", "", fmt.Errorf("error populating environment: %w", err)
	}

	gatewayURL, err = url.JoinPath(gatewayURL, "ui/")
	return portalURL, gatewayURL, err
}
