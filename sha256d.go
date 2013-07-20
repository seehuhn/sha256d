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

// Package sha256d implements the SHA-256d hash algorithm.  This
// algorithm, obtained by applying SHA-256 to the data and then again
// to the resulting hash, was first proposed by Ferguson and Schneier
// in the book "Practical Cryptography".  It is used in the Bitcoin
// protocol and in the Fortuna random number generator.
package sha256d

import (
	"crypto/sha256"
	"hash"
)

const (
	// Size is the size of a SHA-256d checksum in bytes.
	Size = 32

	// BlockSize is the internal blocksize of SHA-256d in bytes.
	BlockSize = 64
)

type sha256d struct {
	h1, h2 hash.Hash
}

func (hash *sha256d) Write(p []byte) (n int, err error) {
	return hash.h1.Write(p)
}

func (hash *sha256d) Sum(b []byte) []byte {
	hash.h2.Reset()
	hash.h2.Write(hash.h1.Sum(nil))
	return hash.h2.Sum(b)
}

func (hash *sha256d) Reset() {
	hash.h1.Reset()
}

func (hash *sha256d) Size() int {
	return Size
}

func (hash *sha256d) BlockSize() int {
	return BlockSize
}

// New returns a new hash.Hash computing the SHA-256d checksum.
func New() hash.Hash {
	return &sha256d{
		h1: sha256.New(),
		h2: sha256.New(),
	}
}
