package commands

import (
	"bytes"
	"go/format"
	"gtag"
	"gtag/fileparser"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

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

		absPath, err := filepath.Abs(src)
		if err != nil {
			log.Fatalf("directory is invalid, err:%v", err)
		}

		file, fs, f, err := fileparser.ParseFile(absPath)
		if err != nil {
			log.Fatalf("failed to parse file %v", err)
		}

		var buf bytes.Buffer
		for _, m := range file {
			gtag.BeautifyTag(&m)
		}
		err = format.Node(&buf, fs, f)
		if err != nil {
			log.Fatalf("failed to format node: %v\n", err)
		}

		if isWrite {
			if err := ioutil.WriteFile(absPath, buf.Bytes(), 0); err != nil {
				log.Fatalf("failed to write file, err :%v", err)
			}
		} else {
			_, err := buf.WriteTo(os.Stdout)
			if err != nil {
				log.Fatalf("failed to write file, err :%v", err)
			}
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
