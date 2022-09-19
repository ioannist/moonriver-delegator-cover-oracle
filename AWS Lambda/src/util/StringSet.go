package util

import (
	"math/rand"
	"time"
	"unsafe"
)

type StringSet map[string]bool

var src = rand.NewSource(time.Now().UnixNano())

const letterBytes = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func NewStringSetFromSlice(slice []string) StringSet {
	s := make(StringSet, len(slice))
	for _, value := range slice {
		s[value] = true
	}
	return s
}

func (s StringSet) Difference(other StringSet) StringSet {
	result := make(StringSet)
	for value := range s {
		if !other[value] {
			result[value] = true
		}
	}
	return result
}

func (s StringSet) ToSlice() []string {
	result := make([]string, 0, len(s))
	for value := range s {
		result = append(result, value)
	}
	return result
}

func RandStringBytesMaskImprSrcUnsafe(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}
