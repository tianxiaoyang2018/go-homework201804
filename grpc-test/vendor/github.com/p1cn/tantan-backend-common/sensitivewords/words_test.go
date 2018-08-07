package sensitivewords

import (
	"fmt"
	"testing"
	"time"

	"github.com/p1cn/tantan-backend-common/sensitivewords/matcher"
)

//  run consul as "consul agent -dev" first
func TestWords(t *testing.T) {
	test_setSensitiveWords(t)

	exist := Match("A Man A Plan A 使用的是列初始化, 并不实现白皮书中的行初始化. 列初始化但需要插入key时候, 需要重建DAT. 行初始化产生的冲突较多, 并且数组使用效率低. 中华名人共和国 共产中年Canal: Panama, w台湾政论区which Man Planned The Canal")
	fmt.Println(exist)
	if !exist {
		t.Fatal("error")
	}
	exist = Match("中华名人共和国 共产中年Canal: Panama, w台湾政论区which Man Planned The Canal")
	if exist {
		t.Fatal("error")
	}
}

func TestWordsUpdate(t *testing.T) {
	test_setSensitiveWords(t)

	exist := Match("A Man A Plan A 使用的是列初始化, 并不实现白皮书中的行初始化. 列初始化但需要插入key时候, 需要重建DAT. 行初始化产生的冲突较多, 并且数组使用效率低. 中华名人共和国 共产中年Canal: Panama, w台湾政论区which Man Planned The Canal")
	if !exist {
		t.Fatal("error")
	}
	test_clearSensitiveWords(t)
	time.Sleep(200 * time.Millisecond)
	exist = Match("A Man A Plan A 使用的是列初始化, 并不实现白皮书中的行初始化. 列初始化但需要插入key时候, 需要重建DAT. 行初始化产生的冲突较多, 并且数组使用效率低. 中华名人共和国 共产中年Canal: Panama, w台湾政论区which Man Planned The Canal")
	if exist {
		t.Fatal("error")
	}
}

func test_setSensitiveWords(t *testing.T) error {

	words := []string{"使用", "的是", "列初始化", "并不", "实现", "白皮书中的行初始化. 列初始化但需要"}

	dict := new(matcher.Dictionary)
	for _, w := range words {
		dict.AddWord(w, 0)
	}

	err := InitSensitiveWords(Config{
		Dict: dict,
	})
	if err != nil {
		t.Fatal(err)
	}
	return nil
}

func test_clearSensitiveWords(t *testing.T) error {

	err := InitSensitiveWords(Config{
		Dict: &matcher.Dictionary{},
	})
	if err != nil {
		t.Fatal(err)
	}
	return nil
}
