package util

import (
	"regexp"
	"strings"
	"testing"
)

func TestSanitizeHyperLinkTags(t *testing.T) {
	data := []struct {
		text      string
		sanitized string
	}{
		{`恭喜你成功完成学生身份认证`, `恭喜你成功完成学生身份认证`},
		{``, ``},
		{`抱歉，你的资料没通过审核，<a href = "tantan://verification/school/rejected">戳我重新提交资料。</a>`, `抱歉，你的资料没通过审核，戳我重新提交资料。`},
		{`<a href = "tantan://verification/school/rejected">戳我重新提交资料。</a>`, `戳我重新提交资料。`},
		{`<a href = "tantan://verification/school/rejected" > 戳我重新提交资料</a>。`, ` 戳我重新提交资料。`},
		{`<a  href = "tantan://verification/school/rejected" >戳我重新提交资料。</a>`, `戳我重新提交资料。`},
		{`<a href = "tantan://verification/school/rejected">戳我</a>重新<a href = "tantan://verification/school/rejected">提交资料。</a>`, `戳我重新提交资料。`},
	}

	for _, v := range data {
		if SanitizeHyperLinkTags(v.text) != v.sanitized {
			t.Errorf("text: %s, expected sanitization: %s, actual: %s", v.text, v.sanitized, SanitizeHyperLinkTags(v.text))
		}
	}
}

func BenchmarkSanitizeHyperLinkTags(b *testing.B) {
	str := `抱歉，你的资料没通过审核，<a href = "tantan://verification/school/rejected">戳我重新提交资料。</a>`
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		SanitizeHyperLinkTags(str)
	}
}

func BenchmarkSanitizeHyperLinkTagsWithoutCompiledRegexp(b *testing.B) {
	str := `抱歉，你的资料没通过审核，<a href = "tantan://verification/school/rejected">戳我重新提交资料。</a>`
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		re := regexp.MustCompile(`<a\s+.*?>(.*?)</a>`)
		re.ReplaceAllString(str, "$1")
	}
}

func BenchmarkSanitizeHyperLinkTagsWihtoutStringsContainsDetection(b *testing.B) {
	str := `抱歉，你的资料没通过审核，<a href = "tantan://verification/school/rejected">戳我重新提交资料。</a>`
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		hyperLinkTagsRegexp.ReplaceAllString(str, "$1")
	}
}

func BenchmarkSanitizeHyperLinkTagsPlainTextWithStringsContains(b *testing.B) {
	str := `抱歉，你的资料没通过审核，戳我重新提交资料。`
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if strings.Contains(str, "</a>") {
			hyperLinkTagsRegexp.ReplaceAllString(str, "$1")
		}
	}
}

func BenchmarkSanitizeHyperLinkTagsPlainTextWithoutStringsContains(b *testing.B) {
	str := `抱歉，你的资料没通过审核，戳我重新提交资料。`
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		hyperLinkTagsRegexp.ReplaceAllString(str, "$1")
	}
}
