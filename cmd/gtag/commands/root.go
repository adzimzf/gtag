package commands

import (
	"gtag"
	"gtag/fileparser"
	"gtag/formatter"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gtag",
	Short: "gtag is Golang Tag tools",
	Long:  `gtag is a Golang Tag tools with fast speed`,
	Run: func(cmd *cobra.Command, args []string) {
		src, err := cmd.Flags().GetString("input")
		if err != nil {
			log.Fatalf("failed to get bool: %v", err)
		}
		isWrite, err := cmd.Flags().GetBool("write")
		if err != nil {
			log.Fatalf("failed to get bool: %v", err)
		}

		cfg := &gtag.Config{
			FileSource:  src,
			IsOverWrite: isWrite,
		}

		fileParser, err := fileparser.NewFileParser(cfg)
		if err != nil {
			cmd.PrintErrf("failed to parse file due to %v", err)
			return
		}
		for _, m := range fileParser.FindStructs() {
			formatter.Format(m)
		}
		err = fileParser.Write()
		if err != nil {
			cmd.PrintErrf("failed to write due to %v", err)
			return
		}
	},
}

func Execute() {
	rootCmd.Flags().StringP("input", "i", "", "input is a Golang file you want to format")
	rootCmd.Flags().BoolP("write", "w", false, "write will overwrite your input file")
	_ = rootCmd.MarkFlagRequired("input")
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
