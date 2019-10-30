# caesar

_a classic cipher cli_

## Installation
```shell
$ go install github.com/stripedpajamas/caesar/cmd/...
$ caesar help
```

generally the syntax is

```shell
caesar -c <cipher> -k <key> -t <text> <encrypt|decrypt>
```

leave off `-t <text>` and caesar will read from stdin.

optionally you can also provide a `--groups` or `-g` parameter to specify how the output should be grouped. the default value is 5. groups setting is ignored for decrypt operations.

```shell
# encrypting using caesar cipher
$ caesar -c caesar -k m -t "peter" encrypt
BQFQD
$ caesar -c caesar -k m -t "BQFQD" decrypt
PETER
```

ciphers implemented:
- caesar https://en.wikipedia.org/wiki/Caesar_cipher
- playfair https://en.wikipedia.org/wiki/Playfair_cipher
- vigenere https://en.wikipedia.org/wiki/Vigenère_cipher
- ADFGX https://en.wikipedia.org/wiki/ADFGVX_cipher (ADFGVX is not implemented)
- Bifid https://en.wikipedia.org/wiki/Bifid_cipher

generally all the ciphers operate on english alphabetic letters (a-z). plaintext and keys are both specified as letters (although a number works for the caesar cipher as well, see below).

## Using Different Ciphers

### Caesar
- encryption removes all non-alphabetic characters (encrypting `abc123` == encrypting `abc`)
- the key can be a letter or a number
  - the letter _x_ is interpreted as _a = x_ for the shift
  - the number _n_ is interpreted as _shift the alphabet n places_. concretely, this means keys _a, b, c, ..._ are equivalent to keys _0, 1, 2, ..._.
  - any number provided is wrapped around mod 26
  - if a multi-letter key is provided, all but the first letter is ignored
- example:
  ```shell
  $ caesar -c caesar -k m -t "peter" encrypt
  BQFQD
  $ caesar -c caesar -k m -t "BQFQD" decrypt
  PETER
  ```

### Playfair
- encryption removes all non-alphabetic characters (encrypting "abc123" == encrypting "abc")
- the key should be a keyword (anything else would be equivalent encrypting without a key)
- `j` is merged with `i`, so decrypting an encryption of "Joel" will result in "Ioel"
- odd-length plaintext and double-letters are padded with `X`, unless the double letter is an `X`, then it is padded with `Q`
- example:
  ```shell
  $ caesar -c playfair -k secret -t "peter" encrypt
  OCSCC Y
  $ caesar -c playfair -k secret -t "OCSCCY" decrypt
  PETERX
  ```

### Vigenère
- encryption removes all non-alphabetic characters (encrypting "abc123" == encrypting "abc")
- non-alphabetic characters are ignored if present in the key
- an empty key will return an error
- example:
  ```shell
  $ caesar -c vigenere -k hotsauce -t "let's eat some cheerios" encrypt
  SSMKE UVWVA XUHYG VPCL
  $ caesar -c vigenere -k hotsauce -t "SSM KEU VWV AXU HYG VPC L" decrypt
  LETSEATSOMECHEERIOS
  ```

### ADFGX
- encryption removes all non-alphabetic characters (encrypting "abc123" == encrypting "abc")
- this is not ADFGVX, so numbers are no supported
- two keys are required for this cipher, one for the polybius square and one for the transposition step
  - the keys should be comma delimited when running (e.g. `-k keyOne,keyTwo`)
- non-alphabetic characters are ignored if present in either key
- if either key is empty, it will return an error
- example:
  ```shell
  $ caesar -c adfgx -k help,me -t "this is a test" encrypt
  GAAFA FXGDF GGAFG FGAGA GG
  $ caesar -c adfgx -k help,me -t "GAAFA FXGDF GGAFG FGAGA GG" decrypt
  THISISATEST
  ```

### Bifid
- encryption removes all non-alphabetic characters (encrypting "abc123" == encrypting "abc")
- non-alphabetic characters are ignored if present in either key
- example:
  ```shell
  $ caesar -c bifid -k golden -t spandex encrypt
  SAGXW DX
  $ caesar -c bifid -k golden -t sagxwdx decrypt
  SPANDEX
  ```

## Other options
- group output in chunks
  ```shell
  $ caesar -g 10 -c caesar -k m -t "mary had a little lamb whose fleece was white as snow" encrypt
  YMDKTMPMXU FFXQXMYNIT AEQRXQQOQI MEITUFQMEE ZAI
  ```
- read from stdin
  ```shell
  $ cat my_love_letter.txt | caesar -c caesar -k 5 encrypt
  NKDTZ WJWJF INSLY MNXYM FYXUW JYYDH TTQ
  ```
- chaining (careful when chaining ciphers that use a polybius square with ciphers that don't, as the I/J situation can get wonky)
  ```shell
  $ echo "potatoes or potatos" | caesar -c caesar -k 7 e | caesar -c vigenere -k "poptarts" e
  LJPAA MERKM LOAYT NO
  $ echo "LJPAA MERKM LOAYT NO" | caesar -c vigenere -k "poptarts" d | caesar -c caesar -k 7 d
  POTATOESORPOTATOS
  ```

# License
MIT
