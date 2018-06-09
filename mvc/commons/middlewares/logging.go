package middlewares

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return LoggerWithWriter(os.Stdout)
}

func LoggerWithWriter(out io.Writer, notlogged ...string) gin.HandlerFunc {
	var (
		format = strings.Replace(
			"time:%v method:%s path:%s query:%s body:%s status:%d took:%v\n",
			" ",
			"\t",
			-1,
		)
	)
	return func(c *gin.Context) {
		var (
			method  = c.Request.Method
			path    = c.Request.URL.Path
			query   = c.Request.URL.RawQuery
			body, _ = ioutil.ReadAll(c.Request.Body)
			status  = c.Writer.Status()
		)
		defer func(begin time.Time) {
			var (
				took = time.Since(begin).Round(time.Millisecond)
			)
			fmt.Fprintf(out, format,
				begin.Format(time.RFC3339),
				method,
				path,
				query,
				string(body),
				status,
				took,
			)
		}(time.Now())

		// set nopcloser request body
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		// Process request
		c.Next()
	}
}
