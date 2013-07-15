SHA-256d
========

An implementation of the SHA-256d hash for the Go programming language.

Copyright (C) 2013  Jochen Voss

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

The homepage of this package is at http://www.seehuhn.de/pages/sha256d .
Please send any comments or bug reports to the program's author,
Jochen Voss <voss@seehuhn.de> .

Overview
--------

This package provides an implementation of the SHA-256d hash algorithm
(also known as "double SHA-256") for the Go programming language.
SHA-256d is a cryptographic hash, first proposed by Ferguson and
Schneier in the book "Practical Cryptography".  The SHA-256d hash is
used in the Bitcoin protocol and in the Fortuna random number
generator.

The SHA-256d hash is obtained by applying the SHA-256 hash twice,
i.e. by first applying SHA-256 to the data and then again to the
resulting hash.

Installation
------------

::

    go get github.com/seehuhn/sha256d

Usage
-----

The sha256d package implements the standard hash.Hash interface.
Example::

    hash := sha256d.New()
    n, _ := hash.Write([]byte("hello"))
    hashVal := hash.Sum(nil)
