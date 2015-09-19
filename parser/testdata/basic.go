package testdata

// Basic is a representation of basic testdata espresso file.
var Basic = map[string]interface{}{
	"distros": []map[string]interface{}{
		{"name": "Ubuntu", "version": "x.x.x"},
		{"name": "Debian"},
		{"name": "Mint", "motto": "from freedom came elegance"},
	},
	"greetings": map[string]string{
		"普通话":     "你好",
		"русский": "привет",
	},
	"numbers": map[string]string{
		"1": "One",
		"2": "Two",
		"3": "Three",
	},
	"bools": map[string]string{
		"true":  "Yes",
		"false": "No",
	},
	"dbs": map[string]interface{}{
		"MySQL": map[string]string{
			"username": "test",
			"password": "test",
		},
		"Another Database": map[string]interface{}{
			"active":   false,
			"port":     1234,
			"username": "unknown",
			"password": "unknown",
		},
		"Redis": "inactive",
	},
}
