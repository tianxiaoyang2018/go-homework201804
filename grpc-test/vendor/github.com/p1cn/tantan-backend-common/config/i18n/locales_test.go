package i18n

import (
	"testing"

	gotext "gopkg.in/leonelquinteros/gotext.v1"
)

// Benchmarking results:
// $ go test -bench="Benchmark" locales_test.go
// BenchmarkLocaleGetForEnglish-8	    100000	     15814 ns/op
// BenchmarkPoGetForEnglish-8		   5000000	       310 ns/op
// BenchmarkLocaleGetForChinese-8	    200000	     10057 ns/op
// BenchmarkPoGetForChinese-8		   5000000	       315 ns/op
//
// The reason why using locale object directly is tremendously slow is extra parsing
// for "(N)th plural form" caused by:
// locale.GetD(...) calls locale.GetND(...) with plural form index 1 by default

// Solutions:
// 1) Modify Locales struct to use "po" files now.
// 2) Will submit issues and create pull request for gotext.v1 github project afterwards.

const (
	benchLocalesDir    = "../locales"
	benchLocalesDomain = "backend"
	benchMessageId     = "TEAM_ACCOUNT_NAME"
	benchLangEn        = "en-US"
	benchLangZh        = "zh-CN"
)

// English has two plural forms.
func BenchmarkLocaleGetForEnglish(b *testing.B) {
	locale := gotext.NewLocale(benchLocalesDir, benchLangEn)
	locale.AddDomain(benchLocalesDomain)

	for i := 0; i < b.N; i++ {
		_ = locale.GetD(benchLocalesDomain, benchMessageId)
	}
}

func BenchmarkPoGetForEnglish(b *testing.B) {
	po := new(gotext.Po)
	po.ParseFile(benchLocalesDir + "/" + benchLangEn + "/LC_MESSAGES/" + benchLocalesDomain + ".po")

	for i := 0; i < b.N; i++ {
		_ = po.Get(benchMessageId)
	}
}

// Chinse has only one singular form.
func BenchmarkLocaleGetForChinese(b *testing.B) {
	locale := gotext.NewLocale(benchLocalesDir, benchLangZh)
	locale.AddDomain(benchLocalesDomain)

	for i := 0; i < b.N; i++ {
		_ = locale.GetD(benchLocalesDomain, benchMessageId)
	}
}

func BenchmarkPoGetForChinese(b *testing.B) {
	po := new(gotext.Po)
	po.ParseFile(benchLocalesDir + "/" + benchLangZh + "/LC_MESSAGES/" + benchLocalesDomain + ".po")

	for i := 0; i < b.N; i++ {
		_ = po.Get(benchMessageId)
	}
}
