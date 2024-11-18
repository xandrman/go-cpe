package routeros

import (
	"log"
	"sync"
	"time"

	"github.com/go-routeros/routeros"
)

type Connection struct {
	Client  *routeros.Client
	Address string
	Timer   *time.Timer
}

type ConnectionManager struct {
	mu          sync.Mutex
	connections map[string]*Connection
	username    string
	password    string
}

func NewConnectionManager(username, password string) *ConnectionManager {
	return &ConnectionManager{
		connections: make(map[string]*Connection),
		username:    username,
		password:    password,
	}
}

func (cm *ConnectionManager) GetConnection(address string) (*routeros.Client, error) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	if conn, exists := cm.connections[address]; exists {
		conn.Timer.Reset(time.Minute)
		return conn.Client, nil
	}

	client, err := routeros.Dial(address, cm.username, cm.password)
	if err != nil {
		return nil, err
	}

	timer := time.AfterFunc(time.Minute, func() {
		cm.CloseConnection(address)
	})
	cm.connections[address] = &Connection{
		Client:  client,
		Address: address,
		Timer:   timer,
	}

	return client, nil
}

func (cm *ConnectionManager) CloseConnection(address string) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	if conn, exists := cm.connections[address]; exists {
		conn.Client.Close()
		conn.Timer.Stop()
		delete(cm.connections, address)
		log.Printf("Connection to %s closed due to inactivity", address)
	}
}
