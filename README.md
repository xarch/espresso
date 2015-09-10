# espresso
Human friendly format that is based on a subset of CoffeeScript Language, an alternative to
[YAML](http://yaml.org), [JSON](http://json.org), and [CSON](https://github.com/bevry/cson).
Reference Implementation.

## Version
**0.0.1**

## Example
```coffee
# This is a sample comment.
distros: [ # A list of ISOs we have.
  $ name: "Ubuntu", version: "x.x.x"
  $ name: "Debian"
  $
    name: "Mint"
    motto: """
      from freedom
      came elegance
    """
    nameTranslations:
      русский: "Минт"
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

## License
Distributed under the BSD 2-clause "Simplified" License unless otherwise noted.
