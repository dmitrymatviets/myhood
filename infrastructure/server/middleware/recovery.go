package middleware

import (
	"bytes"
	"fmt"
	"github.com/dmitrymatviets/myhood/infrastructure"
	"github.com/dmitrymatviets/myhood/infrastructure/logger"
	config "github.com/dmitrymatviets/myhood/infrastructure/server/config"
	"io/ioutil"
	"net/http/httputil"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	dunno     = []byte("???")
	centerDot = []byte("·")
	dot       = []byte(".")
	slash     = []byte("/")
)

// RecoveryMiddleware returns a middleware that recovers from any panics and writes a 500 if there was one.
func RecoveryMiddleware(cfg *config.ServerConfig, logger *logger.Logger) gin.HandlerFunc {
	return RecoveryWithWriter(cfg, logger)
}

// RecoveryWithWriter returns a middleware for a given writer that recovers from any panics and writes a 500 if there was one.
func RecoveryWithWriter(cfg *config.ServerConfig, logger *logger.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				stack := stack(3)
				httprequest, _ := httputil.DumpRequest(ctx.Request, false)

				logger.Error(ctx, fmt.Sprintf("[recovery] %s panic recovered:\n%s\n%s\n%s", timeFormat(time.Now()), string(httprequest), err, stack))

				err = ctx.Error(fmt.Errorf("%s", err))

				ctx.Set(infrastructure.CtxKeyResponse, err)
				sendDecoratedJsonResponse(ctx, cfg, logger)

				ctx.Abort()
			}
		}()
		ctx.Next()
	}
}

// stack returns a nicely formatted stack frame, skipping skip frames.
func stack(skip int) []byte {
	buf := new(bytes.Buffer) // the returned data
	// As we loop, we open files and read them. These variables record the currently
	// loaded file.
	var lines [][]byte
	var lastFile string
	for i := skip; ; i++ { // Skip the expected number of frames
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		// Print this much at least.  If we can't find the source, it won't show.
		_, _ = fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
		if file != lastFile {
			data, err := ioutil.ReadFile(file)
			if err != nil {
				continue
			}
			lines = bytes.Split(data, []byte{'\n'})
			lastFile = file
		}
		_, _ = fmt.Fprintf(buf, "    %s: %s\n", function(pc), source(lines, line))
	}
	return buf.Bytes()
}

// source returns a space-trimmed slice of the n'th line.
func source(lines [][]byte, n int) []byte {
	n-- // in stack trace, lines are 1-indexed but our array is 0-indexed
	if n < 0 || n >= len(lines) {
		return dunno
	}
	return bytes.TrimSpace(lines[n])
}

// function returns, if possible, the name of the function containing the PC.
func function(pc uintptr) []byte {
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return dunno
	}
	name := []byte(fn.Name())
	// The name includes the path name to the package, which is unnecessary
	// since the file name is already included.  Plus, it has center dots.
	// That is, we see
	//	runtime/debug.*T·ptrmethod
	// and want
	//	*T.ptrmethod
	// Also the package path might contains dot (e.g. code.google.com/...),
	// so first eliminate the path prefix
	if lastslash := bytes.LastIndex(name, slash); lastslash >= 0 {
		name = name[lastslash+1:]
	}
	if period := bytes.Index(name, dot); period >= 0 {
		name = name[period+1:]
	}
	name = bytes.Replace(name, centerDot, dot, -1)
	return name
}

func timeFormat(t time.Time) string {
	var timeString = t.Format("02.01.2006 15:04:05")
	return timeString
}
