// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"

	gokitlog "github.com/go-kit/log"
	"github.com/go-stack/stack"
	"github.com/mattn/go-isatty"
	"gopkg.in/ini.v1"

	"github.com/grafana/grafana/pkg/infra/log/level"
	"github.com/grafana/grafana/pkg/infra/log/term"
	"github.com/grafana/grafana/pkg/util"
	"github.com/grafana/grafana/pkg/util/errutil"
)

var loggersToClose []DisposableHandler
var loggersToReload []ReloadableHandler
var root *logManager

const (
	// top 7 calls in the stack are within logger
	DefaultCallerDepth = 7
	CallerContextKey   = "caller"
)

func init() {
	loggersToClose = make([]DisposableHandler, 0)
	loggersToReload = make([]ReloadableHandler, 0)

	// Use console by default
	format := getLogFormat("console")
	logger := level.NewFilter(format(os.Stderr), level.AllowInfo())
	root = newManager(logger)
}

// logManager manage loggers
type logManager struct {
	*ConcreteLogger
	loggersByName map[string]*ConcreteLogger
	logFilters    []LogWithFilters
	mutex         sync.RWMutex
}

func newManager(logger gokitlog.Logger) *logManager {
	return &logManager{
		ConcreteLogger: newConcreteLogger(logger),
		loggersByName:  map[string]*ConcreteLogger{},
	}
}

func (lm *logManager) initialize(loggers []LogWithFilters) {
	lm.mutex.Lock()
	defer lm.mutex.Unlock()

	defaultLoggers := make([]gokitlog.Logger, len(loggers))
	for index, logger := range loggers {
		defaultLoggers[index] = level.NewFilter(logger.val, logger.maxLevel)
	}

	lm.ConcreteLogger.SetLogger(&compositeLogger{loggers: defaultLoggers})
	lm.logFilters = loggers

	loggersByName := []string{}
	for k := range lm.loggersByName {
		loggersByName = append(loggersByName, k)
	}
	sort.Strings(loggersByName)

	for _, name := range loggersByName {
		ctxLoggers := make([]gokitlog.Logger, len(loggers))

		for index, logger := range loggers {
			if filterLevel, exists := logger.filters[name]; !exists {
				ctxLoggers[index] = level.NewFilter(logger.val, logger.maxLevel)
			} else {
				ctxLoggers[index] = level.NewFilter(logger.val, filterLevel)
			}
		}

		lm.loggersByName[name].SetLogger(&compositeLogger{loggers: ctxLoggers})
	}
}

func (lm *logManager) SetLogger(logger gokitlog.Logger) {
	lm.ConcreteLogger.SetLogger(logger)
}

func (lm *logManager) GetLogger() gokitlog.Logger {
	return lm.ConcreteLogger.GetLogger()
}

func (lm *logManager) Log(args ...interface{}) error {
	lm.mutex.RLock()
	defer lm.mutex.RUnlock()
	if err := lm.ConcreteLogger.Log(args...); err != nil {
		log.Println("Logging error", "error", err)
	}

	return nil
}

func (lm *logManager) New(ctx ...interface{}) *ConcreteLogger {
	lm.mutex.Lock()
	defer lm.mutex.Unlock()
	if len(ctx) == 0 {
		return lm.ConcreteLogger
	}

	loggerName, ok := ctx[0].(string)
	if !ok {
		return lm.ConcreteLogger
	}

	if logger, exists := lm.loggersByName[loggerName]; exists {
		return logger
	}

	ctx = append([]interface{}{"logger"}, ctx...)

	if len(lm.logFilters) == 0 {
		ctxLogger := newConcreteLogger(lm.logger, ctx...)
		lm.loggersByName[loggerName] = ctxLogger
		return ctxLogger
	}

	compositeLogger := newCompositeLogger()
	for _, logWithFilter := range lm.logFilters {
		filterLevel, ok := logWithFilter.filters[loggerName]
		if ok {
			logWithFilter.val = level.NewFilter(logWithFilter.val, filterLevel)
		} else {
			logWithFilter.val = level.NewFilter(logWithFilter.val, logWithFilter.maxLevel)
		}

		compositeLogger.loggers = append(compositeLogger.loggers, logWithFilter.val)
	}

	ctxLogger := newConcreteLogger(compositeLogger, ctx...)
	lm.loggersByName[loggerName] = ctxLogger
	return ctxLogger
}

type ConcreteLogger struct {
	ctx    []interface{}
	logger gokitlog.Logger
	mutex  sync.RWMutex
}

func newConcreteLogger(logger gokitlog.Logger, ctx ...interface{}) *ConcreteLogger {
	if len(ctx) == 0 {
		ctx = []interface{}{}
	} else {
		logger = gokitlog.With(logger, ctx...)
	}

	return &ConcreteLogger{
		ctx:    ctx,
		logger: logger,
	}
}

