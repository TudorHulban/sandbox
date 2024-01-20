package main

import (
	"fmt"
	"os"

	"github.com/TudorHulban/GolangSandbox/apperrors"
	"github.com/TudorHulban/GolangSandbox/cmd/09_Templates/04_StaticRendered/app"
)

func main() {
	application, errApp := app.NewApp(_configFilePath)
	if errApp != nil {
		fmt.Println(errApp)

		os.Exit(apperrors.OSExitForAppInitialization)
	}

	defer application.SaveConfigurationTo(_configFilePathBackup)

	if errStart := application.Start(); errStart != nil {
		fmt.Println(errApp)

		os.Exit(apperrors.OSExitForAppStart)
	}
}
