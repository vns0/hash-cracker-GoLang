# hash-cracker-GoLang

- Simple multithreading bruteforce hash:

    |    Algorithm    | 128 | 160 | 192 | 256 | 512 |Performance|
    |:----------------|:---:|:---:|:---:|:---:|:---:|:----:|
    | BLAKE-2B        |     |     |     | O   | O   |Fast  |
    | BLAKE-2S        | O   |     |     | O   |     |Fast  |
    | GOST94 CryptoPro|     |     |     | O   |     |Slow  |
    | Grøstl          |     |     |     | O   |     |Slow  |
    | JH              |     |     |     | O   |     |Slow  |
    | Keccak          |     |     |     | O   | O   |Fast  |
    | MD5 [Obsolete]  | O   |     |     |     |     |Fast  |
    | RIPEMD          |     | O   |     |     |     |Fast  |
    | SHA1 [Obsolete] |     | O   |     |     |     |Fast  |
    | SHA2 (default)  |     |     |     | O   | O   |Fast  | 
    | SHA3            |     |     |     | O   | O   |Fast  |
    | Skein512        |     |     |     | O   | O   |Medium|
    | SM3             |     |     |     | O   |     |Fast  |
    | Streebog        |     |     |     | O   | O   |Slow  | 
    | Tiger           |     |     | O   |     |     |Fast  | 
    | Whirlpool       |     |     |     |     | O   |Slow  |
    
## Usage
    ./dehash -hash hexstring|-file path [-type sha1|sha256|sha512|md5] [-max val] [-min val] [-charset chars] [-threads num] 

    Usage of ./dehash:
    -charset string
          char set for possible message (default "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~ \t\n\r\v\f")
    -file string
          path to file with hashes
    -hash string
          hash to be decrypted
    -max int
          max length of message (default 10)
    -min int
          min length of message (default 1)
    -threads int
          max number of threads (default 256)
    -type string
          hash algorithm (default "sha1")

## Examples

Crack **sha1** hash
  
    ./dehash -hash fc19318dd13128ce14344d066510a982269c241b

Crack **md5** hash with message length between 3 and 5 characters
    
    ./dehash -type md5 -min 2 -max 5 -hash 6b2ded51d81a4403d8a4bd25fa1e57ee

Crack **sha1** hash with custom char set
    
    ./dehash -hash b3da1b8c56c939e616aa3c0934bce72cb1e82b32 -charset abcdhijklm

Crack **sha256** with 512 threads
    
    ./dehash -type sha256 -threads 512 -hash 97c10efe01d5c9c88704a12d361d8429b3a6aa2412290a0773109d5d2d603d5e

Crack **sha1** hashes from file
    
    ./dehash -type sha1 -file ./hashes.txt

## Get elapsed time

Just use **time** command

    $ time ./dehash -hash 2e96e89125f4c1aef797410a4bfdb32c0632ef0c
    Start cracking hash 2e96e89125f4c1aef797410a4bfdb32c0632ef0c
    Check mesages with length: 1 | Possible variants: 100
    Check mesages with length: 2 | Possible variants: 10000
    Check mesages with length: 3 | Possible variants: 1000000
    Check mesages with length: 4 | Possible variants: 100000000
    =========> Message: emit

    real    0m29,028s
    user    1m47,758s
    sys 0m1,119s

## Output examples

Ex. 1

    $ ./dehash -hash fc19318dd13128ce14344d066510a982269c241b
    Start cracking hash fc19318dd13128ce14344d066510a982269c241b
    Check mesages with length: 1 | Possible variants: 100
    Check mesages with length: 2 | Possible variants: 10000
    Check mesages with length: 3 | Possible variants: 1000000
    Check mesages with length: 4 | Possible variants: 100000000
    =========> Message: good

Ex. 2

    $ ./dehash -max 5 -file ./hashes.txt -charset abcdefghijklmnopqrstuvwxyz/:.
    Start cracking hash c3437dbc7c1255d3a21d444d86ebf2e9234c22bd
    Check mesages with length: 1 | Possible variants: 29
    Check mesages with length: 2 | Possible variants: 841
    Check mesages with length: 3 | Possible variants: 24389
    Check mesages with length: 4 | Possible variants: 707281
    Check mesages with length: 5 | Possible variants: 20511149
    =========> Message: https

    Start cracking hash ef81042e1e86acb765718ea37393a1292452bbcc
    Check mesages with length: 1 | Possible variants: 29
    Check mesages with length: 2 | Possible variants: 841
    Check mesages with length: 3 | Possible variants: 24389
    =========> Message: ://

    Start cracking hash a3c1509bd8df6d72992b312e4f6b7f4ce7fd3f3d
    Check mesages with length: 1 | Possible variants: 29
    Check mesages with length: 2 | Possible variants: 841
    Check mesages with length: 3 | Possible variants: 24389
    Check mesages with length: 4 | Possible variants: 707281
    Check mesages with length: 5 | Possible variants: 20511149
    =========> Message: docs.

    Start cracking hash 3f95edc0399d06d4b84e7811dd79272c69c8ed3a
    Check mesages with length: 1 | Possible variants: 29
    Check mesages with length: 2 | Possible variants: 841
    Check mesages with length: 3 | Possible variants: 24389
    =========> Message: goo

```
Donate:

    BTC:  192TC7d7ZRYJQbQnAWvMpkccnBNQN1ae6R
    ETH:  0x7d1082d952f4d584ae2910e14018f4dce7495c74
    LTC:  MLx6wmFjXfBTKj6JfB5NXaiKjNLeEntRoZ
    DOGE: DHCjW71EWBzvv43XPXVJc491brcBJXXq88
```
    author: 
    
    Name:          Nikita
    Company:       SmartWorld
    Position:      TeamLead
    Mail:          n.vtorushin@inbox.ru
    TG:            @nikitavoryet
    Year of birth: 1999
    FullStack:     JS/GO
