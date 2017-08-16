Constants
=========

`constants` is a tool to extract constants from within a Golang file or package.

Example Usage
-------------

Example package: accounts


```console
$ constants accounts
name="John Woo"
age="10"
customer="John Woo"

$ constants accounts --duplicates
constant: "John Woo"
	people.go: name
	sales.go: customer
```
