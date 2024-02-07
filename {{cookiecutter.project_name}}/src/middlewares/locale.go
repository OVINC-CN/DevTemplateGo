package middlewares

import (
	"encoding/json"
	"fmt"
	"github.com/OVINC-CN/DevTemplateGo/src/utils"
	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
	"path/filepath"
	"runtime"
)

func Locale() gin.HandlerFunc {
	_, currentPath, _, ok := runtime.Caller(0)
	if !ok {
		utils.Logger.Warningf("[LoadLocaleFileFailed] %T", ok)
		panic("LoadLocaleFileFailed")
	}
	return ginI18n.Localize(
		ginI18n.WithBundle(&ginI18n.BundleCfg{
			RootPath:         fmt.Sprintf("%s/locale", filepath.Dir(filepath.Dir(currentPath))),
			AcceptLanguage:   []language.Tag{language.Chinese, language.English},
			DefaultLanguage:  language.Chinese,
			UnmarshalFunc:    json.Unmarshal,
			FormatBundleFile: "json",
		}),
	)
}
