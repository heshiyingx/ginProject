package utils

import (
	"fmt"
	"log"
)

// DBREventReceiver is a sentinel EventReceiver.
// Use it if the caller doesn't supply one.
type DBREventReceiver struct {
	Debug bool
}

func NewDBLogger(debug bool) *DBREventReceiver {
	return &DBREventReceiver{
		Debug: debug,
	}
}

// Event receives a simple notification when various events occur.
func (n *DBREventReceiver) Event(eventName string) {
	if !n.Debug {
		return
	}
	fmt.Printf("\x1b[31m%v\x1b[0m\n", eventName)
}

// EventKv receives a notification when various events occur along with
// optional key/value data.
func (n *DBREventReceiver) EventKv(eventName string, kvs map[string]string) {
	if !n.Debug {
		return
	}
	s1 := fmt.Sprintf("eventName=%v", eventName)
	s2 := ""
	for k, v := range kvs {
		s2 += fmt.Sprintf("\n%v: %v", k, v)
	}
	fmt.Printf("%v \x1b[32m%v\x1b[0m\n", s1, s2)
}

// EventErr receives a notification of an error if one occurs.
func (n *DBREventReceiver) EventErr(eventName string, err error) error { return err }

// EventErrKv receives a notification of an error if one occurs along with
// optional key/value data.
func (n *DBREventReceiver) EventErrKv(eventName string, err error, kvs map[string]string) error {
	s1 := fmt.Sprintf("eventName=%v error=%v", eventName, err)
	s2 := ""
	for k, v := range kvs {
		s2 += fmt.Sprintf("\n%v: %v", k, v)
	}
	log.Printf("%v \x1b[31m%v\x1b[0m", s1, s2)
	return err
}

// Timing receives the time an event took to happen.
func (n *DBREventReceiver) Timing(eventName string, nanoseconds int64) {
	if !n.Debug {
		return
	}
	fmt.Printf("eventName=%v, cost time=%v(us)\n", eventName, nanoseconds/1000)
}

// TimingKv receives the time an event took to happen along with optional key/value data.
func (n *DBREventReceiver) TimingKv(eventName string, nanoseconds int64, kvs map[string]string) {
	if n.Debug {
		s1 := fmt.Sprintf("eventName=%v, cost time=\x1b[44m%v\x1b[0m(us)", eventName, nanoseconds/1000)
		s2 := ""
		for k, v := range kvs {
			s2 += fmt.Sprintf("\n%v: %v", k, v)
		}
		fmt.Printf("%v \x1b[32m%v\x1b[0m\n", s1, s2)
	}
}
