package cmd

import (
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type DeleteCmdArgs struct {
	Key string
}

func parseDeleteCmdArgs(args []string) DeleteCmdArgs {
	key := args[0]

	return DeleteCmdArgs{
		Key: key,
	}
}

var deleteExamples = Examples(`
		# 删除 key 为 test.txt 的文件
		go run main.go delete test.txt`)

var deleteCmd = &cobra.Command{
	Use:     "delete <key>",
	Short:   "删除指定文件",
	Long:    "删除指定文件",
	Example: deleteExamples,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		deleteCmdArgs := parseDeleteCmdArgs(args)

		exists, err := app.Exists(deleteCmdArgs.Key)
		if err != nil {
			log.Fatal(err)
		}
		if !exists {
			color.Cyan(fmt.Sprintf("key 不存在：%s！", deleteCmdArgs.Key))
			return
		}

		err = app.Delete(deleteCmdArgs.Key)
		if err != nil {
			log.Fatal(err)
		}

		color.Green(fmt.Sprintf("删除成功：%s！", deleteCmdArgs.Key))
	},
}
