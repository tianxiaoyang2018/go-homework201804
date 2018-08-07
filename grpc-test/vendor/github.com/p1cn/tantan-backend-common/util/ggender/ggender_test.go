package ggender

import "testing"

func TestNew(t *testing.T) {
	gg, err := New("charfreq.csv")
	if err != nil {
		t.Fail()
	}
	t.Logf("%f,%f\n", gg.pTotal[0], gg.pTotal[1])
}

func TestGuest(t *testing.T) {
	gg, err := New("charfreq.csv")
	if err != nil {
		t.Fail()
	}
	names := []struct {
		name   string
		gender string
	}{
		{
			name:   "祖广乐",
			gender: "male",
		},
		{
			name:   "毛格格",
			gender: "female",
		},
		{
			name:   "李赛男",
			gender: "female",
		},
		{
			name:   "李胜男",
			gender: "female",
		},
		{
			name:   "李亚男",
			gender: "female",
		}, {
			name:   "杨利松",
			gender: "male",
		},
	}

	for _, v := range names {
		g, p := gg.Guess(v.name)
		if g != v.gender {
			t.Error(v, g, p)
		}
	}

}
