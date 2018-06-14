package discovery

import (
	"github.com/coreos/etcd/client"
)

type Worker struct {
	Name    string
	IP      string
	KeysAPI client.KeysAPI
}

// workerInfo is the service register information to etcd
type WorkerInfo struct {
	Name string
	IP   string
	CPU  int
}
