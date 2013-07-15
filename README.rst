SHA-256d
========

This package provides an implementation of the SHA-256d hash algorithm
for the Go programmin language.  SHA-256d is a cryptographic hash,
first proposed by Ferguson and Schneier in the book "Practical
Cryptography".  The SHA-256d hash is used in the Bitcoin protocol and
in the Fortuna random number generator.

The SHA-256d hash is obtained by applying the SHA-256 hash twice,
i.e. by first applying SHA-256 to the data and then again to the
resulting hash.
