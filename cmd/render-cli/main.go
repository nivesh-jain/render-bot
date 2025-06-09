package main

import (
	"fmt"
	"os"

	"github.com/nivesh-jain/render-bot/config"
	"github.com/nivesh-jain/render-bot/helper"
	"github.com/spf13/cobra"
)

var (
	configFile string
	doToken    string
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "render-cli",
		Short: "GPU Render Orchestrator CLI",
		Run: func(cmd *cobra.Command, args []string) {
			// Load configuration
			cfg, err := config.LoadConfig(configFile)
			if err != nil {
				fmt.Println("Error loading config:", err)
				os.Exit(1)
			}

			// Validate folder
			if err := helper.ValidateInputFolder(cfg.InputFolder); err != nil {
				fmt.Printf("❌ Input folder error: %v\n", err)
				os.Exit(1)
			}

			// Create Droplet via GODO
			droplet, err := helper.CreateDroplet(doToken, cfg.DropletName, cfg.Region, cfg.Size, cfg.Image)
			if err != nil {
				fmt.Printf("Failed to create droplet: %v\n", err)
				os.Exit(1)
			}

			// Handle remaining actions like cloud-init, file transfer, etc.
			fmt.Printf("Droplet created with ID: %d, IP: %s\n", droplet.ID, droplet.Networks.V4[0].IPAddress)
		},
	}

	// Set flags for config and DO token
	rootCmd.Flags().StringVarP(&configFile, "config", "c", "config.yaml", "Config file (YAML/JSON/Env)")
	rootCmd.Flags().StringVarP(&doToken, "token", "t", "", "DigitalOcean API Token")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
