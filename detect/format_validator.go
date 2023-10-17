package detect

import (
	"strings"

	"github.com/gabriel-vasile/mimetype"

	"github.com/zricethezav/gitleaks/v8/config"
)

type FormatValidator struct {
	WhiteList   []string
	ScipList    []string
	PathProject string
}

func NewFormatValidator(configFile string) *FormatValidator {
	mimetype.SetLimit(1024 * 1024)
	conf, _ := config.NewConfigLists(configFile)
	return &FormatValidator{
		WhiteList: config.GetList(conf.WhiteList),
		ScipList:  config.GetList(conf.ScipList),
	}
}

func (fv *FormatValidator) IsValidFormat(mimeType *mimetype.MIME) bool {
	if mimeType == nil {
		return false
	}
	for _, validMemType := range fv.WhiteList {
		if strings.HasPrefix(mimeType.String(), validMemType) {
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
		if strings.HasSuffix(path, scipFile) {
			return true
		}
	}

	return false
}
