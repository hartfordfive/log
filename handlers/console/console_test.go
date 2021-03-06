package console

import (
	"bytes"
	"testing"
	"time"

	"github.com/go-playground/log"
	. "gopkg.in/go-playground/assert.v1"
)

// NOTES:
// - Run "go test" to run tests
// - Run "gocov test | gocov report" to report on test converage by file
// - Run "gocov test | gocov annotate -" to report on all code and functions, those ,marked with "MISS" were never called

// or

// -- may be a good idea to change to output path to somewherelike /tmp
// go test -coverprofile cover.out && go tool cover -html=cover.out -o cover.html

func TestConsoleLogger(t *testing.T) {

	buff := new(bytes.Buffer)

	cLog := New()
	cLog.SetWriter(buff)
	cLog.DisplayColor(false)
	cLog.SetChannelBuffer(3)
	cLog.SetTimestampFormat(time.RFC3339)
	cLog.UseMiniTimestamp(true)
	cLog.SetANSIReset(log.Reset)

	log.RegisterHandler(cLog, log.AllLevels...)

	log.Debug("debug")
	Equal(t, buff.String(), " DEBUG[0000] debug\n")
	buff.Reset()

	log.Debugf("%s", "debugf")
	Equal(t, buff.String(), " DEBUG[0000] debugf\n")
	buff.Reset()

	log.Info("info")
	Equal(t, buff.String(), "  INFO[0000] info\n")
	buff.Reset()

	log.Infof("%s", "infof")
	Equal(t, buff.String(), "  INFO[0000] infof\n")
	buff.Reset()

	log.Notice("notice")
	Equal(t, buff.String(), "NOTICE[0000] notice\n")
	buff.Reset()

	log.Noticef("%s", "noticef")
	Equal(t, buff.String(), "NOTICE[0000] noticef\n")
	buff.Reset()

	log.Warn("warn")
	Equal(t, buff.String(), "  WARN[0000] warn\n")
	buff.Reset()

	log.Warnf("%s", "warnf")
	Equal(t, buff.String(), "  WARN[0000] warnf\n")
	buff.Reset()

	log.Error("error")
	Equal(t, buff.String(), " ERROR[0000] error\n")
	buff.Reset()

	log.Errorf("%s", "errorf")
	Equal(t, buff.String(), " ERROR[0000] errorf\n")
	buff.Reset()

	log.Alert("alert")
	Equal(t, buff.String(), " ALERT[0000] alert\n")
	buff.Reset()

	log.Alertf("%s", "alertf")
	Equal(t, buff.String(), " ALERT[0000] alertf\n")
	buff.Reset()

	log.Print("print")
	Equal(t, buff.String(), "  INFO[0000] print\n")
	buff.Reset()

	log.Printf("%s", "printf")
	Equal(t, buff.String(), "  INFO[0000] printf\n")
	buff.Reset()

	log.Println("println")
	Equal(t, buff.String(), "  INFO[0000] println\n")
	buff.Reset()

	PanicMatches(t, func() { log.Panic("panic") }, "panic")
	Equal(t, buff.String(), " PANIC[0000] panic\n")
	buff.Reset()

	PanicMatches(t, func() { log.Panicf("%s", "panicf") }, "panicf")
	Equal(t, buff.String(), " PANIC[0000] panicf\n")
	buff.Reset()

	PanicMatches(t, func() { log.Panicln("panicln") }, "panicln")
	Equal(t, buff.String(), " PANIC[0000] panicln\n")
	buff.Reset()

	// WithFields
	log.WithFields(log.F("key", "value")).Debug("debug")
	Equal(t, buff.String(), " DEBUG[0000] debug                     key=value\n")
	buff.Reset()

	log.WithFields(log.F("key", "value")).Debugf("%s", "debugf")
	Equal(t, buff.String(), " DEBUG[0000] debugf                    key=value\n")
	buff.Reset()

	log.WithFields(log.F("key", "value")).Info("info")
	Equal(t, buff.String(), "  INFO[0000] info                      key=value\n")
	buff.Reset()

	log.WithFields(log.F("key", "value")).Infof("%s", "infof")
	Equal(t, buff.String(), "  INFO[0000] infof                     key=value\n")
	buff.Reset()

	log.WithFields(log.F("key", "value")).Notice("notice")
	Equal(t, buff.String(), "NOTICE[0000] notice                    key=value\n")
	buff.Reset()

	log.WithFields(log.F("key", "value")).Noticef("%s", "noticef")
	Equal(t, buff.String(), "NOTICE[0000] noticef                   key=value\n")
	buff.Reset()

	log.WithFields(log.F("key", "value")).Warn("warn")
	Equal(t, buff.String(), "  WARN[0000] warn                      key=value\n")
	buff.Reset()

	log.WithFields(log.F("key", "value")).Warnf("%s", "warnf")
	Equal(t, buff.String(), "  WARN[0000] warnf                     key=value\n")
	buff.Reset()

	log.WithFields(log.F("key", "value")).Error("error")
	Equal(t, buff.String(), " ERROR[0000] error                     key=value\n")
	buff.Reset()

	log.WithFields(log.F("key", "value")).Errorf("%s", "errorf")
	Equal(t, buff.String(), " ERROR[0000] errorf                    key=value\n")
	buff.Reset()

	log.WithFields(log.F("key", "value")).Alert("alert")
	Equal(t, buff.String(), " ALERT[0000] alert                     key=value\n")
	buff.Reset()

	log.WithFields(log.F("key", "value")).Alertf("%s", "alertf")
	Equal(t, buff.String(), " ALERT[0000] alertf                    key=value\n")
	buff.Reset()

	PanicMatches(t, func() { log.WithFields(log.F("key", "value")).Panicf("%s", "panicf") }, "panicf key=value")
	Equal(t, buff.String(), " PANIC[0000] panicf                    key=value\n")
	buff.Reset()

	PanicMatches(t, func() { log.WithFields(log.F("key", "value")).Panic("panic") }, "panic key=value")
	Equal(t, buff.String(), " PANIC[0000] panic                     key=value\n")
	buff.Reset()

	func() {
		defer log.Trace("trace").End()
	}()

	// TODO: finish up regex
	MatchRegex(t, buff.String(), "^\\sTRACE\\[0000\\]\\strace\\s+\\.*")
	buff.Reset()

	func() {
		defer log.Tracef("tracef").End()
	}()

	// TODO: finish up regex
	MatchRegex(t, buff.String(), "^\\sTRACE\\[0000\\]\\stracef\\s+\\.*")
	buff.Reset()

	func() {
		defer log.WithFields(log.F("key", "value")).Trace("trace").End()
	}()

	// TODO: finish up regex
	MatchRegex(t, buff.String(), "^\\sTRACE\\[0000\\]\\strace\\s+\\.*")
	buff.Reset()

	func() {
		defer log.WithFields(log.F("key", "value")).Tracef("tracef").End()
	}()

	// TODO: finish up regex
	MatchRegex(t, buff.String(), "^\\sTRACE\\[0000\\]\\stracef\\s+\\.*")
	buff.Reset()

	year := time.Now().Format("2006")
	cLog.UseMiniTimestamp(false)
	cLog.SetTimestampFormat("2006")

	log.Info("info")
	Equal(t, buff.String(), "  INFO["+year+"] info\n")
	buff.Reset()
}

