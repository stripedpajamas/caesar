# caesar

## use
```shell
$ go install github.com/stripedpajamas/caesar/cmd/...
$ caesar help
```

generally the syntax is `caesar -c <cipher> -k <key> -t <text> <encrypt|decrypt>`

```shell
$ caesar -c caesar -k m -t "peter" encrypt
BQFQD
$ caesar -c caesar -k m -t "BQFQD" decrypt
peter

$ caesar -c playfair -k secret -t "peter" encrypt
OCSCCY
$ caesar -c playfair -k secret -t "OCSCCY" decrypt
PETERX

$ caesar -k hotsauce -t "let's eat some cheerios" -c vigenere encrypt
SSMKEUVWVAXUHYGVPCL
$ caesar -k hotsauce -t "SSMKEUVWVAXUHYGVPCL" -c vigenere decrypt
LETSEATSOMECHEERIOS
```


ciphers implemented:
- caesar https://en.wikipedia.org/wiki/Caesar_cipher
- playfair https://en.wikipedia.org/wiki/Playfair_cipher
- vigenere https://en.wikipedia.org/wiki/Vigenère_cipher

generally all the ciphers operate on english alphabetic letters (a-z). plaintext and keys are both specified as letters (although a number works for the caesar cipher as well, see below).

### caesar cipher
the key can be a letter or a number. the letter _x_ is interpreted as _a = x_ for the shift. the number _n_ is interpreted as _shift the alphabet n places_. concretely, this means a = 0, b = 1, ..., z = 25. any number provided is wrapped around mod 26 and with any multi-letter key all but the first letter is ignored.

### playfair cipher
the key should be a keyword (anything else would be equivalent encrypting without a key). `j` is merged with `i`, so decrypting an encryption of "Joel" will result in "Ioel". odd-length plaintext and double-letters are padded with X, unless the double letter is X, then we pad with Q.

### vigenere cipher
the key should be a keyword (anything else would be equivalent encrypting without a key). non-alphabetic characters are ignored in the plaintext, and throw errors if present in the key.

# license
MIT
