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
type SSLWatcherParams struct {
	ID       int
	Location string
	Interval int
}
type WatcherResult struct {
	ID           int
	ResponseTime int
	Remark       string
}
type SSLWatcherResult struct {
	ID         int
	IsValid    bool
	ExpiryDate time.Time
	Remark     string
}
type Watcher struct {
	Result          chan WatcherResult
	SSLResult       chan SSLWatcherResult
	watchQuitter    map[int]chan struct{}
	sslWatchQuitter map[int]chan struct{}
}

func New() *Watcher {
	return &Watcher{
		Result:          make(chan WatcherResult),
		SSLResult:       make(chan SSLWatcherResult),
		watchQuitter:    make(map[int]chan struct{}),
		sslWatchQuitter: make(map[int]chan struct{}),
	}
}

// start a performer for a uptime watch
func (w *Watcher) Register(arg WatcherParams) {
	
	// check to prevent panic in case the interval is not set
	if arg.Interval <= 0 {
		return
	}
	// first un-registering, if already registered
	if w.IfAlreadyRegistered(arg.ID) {
		w.UnRegsiter(arg.ID)
	}

	// creating a quitter for each register request
	quitter := make(chan struct{})
	w.watchQuitter[arg.ID] = quitter

	// creating a go routine for each register request
	go Performer(arg, w.Result, quitter)
}

// start a performer for a ssl watch
func (w *Watcher) RegisterSSL(arg SSLWatcherParams) {

	// check to prevent panic in case the interval is not set
	if arg.Interval <= 0 {
		return
	}
	if w.IfSSLAlreadyRegistered(arg.ID) {
		w.UnRegsiterSSL(arg.ID)
	}

	quitter := make(chan struct{})
	w.sslWatchQuitter[arg.ID] = quitter

	go SSLPerformer(arg, w.SSLResult, quitter)
}

// removes a single the uptime watch performer
func (w *Watcher) UnRegsiter(id int) {
	if quitterChan, ok := w.watchQuitter[id]; ok {
		quitterChan <- struct{}{}
	}
}

// removes a single the ssl watch performer
func (w *Watcher) UnRegsiterSSL(id int) {
	if quitterChan, ok := w.sslWatchQuitter[id]; ok {
		quitterChan <- struct{}{}
	}
}

// removes all the uptime watch performer
func (w *Watcher) UnRegsiterAll() {
	for k := range w.watchQuitter {
		w.UnRegsiter(k)
	}
}

// removes all the ssl watch performer
func (w *Watcher) UnRegsiterAllSSL() {
	for k := range w.sslWatchQuitter {
		w.UnRegsiterSSL(k)
	}
}

// check if the uptime watch performer already running
func (w *Watcher) IfAlreadyRegistered(id int) bool {
	_, ok := w.watchQuitter[id]
	return ok
}

// check if the ssl performer already running
func (w *Watcher) IfSSLAlreadyRegistered(id int) bool {
	_, ok := w.sslWatchQuitter[id]
	return ok
}

// ticker performer for uptime watch
func Performer(arg WatcherParams, result chan<- WatcherResult, quitter <-chan struct{}) {
	ticker := time.NewTicker(time.Second * time.Duration(arg.Interval))
	for {
		select {
		case <-ticker.C:
			go GetWatcherResult(&arg, result)
		case <-quitter:
			return
		}
	}
}

// ticker performer for SSL watch
func SSLPerformer(arg SSLWatcherParams, result chan<- SSLWatcherResult, quitter <-chan struct{}) {
	ticker := time.NewTicker(time.Second * time.Duration(arg.Interval))
	for {
		select {
		case <-ticker.C:
			go GetSSLResult(&arg, result)
		case <-quitter:
			return
		}
	}
}

// function to perform uptime watch request and send result in channel
func GetWatcherResult(arg *WatcherParams, result chan<- WatcherResult) {
	genRes := WatcherResult{
		ID:           arg.ID,
		ResponseTime: -1,
		Remark:       "",
	}
	start := time.Now()
	res, err := util.Fetch("GET", arg.Location, nil, nil, arg.MaxResponseTime)
	escaped := time.Since(start).Milliseconds()
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

// function to perform ssl certificate check and send result in channel
func GetSSLResult(arg *SSLWatcherParams, result chan<- SSLWatcherResult) {
	genRes := SSLWatcherResult{
		ID:      arg.ID,
		IsValid: false,
		Remark:  "",
	}

	// TODO: Need to pass this duration from config.
	inRes := util.GetSSLCertificateDetails(arg.Location, 5000)
	genRes.ExpiryDate = inRes.Expiry.UTC()
	genRes.IsValid = inRes.IsValid
	genRes.Remark = inRes.Remark
	result <- genRes
}
