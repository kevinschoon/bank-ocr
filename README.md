# bankocr

This is a quick solution for the [BankOCR](https://codingdojo.org/kata/BankOCR/)
Kata challenge.

## Installation & Usage

You need to have a modern version of [Go](https://golang.org/) installed
and also [Make](https://www.gnu.org/software/make/manual/make.html).

```sh
git clone git@github.com:kevinschoon/bank-ocr
cd bank-ocr
make
# now you should be able to execute the bank-ocr program
bin/bank-ocr -path examples/sample.txt
# run tests
make test
```
