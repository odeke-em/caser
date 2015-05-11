package caser

import (
	"fmt"
)

type SyllableType int

type Encoding []uint64

const (
	Consonant SyllableType = iota
	Vowel
)

var (
	MaxIntMask uint64 = ^uint64(0)
)

func popCount64(i uint64) (count uint) {
	count = 0
	for {
		if i < 1 {
			break
		}

		if (i & 1) == 1 {
			count += 1
		}

		i >>= 1
	}
	return count
}

var vowels = map[uint8]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
}

func (e Encoding) String() string {
	repr := ""
	for _, v := range e {
		repr = fmt.Sprintf("%b%s", v, repr)
	}
	return repr
}

func vowely(r uint8) bool {
	_, ok := vowels[r]
	return ok
}

func consonanty(r uint8) bool {
	return !vowely(r)
}

func byTypeEncoder(s string) (encoded Encoding) {
	sLen := len(s)

	maskInit := ^(MaxIntMask >> 1)
	shaver := MaxIntMask >> 1

	maxPopCount := popCount64(MaxIntMask)
	mask := maskInit

	i := 0
	bitIndex := uint(0)

	for {
		done := i >= sLen

		if done || bitIndex >= maxPopCount {
			encoded = append(encoded, mask&shaver)
			mask = maskInit
			bitIndex = 0
		}

		if done {
			break
		}

		masked := Consonant
		if vowely(s[i]) {
			masked = Vowel
		}

		mask |= (uint64(masked) << bitIndex)
		i += 1
		bitIndex += 1
	}

	return
}
