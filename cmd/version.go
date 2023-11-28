/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/jiaozifs/jiaozifs/api"
	"github.com/jiaozifs/jiaozifs/api/api_impl"
	"github.com/jiaozifs/jiaozifs/config"
	"github.com/jiaozifs/jiaozifs/version"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version of jiaozifs",
	Long:  `jiaozifs version`,
	RunE: func(cmd *cobra.Command, args []string) error {
		swagger, err := api.GetSwagger()
		if err != nil {
			return err
		}
		fmt.Println("Version ", version.UserVersion())
		fmt.Println("API Version ", swagger.Info.Version)

		//get runtime version
		cfg, err := config.LoadConfig(cfgFile)
		if err != nil {
			return err
		}
		client, err := api.NewClient(cfg.API.Listen + api_impl.APIV1Prefix)
		if err != nil {
			return err
		}

		versionResp, err := client.GetVersion(cmd.Context())
		if err != nil {
			return err
		}
		okResp, err := api.ParseGetVersionResponse(versionResp)
		if err != nil {
			return err
		}
		if okResp.JSON200 == nil {
			return fmt.Errorf("request version fail %d %s", okResp.HTTPResponse.StatusCode, okResp.HTTPResponse.Body)
		}
		fmt.Println("Runtime Version ", okResp.JSON200.Version)
		fmt.Println("Runtime API Version ", okResp.JSON200.ApiVersion)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
