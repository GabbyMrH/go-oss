package handlers

import (
	"crypto/md5"
	"encoding/hex"
)

// md5散列(32位)
func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}
