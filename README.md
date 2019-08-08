# caesar

## use
```shell
$ go install github.com/stripedpajamas/caesar/cmd/...
$ caesar help
```

generally the syntax is `caesar -c <cipher> -k <key> -t <text> <encrypt|decrypt>`

leave off `-t <text>` and caesar will read from stdin.

optionally you can also provide a `--groups` or `-g` parameter to specify how the output should be grouped. the default value is 5. groups setting is ignored for decrypt operations.

```shell
# Caesar cipher
$ caesar -c caesar -k m -t "peter" encrypt
BQFQD
$ caesar -c caesar -k m -t "BQFQD" decrypt
PETER

# Playfair cipher
$ caesar -c playfair -k secret -t "peter" encrypt
OCSCC Y
$ caesar -c playfair -k secret -t "OCSCCY" decrypt
PETERX

# Vignere cipher
$ caesar -c vigenere -k hotsauce -t "let's eat some cheerios" encrypt
SSMKE UVWVA XUHYG VPCL
$ caesar -c vigenere -k hotsauce -t "SSM KEU VWV AXU HYG VPC L" decrypt
LETSEATSOMECHEERIOS

# ADFGX cipher (takes two keys)
$ caesar -c adfgx -k help,me -t "this is a test" encrypt
GAAFA FXGDF GGAFG FGAGA GG
$ caesar -c adfgx -k help,me -t "GAAFA FXGDF GGAFG FGAGA GG" decrypt
THISISATEST

# group output in chunks of 10 chars
$ caesar -g 10 -c caesar -k m -t "mary had a little lamb whose fleece was white as snow" encrypt
YMDKTMPMXU FFXQXMYNIT AEQRXQQOQI MEITUFQMEE ZAI

# read from stdin
$ cat my_love_letter.txt | caesar -c caesar -k 5 encrypt
NKDTZ WJWJF INSLY MNXYM FYXUW JYYDH TTQ
```


ciphers implemented:
- caesar https://en.wikipedia.org/wiki/Caesar_cipher
- playfair https://en.wikipedia.org/wiki/Playfair_cipher
- vigenere https://en.wikipedia.org/wiki/Vigenère_cipher
- ADFGX https://en.wikipedia.org/wiki/ADFGVX_cipher (ADFGVX is not implemented)

generally all the ciphers operate on english alphabetic letters (a-z). plaintext and keys are both specified as letters (although a number works for the caesar cipher as well, see below).

### caesar cipher
the key can be a letter or a number. the letter _x_ is interpreted as _a = x_ for the shift. the number _n_ is interpreted as _shift the alphabet n places_. concretely, this means a = 0, b = 1, ..., z = 25. any number provided is wrapped around mod 26 and with any multi-letter key all but the first letter is ignored.

### playfair cipher
the key should be a keyword (anything else would be equivalent encrypting without a key). `j` is merged with `i`, so decrypting an encryption of "Joel" will result in "Ioel". odd-length plaintext and double-letters are padded with X, unless the double letter is X, then we pad with Q.

### vigenere cipher
the key should be a keyword (anything else would be equivalent encrypting without a key). non-alphabetic characters are ignored in the plaintext, and throw errors if present in the key.

### ADFGX cipher
two keys are required for this cipher, one for the polybius square and one for the transposition step. these keys should be comma delimited when running (e.g. `-k keyOne,keyTwo`).

# license
MIT
