// flowlimiter 是一个限流的工具, 限流有很多方式，
// 例如某一个版本的 App 的发布节奏为 24h 分10次发布, 即每次会增加1/10的用户可以获取该版本, 这种方式我们称为 RatioLimit
// 又比如某一个版本的 App 的发布节奏为前24小时每分钟允许有至多1000人下载，24小时之后不做限制等等, 这种方式我们成为 QuantityLimit, 该方式目前没有实现
package flowlimiter

import (
	"math/rand"
	"time"
)

// 例如某一个版本的 App 的发布节奏为 24h 分10次发布, 即每次会增加1/10的用户可以获取该版本, 这种方式我们称为 RatioLimit
func RatioLimit(steps uint, seed int64, duration time.Duration, start time.Time) bool {
	cur := time.Since(start)
	if cur <= 0 {
		return true
	} else if cur >= duration {
		return false
	}
	if steps <= 1 {
		return false
	}
	rand.Seed(seed)
	num := rand.Intn(int(steps))
	if num <= int(float64(time.Duration(steps)*cur)/float64(duration)) {
		return false
	}
	return true
}
