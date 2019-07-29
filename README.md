# caesar

## use
```bash
$ go install github.com/stripedpajamas/caesar/cmd/...
$ caesar help

$ caesar help encrypt
$ caesar help decrypt
```

generally the syntax is `caesar <en|de>crypt -k <key> -t <text>`

the key can be a letter or a number. the letter _x_ is interpreted as _a = x_ for the shift. the number _n_ is interpreted as _shift the alphabet n places_. concretly, this means a = 0, b = 1, ..., z = 25. any number provided is wrapped around mod 26 and with any multi-letter key all but the first letter is ignored.

# license
MIT
