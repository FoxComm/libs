// Package announcer enables microservices to publish their location to etcd, primarily to be read by the router.
package announcer

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/FoxComm/FoxComm/configs"
	"github.com/FoxComm/libs/announcer/toml"
	"github.com/FoxComm/libs/logger"
)

type endpoint struct {
	name string
	port string
}

type endpoints struct {
	sync.RWMutex
	endpoints map[endpoint]bool
}

type Announcer interface {
	AnnounceStart(endpoint, host, port string) error
	AnnounceStop(endpoint, host, port string) error
}

var knownEndpoints endpoints
var interruptInstalled bool
var announcer Announcer

func init() {
	knownEndpoints = endpoints{endpoints: make(map[endpoint]bool)}
	// fix usage of toml by now
	announcer = toml.NewAnnouncer()
}

func (r *endpoints) Add(serviceName, port string) {
	r.Lock()
	defer r.Unlock()
	ep := endpoint{name: serviceName, port: port}
	r.endpoints[ep] = true
}

func (r *endpoints) Remove(serviceName, port string) {
	r.Lock()
	defer r.Unlock()
	ep := endpoint{name: serviceName, port: port}
	delete(r.endpoints, ep)
}

func (r *endpoints) Clear() {
	r.Lock()
	defer r.Unlock()
	keys := make([]endpoint, len(r.endpoints))
	for ep := range r.endpoints {
		keys = append(keys, ep)
		if err := announcer.AnnounceStop(ep.name, configs.Get("PRIVATE_IPV4"), ep.port); err != nil {
			logger.Error("[Announcer:Stop] error: %s", err.Error())
			return
		}
		delete(r.endpoints, ep)
	}
}

// Setup initializes the interruption handler and the announcer.
func Setup(endpoint, port string) {
	if !interruptInstalled {
		HandleInterruption()
		interruptInstalled = true
	}
	if err := announcer.AnnounceStart(endpoint, configs.Get("PRIVATE_IPV4"), port); err != nil {
		logger.Error("[Announcer:Start] error: %s", err.Error())
		return
	}
	knownEndpoints.Add(endpoint, port)
}

func Cleanup() {
	knownEndpoints.Clear()
	if r := recover(); r != nil {
		logger.Error("[Announcer] App crashed: %+v", r)
	}
}

// HandleInterruption listens to OS termination signals and logs them as well as calls AnnounceStop if a service detects that it is being terminated.
func HandleInterruption() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP)
	go func() {
		sig := <-c
		logger.Info("captured %v ...", sig)
		knownEndpoints.Clear()
		os.Exit(1)
	}()
}
