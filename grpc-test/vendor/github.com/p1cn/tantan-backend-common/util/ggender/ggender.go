package ggender

import (
	"bufio"
	"encoding/csv"
	"errors"
	"os"
	"strconv"

	"github.com/p1cn/tantan-backend-common/util"
)

type GGender struct {
	pTotal [2]float64            // 0: male, 1: female
	freq   map[string][2]float64 // 0: male, 1: female
}

// New is to init a ggender
func New(charFreq string) (*GGender, error) {
	f, err := os.Open(charFreq)
	if err != nil {
		return nil, errors.New("char frequence file doesn't exist")
	}
	defer f.Close()
	r := csv.NewReader(bufio.NewReader(f))
	lines, err := r.ReadAll()
	if err != nil {
		return nil, errors.New("read csv failed")
	}
	gg := &GGender{
		freq: make(map[string][2]float64),
	}
	var mtotal, ftotal int
	for _, l := range lines[1:] {
		male, _ := strconv.Atoi(l[1])
		female, _ := strconv.Atoi(l[2])
		mtotal += male
		ftotal += female
	}
	gg.pTotal = [2]float64{float64(mtotal) / float64(ftotal+mtotal), float64(ftotal) / float64(ftotal+mtotal)}
	for _, l := range lines[1:] {
		name := l[0]
		male, _ := strconv.Atoi(l[1])
		female, _ := strconv.Atoi(l[2])
		gg.freq[name] = [2]float64{float64(male) / float64(mtotal), float64(female) / float64(ftotal)}
	}
	return gg, nil
}

func (gg *GGender) Guess(firstname string) (string, float64) {
	if !util.IsChinese(firstname) {
		return "both", 0
	}
	pm := gg.probeGender(firstname, 0)
	pf := gg.probeGender(firstname, 1)
	switch {
	case pm > pf:
		return "male", pm / (pm + pf)
	case pm < pf:
		return "female", pf / (pm + pf)
	default:
		return "both", 0
	}
}

func (gg *GGender) probeGender(firstname string, gender int) float64 {
	p := gg.pTotal[gender]
	fn := []rune(firstname)[1:]
	for _, v := range fn {
		f, has := gg.freq[string(v)]
		if has {
			p *= f[gender]

		}
	}
	return p
}
