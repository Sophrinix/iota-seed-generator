IOTA Random Seed Generator
=====

This is a offline, crytographically secure pseudo-random number generator [(CSPRNG)](https://en.wikipedia.org/wiki/Cryptographically_secure_pseudorandom_number_generator) that generate [IOTA](http://iota.org/) seeds. It's written in [Go](https://golang.org) and it's using [crypto/rand](https://golang.org/pkg/crypto/rand/
) which uses `/dev/urandom/` and Windows’ `CryptGenRandom` API.

### Install

```
$ go install github.com/0x13a/iota-seed-generator
```

### Use it

```
$ iota-seed-generator
```

