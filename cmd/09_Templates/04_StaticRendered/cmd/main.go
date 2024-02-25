package main

import (
	"fmt"
	"os"

	"github.com/TudorHulban/GolangSandbox/apperrors"
	"github.com/TudorHulban/GolangSandbox/cmd/09_Templates/04_StaticRendered/app"
	"github.com/TudorHulban/GolangSandbox/cmd/09_Templates/04_StaticRendered/app/constants"
)

func main() {
	application, errApp := app.NewApp(constants.ConfigFilePath)
	if errApp != nil {
		fmt.Println(errApp)

		os.Exit(apperrors.OSExitForAppInitialization)
	}

	defer func() {
		application.SaveConfigurationTo(constants.ConfigFilePathBackup)
	}()

	if errStart := application.Start(); errStart != nil {
		fmt.Println(errStart)

		os.Exit(apperrors.OSExitForAppStart)
	}
}
