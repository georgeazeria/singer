# Singer

cross-platform sanitisation for file and folder names in Go

## Installation

```
$ go get github.com/georgeazeria/singer
```

## Usage

singer ensures any string can be used as a file or folder (directory) name on any platform by normalising white space characters, filtering out illegal characters, and truncating the input to 255 characters

```go
package main

import (
	"fmt"
	"github.com/azeria/singer"
)

func main() {
	name := singer.Sanitise("  what\\ēver//wëird:user:înput:")
	fmt.Println(name) // => "whatēverwëirduserînput"
}
```

you can add extra room for filename by using `SanitiseLength`:

```go
// import "strings"

name := strings.Repeat("A", 400)

singer.Sanitise(name)
// => resulting filename is 255 characters long

singer.SanitiseLength(name, 100)
// => resulting filename is 155 characters long
```

and yes, the American spelling also works:

```go

func main() {
	name := singer.Sanitize("  what\\ēver//wëird:user:înput:")
	username := singer.SanitizeLength("  what\\ēver//wëird:user:înput:", 100)
	fmt.Println(name) // => "whatēverwëirduserînput"
	fmt.Println(username) // => "whatēverwëirduserînput"
}
```

## Valid file and folder names

follows these principles for ensuring best practices for having a safe and cross-platform filenames are:

- Does not contain [ASCII control characters](http://en.wikipedia.org/wiki/ASCII#ASCII_control_characters) (hexadecimal `00` to `1f`)
- Does not contain [Unicode whitespace](http://en.wikipedia.org/wiki/Whitespace_character#Unicode) at the beginning and the end of filename
- Does not contain multiple Unicode whitespaces within the filename
- Does not contain [reserved filenames in Windows](http://msdn.microsoft.com/en-us/library/windows/desktop/aa365247%28v=vs.85%29.aspx)
- Does not contain following characters (according to [wikipedia](http://en.wikipedia.org/wiki/Filename)): `/ \ ? * : | " < >`
- Does not end in a full stop `'.'` or comma `','` (in singer, final rune is replaced with '\_' if it matches either of these)

## Credits

singer is a fork of [gozaru](https://subosito/gozaru) by [@subosito](https://subosito); itself a Go port of [zaru](https://github.com/madrobby/zaru) by [@madrobby](https://github.com/madrobby). Appreciate you both.
