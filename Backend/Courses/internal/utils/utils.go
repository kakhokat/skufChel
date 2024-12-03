package utils

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"time"
)

const salt = "wej,fnwekjfnlwjfnwelrjgnvdfjngwejnfakedalskmnedlaksdnkjsdbfkjanrkjavjksdbkjwdnfjklq"

func HashPass(pass string) string {
	hash := sha256.New()
	hash.Write([]byte(pass))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func Random(min, max int) int {
	rand.Seed(time.Now().Unix())
	if min > max {
		return min
	} else {
		return rand.Intn(max-min) + min
	}
}