func TestConsoleLoggerColor(t *testing.T) {

	buff := new(bytes.Buffer)

	cLog := New()
	cLog.SetWriter(buff)
	cLog.DisplayColor(true)
	cLog.SetChannelBuffer(3)
	cLog.SetTimestampFormat(time.RFC3339)
	cLog.UseMiniTimestamp(true)

	log.RegisterHandler(cLog, log.AllLevels...)

	log.Debug("debug")
	Equal(t, buff.String(), "[32m DEBUG[0m[0000] debug\n")
	buff.Reset()

	log.Debugf("%s", "debugf")
	Equal(t, buff.String(), "[32m DEBUG[0m[0000] debugf\n")
	buff.Reset()

	log.Info("info")
	Equal(t, buff.String(), "[34m  INFO[0m[0000] info\n")
	buff.Reset()

	log.Infof("%s", "infof")
	Equal(t, buff.String(), "[34m  INFO[0m[0000] infof\n")
	buff.Reset()

	log.Notice("notice")
	Equal(t, buff.String(), "[36;1mNOTICE[0m[0000] notice\n")
	buff.Reset()

	log.Notice("%s", "noticef")
	Equal(t, buff.String(), "[36;1mNOTICE[0m[0000] %snoticef\n")
	buff.Reset()

	log.Warn("warn")
	Equal(t, buff.String(), "[33;1m  WARN[0m[0000] warn\n")
	buff.Reset()

	log.Warnf("%s", "warnf")
	Equal(t, buff.String(), "[33;1m  WARN[0m[0000] warnf\n")
	buff.Reset()

	log.Error("error")
	Equal(t, buff.String(), "[31;1m ERROR[0m[0000] error\n")
	buff.Reset()

	log.Errorf("%s", "errorf")
	Equal(t, buff.String(), "[31;1m ERROR[0m[0000] errorf\n")
	buff.Reset()

	log.Alert("alert")
	Equal(t, buff.String(), "[31m[4m ALERT[0m[0000] alert\n")
	buff.Reset()

	log.Alertf("%s", "alertf")
	Equal(t, buff.String(), "[31m[4m ALERT[0m[0000] alertf\n")
	buff.Reset()

	log.Print("print")
	Equal(t, buff.String(), "[34m  INFO[0m[0000] print\n")
	buff.Reset()

	log.Printf("%s", "printf")
	Equal(t, buff.String(), "[34m  INFO[0m[0000] printf\n")
	buff.Reset()

	log.Println("println")
	Equal(t, buff.String(), "[34m  INFO[0m[0000] println\n")
	buff.Reset()

	PanicMatches(t, func() { log.Panic("panic") }, "panic")
	Equal(t, buff.String(), "[31m PANIC[0m[0000] panic\n")
	buff.Reset()

	PanicMatches(t, func() { log.Panicf("%s", "panicf") }, "panicf")
	Equal(t, buff.String(), "[31m PANIC[0m[0000] panicf\n")
	buff.Reset()

	PanicMatches(t, func() { log.Panicln("panicln") }, "panicln")
	Equal(t, buff.String(), "[31m PANIC[0m[0000] panicln\n")
	buff.Reset()

	// WithFields
	log.WithFields(log.F("key", "value")).Debug("debug")
	Equal(t, buff.String(), "[32m DEBUG[0m[0000] debug                     [32mkey[0m=value\n")
	buff.Reset()

	log.WithFields(log.F("key", "value")).Debugf("%s", "debugf")
	Equal(t, buff.String(), "[32m DEBUG[0m[0000] debugf                    [32mkey[0m=value\n")
	buff.Reset()

	log.WithFields(log.F("key", "value")).Info("info")
	Equal(t, buff.String(), "[34m  INFO[0m[0000] info                      [34mkey[0m=value\n")
	buff.Reset()

	log.WithFields(log.F("key", "value")).Infof("%s", "infof")
	Equal(t, buff.String(), "[34m  INFO[0m[0000] infof                     [34mkey[0m=value\n")
	buff.Reset()

	log.WithFields(log.F("key", "value")).Notice("notice")
	Equal(t, buff.String(), "[36;1mNOTICE[0m[0000] notice                    [36;1mkey[0m=value\n")
	buff.Reset()

	log.WithFields(log.F("key", "value")).Noticef("%s", "noticef")
	Equal(t, buff.String(), "[36;1mNOTICE[0m[0000] noticef                   [36;1mkey[0m=value\n")
	buff.Reset()

	log.WithFields(log.F("key", "value")).Warn("warn")
	Equal(t, buff.String(), "[33;1m  WARN[0m[0000] warn                      [33;1mkey[0m=value\n")
	buff.Reset()

	log.WithFields(log.F("key", "value")).Warnf("%s", "warnf")
	Equal(t, buff.String(), "[33;1m  WARN[0m[0000] warnf                     [33;1mkey[0m=value\n")
	buff.Reset()

	log.WithFields(log.F("key", "value")).Error("error")
	Equal(t, buff.String(), "[31;1m ERROR[0m[0000] error                     [31;1mkey[0m=value\n")
	buff.Reset()

	log.WithFields(log.F("key", "value")).Errorf("%s", "errorf")
	Equal(t, buff.String(), "[31;1m ERROR[0m[0000] errorf                    [31;1mkey[0m=value\n")
	buff.Reset()

	log.WithFields(log.F("key", "value")).Alert("alert")
	Equal(t, buff.String(), "[31m[4m ALERT[0m[0000] alert                     [31m[4mkey[0m=value\n")
	buff.Reset()

	log.WithFields(log.F("key", "value")).Alertf("%s", "alertf")
	Equal(t, buff.String(), "[31m[4m ALERT[0m[0000] alertf                    [31m[4mkey[0m=value\n")
	buff.Reset()

	PanicMatches(t, func() { log.WithFields(log.F("key", "value")).Panicf("%s", "panicf") }, "panicf key=value")
	Equal(t, buff.String(), "[31m PANIC[0m[0000] panicf                    [31mkey[0m=value\n")
	buff.Reset()

	PanicMatches(t, func() { log.WithFields(log.F("key", "value")).Panic("panic") }, "panic key=value")
	Equal(t, buff.String(), "[31m PANIC[0m[0000] panic                     [31mkey[0m=value\n")
	buff.Reset()

	cLog.SetLevelColor(log.DebugLevel, log.LightGreen)

	log.Debug("debug")
	Equal(t, buff.String(), "[32;1m DEBUG[0m[0000] debug\n")
	buff.Reset()

	year := time.Now().Format("2006")
	cLog.UseMiniTimestamp(false)
	cLog.SetTimestampFormat("2006")

	log.Info("info")
	Equal(t, buff.String(), "[34m  INFO[0m["+year+"] info\n")
	buff.Reset()

	log.WithFields(log.F("key", "value")).Info("info")
	Equal(t, buff.String(), "[34m  INFO[0m["+year+"] info                      [34mkey[0m=value\n")
	buff.Reset()
}
