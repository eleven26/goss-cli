package cmd

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/schollz/progressbar/v3"

	fs "github.com/eleven26/go-filesystem"
	"github.com/spf13/cobra"
)

type PutCmdArgs struct {
	Key   string
	Path  string
	Force bool
}

func filename(path string) string {
	if filepath.Dir(path) == "." {
		return path
	}

	return filepath.Base(path)
}

func parsePutCmdArgs(cmd *cobra.Command, args []string) PutCmdArgs {
	key, _ := cmd.Flags().GetString("key")
	if key == "" {
		key = filename(args[0])
	}

	force, _ := cmd.Flags().GetBool("force")

	path := args[0]

	return PutCmdArgs{
		Key:   key,
		Path:  path,
		Force: force,
	}
}

func NewReaderWithProgress(path string, description string) io.Reader {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	size, err := fs.Size(path)
	if err != nil {
		log.Fatal(err)
	}

	bar := progressbar.DefaultBytes(size, description)

	r := progressbar.NewReader(bufio.NewReader(file), bar)

	return &r
}

var putExamples = Examples(`
		# 将 test.txt 上传到 oss，保存路径是 test.txt
		go run main.go put test.txt`)

var putCmd = &cobra.Command{
	Use:     "put <path>",
	Short:   "上传文件",
	Long:    "上传文件，将本地路径 path 指向的文件上传到 oss",
	Example: putExamples,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		putCmdArgs := parsePutCmdArgs(cmd, args)

		if exist, _ := fs.Exists(putCmdArgs.Path); !exist {
			log.Fatalf("文件不存在：%s\n", putCmdArgs.Path)
		}

		exist, err := app.Storage.Exists(putCmdArgs.Key)
		if err != nil {
			log.Fatal(err)
		}

		if exist && !putCmdArgs.Force {
			log.Fatalf("文件已存在，需要覆盖请使用 -f 参数强制覆盖原有文件")
		}

		// err = app.Storage.PutFromFile(putCmdArgs.Key, putCmdArgs.Path)
		r := NewReaderWithProgress(putCmdArgs.Path, fmt.Sprintf("\"%s\" -> \"%s\"", putCmdArgs.Path, putCmdArgs.Key))
		err = app.Storage.Put(putCmdArgs.Key, r)
		if err != nil {
			log.Fatal(err)
		}

		// color.Green(fmt.Sprintf("上传成功！\"%s\" -> \"%s\"", putCmdArgs.Path, putCmdArgs.Key))
	},
}

func init() {
	putCmd.PersistentFlags().StringP("key", "k", "", "保存到 oss 的 key")
	putCmd.PersistentFlags().BoolP("force", "f", false, "如果文件已经存在，是否覆盖，默认否")
}
