package util

import "testing"

func TestCalSearchRadius(t *testing.T) {
	suggestSearchLevels := []int{5000, 10000, 20000, 50000, 100000}
	if val := CalSearchRadius(1000,
		0,
		suggestSearchLevels); val != 5000 {
		t.Errorf("not get expected val:%d realVal:%d", val, 5000)
	}

	if val := CalSearchRadius(1000,
		5000,
		suggestSearchLevels); val != 10000{
		t.Errorf("not get expected val:%d realVal:%d", val, 10000)
	}

	if val := CalSearchRadius(1000,
		100000,
		suggestSearchLevels); val != 100000{
		t.Errorf("not get expected val:%d realVal:%d", val, 100000)
	}

	if val := CalSearchRadius(500000,
		0,
		suggestSearchLevels); val != 500000{
		t.Errorf("not get expected val:%d realVal:%d", val, 500000)
	}
}
