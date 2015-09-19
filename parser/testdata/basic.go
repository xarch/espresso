package testdata

// Basic is a representation of basic testdata espresso file.
var Basic = map[interface{}]interface{}{
	"distros": []interface{}{
		map[interface{}]interface{}{"name": "Ubuntu", "version": "x.x.x"},
		map[interface{}]interface{}{"name": "Debian"},
		map[interface{}]interface{}{"name": "Mint", "motto": `from freedom
came elegance`},
	},
	"greetings": map[interface{}]interface{}{
		"普通话":     "你好",
		"русский": "привет",
	},
	"numbers": map[interface{}]interface{}{
		"1": "One",
		"2": "Two",
		"3": "Three",
	},
	"bools": map[interface{}]interface{}{
		"true":  "Yes",
		"false": "No",
	},
	"dbs": map[interface{}]interface{}{
		"MySQL": map[interface{}]interface{}{
			"username": "test",
			"password": "test",
		},
		"Another Database": map[interface{}]interface{}{
			"active":    false,
			"port":      int64(1234),
			"usernames": []interface{}{"admin", "root", "userx"},
			"password":  "unknown",
		},
		"Redis": "inactive",
	},
}
