// An implementation of the SHA-256d hash for the Go programming language.
// Copyright (C) 2013  Jochen Voss <voss@seehuhn.de>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package sha256d

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"testing"
)

func TestSha256d(t *testing.T) {
	hash := New()
	h1 := hash.(*sha256d).h1
	orig := sha256.New()
	if hash.Size() != Size ||
		orig.Size() != Size ||
		h1.Size() != Size {
		t.Error("wrong hash size")
	}
	if hash.BlockSize() != BlockSize ||
		orig.BlockSize() != BlockSize ||
		h1.BlockSize() != BlockSize {
		t.Error("wrong block size")
	}

	for _, l := range []int{0, 1, hash.Size(), hash.BlockSize(), 1024 * 1024} {
		buf := make([]byte, l)
		n, err := hash.Write(buf)
		if err != nil {
			t.Error("Write() returned an error")
		} else if n != l {
			t.Error("Write() returned wrong length")
		}
	}

	// test case from https://en.bitcoin.it/wiki/Protocol_specification
	in := []byte("hello")
	hash.Reset()
	n, err := hash.Write(in)
	if err != nil {
		t.Error("Write() returned an error")
	} else if n != len(in) {
		t.Error("Write() returned wrong length")
	}
	out := hash.Sum(nil)
	correct, err := hex.DecodeString(
		"9595c9df90075148eb06860365df33584b75bff782a510c6cd4883a419833d50")
	if err != nil {
		t.Fatal("test is corrupted")
	}
	if bytes.Compare(out, correct) != 0 {
		t.Error("wrong hash value")
	}
}

func BenchmarkHash(b *testing.B) {
	hash := New()
	data := []byte("This is a test string, containing data to hash.")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hash.Write(data)
		hash.Write(data)
		hash.Sum(nil)
	}
}
