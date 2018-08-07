package trie

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/p1cn/tantan-backend-common/sensitivewords/matcher"
)

func assert(t *testing.T, b bool) {
	if !b {
		t.Fail()
	}
}

func newStringMatcher(strs []string) *TrieMatcher {
	dict := new(matcher.Dictionary)
	for _, str := range strs {
		dict.AddWord(str, 0)
	}
	return NewMatcher(dict)
}

func TestNoPatterns(t *testing.T) {
	m := newStringMatcher([]string{})
	hits := m.match(("foo bar baz"))
	assert(t, len(hits) == 0)
}

func TestNoData(t *testing.T) {
	m := newStringMatcher([]string{"foo", "baz", "bar"})
	hits := m.match((""))
	assert(t, len(hits) == 0)
}

func TestSuffixesWhiteList(t *testing.T) {
	//m := newStringMatcher([]string{"Superman", "uperman", "perman", "erman"})
	dict := new(matcher.Dictionary)
	dict.AddWord("Superman", 0)
	dict.AddWord("uperman", 0)
	dict.AddWord("perman", 0x20)
	dict.AddWord("erman", 0)
	dict.AddWord("Ste", 0)
	dict.AddWord("Stee", 0)
	dict.AddWord("Steel", 0x20)
	//dict.AddWord("Steels", 0)

	str := "The Man Of Steels: Superman"
	// "The Man Of Steels: ********"
	m := NewMatcher(dict)
	hits := m.match(str)

	assert(t, len(hits) == 2)
	assert(t, hits[0] == 0)
	assert(t, hits[1] == 1)

	ret, count := m.Replace(str, "*")
	assert(t, ret == "The Man Of Steels: ********")
	assert(t, count == 1)
}

func TestSuffixesWhiteList2(t *testing.T) {
	//m := newStringMatcher([]string{"Superman", "uperman", "perman", "erman"})
	dict := new(matcher.Dictionary)
	dict.AddWord("Superman", 0x20)
	dict.AddWord("uperman", 0)
	dict.AddWord("perman", 0x20)
	dict.AddWord("erman", 0)
	dict.AddWord("Ste", 0)
	dict.AddWord("Stee", 0x20)
	dict.AddWord("Steel", 0)
	//dict.AddWord("Steels", 0)

	str := "The Man Of Steels: Superman"
	// "The Man Of Steels: ********"
	m := NewMatcher(dict)
	hits := m.match(str)

	assert(t, len(hits) == 2)
	assert(t, hits[0] == 4)
	assert(t, hits[1] == 6)

	ret, count := m.Replace(str, "*")
	fmt.Println(ret)
	assert(t, ret == "The Man Of ***els: Superman")
	assert(t, count == 1)
}

func TestSuffixesReplaceWhiteList(t *testing.T) {
	dict := new(matcher.Dictionary)
	dict.AddWord("Superman", 0)
	dict.AddWord("uperman", 0)
	dict.AddWord("perman", 0x20)
	dict.AddWord("erman", 0)
	dict.AddWord("Ste", 0)
	dict.AddWord("Stee", 0x20)
	dict.AddWord("Steel", 0x0)
	//dict.AddWord("Steels", 0)
	m := NewMatcher(dict)

	str := "The Man Of Steels: Superman"

	ret, count := m.Replace(str, "*")
	assert(t, ret == "The Man Of ***els: ********")
	assert(t, count == 2)
}

func TestSuffixesReplaceWhiteListChinese(t *testing.T) {
	dict := new(matcher.Dictionary)
	dict.AddWord("中国", 0)
	dict.AddWord("中华", 0)
	dict.AddWord("中华民国", 0x20)
	dict.AddWord("中华人民共和国", 0)
	dict.AddWord("共和国", 0x20)
	dict.AddWord("和国", 0x00)
	dict.AddWord("Steel", 0x0)
	dict.AddWord("戴黑框眼镜", 0x20)
	dict.AddWord("黑框眼镜", 0x0)

	//dict.AddWord("Steels", 0)
	m := NewMatcher(dict)

	str := "新中华人民共和国成立了，那个时候他戴黑框眼镜站在那里"

	ms := m.match(str)
	assert(t, len(ms) == 2)
	assert(t, ms[0] == 1)
	assert(t, ms[1] == 3)
	ret, count := m.Replace(str, "*")
	assert(t, count == 1)
	assert(t, ret == "新**人民共和国成立了，那个时候他戴黑框眼镜站在那里")

}
func TestSuffixesReplaceWhiteListChinesePrefix(t *testing.T) {
	dict := new(matcher.Dictionary)
	dict.AddWord("黑框眼镜", 0)
	dict.AddWord("戴黑框眼镜", 0x20)
	m := NewMatcher(dict)

	str := "是我戴黑框眼镜"
	ret := m.Match(str)
	assert(t, ret == false)

	dict.AddWord("我戴黑框眼镜", 0)
	m = NewMatcher(dict)
	ret = m.Match(str)
	assert(t, ret == true)
}

