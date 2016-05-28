package monitor_db

import (
	"encoding/json"
	"fmt"
	"time"
)

type Service struct {
	Name    string // name of the service
	Status  string //
	Updated time.Time
}

type Computer struct {
	Name     string // name of the server
	Status   string
	IP       string
	Services []*Service
	Updated  time.Time
}

func NewComputer(name string, status string, updated time.Time) *Computer {
	return &Computer{Name: name, Status: status, Services: make([]*Service, 0), Updated: updated}
}

func NewComputerFromJSON(js []byte) *Computer {
	var c Computer
	err := json.Unmarshal(js, &c)
	if err != nil {
		fmt.Println(err)
		return nil
	} else {
		return &c
	}
}

func (s *Service) String() string {
	return fmt.Sprintf("\tName: %s\n\tStatus: %s\n\tUpdated: %v\n", s.Name, s.Status, s.Updated)
}

func (c *Computer) String() string {
	retval := fmt.Sprintf("Computer\nName: %s\nStatus: %s\nServices:\n", c.Name, c.Status)
	for _, s := range c.Services {
		retval = retval + s.String()
	}
	return retval
}

func (c *Computer) JSON() []byte {
	b, err := json.Marshal(c)
	if err != nil {
		fmt.Println(b)
	}
	return b
}

func (c *Computer) GetService(name string) (*Service, error) {
	for i := 0; i < len(c.Services); i++ {
		if c.Services[i].Name == name {
			// fmt.Println("Found a match")
			return c.Services[i], nil
		}
	}
	fmt.Println("Not found")
	return nil, fmt.Errorf("Not found")
}

func (c *Computer) SetService(svc *Service) {
	for _, psvc := range c.Services {
		// fmt.Println("Checking <<%s>>", psvc.name)
		if psvc.Name == svc.Name {
			psvc.Status = svc.Status
			psvc.Updated = svc.Updated
			// fmt.Println("Updated existing service")
			return
		}
	}
	c.Services = append(c.Services, svc)
	// fmt.Printf("Added new service. Length of array is now %d\n\r", len(c.services))
}

func (c *Computer) ListServices() []string {
	retval := make([]string, 0)
	for _, psvc := range c.Services {
		retval = append(retval, psvc.Name)
	}
	return retval
}

func (c *Computer) PrintServices() {
	fmt.Println(c.String())
}
