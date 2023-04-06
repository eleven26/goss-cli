package cmd

import (
	"fmt"
	"log"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

type ListCmdArgs struct {
	Prefix string
}

func parseListCmdArgs(cmd *cobra.Command, args []string) ListCmdArgs {
	var prefix string
	if len(args) > 0 {
		prefix = args[0]
	} else {
		prefix = ""
	}

	return ListCmdArgs{
		Prefix: prefix,
	}
}

var listExamples = Examples(`
		# 列出 test 目录下的文件
		go run main.go list test`)

var listCmd = &cobra.Command{
	Use:     "list <prefix>",
	Short:   "列出指定目录下的文件",
	Long:    "列出指定目录下的文件",
	Example: listExamples,
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		listCmdArgs := parseListCmdArgs(cmd, args)

		files, err := app.Files(listCmdArgs.Prefix)
		if err != nil {
			log.Fatal(err)
		}

		if len(files) == 0 {
			fmt.Println("No files found.")
			color.Warnln("No files found.")
			return
		}

		fmt.Printf("%9s %19s %s\n", "Size", "LastModified", "Key")
		for _, file := range files {
			fmt.Printf("%9s %s %s\n",
				Size(file.Size()),
				file.LastModified().Format("2006-01-02 15:04:05"),
				file.Key())
		}

		// spew.Dump(files)
	},
}

func Size(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%dB", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f%cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}
