# caesar

## use
```shell
$ go install github.com/stripedpajamas/caesar/cmd/...
$ caesar help
```

generally the syntax is `caesar -c <cipher> -k <key> -t <text> <encrypt|decrypt>`

```shell
$ caesar -c caesar -k m -t "peter" encrypt
bqfqd
$ caesar -c caesar -k m -t "bqfqd" decrypt
peter

$ caesar -c playfair -k secret -t "peter" encrypt
OCSCCY
$ caesar -c playfair -k secret -t "OCSCCY" decrypt
PETERX
```


ciphers implemented:
- caesar https://en.wikipedia.org/wiki/Caesar_cipher
- playfair https://en.wikipedia.org/wiki/Playfair_cipher

### caesar cipher
the key can be a letter or a number. the letter _x_ is interpreted as _a = x_ for the shift. the number _n_ is interpreted as _shift the alphabet n places_. concretly, this means a = 0, b = 1, ..., z = 25. any number provided is wrapped around mod 26 and with any multi-letter key all but the first letter is ignored.

### playfair cipher
the key should be a keyword (anything else would be equivalent encrypting without a key). `j` is merged with `i`, so decrypting an encryption of "Joel" will result in "Ioel". odd-length plaintext and double-letters are padded with X, unless the double letter is X, then we pad with Q.

# license
MIT
