package random

import (
	"fmt"
	"math/rand"
	"strings"
)

const NumericRunes = "0123456789"
const LowerAlphaRunes = "abcdefghijklmnopqrstuvwxyz"
const UpperAlphaRunes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const AlphaRunes = UpperAlphaRunes + LowerAlphaRunes

// Return a random string made up of characters passed
func Runes(prefix string, length int, runes ...string) string {
	chars := strings.Join(runes, "")
	var bytes = make([]byte, length)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = chars[b%byte(len(chars))]
	}
	return prefix + string(bytes)
}

// Return a random string of alpha characters
func Alpha(prefix string, length int) string {
	return Runes(prefix, length, AlphaRunes)
}

// Return a random string of alpha and numeric characters
func String(prefix string, length int) string {
	return Runes(prefix, length, AlphaRunes, NumericRunes)
}

// Given a list of strings, return one of the strings randomly
func Item(items ...string) string {
	var bytes = make([]byte, 1)
	rand.Read(bytes)
	return items[bytes[0]%byte(len(items))]
}

// Return a random domain name in the form "randomAlpha.net"
func DomainName() string {
	return fmt.Sprintf("%s.%s",
		Alpha("", 14),
		Item("net", "com", "org", "io", "gov"))
}
