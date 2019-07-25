package monitorsuptime

import (
	"sync"
	"time"
)

type ServerStatusData struct {
	rwmu sync.RwMutex
	ServerStatus map[*Server][]*statusAtTime `json:"serverStatus"`
}

type statusAtTime struct {
	Time time.Time `json:"time"`
	Status bool `json:"online"`
}

func NewServerStatusData(servers Servers) *ServerStatusData {
  serverStatusData := &ServerStatusData{
	  ServerStatus: make(map[*Server][]*statusAtTime)
  }
  for _, server := range servers {
	  serverStatusData.ServerStatus[server] = make([]*statusAtTime, 0, 100)
  }
 return serverStatusData 
}
func (s *ServerStatusData) SetStatusAtTimeForServer(server *Server, timeNow time.Time, status bool) {
	s.rwmu.Lock()
	defer s.rwmu.Unlock()
	s.ServerStatus[server] = append(s.ServerStatus[server], &statusAtTime{Time: timeNow, Status: status})
}

func (s *ServerStatusData) GetServerStatus() map[*Server][]*statusAtTime {
	s.rwmu.RLock()
	defer s.rwmu.RUnlock()
	return s.ServerStatus
}