package static

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func HTMLFileHandler(c *gin.Context) {
	r := c.Request
	r.URL.Path = "/html" + r.URL.Path
	FileHandler(c)
}

func FileHandler(c *gin.Context) {
	r := c.Request
	w := c.Writer
	f, err := os.Open("./web" + r.URL.Path)
	if err != nil {
		log.Print(err)
		return
	}
	defer f.Close()

	d, err := f.Stat()
	if err != nil {
		log.Print(err)
		return
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Print(err)
		return
	}

	w.Header().Set("Content-Length", strconv.FormatInt(d.Size(), 10))
	w.Write(data)
}

