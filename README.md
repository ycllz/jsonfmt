#jsonfmt
Beautify json string data.

# Installation
```
go get -u github.com/voidint/jsonfmt
```

# How to use?
- JSON data from file
```
$ jsonfmt --input test.json
{
	"name": "voidint",
	"age": 25
}
```

- JSON data from pipeline
```
$ echo '{"name":"voidint","age":25}' | jsonfmt
{
	"name": "voidint",
	"age": 25
}
```

- Output to file
```
$ jsonfmt --output ~/test2.json
{"name":"voidint","age":25}
// EOF by CTRL+D(Unix/Linux) or CTRL+Z(Win)
```

- Output to terminal
```
$ jsonfmt
{"name":"voidint","age":25}
// EOF by CTRL+D(Unix/Linux) or CTRL+Z(Win)
{
	"name": "voidint",
	"age": 25
}
```


