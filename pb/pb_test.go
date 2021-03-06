// Copyright 2013-2014 Vasiliy Gorin. All rights reserved.
// Use of this source code is governed by a GNU-style
// license that can be found in the LICENSE file.

package pb

import "testing"
import "bytes"

import "github.com/vgorin/cryptogo/rnd"

func TestPBAes(t *testing.T) {
	message := "this is my tiny message"
	t.Logf("message: %s", message)
	password := "this is my easy to guess password"
	message_bytes := []byte(message)
	encrypted_bytes, err := PBAesEncrypt(message_bytes, password)
	if err != nil {
		t.Fatal(err)
	}
	decrypted_bytes, err := PBAesDecrypt(encrypted_bytes, password)
	decrypted := string(decrypted_bytes)
	t.Logf("decrypted: %s", decrypted)
	if bytes.Compare(message_bytes, decrypted_bytes) != 0 {
		t.Error("original and decryoted messages mismatch")
	}
}

func TestRandom(t *testing.T) {
	buffer_length := 1 << 20
	cycles := 1

	for i := 0; i < cycles; i++ {
		buffer, err := rnd.RandomBytes(buffer_length)
		if err != nil {
			t.Fatal(err)
		}
		password_bytes, err := rnd.RandomBytes(10)
		if err != nil {
			t.Fatal(err)
		}
		password := string(password_bytes)

		encrypted, err := PBAesEncrypt(buffer, password)
		if err != nil {
			t.Fatal(err)
		}
		decrypted, err := PBAesDecrypt(encrypted, password)
		if bytes.Compare(buffer, decrypted) != 0 {
			t.Error("original and decryoted messages mismatch")
		}
	}
}

func BenchmarkPBAes(b *testing.B) {
	buffer_length := b.N
	cycles := 10

	buffer, err := rnd.RandomBytes(buffer_length)
	if err != nil {
		b.Fatal(err)
	}

	password := "my easy pass"
	for i := 0; i < cycles; i++ {
		encrypted, err := PBAesEncrypt(buffer, password)
		if err != nil {
			b.Fatal(err)
		}
		decrypted, err := PBAesDecrypt(encrypted, password)
		if bytes.Compare(buffer, decrypted) != 0 {
			b.Error("original and decryoted messages mismatch")
		}
	}
}
