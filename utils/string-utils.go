package utils
 
import (
	"math/rand"
	"time"
)
 
func RandomString(n int) string {
    var charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
    s := make([]byte, n)
    for i := range s {
        s[i] = charset[seededRand.Intn(len(charset))]
    }
    return string(s)
}