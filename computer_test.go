package monitor_db

import "testing"
import "time"

func TestComputer(t *testing.T) {
	s := Service{Name: "test", Status: "ok", Updated: time.Now()}
	c := NewComputer("pizero", "ok", time.Now())
	if len(c.ListServices()) != 0 {
		t.Error("Failed to initialise Computer")
	}
	c.SetService(&s)
	if len(c.ListServices()) != 1 {
		t.Error("Failed to add a Computer")
	}
	s1, err := c.GetService("test")
	if err != nil {
		t.Error("Failed to retrieve a Service")
	}
	if s1.Name != "test" {
		t.Error("Failed to retrieve the correct Service")
	}

	_, err = c.GetService("foobar")
	if err == nil {
		t.Error("Should throw an error when Service not found")
	}
}

func TestJSON(t *testing.T) {
	s := Service{Name: "test", Status: "ok", Updated: time.Now()}
	c := NewComputer("pizero", "ok", time.Now())
	c.SetService(&s)
	ba := c.JSON()
	c1 := NewComputerFromJSON(ba)
	if c1.Name != c.Name {
		t.Error("Failed to create from JSON")
	}
}
