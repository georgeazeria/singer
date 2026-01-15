# Singer

cross-platform sanitisation for file and folder names in go

## Installation

```
$ go get github.com/georgeazeria/singer
```

## Usage

singer ensures any string can be used as a file or folder (directory) name on any platform by normalising white space characters, filtering out illegal characters, and truncating the input to 255 characters

use `singer.File` to santise inputs if they'll only ever be used as a file name;
for folders `singer.Folder`, the only difference is string ending `.` or `,` becomes `_`.

```go
package main

import (
	"fmt"
	"github.com/azeria/singer"
)

func main() {
	fileName := singer.File("Feel Good Inc.")+".flac"
	fmt.Println(fileName) // => "Feel Good Inc..flac"

    folderName := singer.Folder("Convenient, Trash. - Convenient, Trash.")
	fmt.Println(folderName) // => "Convenient, Trash. - Convenient, Trash_"
}
```

## Valid file and folder names

follows these principles for ensuring best practices for having a safe and cross-platform file and folder names:

- Does not contain [ASCII control characters](http://en.wikipedia.org/wiki/ASCII#ASCII_control_characters) (hexadecimal `00` to `1f`)
- Does not contain [Unicode whitespace](http://en.wikipedia.org/wiki/Whitespace_character#Unicode) at the beginning and the end of filename
- Does not contain multiple Unicode whitespaces within the filename
- Does not contain [reserved filenames in Windows](http://msdn.microsoft.com/en-us/library/windows/desktop/aa365247%28v=vs.85%29.aspx)
- Does not contain following characters (according to [wikipedia](http://en.wikipedia.org/wiki/Filename)): `/ \ ? * : | " < >`
- Folder names cannot end in a full stop `.` or comma `,` (`singer.Folder()` changes the final rune to `_` if it matches either of these)

## Credits

singer is a fork of [gozaru](https://subosito/gozaru) by [@subosito](https://subosito); itself a go port of [zaru](https://github.com/madrobby/zaru) by [@madrobby](https://github.com/madrobby).
appreciate you both.
