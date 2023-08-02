package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"github.com/spf13/cobra"

	"github.com/lixoi/prjct_validator/config"
)

var configFile string

type FormatValidator struct {
	WhiteList   []string
	ScipList    []string
	PathProject string
	MimeType    *mimetype.MIME
}

func NewFormatValidator() *FormatValidator {
	mimetype.SetLimit(1024 * 1024)
	conf, _ := config.NewConfig(configFile)
	return &FormatValidator{
		WhiteList:   config.GetList(conf.WhiteList),
		ScipList:    config.GetList(conf.ScipList),
		PathProject: conf.Project,
	}
}

func (fv *FormatValidator) SetContextType(fpath string) error {
	data, err := ioutil.ReadFile(fpath)
	if err != nil {
		return err
	}

	fv.MimeType, err = mimetype.DetectReader(bytes.NewReader(data))

	return err
}

func (fv *FormatValidator) IsValidFormat() bool {
	if fv.MimeType == nil {
		return false
	}
	for _, validMemType := range fv.WhiteList {
		if strings.HasPrefix(fv.MimeType.String(), validMemType) {
			return true
		}
	}

	return false
}

func (fv *FormatValidator) IsScip(path string) bool {
	if path == "" {
		return false
	}
	for _, scipFile := range fv.ScipList {
		if strings.HasPrefix(path, scipFile) {
			return true
		}
	}

	return false
}

var StartValidator = &cobra.Command{
	Use:   "cmd",
	Short: "Run project validator",
	Run: func(cmd *cobra.Command, args []string) {
		fv := NewFormatValidator()
		err := filepath.Walk(fv.PathProject, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			if info.IsDir() || fv.IsScip(path) {
				return nil
			}

			if err := fv.SetContextType(path); err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			if !fv.IsValidFormat() {
				fmt.Printf("invalid format for file %s\n", path)
				os.Exit(1)
			}

			return nil
		})
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	StartValidator.Flags().StringVar(&configFile, "configFile", "config.json", "path to configuration file for Format Validator")
}
