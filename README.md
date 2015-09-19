# espresso
CoffeeScript family Object Notation, an alternative to
[YAML](http://yaml.org), [JSON](http://json.org), and [CSON](https://github.com/bevry/cson)
for configuration files. Reference Implementation.

* Easy to read and write.
* Easy to parse.
* Whitespace is not significant.
* CoffeeScript syntax highlighting can be used.

## Version
**0.0.1**

## Example
```coffee
# Sample espresso file.
{
	distros: [ # A list of ISOs we have.
		name: "Ubuntu" version: "x.x.x"
		name: "Debian"
		{
			name: "Mint",
			motto: `from freedom
came elegance`
		}
		name: "Gentoo"
	]

	greetings: {
		普通话: "你好" # Comments at the end of lines are allowed.
		русский: "привет"
	}

	dbs: { # DB parameters.
		MySQL: username: "test" password: "test"
		"Another Database": {
			active: false
			port: 1234
			username: "unknown"
			password: "unknown"
		}
		Redis: smth: "here"
	}
}
```

## Spec
Read [PEG file](grammar/espresso_v0_0_1.peg) for more formal description of the format.

### Booleans
```coffee
active: true
deprecated: false
```
### Comments
```coffee
# A comment may start at the beginning of line.
lst: [ # Or at the end of line.
	1 2 3 4 5 # Like here.
] # Or here.
```
### Lists
```coffee
strings: [
	"s1" "s2" "s3"
]
ints: [1 2 3]
```
### Numbers
#### Integers
```coffee
a: 100
b: -200
```
### Objects
#### Inline
```coffee
somekey: anotherKey1: "value1" anotherKey2: "value2"
"Key with spaces": 123
```
#### Multiline
```coffee
somekey: {
	anotherKey1: "value1"
	anotherKey2: "value2"
}
```
### Strings
```coffee
"Some non-breakable string."
```
```coffee
`Multiline raw
string
is here...`
```

## License
Distributed under the BSD 2-clause "Simplified" License unless otherwise noted.
