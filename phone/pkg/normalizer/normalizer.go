package normalizer

import (
	"bytes"
	"regexp"
)

func Nomarlize(phone string) string {
	var buff bytes.Buffer
	for _, ch := range phone {
		if ch >= '0' && ch <= '9' {
			buff.WriteRune(ch)
		}
	}
	return buff.String()
}

func NormalizeRegex(phone string) string {
	re := regexp.MustCompile("[^0-9]")
	return re.ReplaceAllString(phone, "")
}
