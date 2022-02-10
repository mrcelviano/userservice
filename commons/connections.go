package commons

import (
	"encoding/json"
	"fmt"
	gocraft "github.com/gocraft/dbr"
	"sync"
)

func InitGocraftDBRConnectionPG() *gocraft.Connection {
	connectionString := "host=" + settings.DataBase.Host + " port=" + settings.DataBase.Port + " user=" + settings.DataBase.User + " dbname=" + settings.DataBase.DBName + " sslmode=disable password=" + settings.DataBase.Password
	conn, err := gocraft.Open("postgres", connectionString, nil)
	if err != nil {
		fmt.Printf("Error dbr: %v \n", err)
	}
	conn.SetMaxOpenConns(6)
	conn.SetMaxIdleConns(2)
	return conn
}

type ServicesList struct {
	lock        sync.RWMutex
	ServicesMap map[string]string `json:"services"`
}

func (s *ServicesList) UnmarshalJSON(b []byte) error {
	s.ServicesMap = make(map[string]string)
	tmp := make(map[string]string)
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	for key, val := range tmp {
		s.ServicesMap[key] = val
	}
	return nil
}

func (s *ServicesList) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.ServicesMap)
}

func (s *ServicesList) Get(key string) (string, bool) {
	s.lock.RLock()
	val, ok := s.ServicesMap[key]
	s.lock.RUnlock()
	return val, ok
}

//GetServiceAddr получить адрес сервиса (ip+port)
func GetServiceAddr(serviceName string) string {
	name, _ := settings.Services.Get(serviceName)
	return name
}