func TestSuffixesTraditionalChinese(t *testing.T) {
	dict := new(matcher.Dictionary)
	dict.AddWord("中國人", 0)
	dict.AddWord("小東", 0x0)
	dict.AddWord("華文明", 0x0)
	dict.AddWord("華文明lala", 0x20)

	m := NewMatcher(dict)

	str := "我是z中國中国人，小东华文明lala華文明"
	ret := m.match(str)

	assert(t, len(ret) == 3)
	assert(t, ret[0] == 0)
	assert(t, ret[1] == 1)
	assert(t, ret[2] == 2)

	reStr, count := m.Replace(str, "*")
	assert(t, count == 3)
	assert(t, reStr == "我是z中國***，**华文明lala***")

}

func TestSuffixes(t *testing.T) {
	m := newStringMatcher([]string{"Superman", "uperman", "perman", "erman"})
	hits := m.match(("The Man Of Steel: Superman"))
	assert(t, len(hits) == 4)
	assert(t, hits[0] == 0)
	assert(t, hits[1] == 1)
	assert(t, hits[2] == 2)
	assert(t, hits[3] == 3)
}

func TestPrefixes(t *testing.T) {
	m := newStringMatcher([]string{"Superman", "Superma", "Superm", "Super"})
	hits := m.match(("The Man Of Steel: Superman"))
	fmt.Println(hits)
	assert(t, len(hits) == 4)
	assert(t, hits[0] == 3)
	assert(t, hits[1] == 2)
	assert(t, hits[2] == 1)
	assert(t, hits[3] == 0)
}

func TestInterior(t *testing.T) {
	m := newStringMatcher([]string{"Steel", "tee", "e"})
	hits := m.match(("The Man Of Steel: Superman"))
	assert(t, len(hits) == 6)
	//fmt.Println(hits)
	assert(t, hits[2] == 1)
	assert(t, hits[1] == 0)
	assert(t, hits[0] == 2)
}

func TestMatchAtStart(t *testing.T) {
	m := newStringMatcher([]string{"The", "Th", "he"})
	hits := m.match(("The Man Of Steel: Superman"))
	assert(t, len(hits) == 3)
	assert(t, hits[0] == 1)
	assert(t, hits[1] == 0)
	assert(t, hits[2] == 2)
}

func TestMatchAtEnd(t *testing.T) {
	m := newStringMatcher([]string{"teel", "eel", "el"})
	hits := m.match(("The Man Of Steel"))
	assert(t, len(hits) == 3)
	assert(t, hits[0] == 0)
	assert(t, hits[1] == 1)
	assert(t, hits[2] == 2)
}

func TestOverlappingPatterns(t *testing.T) {
	m := newStringMatcher([]string{"Man ", "n Of", "Of S"})
	hits := m.match(("The Man Of Steel"))
	assert(t, len(hits) == 3)
	assert(t, hits[0] == 0)
	assert(t, hits[1] == 1)
	assert(t, hits[2] == 2)
}

func TestMultipleMatches(t *testing.T) {
	m := newStringMatcher([]string{"The", "Man", "an"})
	hits := m.match(("A Man A Plan A Canal: Panama, which Man Planned The Canal"))
	//fmt.Println(hits)
	//[1 2 2 2 2 1 2 2 0 2]
	assert(t, len(hits) == 10)
	assert(t, hits[0] == 1)
	assert(t, hits[1] == 2)
	assert(t, hits[2] == 2)
	assert(t, hits[5] == 1)

}

func TestSingleCharacterMatches(t *testing.T) {
	m := newStringMatcher([]string{"a", "M", "z"})
	hits := m.match(("A Man A Plan A Canal: Panama, which Man Planned The Canal"))
	assert(t, len(hits) == 13)
	assert(t, hits[0] == 1)
	assert(t, hits[1] == 0)
	assert(t, hits[8] == 1)

}

