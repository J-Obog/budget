package workers

import (
	"log"
	"strconv"
	"sync"

	"github.com/J-Obog/paidoff/cache"
	"github.com/J-Obog/paidoff/clock"
	"github.com/J-Obog/paidoff/types"
)

type WorkerRunner struct {
	cache   cache.Cache
	clock   clock.Clock
	configs map[string]*workerConfig
}

type workerConfig struct {
	sync.Mutex
	worker    Worker
	lastRanAt *int64
	cadence   int64
	isIdle    bool
}

func lastRanAtKey(workerId string) string {
	return "workers.lastRanAt." + workerId
}

func (wr *WorkerRunner) RegisterWorker(worker Worker, workerId string, cadence int64) {
	cfg := &workerConfig{
		worker:  worker,
		cadence: cadence,
		isIdle:  true,
	}

	// TODO: handle error
	res, _ := wr.cache.Get(lastRanAtKey(workerId))

	if res != nil {
		val, err := strconv.Atoi(*res)

		if err != nil {
			log.Fatal("value must be an int")
		}

		*cfg.lastRanAt = int64(val)
	}

	wr.configs[workerId] = cfg
}

func (wr *WorkerRunner) runWorker(workerId string, cfg *workerConfig) {
	err := cfg.worker.Work()
	cfg.Lock()
	timestamp := wr.clock.Now()

	if err != nil {
		if err := wr.cache.Set(lastRanAtKey(workerId), strconv.Itoa(int(timestamp))); err != nil {
			log.Fatal(err)
		}
		cfg.lastRanAt = types.Int64Ptr(timestamp)
	}

	cfg.isIdle = true
	cfg.Unlock()
}

func (wr *WorkerRunner) shouldRun(lastRanAt *int64, cadence int64) bool {
	timestamp := wr.clock.Now()
	return (lastRanAt == nil) || ((timestamp - *lastRanAt) >= cadence)
}

func (wr *WorkerRunner) Start() {
	for {
		for workerId, cfg := range wr.configs {
			cfg.Lock()

			if cfg.isIdle {
				if wr.shouldRun(cfg.lastRanAt, cfg.cadence) {
					cfg.isIdle = false
					go wr.runWorker(workerId, cfg)
				}
			}

			cfg.Unlock()
		}
	}
}

func (wr *WorkerRunner) Stop() {

}
