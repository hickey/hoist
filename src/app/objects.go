package app

import (
	"github.com/op/go-logging"
	"github.com/spf13/viper"
	"sync"
	"sync/atomic"
)

type Once struct {
	m    sync.Mutex
	done uint32
	wg   sync.WaitGroup
	//settings Settings
	logger logging.Logger
}

// Variable for finding application objects
var Objects Once

func init() {
	Objects.Start()
}

func (o *Once) Start() {
	//var err error
	var log_level logging.Level

	if atomic.LoadUint32(&o.done) == 1 { // <-- Check
		return
	}
	// Slow-path.
	o.m.Lock() // <-- Lock
	defer o.m.Unlock()
	if o.done == 0 { // <-- Check
		defer atomic.StoreUint32(&o.done, 1)

		// Application initialization
		o.logger = *logging.MustGetLogger("hoist")
		format := logging.MustStringFormatter(
			`%{color}%{shortpkg}.%{longfunc} ▶ %{level:.5s} %{color:reset} %{message}`,
		)
		logging.SetFormatter(format)
		switch viper.Get("logging.level") {
		case "CRITICAL":
			log_level = logging.CRITICAL
		case "ERROR":
			log_level = logging.ERROR
		case "WARNING":
			log_level = logging.WARNING
		case "NOTICE":
			log_level = logging.NOTICE
		case "INFO":
			log_level = logging.INFO
		case "DEBUG":
			log_level = logging.DEBUG
		default:
			log_level = logging.WARNING
		}

		logging.SetLevel(log_level, "")
	}
}

func GetWaitGroup() sync.WaitGroup {
	return Objects.wg
}

func GetLogger() *logging.Logger {
	return &Objects.logger
}