func TestNothingMatches(t *testing.T) {
	m := newStringMatcher([]string{"baz", "bar", "foo"})
	hits := m.match(("A Man A Plan A Canal: Panama, which Man Planned The Canal"))
	assert(t, len(hits) == 0)
}

func TestWikipedia(t *testing.T) {
	m := newStringMatcher([]string{"a", "ab", "bc", "bca", "c", "caa"})
	hits := m.match(("abccab"))
	//fmt.Println(hits)
	//[0 1 2 4 4 0 1]

	assert(t, len(hits) == 7)
	assert(t, hits[0] == 0)
	assert(t, hits[1] == 1)
	assert(t, hits[2] == 2)
	assert(t, hits[3] == 4)

	hits = m.match(("bccab"))
	//[2 4 4 0 1]
	//fmt.Println(hits)
	assert(t, len(hits) == 5)
	assert(t, hits[0] == 2)
	assert(t, hits[1] == 4)
	assert(t, hits[2] == 4)
	assert(t, hits[3] == 0)

	hits = m.match(("bccb"))
	//fmt.Println(hits)
	assert(t, len(hits) == 3)
	assert(t, hits[0] == 2)
	assert(t, hits[1] == 4)
	assert(t, hits[2] == 4)

}

func TestMatch(t *testing.T) {
	m := newStringMatcher([]string{"Mozilla", "Mac", "Macintosh", "Safari", "Sausage"})
	hits := m.match(("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/30.0.1599.101 Safari/537.36"))
	//fmt.Println(hits)
	//[0 1 2 1 3]
	assert(t, len(hits) == 5)
	assert(t, hits[0] == 0)
	assert(t, hits[1] == 1)
	assert(t, hits[2] == 2)
	assert(t, hits[3] == 1)
	assert(t, hits[4] == 3)

	hits = m.match(("Mozilla/5.0 (Mac; Intel Mac OS X 10_7_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/30.0.1599.101 Safari/537.36"))
	//fmt.Println(hits)
	//[0 1 1 3]
	assert(t, len(hits) == 4)
	assert(t, hits[0] == 0)
	assert(t, hits[1] == 1)
	assert(t, hits[2] == 1)
	assert(t, hits[3] == 3)

	hits = m.match(("Mozilla/5.0 (Moc; Intel Computer OS X 10_7_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/30.0.1599.101 Safari/537.36"))
	assert(t, len(hits) == 2)
	assert(t, hits[0] == 0)
	assert(t, hits[1] == 3)

	hits = m.match(("Mozilla/5.0 (Moc; Intel Computer OS X 10_7_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/30.0.1599.101 Sofari/537.36"))
	assert(t, len(hits) == 1)
	assert(t, hits[0] == 0)

	hits = m.match(("Mazilla/5.0 (Moc; Intel Computer OS X 10_7_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/30.0.1599.101 Sofari/537.36"))
	assert(t, len(hits) == 0)
}

var bytes = []byte("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/30.0.1599.101 Safari/537.36")
var sbytes = string(bytes)
var dictionary = []string{"Mozilla", "Mac", "Macintosh", "Safari", "Sausage"}
var precomputed = newStringMatcher(dictionary)

func BenchmarkMatchWorks(b *testing.B) {
	for i := 0; i < b.N; i++ {
		precomputed.Match(string(bytes))
	}
}

func BenchmarkContainsWorks(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hits := make([]int, 0)
		for i, s := range dictionary {
			if strings.Contains(sbytes, s) {
				hits = append(hits, i)
			}
		}
	}
}

var re = regexp.MustCompile("(" + strings.Join(dictionary, "|") + ")")

func BenchmarkRegexpWorks(b *testing.B) {
	for i := 0; i < b.N; i++ {
		re.FindAllIndex(bytes, -1)
	}
}

var dictionary2 = []string{"Googlebot", "bingbot", "msnbot", "Yandex", "Baiduspider"}
var precomputed2 = newStringMatcher(dictionary2)

func BenchmarkMatchFails(b *testing.B) {
	for i := 0; i < b.N; i++ {
		precomputed2.Match(string(bytes))
	}
}

func BenchmarkContainsFails(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hits := make([]int, 0)
		for i, s := range dictionary2 {
			if strings.Contains(sbytes, s) {
				hits = append(hits, i)
			}
		}
	}
}

var re2 = regexp.MustCompile("(" + strings.Join(dictionary2, "|") + ")")

