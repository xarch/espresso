# espresso
Human friendly format that is based on a CoffeeScript Language's syntax, an alternative to
[YAML](http://yaml.org), [JSON](http://json.org), and [CSON](https://github.com/bevry/cson).
Reference Implementation.

## Version
**0.0.1**

## Example
```coffee
# Sample espresso file.

distros: [ # A list of ISOs we have.
	$ name: "Ubuntu", version: "x.x.x"
	$ name: "Debian"
	$
		name: "Mint"
		motto: """from freedom
came elegance"""
		nameTranslations:
			русский: "Минт"
	$ name: "Gentoo"
]

dbs: # DB parameters.
	MySQL: username: "test", password: "test",
	"Another Database":
		active: false
		port: 1234
		username: "unknown"
		password: "unknown"
	Redis:
		smth: "here"
```

## Spec
TODO: PEG file with a formal specification.

### Booleans
### Comments
### Lists
### Numbers
### Objects
### Strings
```coffee
"Some non-breakable string."
```
```coffee
"""Multiline
string
is here..."""
```
**NB**: For simplicity, as opposed to CoffeeScript:

1. No multiline strings of form `"..."` are allowed.
2. Body of `"""..."""` is not processed in any special way. To illustrate:
```coffee
"""
	Some
	string...
"""
```
The string above is equivalent to `"\n\tSome\n\tstring...\n"`, **NOT**
`"Some string..."` as in CoffeeScript.

## License
Distributed under the BSD 2-clause "Simplified" License unless otherwise noted.
