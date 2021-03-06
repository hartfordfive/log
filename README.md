## log
<img align="right" src="https://raw.githubusercontent.com/go-playground/log/master/logo.png">
![Project status](https://img.shields.io/badge/version-1.3-green.svg)
[![Build Status](https://semaphoreci.com/api/v1/joeybloggs/log/branches/master/badge.svg)](https://semaphoreci.com/joeybloggs/log)
[![Coverage Status](https://coveralls.io/repos/github/go-playground/log/badge.svg?branch=master)](https://coveralls.io/github/go-playground/log?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-playground/log)](https://goreportcard.com/report/github.com/go-playground/log)
[![GoDoc](https://godoc.org/github.com/go-playground/log?status.svg)](https://godoc.org/github.com/go-playground/log)
![License](https://img.shields.io/dub/l/vibe-d.svg)

Log is a simple,highly configurable, Structured Logging that is a near drop in replacement for the std library log

Why another logging library?
----------------------------
There's allot of great stuff out there, but also thought a log library could be made easier to use, more efficient by reusing objects and more performant using channels.

Features
--------
- [x] Logger is simple, only logic to create the log entry and send it off to the handlers and they take it from there.
- [x] Sends the log entry to the handlers asynchronously, but waits for all to complete; meaning all your handlers can be dealing with the log entry at the same time, but log will wait until all have completed before moving on.
- [x] Ability to specify which log levels get sent to each handler
- [x] Built-in console + syslog
- [x] Handlers are simple to write + easy to register
- [x] Logger is a singleton ( one of the few instances a singleton is desired ) so the root package registers which handlers are used and any libraries just follow suit.

Installation
-----------

Use go get 

```go
go get github.com/go-playground/log
``` 

Replacing std log
-----------------
change import from
```go
import "log"
```
to
```go
import "github.com/go-playground/log"
```

Usage
------
import the log package, setup at least one handler
```go
package main

import (
	"github.com/go-playground/log"
	"github.com/go-playground/log/handlers/console"
)

func main() {

	cLog := console.New()

	log.RegisterHandler(cLog, log.AllLevels...)

	// Trace
	defer log.Trace("trace").End()

	log.Debug("debug")
	log.Info("info")
	log.Notice("notice")
	log.Warn("warn")
	log.Error("error")
	// log.Panic("panic") // this will panic
	log.Alert("alert")
	// log.Fatal("fatal") // this will call os.Exit(1)

	// logging with fields can be used with any of the above
	log.WithFields(log.F("key", "value")).Info("test info")
}
```

Adding your own Handler
```go
package main

import (
	"bytes"
	"fmt"

	"github.com/go-playground/log"
)

// CustomHandler is your custom handler
type CustomHandler struct {
	// whatever properties you need
	buffer uint // channel buffer
}

// Run starts the logger consuming on the returned channed
func (c *CustomHandler) Run() chan<- *log.Entry {

	// in a big high traffic app, set a higher buffer
	ch := make(chan *log.Entry, c.buffer)

	// can run as many consumers on the channel as you want,
	// depending on the buffer size or your needs
	go func(entries <-chan *log.Entry) {

		// below prints to os.Stderr but could marshal to JSON
		// and send to central logging server
		//																						       ---------
		// 				                                                                 |----------> | console |
		//                                                                               |             ---------
		// i.e. -----------------               -----------------     Unmarshal    -------------       --------
		//     | app log handler | -- json --> | central log app | --    to    -> | log handler | --> | syslog |
		//      -----------------               -----------------       Entry      -------------       --------
		//      																         |             ---------
		//                                  									         |----------> | DataDog |
		//          																	        	   ---------
		var e *log.Entry
		b := new(bytes.Buffer)

		for e = range entries {

			b.Reset()
			b.WriteString(e.Message)

			for _, f := range e.Fields {
				fmt.Fprintf(b, " %s=%v", f.Key, f.Value)
			}

			fmt.Println(b.String())
			e.WG.Done() // done writing the entry
		}

	}(ch)

	return ch
}

func main() {

	cLog := &CustomHandler{
		buffer: 0,
	}

	log.RegisterHandler(cLog, log.AllLevels...)

	// Trace
	defer log.Trace("trace").End()

	log.Debug("debug")
	log.Info("info")
	log.Notice("notice")
	log.Warn("warn")
	log.Error("error")
	// log.Panic("panic") // this will panic
	log.Alert("alert")
	// log.Fatal("fatal") // this will call os.Exit(1)

	// logging with fields can be used with any of the above
	log.WithFields(log.F("key", "value")).Info("test info")
}
```

Log Level Definitions
---------------------

**DebugLevel** - Info useful to developers for debugging the application, not useful during operations.

**TraceLevel** - Info useful to developers for debugging the application and reporting on possible bottlenecks.

**InfoLevel** - Normal operational messages - may be harvested for reporting, measuring throughput, etc. - no action required.

**NoticeLevel** - Normal but significant condition. Events that are unusual but not error conditions - might be summarized in an email to developers or admins to spot potential problems - no immediate action required.

**WarnLevel** - Warning messages, not an error, but indication that an error will occur if action is not taken, e.g. file system 85% full - each item must be resolved within a given time.

**ErrorLevel** - Non-urgent failures, these should be relayed to developers or admins; each item must be resolved within a given time.

**PanicLevel** - A "panic" condition usually affecting multiple apps/servers/sites. At this level it would usually notify all tech staff on call.

**AlertLevel** - Action must be taken immediately. Should be corrected immediately, therefore notify staff who can fix the problem. An example would be the loss of a primary ISP connection.

**FatalLevel** - Should be corrected immediately, but indicates failure in a primary system, an example is a loss of a backup ISP connection. ( same as SYSLOG CRITICAL )

More Handlers
-------------
Pull requests for new handlers are welcome, please provide test coverage is all I ask.

Special Thanks
--------------
Special thanks to the following libraries that inspired
* [logrus](https://github.com/Sirupsen/logrus) - Structured, pluggable logging for Go.
* [apex log](https://github.com/apex/log) - Structured logging package for Go.
