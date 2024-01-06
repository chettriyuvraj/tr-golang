# README


A basic implementation of tr in Golang - because I was bored :)



Supports the following: 
- Simple substitutions (tr a b)
- Unicode substitutions (tr 🔥 a  OR  tr a 🔥)
- Simple range substitutions (a-z A-Z)
- Reverse ranges (a-z z-a)
- Some class specifiers ("[:alpha:]" "[:lower:]" "[:upper:]")


## How to run

In the root directory:
- _go build_
- _./tr a-z A-Z_