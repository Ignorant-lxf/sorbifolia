//go:build ignore

package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

const (
	toLower = 'a' - 'A'
)

func main() {
	hex2intTable := func() [256]byte {
		var b [256]byte
		for i := 0; i < 256; i++ {
			c := byte(16)
			if i >= '0' && i <= '9' {
				c = byte(i) - '0'
			} else if i >= 'a' && i <= 'f' {
				c = byte(i) - 'a' + 10
			} else if i >= 'A' && i <= 'F' {
				c = byte(i) - 'A' + 10
			}
			b[i] = c
		}
		return b
	}()

	toLowerTable := func() [256]byte {
		var a [256]byte
		for i := 0; i < 256; i++ {
			c := byte(i)
			if c >= 'A' && c <= 'Z' {
				c += toLower
			}
			a[i] = c
		}
		return a
	}()

	toUpperTable := func() [256]byte {
		var a [256]byte
		for i := 0; i < 256; i++ {
			c := byte(i)
			if c >= 'a' && c <= 'z' {
				c -= toLower
			}
			a[i] = c
		}
		return a
	}()

	quotedArgShouldEscapeTable := func() [256]byte {
		// According to RFC 3986 §2.3
		var a [256]byte
		for i := 0; i < 256; i++ {
			a[i] = 1
		}

		// ALPHA
		for i := int('a'); i <= int('z'); i++ {
			a[i] = 0
		}
		for i := int('A'); i <= int('Z'); i++ {
			a[i] = 0
		}

		// DIGIT
		for i := int('0'); i <= int('9'); i++ {
			a[i] = 0
		}

		// Unreserved characters
		for _, v := range `-_.~` {
			a[v] = 0
		}

		return a
	}()

	quotedPathShouldEscapeTable := func() [256]byte {
		// The implementation here equal to net/url shouldEscape(s, encodePath)
		//
		// The RFC allows : @ & = + $ but saves / ; , for assigning
		// meaning to individual path segments. This package
		// only manipulates the path as a whole, so we allow those
		// last three as well. That leaves only ? to escape.
		var a = quotedArgShouldEscapeTable

		for _, v := range `$&+,/:;=@` {
			a[v] = 0
		}

		return a
	}()

	w := new(bytes.Buffer)
	w.WriteString(pre)
	fmt.Fprintf(w, "const hex2intTable = %q\n", hex2intTable)
	fmt.Fprintf(w, "const toLowerTable = %q\n", toLowerTable)
	fmt.Fprintf(w, "const toUpperTable = %q\n", toUpperTable)
	fmt.Fprintf(w, "const quotedArgShouldEscapeTable = %q\n", quotedArgShouldEscapeTable)
	fmt.Fprintf(w, "const quotedPathShouldEscapeTable = %q\n", quotedPathShouldEscapeTable)

	if err := os.WriteFile("byte_table.go", w.Bytes(), 0660); err != nil {
		log.Fatal(err)
	}
}

const pre = `package util

// Code generated by go run byte_table_gen.go; DO NOT EDIT.
// See byte_table_gen.go for more information about these tables.

`
