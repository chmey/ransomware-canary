package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/chmey/ransomware_canary/canary"
	"github.com/chmey/ransomware_canary/cfg"
	"github.com/spf13/cobra"
)

const defaultCfgPath = "/usr/local/etc/ranscanary/config.toml"

var (
	cfgFile string
	config  *cfg.CanaryConfig
)

func initConfig() {

	if cfgFile == "" {
		log.Println("no -config PATH option specified, using default")
		cfgFile = defaultCfgPath
	}
	var err error
	config, err = cfg.NewConfig(cfgFile)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().StringVar(&cfgFile, "config", "", fmt.Sprintf("config file (default: %s)", defaultCfgPath))
}

var rootCmd = &cobra.Command{
	Use:   "ranscanary",
	Short: "Ransomware canary",
	Long: `The ransomware canary watches a file.
Upon modification or deletion it logs and sends an alert email.
See https://github.com/chmey/ransomware-canary`,
	Run: func(cmd *cobra.Command, args []string) {
		c := canary.NewCanary(config)
		c.Start()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
