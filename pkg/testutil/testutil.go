package testutil

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

// JSONDeepEquals compares the JSON marshaled results.
//
// a little miffed I had to make this. basically routes around comparing
// different pointer values with reflect.DeepEqual.
func JSONDeepEquals(r1, r2 interface{}) (bool, error) {
	r1b := bytes.NewBuffer(nil)
	r2b := bytes.NewBuffer(nil)

	if err := json.NewEncoder(r1b).Encode(r1); err != nil {
		return false, fmt.Errorf("while encoding first object: %v", err)
	}

	if err := json.NewEncoder(r2b).Encode(r2); err != nil {
		return false, fmt.Errorf("while encoding second object: %v", err)
	}

	return bytes.Equal(r1b.Bytes(), r2b.Bytes()), nil
}

// RandomSHA generates a random hex-encoded git-looking thing.
func RandomSHA() string {
	return randomSHA(20)
}

// RandomUploadSHA generates a random hex-encoded SHA256-looking thing.
func RandomUploadSHA() string {
	return randomSHA(32)
}

func randomSHA(sz int) string {
	buf := make([]byte, sz)

	n, err := rand.Reader.Read(buf)
	if err != nil {
		panic(err)
	}

	if n != sz {
		panic("short read in random device")
	}

	return hex.EncodeToString(buf)
}

// RandomString generates a random string which has a length of the provided
// amount with a guaranteed minimum. It chooses from a-zA-Z0-9. It's very
// simple.
func RandomString(length, min uint) string {
	buf := make([]byte, length+1)
	n, err := rand.Reader.Read(buf)
	if err != nil {
		panic(err)
	}

	if uint(n) != length+1 {
		panic("short read on random device")
	}

	// we use the first random byte to determine the length of the rest of the
	// string. We cap it by subtracting the minimum before the mod then adding it
	// back in.

	l := length
	if min < length {
		l = length - min
	}

	l = uint(buf[0]) % l

	if min < length {
		length = l + min
	}

	str := []byte{}

	// Because we're using the first byte above, the offsets must be calculated.
	for _, b := range buf[1 : length+1] {
		c := b % 62
		// above 52 is a number; calculate from '0'. above 26 is a capital letter,
		// so calculate the offset from 'A'. Remember, in ascii, latin alphabet is
		// listed in order, so we can just add to get the letters.
		if c >= 52 {
			str = append([]byte(str), byte('0')+byte(c-52))
		} else if c >= 26 {
			str = append([]byte(str), byte('A')+byte(c-26))
		} else {
			str = append([]byte(str), byte('a')+byte(c))
		}
	}

	return string(str)
}
