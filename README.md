# espresso
Simple configuration file format that aims to be easy to read, write, and parse.
Reference Implementation.

* Whitespace is not significant.
* CoffeeScript syntax highlighting can be used.

## Version
**0.0.1**

## Story of espresso
#### In the beginning: JSON
Initially there was a JSON configuration file (*fragment*):
```json
"watch": {
	"_comment": [
		"This section contains paths to directories that should be watched and commands",
		"that should be executed when files in those directories are changed."
	],

	"path/to/controllers": [
		{"_comment": "These commands will be started every time you modify your controllers."}

		{"run": "go get ./controllers"},
		{"run": "go build ./...", "async": true},
		{"run": "go test ./... -v"},
		{"run": "app", "single_instance": true}
	]
}
```

#### Comments
It was almost ideal except, there were no comments. So, **espresso** brought support of them:
```coffee
# This section contains paths to directories that should be watched and commands
# that should be executed when files in those directories are changed.
```

#### Commas
Then commas... It was a bit boring to type them every time and it was a bit irritating to control
there were no trailing commas somewhere.
A decision to get rid of them altogether was made, space is enough for all!
```coffee
"run": "go build ./..." "async": true
```

#### "Keys"
But how about keys? Typing those double quotes (`""`) around them was not big fun, too.
So, quotes became optional:
```coffee
key_without_quotes: "some value"
"key with quotes": "some value"
```

#### Curly braces
We were cool. But it felt like those curly braces (`{}`) were superfluous sometimes.
That's why they were allowed to be dropped in case of simple one-line objects.
And what we finally got is:
```coffee
# This section contains paths to directories that should be watched and commands
# that should be executed when files in those directories are changed.
watch: {
	# These commands will be started every time you modify your controllers.
	"path/to/controllers": [
		run: "go get ./controllers"
		run: "go build ./..." async: true
		run: "go test ./... -v"
		run: "app" single_instance: true
	]
}
```
Much easier to read and write, right? And is almost as simple to parse as JSON.

## Full sample
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

> Note: Keys are always parsed as strings.
> ```coffee
> 123: "value" false: "value"
> ```
> The example above is an equivalent of:
> ```go
> map[interface{}]interface{}{"123": "value", "false": "value"}
> ```

### Strings
```coffee
"Some non-breakable string."
```
```coffee
`Multiline
string
is here...`
```

## License
Distributed under the BSD 2-clause "Simplified" License unless otherwise noted.
