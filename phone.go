package gomini

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	encryptRegPhone = regexp.MustCompile("([0-9]{3})[0-9]{4}([0-9]{4})")
	regPhone = regexp.MustCompile(`[+]*[\d]+`)
	regPhoneNumber = regexp.MustCompile(`\\u202a|\\u202b|\\u202c|\\u202d|\\u202e| |\\n|\\r\\n|\?|\*`)
)

func EncryptPhone(phone string) string {
	return encryptRegPhone.ReplaceAllString(phone, "$1****$2")
}

func FormatPhone(mobile string) string  {
	allString := regPhone.FindAllString(mobile, -1)
	return strings.Join(allString , "")
}
// U+202A:   LEFT-TO-RIGHT EMBEDDING (LRE)
// U+202B:   RIGHT-TO-LEFT EMBEDDING (RLE)
// U+202D:   LEFT-TO-RIGHT OVERRIDE (LRO)
// U+202E:   RIGHT-TO-LEFT OVERRIDE (RLO)
// U+202C:   POP DIRECTIONAL FORMATTING (PDF)
// U+202C:   POP DIRECTIONAL FORMATTING (PDF)
func FormatPhoneNumber(phone string) string {
	textQuoted := strconv.QuoteToASCII(phone)
	textUnquoted := textQuoted[1 : len(textQuoted)-1]
	return regPhoneNumber.ReplaceAllString(textUnquoted,"")
}