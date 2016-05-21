package monitor_db

import "testing"
import "time"

func TestDatabase(t *testing.T) {
	s := Service{Name: "test", Status: "ok", Updated: time.Now()}
	c := NewComputer("pizero", "ok", time.Now())
	c.SetService(&s)
	d := NewDatabase()
	if len(d.ListComputers()) != 0 {
		t.Error("Failed to initialise Database")
	}
	d.AddComputer(c)
	if len(d.ListComputers()) != 1 {
		t.Error("Failed to add Computer")
	}
	c1, err := d.GetComputer("pizero")
	if err != nil {
		t.Error("Failed to find Computer")
	}
	if c1.Name != "pizero" {
		t.Error("Failed to find the correct Computer")
	}
	_, err = d.GetComputer("foobar")
	if err == nil {
		t.Error("Found the incorrect computer")
	}
}