func (cl *ConcreteLogger) SetLogger(logger gokitlog.Logger) {
	cl.mutex.Lock()
	cl.logger = gokitlog.With(logger, cl.ctx...)
	cl.mutex.Unlock()
}

func (cl *ConcreteLogger) GetLogger() gokitlog.Logger {
	cl.mutex.Lock()
	defer cl.mutex.Unlock()
	return cl.logger
}

func (cl *ConcreteLogger) Warn(msg string, args ...interface{}) {
	_ = cl.log(msg, level.WarnValue(), args...)
}

func (cl *ConcreteLogger) Debug(msg string, args ...interface{}) {
	// args = append([]interface{}{level.Key(), level.DebugValue(), "msg", msg}, args...)
	_ = cl.log(msg, level.DebugValue(), args...)
}

func (cl *ConcreteLogger) Error(msg string, args ...interface{}) {
	_ = cl.log(msg, level.ErrorValue(), args...)
}

func (cl *ConcreteLogger) Info(msg string, args ...interface{}) {
	_ = cl.log(msg, level.InfoValue(), args...)
}

func (cl *ConcreteLogger) log(msg string, logLevel level.Value, args ...interface{}) error {
	cl.mutex.RLock()
	logger := gokitlog.With(cl.logger, "t", gokitlog.TimestampFormat(time.Now, "2006-01-02T15:04:05.99-0700"))
	cl.mutex.RUnlock()

	args = append([]interface{}{level.Key(), logLevel, "msg", msg}, args...)

	return logger.Log(args...)
}

func (cl *ConcreteLogger) Log(keyvals ...interface{}) error {
	cl.mutex.RLock()
	defer cl.mutex.RUnlock()
	return cl.logger.Log(keyvals...)
}

func (cl *ConcreteLogger) New(ctx ...interface{}) *ConcreteLogger {
	if len(ctx) == 0 {
		root.New()
	}

	keyvals := []interface{}{}

	if len(cl.ctx)%2 == 1 {
		cl.ctx = append(cl.ctx, nil)
	}

	for i := 0; i < len(cl.ctx); i += 2 {
		k, v := cl.ctx[i], cl.ctx[i+1]

		if k == "logger" {
			continue
		}

		keyvals = append(keyvals, k, v)
	}

	keyvals = append(keyvals, ctx...)

	return root.New(keyvals...)
}

func New(ctx ...interface{}) *ConcreteLogger {
	return root.New(ctx...)
}

type LogWithFilters struct {
	val      gokitlog.Logger
	filters  map[string]level.Option
	maxLevel level.Option
}

func with(ctxLogger *ConcreteLogger, withFunc func(gokitlog.Logger, ...interface{}) gokitlog.Logger, ctx []interface{}) *ConcreteLogger {
	if len(ctx) == 0 {
		return ctxLogger
	}

	ctxLogger.logger = withFunc(ctxLogger.logger, ctx...)
	return ctxLogger
}

// WithPrefix adds context that will be added to the log message
func WithPrefix(ctxLogger *ConcreteLogger, ctx ...interface{}) *ConcreteLogger {
	return with(ctxLogger, gokitlog.WithPrefix, ctx)
}

// WithSuffix adds context that will be appended at the end of the log message
func WithSuffix(ctxLogger *ConcreteLogger, ctx ...interface{}) *ConcreteLogger {
	return with(ctxLogger, gokitlog.WithSuffix, ctx)
}

var logLevels = map[string]level.Option{
	"trace":    level.AllowDebug(),
	"debug":    level.AllowDebug(),
	"info":     level.AllowInfo(),
	"warn":     level.AllowWarn(),
	"error":    level.AllowError(),
	"critical": level.AllowError(),
}

func getLogLevelFromConfig(key string, defaultName string, cfg *ini.File) (string, level.Option) {
	levelName := cfg.Section(key).Key("level").MustString(defaultName)
	levelName = strings.ToLower(levelName)
	level := getLogLevelFromString(levelName)
	return levelName, level
}

func getLogLevelFromString(levelName string) level.Option {
	loglevel, ok := logLevels[levelName]

	if !ok {
		_ = level.Error(root).Log("Unknown log level", "level", levelName)
		return level.AllowError()
	}

	return loglevel
}

// the filter is composed with logger name and level
func getFilters(filterStrArray []string) map[string]level.Option {
	filterMap := make(map[string]level.Option)

	for _, filterStr := range filterStrArray {
		parts := strings.Split(filterStr, ":")
		if len(parts) > 1 {
			filterMap[parts[0]] = getLogLevelFromString(parts[1])
		}
	}

	return filterMap
}

func Stack(skip int) string {
	call := stack.Caller(skip)
	s := stack.Trace().TrimBelow(call).TrimRuntime()
	return s.String()
}

