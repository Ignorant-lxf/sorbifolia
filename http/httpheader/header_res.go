package httpheader

import (
	"bytes"

	"go.x2ox.com/sorbifolia/http/internal/char"
	"go.x2ox.com/sorbifolia/http/kv"
)

type ResponseHeader struct {
	kv.KVs

	ContentLength ContentLength
	ContentType   ContentType
	SetCookies    SetCookies

	Close bool
}

func (rh *ResponseHeader) RawParse() error {
	rh.Each(func(kv kv.KV) bool {
		switch {
		case bytes.EqualFold(kv.K, char.Connection):
			if bytes.EqualFold(kv.V, char.Close) {
				rh.Close = true
			}
		case bytes.EqualFold(kv.K, char.ContentLength):
			rh.ContentLength = kv.V
		case bytes.EqualFold(kv.K, char.SetCookie):
			rh.SetCookies = append(rh.SetCookies, kv.V)
		}
		return true
	})

	return nil
}

func (rh *ResponseHeader) Reset() {
	rh.KVs.Reset()
	rh.ContentLength = rh.ContentLength[:0]
	rh.ContentType = rh.ContentType[:0]
	rh.Close = false

	for i := range rh.SetCookies {
		rh.SetCookies[i] = rh.SetCookies[i][:0]
	}
	rh.SetCookies = rh.SetCookies[:0]
}
