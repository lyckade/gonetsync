//Package walk implements a recursive filewalk
package walk

import (
	"fmt"
	"os"
	"path/filepath"
	"raspi/apxp/golib/mylogger"
)

var myLogger = mylogger.NewFileLogger("../log.txt", "")

//Folder implements the filepath.Walk
func Folder(f string) {
	filepath.Walk(f, func(path string, info os.FileInfo, err error) error {
		fmt.Printf("%v\n", path)
		fmt.Printf("%v\n", info.IsDir())
		fmt.Printf("%v\n\n", info.ModTime())
		return err
	})
}