// StackCaller returns a go-kit Valuer function that returns the stack trace from the place it is called. Argument `skip` allows skipping top n lines from the stack.
func StackCaller(skip int) gokitlog.Valuer {
	return func() interface{} {
		return Stack(skip + 1)
	}
}

// Caller proxies go-kit/log Caller and returns a Valuer function that returns a file and line from a specified depth
// in the callstack
func Caller(depth int) gokitlog.Valuer {
	return gokitlog.Caller(depth)
}

type Formatedlogger func(w io.Writer) gokitlog.Logger

func getLogFormat(format string) Formatedlogger {
	switch format {
	case "console":
		if isatty.IsTerminal(os.Stdout.Fd()) {
			return func(w io.Writer) gokitlog.Logger {
				return term.NewTerminalLogger(w)
			}
		}
		return func(w io.Writer) gokitlog.Logger {
			return gokitlog.NewLogfmtLogger(w)
		}
	case "text":
		return func(w io.Writer) gokitlog.Logger {
			return gokitlog.NewLogfmtLogger(w)
		}
	case "json":
		return func(w io.Writer) gokitlog.Logger {
			return gokitlog.NewJSONLogger(gokitlog.NewSyncWriter(w))
		}
	default:
		return func(w io.Writer) gokitlog.Logger {
			return gokitlog.NewLogfmtLogger(w)
		}
	}
}

// this is for file logger only
func Close() error {
	var err error
	for _, logger := range loggersToClose {
		if e := logger.Close(); e != nil && err == nil {
			err = e
		}
	}
	loggersToClose = make([]DisposableHandler, 0)

	return err
}

// Reload reloads all loggers.
func Reload() error {
	for _, logger := range loggersToReload {
		if err := logger.Reload(); err != nil {
			return err
		}
	}

	return nil
}

func ReadLoggingConfig(modes []string, logsPath string, cfg *ini.File) error {
	if err := Close(); err != nil {
		return err
	}

	defaultLevelName, _ := getLogLevelFromConfig("log", "info", cfg)
	defaultFilters := getFilters(util.SplitString(cfg.Section("log").Key("filters").String()))

	var configLoggers []LogWithFilters
	for _, mode := range modes {
		mode = strings.TrimSpace(mode)
		sec, err := cfg.GetSection("log." + mode)
		if err != nil {
			_ = level.Error(root).Log("Unknown log mode", "mode", mode)
			return errutil.Wrapf(err, "failed to get config section log.%s", mode)
		}

		// Log level.
		_, leveloption := getLogLevelFromConfig("log."+mode, defaultLevelName, cfg)
		modeFilters := getFilters(util.SplitString(sec.Key("filters").String()))

		format := getLogFormat(sec.Key("format").MustString(""))

		var handler LogWithFilters

		switch mode {
		case "console":
			handler.val = format(os.Stdout)
		case "file":
			fileName := sec.Key("file_name").MustString(filepath.Join(logsPath, "grafana.log"))
			dpath := filepath.Dir(fileName)
			if err := os.MkdirAll(dpath, os.ModePerm); err != nil {
				_ = level.Error(root).Log("Failed to create directory", "dpath", dpath, "err", err)
				continue
			}
			fileHandler := NewFileWriter()
			fileHandler.Filename = fileName
			fileHandler.Format = format
			fileHandler.Rotate = sec.Key("log_rotate").MustBool(true)
			fileHandler.Maxlines = sec.Key("max_lines").MustInt(1000000)
			fileHandler.Maxsize = 1 << uint(sec.Key("max_size_shift").MustInt(28))
			fileHandler.Daily = sec.Key("daily_rotate").MustBool(true)
			fileHandler.Maxdays = sec.Key("max_days").MustInt64(7)
			if err := fileHandler.Init(); err != nil {
				_ = level.Error(root).Log("Failed to initialize file handler", "dpath", dpath, "err", err)
				continue
			}

			loggersToClose = append(loggersToClose, fileHandler)
			loggersToReload = append(loggersToReload, fileHandler)
			handler.val = fileHandler
		case "syslog":
			sysLogHandler := NewSyslog(sec, format)
			loggersToClose = append(loggersToClose, sysLogHandler)
			handler.val = sysLogHandler.logger
		}
		if handler.val == nil {
			panic(fmt.Sprintf("Handler is uninitialized for mode %q", mode))
		}

		// join default filters and mode filters together
		for key, value := range defaultFilters {
			if _, exist := modeFilters[key]; !exist {
				modeFilters[key] = value
			}
		}

		handler.filters = modeFilters
		handler.maxLevel = leveloption
		configLoggers = append(configLoggers, handler)
	}

	if len(configLoggers) > 0 {
		root.initialize(configLoggers)
	}

	return nil
}
