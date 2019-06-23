package main

import (
	"fmt"
	"strconv"
)

const base91CharTable = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!#$%&()*+,./:;<=>?@[]^_`{|}~\""

// Encode encodes
func Encode(d []byte) []byte {
	encoded := ""
	bits := reverseString(stringToBin(string(d)))
	index := 0
	for {
		done := false
		var s string
		rem := len(bits) - (index + 13)
		if rem > 13 {
			s = bits[index:13]
			index += 13
		} else {
			s = bits[index:]
			done = true
		}
		v := bitStringToInt(s)
		i0 := v % 91
		i1 := v / 91
		encoded = fmt.Sprintf("%s%s%s", encoded, string(base91CharTable[i0]), string(base91CharTable[i1]))
		_ = v
		if done {
			break
		}
	}

	return []byte(encoded)
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Decode decodes
func Decode(d []byte) []byte {
	return nil
}

func stringToBin(s string) (binString string) {
	for _, c := range s {
		binString = fmt.Sprintf("%s%b", binString, c)
	}
	return
}

func bitStringToInt(s string) int64 {
	i, _ := strconv.ParseInt(s, 2, 64)
	return i
}

// EncodeToString decodes a given byte array are returns a string
func EncodeToString(d []byte) string {
	return string(Encode(d))
}

// DecodeToString decodes a given byte array are returns a string
func DecodeToString(d []byte) string {
	return string(Decode(d))
}
