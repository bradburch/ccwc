# Coding Challenges WC #

https://codingchallenges.fyi/challenges/challenge-wc

## How To Build ##

`go build ccwc.go`

## How To Run ##

`./ccwc -<flag> <file-path>`

## Supported Flags ##

`-c` Outputs the number of bytes in a file. 

`-l` Outputs the number of lines in a file.

`-w` Outputs the number of words in a file.

`-m` Outputs the number of characters in a file.

Default outputs `-c`, `-l`, `-w` together. 

## Examples ##

| Command | Output |
| ----------- | ----------- | 
| `./ccwc -w test.txt` | `58164    test.txt` |
| `./ccwc test.txt`    | `342190  7145    58164    test.txt` |






