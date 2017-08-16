Constants
---------

`constants` is a tool to extract constants from Go source files.


### Example Usage

```console
$ constants
name = "John Woo"
age = "10"
customer = "John Woo"

$ constants --dup
constant: "John Woo"
	name: people.go
	customer: sales.go
```
