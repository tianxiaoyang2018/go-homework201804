package runner

import (
	"errors"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/p1cn/tantan-backend-common/config"
	"github.com/p1cn/tantan-backend-common/health"
	"github.com/p1cn/tantan-backend-common/log"
	"github.com/p1cn/tantan-backend-common/util"
)

var (
	ErrTimeout = errors.New("timeout")
)

type Service interface {
	Start() error
	Stop() error
	GetHealthChecks() []health.HealthCheck
}

type ServiceRunner interface {
	Wait()
}

func RunService(s Service) ServiceRunner {
	r := newServiceRunner(s)
	r.run()
	return r
}

func newServiceRunner(s Service) *serviceRunner {
	return &serviceRunner{
		signals: make(chan os.Signal, 1),
		service: s,
	}
}

type serviceRunner struct {
	signals chan os.Signal
	service Service

	stopped bool

	wg sync.WaitGroup
}

func (r *serviceRunner) run() {
	r.wg.Add(1)
	go r.handleSignal()
	go r.handleStart()
}

func (r *serviceRunner) registerHealthChecks() {
	for _, hc := range r.service.GetHealthChecks() {
		health.RegisterHealthCheck(hc)
	}
}

func (r *serviceRunner) handleStart() {
	func() {
		defer util.Recovery()
		err := r.service.Start()
		if err != nil {
			log.Err("%v", err)
		}
	}()
	if !r.stopped {
		r.wg.Done()
	}
}

func (r *serviceRunner) handleSignal() {
	signal.Notify(r.signals, syscall.SIGPIPE, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGABRT)
	for {
		select {
		case sig := <-r.signals:
			log.Info("received signal : %x", sig)
			switch sig {
			case syscall.SIGPIPE:
			default:
				go func() {
					to := 10 * time.Second
					if config.GetCommonConfig().GracefulTimeout.Duration() > 1*time.Second {
						to = config.GetCommonConfig().GracefulTimeout.Duration()
					}

					time.Sleep(to)
					log.Err("graceful timeout")
					os.Exit(1)
				}()
				r.stopped = true
				r.service.Stop()
				r.wg.Done()
			}
		}
	}
}

func (r *serviceRunner) Wait() {
	r.wg.Wait()

	for i := 0; i < 3; i++ {
		err := log.Close()
		if err == nil {
			return
		}
		log.Err("log.Close error : err(%+v)", err)
	}
}
