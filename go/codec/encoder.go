package codec

import (
	"strconv"
)

type Base string

const (
	Base36 Base = "abcdefghijklmnopqrstuvwxyz0123456789"
	Base59 Base = "abcdefghijkmnopqrstuvwxyzABCDEFGHIJKLMNPQRSTUVWXYZ123456789"
	Base62 Base = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var alphabetConfig = Base62

func Encode(num string) string {
	n, _ := strconv.ParseUint(num, 10, 64)
	t := make([]byte, 0)

	/* Special case */
	if n == 0 {
		return string(alphabetConfig[0])
	}

	/* Map */
	for n > 0 {
		r := n % uint64(len(alphabetConfig))
		t = append(t, alphabetConfig[r])
		n = n / uint64(len(alphabetConfig))
	}

	/* Reverse */
	for i, j := 0, len(t)-1; i < j; i, j = i+1, j-1 {
		t[i], t[j] = t[j], t[i]
	}

	return string(t)
}
