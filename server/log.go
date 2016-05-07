package server

import (
	"io/ioutil"
	"os"
	"path"
	"server/utils"
)

// LogErrors logs error bytes to tmp/log directory.
func LogError(bytes []byte) {
	dir := App.Config().String("app.log_dir")
	file := path.Join(dir, utils.DateInt64(utils.Now(), "MMDDHHmmss.log"))
	ioutil.WriteFile(file, bytes, os.ModePerm)
}
