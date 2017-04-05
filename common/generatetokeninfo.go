package common

import (
	"crypto/sha1"
	"fmt"
	"io"
	"sort"
	"strings"
)

//GenerateTokenInfo : func of generate token
func GenerateTokenInfo(UserName, UserPassword, timestamp string) string {
	sl := []string{UserName, UserPassword, timestamp}
	sort.Strings(sl)
	hashcode := sha1.New()
	io.WriteString(hashcode, strings.Join(sl, ""))
	return fmt.Sprintf("%x", hashcode.Sum(nil))
}
