package i18n

import (
	"testing"

	gotext "gopkg.in/leonelquinteros/gotext.v1"

	"github.com/p1cn/tantan-backend-common/util"
)

const (
	testLocalesDir    = "../locales"
	testLocalesDomain = "backend"
)

func TestGetAllMessages(t *testing.T) {
	for _, msg := range msgLists {
		for _, lang := range util.TranslatedLanguages {
			po := new(gotext.Po)
			po.ParseFile(testLocalesDir + "/" + lang + "/LC_MESSAGES/" + testLocalesDomain + ".po")
			text := po.Get(msg.id)
			if text == "" || text == msg.id {
				t.Errorf("expected to have translated text for message: %s, but got: %s", msg.id, text)
			}
		}
	}
}
