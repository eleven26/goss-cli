package cmd

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

type GetCmdArgs struct {
	Key    string
	Target string
}

func parseGetCmdArgs(cmd *cobra.Command, args []string) GetCmdArgs {
	key := args[0]

	target, _ := cmd.Flags().GetString("output")
	if target == "" {
		target = filepath.Base(key)
	}

	return GetCmdArgs{
		Key:    key,
		Target: target,
	}
}

var getExamples = Examples(`
		# 获取 key 为 test.txt 的文件，保存到当前文件夹的 test.txt 文件中
		go run main.go get test.txt`)

var getCmd = &cobra.Command{
	Use:     "get <key>",
	Short:   "获取指定文件",
	Long:    "获取指定文件，保存到当前目录",
	Example: getExamples,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		getCmdArgs := parseGetCmdArgs(cmd, args)

		err := app.Storage.GetToFile(getCmdArgs.Key, getCmdArgs.Target)
		if err != nil {
			log.Fatal(err)
		}

		key := getCmdArgs.Key
		localPath := getCmdArgs.Target

		rc, err := app.Storage.Get(getCmdArgs.Key)
		if err != nil {
			log.Fatal(err)
		}

		defer func() {
			err = rc.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()

		// 获取文件长度
		length, err := app.Storage.Size(key)
		if err != nil {
			log.Fatal(err)
		}

		// 保存到文件 localPath
		f, _ := os.OpenFile(localPath, os.O_CREATE|os.O_WRONLY, 0o644)
		defer func(f *os.File) {
			err = f.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(f)

		// 初始化进度条
		bar := progressbar.DefaultBytes(length, fmt.Sprintf("\"%s\" -> \"%s\"", key, localPath))

		// io.MultiWriter 同时输出到文件和进度条
		_, err = io.Copy(io.MultiWriter(f, bar), rc)

		color.Green(fmt.Sprintf("下载成功！保存路径：\"%s\"", getCmdArgs.Target))
	},
}

func init() {
	getCmd.PersistentFlags().StringP("output", "o", "", "保存到本地的路径")
}
