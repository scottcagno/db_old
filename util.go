// * 
// * Copyright 2013, Scott Cagno. All rights Reserved
// * License: sites.google.com/site/bsdc3license
// * 
// * -------
// * util.go ::: databse utilities
// * -------
// * 

package db

import (
	"encoding/base64"
	"crypto/rand"
	"crypto/md5"
	"bytes"
	"log"
	"io"
)


// return random hash (6*10^49)
func random() []byte {
	e := make([]byte, 32)
	rand.Read(e)
	seed := make([]byte, base64.URLEncoding.EncodedLen(len(e)))
	base64.URLEncoding.Encode(seed, e)
	h := md5.New()
	i := 3
	for i > 0 {
		io.WriteString(h, string(seed))
		i--
	}
	return h.Sum(nil)
}