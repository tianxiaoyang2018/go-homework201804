package sensitivewords

import (
	"sync"

	slog "github.com/p1cn/tantan-backend-common/log"
	"github.com/p1cn/tantan-backend-common/sensitivewords/matcher"
	"github.com/p1cn/tantan-backend-common/sensitivewords/trie"
)

// @todo
// 1. ac for less words ; trie for huge words
var (
	gMatcher matcher.Matcher
	initOnce sync.Once
)

type Config struct {
	Dict *matcher.Dictionary
}

// initialize
func InitSensitiveWords(cfg Config) error {
	gMatcher = trie.NewMatcher(cfg.Dict)

	return nil
}

// match
func Match(str string) bool {
	if gMatcher == nil {
		slog.Warning("matcher is nil")
		return false
	}
	return gMatcher.Match(str)
}

// replace
func Replace(str string, cover string) (string, int) {
	if gMatcher == nil {
		slog.Warning("matcher is nil")
		return str, 0
	}
	return gMatcher.Replace(str, cover)
}
