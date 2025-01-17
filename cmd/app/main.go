package main

import (
	"context"
	"log"

	"github.com/dnsoftware/mpm-save-get-shares/pkg/utils"

	_ "github.com/golang-migrate/migrate/v4/database/clickhouse"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/dnsoftware/mpm-shares-timeseries/config"
	"github.com/dnsoftware/mpm-shares-timeseries/internal/app"
	"github.com/dnsoftware/mpm-shares-timeseries/internal/constants"
)

func main() {
	ctx := context.Background()

	basePath, err := utils.GetProjectRoot(constants.ProjectRootAnchorFile)
	if err != nil {
		log.Fatalf("GetProjectRoot failed: %s", err.Error())
	}
	configFile := basePath + "/config.yaml"
	envFile := basePath + "/.env"

	cfg, err := config.New(configFile, envFile)
	if err != nil {

	}

	app.Run(ctx, cfg)
}
