/*
 *  \
 *  \\,
 *   \\\,^,.,,.                    “Zero to Hero”
 *   ,;7~((\))`;;,,               <zerotohero.dev>
 *   ,(@') ;)`))\;;',    stay up to date, be curious: learn
 *    )  . ),((  ))\;,
 *   /;`,,/7),)) )) )\,,
 *  (& )`   (,((,((;( ))\,
 */

package log

import (
	"fmt"
	"github.com/zerotohero-dev/fizz-env/pkg/env"
	"log"
	"log/syslog"
)

var writer *syslog.Writer
var environment env.FizzEnv



func Init(e env.FizzEnv, appName string) *syslog.Writer {
	e.Log.Sanitize()

	dest := e.Log.Destination

	// Don’t log to Syslog in development mode.
	if e.IsDevelopment() {
		return nil
	}

	w, err := syslog.Dial("udp", dest, syslog.LOG_INFO|syslog.LOG_USER, appName)
	if err != nil {
		Info("failed to dial syslog for log destination '" + dest + "'.")
		return nil
	}

	writer = w

	return writer
}

func Info(s string, args ...interface{}) {
	if environment.Deployment.Type == env.Development || writer == nil {
		log.Printf(s, args...)
		return
	}

	_ = writer.Info(fmt.Sprintf(s, args...))
}

func Err(s string, args ...interface{}) {
	if environment.Deployment.Type == env.Development || writer == nil {
		log.Printf(s, args...)
		return
	}

	_ = writer.Err(fmt.Sprintf(s, args...))
}

func Warning(s string, args ...interface{}) {
	if environment.Deployment.Type == env.Development || writer == nil {
		log.Printf(s, args...)
		return
	}

	_ = writer.Warning(fmt.Sprintf(s, args...))
}

func Fatal(e interface{}) {
	log.Fatal(e)
}