func BenchmarkRegexpFails(b *testing.B) {
	for i := 0; i < b.N; i++ {
		re2.FindAllIndex(bytes, -1)
	}
}

var bytes2 = []byte("Firefox is a web browser, and is Mozilla's flagship software product. It is available in both desktop and mobile versions. Firefox uses the Gecko layout engine to render web pages, which implements current and anticipated web standards. As of April 2013, Firefox has approximately 20% of worldwide usage share of web browsers, making it the third most-used web browser. Firefox began as an experimental branch of the Mozilla codebase by Dave Hyatt, Joe Hewitt and Blake Ross. They believed the commercial requirements of Netscape's sponsorship and developer-driven feature creep compromised the utility of the Mozilla browser. To combat what they saw as the Mozilla Suite's software bloat, they created a stand-alone browser, with which they intended to replace the Mozilla Suite. Firefox was originally named Phoenix but the name was changed so as to avoid trademark conflicts with Phoenix Technologies. The initially-announced replacement, Firebird, provoked objections from the Firebird project community. The current name, Firefox, was chosen on February 9, 2004.")
var sbytes2 = string(bytes2)

var dictionary3 = []string{"Mozilla", "Mac", "Macintosh", "Safari", "Phoenix"}
var precomputed3 = newStringMatcher(dictionary3)

func BenchmarkLongMatchWorks(b *testing.B) {
	for i := 0; i < b.N; i++ {
		precomputed3.Match(string(bytes2))
	}
}

func BenchmarkLongContainsWorks(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hits := make([]int, 0)
		for i, s := range dictionary3 {
			if strings.Contains(sbytes2, s) {
				hits = append(hits, i)
			}
		}
	}
}

var re3 = regexp.MustCompile("(" + strings.Join(dictionary3, "|") + ")")

func BenchmarkLongRegexpWorks(b *testing.B) {
	for i := 0; i < b.N; i++ {
		re3.FindAllIndex(bytes2, -1)
	}
}

var dictionary4 = []string{"12343453", "34353", "234234523", "324234", "33333"}
var precomputed4 = newStringMatcher(dictionary4)

func BenchmarkLongMatchFails(b *testing.B) {
	for i := 0; i < b.N; i++ {
		precomputed4.Match(string(bytes2))
	}
}

func BenchmarkLongContainsFails(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hits := make([]int, 0)
		for i, s := range dictionary4 {
			if strings.Contains(sbytes2, s) {
				hits = append(hits, i)
			}
		}
	}
}

var re4 = regexp.MustCompile("(" + strings.Join(dictionary4, "|") + ")")

func BenchmarkLongRegexpFails(b *testing.B) {
	for i := 0; i < b.N; i++ {
		re4.FindAllIndex(bytes2, -1)
	}
}

var dictionary5 = []string{"12343453", "34353", "234234523", "324234", "33333", "experimental", "branch", "of", "the", "Mozilla", "codebase", "by", "Dave", "Hyatt", "Joe", "Hewitt", "and", "Blake", "Ross", "mother", "frequently", "performed", "in", "concerts", "around", "the", "village", "uses", "the", "Gecko", "layout", "engine"}
var precomputed5 = newStringMatcher(dictionary5)

func BenchmarkMatchMany(b *testing.B) {
	for i := 0; i < b.N; i++ {
		precomputed5.Match(string(bytes))
	}
}

func BenchmarkContainsMany(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hits := make([]int, 0)
		for i, s := range dictionary4 {
			if strings.Contains(sbytes, s) {
				hits = append(hits, i)
			}
		}
	}
}

var re5 = regexp.MustCompile("(" + strings.Join(dictionary5, "|") + ")")

func BenchmarkRegexpMany(b *testing.B) {
	for i := 0; i < b.N; i++ {
		re5.FindAllIndex(bytes, -1)
	}
}

func BenchmarkLongMatchMany(b *testing.B) {
	for i := 0; i < b.N; i++ {
		precomputed5.Match(string(bytes2))
	}
}

func BenchmarkLongContainsMany(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hits := make([]int, 0)
		for i, s := range dictionary4 {
			if strings.Contains(sbytes2, s) {
				hits = append(hits, i)
			}
		}
	}
}

func BenchmarkLongRegexpMany(b *testing.B) {
	for i := 0; i < b.N; i++ {
		re5.FindAllIndex(bytes2, -1)
	}
}
