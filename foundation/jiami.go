package main

import (
	"crypto/sha256"
	"io"
	"fmt"
)


func main()  {
	h := sha256.New()
	// 专家方案
	//dk := scrypt.Key([]byte("some password"), []byte(salt), 16384, 8, 1, 32)
	io.WriteString(h, "His money is twice tainted: 'taint yours and 'taint mine.")
	fmt.Printf("% x", h.Sum(nil))
}
//h := sha256.New()
//io.WriteString(h, "His money is twice tainted: 'taint yours and 'taint mine.")
//fmt.Printf("% x", h.Sum(nil))
//
////import "crypto/sha1"
//h := sha1.New()
//io.WriteString(h, "His money is twice tainted: 'taint yours and 'taint mine.")
//fmt.Printf("% x", h.Sum(nil))
//
////import "crypto/md5"
//h := md5.New()
//io.WriteString(h, "需要加密的密码")
//fmt.Printf("%x", h.Sum(nil))

