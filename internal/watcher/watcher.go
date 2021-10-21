package watcher

import (
	"fmt"
	"time"

	"github.com/orted-org/vyoza/util"
)

type WatcherParams struct {
	ID              int
	Location        string
	Interval        int
	ExpectedStatus  int
	MaxResponseTime int
}

type WatcherResult struct {
	ID           int
	ResponseTime int
	Remark       string
}

type Watcher struct {
	Result          chan WatcherResult
	watchQuitterMap map[int]chan struct{}
}

func New() *Watcher {
	return &Watcher{
		Result:          make(chan WatcherResult),
		watchQuitterMap: make(map[int]chan struct{}),
	}
}
func (w *Watcher) Register(arg WatcherParams) {

	// first un-registering, if already registered
	if w.IfAlreadyRegistered(arg.ID) {
		w.UnRegsiter(arg.ID)
	}

	// creating a quitter for each register request
	quitter := make(chan struct{})
	w.watchQuitterMap[arg.ID] = quitter

	// creating a go routine for each register request
	go Performer(arg, w.Result, quitter)
}

func (w *Watcher) UnRegsiter(id int) {
	quitterChan := w.watchQuitterMap[id]

	// sending quit signal
	quitterChan <- struct{}{}
}

func (w *Watcher) UnRegsiterAll() {
	for k := range w.watchQuitterMap {
		w.UnRegsiter(k)
	}
}

func (w *Watcher) IfAlreadyRegistered(id int) bool {
	_, ok := w.watchQuitterMap[id]
	return ok
}

// function to perform watch at regular interval
func Performer(arg WatcherParams, res chan<- WatcherResult, quitter <-chan struct{}) {
	ticker := time.NewTicker(time.Second * time.Duration(arg.Interval))
	for {
		select {
		case <-ticker.C:
			go GetWatcherResult(&arg, res)
		case <-quitter:
			return
		}
	}
}
func GetWatcherResult(arg *WatcherParams, result chan<- WatcherResult) {
	genRes := WatcherResult{
		ID:           arg.ID,
		ResponseTime: -1,
		Remark:       "",
	}
	start := time.Now()
	res, err := util.Fetch("GET", arg.Location, nil, nil, arg.MaxResponseTime)
	escaped := time.Since(start).Microseconds()
	if err != nil {
		genRes.Remark = err.Error()
	} else {
		if res.StatusCode != arg.ExpectedStatus {
			genRes.Remark = fmt.Sprintf("expected status %d, got %d", arg.ExpectedStatus, res.StatusCode)
		} else {
			genRes.ResponseTime = int(escaped)
		}
	}
	result <- genRes
}
