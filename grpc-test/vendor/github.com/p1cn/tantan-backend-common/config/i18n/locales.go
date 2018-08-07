package i18n

import (
	"fmt"
	"os"
	"path"

	gotext "gopkg.in/leonelquinteros/gotext.v1"

	"github.com/p1cn/tantan-backend-common/log"
	"github.com/p1cn/tantan-backend-common/util"
)

const LocaleDomain = `backend`

var locales = &Locales{}

// TODO: replace Locales definition with below commented version after
// solving the performance issues in using Locale struct from gotext.v1 repo.
// (see locales_test.go for details)

// initialized only once on starting app
type Locales struct {
	m map[string]*gotext.Po
	//sync.RWMutex
}

func (ls *Locales) Add(lang string, p *gotext.Po) {
	if ls.m == nil {
		ls.m = make(map[string]*gotext.Po)
	}
	ls.m[lang] = p
}

func (ls *Locales) Get(lang string) (*gotext.Po, error) {
	if ls.m == nil {
		err := fmt.Errorf("locales not initialized")
		log.Err("%v", err)
		return nil, err
	}
	p, ok := ls.m[lang]
	if !ok {
		err := fmt.Errorf("locale not found for language: %s", lang)
		log.Err("%v", err)
		return nil, err
	}
	return p, nil
}

func AddLocales(dir string) error {
	for _, lang := range util.TranslatedLanguages {
		po := path.Join(dir, lang, "LC_MESSAGES", LocaleDomain+".po")
		if _, err := os.Stat(po); os.IsNotExist(err) {
			return fmt.Errorf("po file not found: %s", po)
		}
		addLocale(po, lang)
	}
	return nil
}

func addLocale(poFile, lang string) {
	po := new(gotext.Po)
	po.ParseFile(poFile)
	locales.Add(lang, po)
}

/*
type Locales struct {
	m map[string]*gotext.Locale
	//sync.RWMutex
}

func (ls *Locales) Add(lang string, l *gotext.Locale) {
	if ls.m == nil {
		ls.m = make(map[string]*gotext.Locale)
	}
	ls.m[lang] = l
}

func (ls *Locales) Get(lang string) (*gotext.Locale, error) {
	if ls.m == nil {
		err := fmt.Errorf("locales not initialized")
		slog.Err("%v", err)
		return nil, err
	}
	l, ok := ls.m[lang]
	if !ok {
		err := fmt.Errorf("locale not found for language: %s", lang)
		slog.Err("%v", err)
		return nil, err
	}
	return l, nil
}

func AddLocales(dir string) error {
	for _, lang := range util.TranslatedLanguages {
		po := path.Join(dir, lang, "LC_MESSAGES", LocaleDomain+".po")
		if _, err := os.Stat(po); os.IsNotExist(err) {
			return fmt.Errorf("po file not found: %s", po)
		}
		addLocale(dir, lang)
	}
	return nil
}

func addLocale(dir, lang string) {
	l := gotext.NewLocale(dir, lang)
	l.AddDomain(LocaleDomain)
	locales.Add(lang, l)
}
*/
