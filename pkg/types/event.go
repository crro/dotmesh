package types

import (
	"encoding/gob"
	"fmt"
	"strings"

	"github.com/dotmesh-io/dotmesh/pkg/container"
)

func init() {
	gob.Register(&Metadata{})
	gob.Register(&container.DockerContainer{})
	gob.Register(&Event{})
}

type EventType int

const (
	EventTypeRequest EventType = iota
	EventTypeResponse
)

type Event struct {
	ID           string
	Name         string
	FilesystemID string
	Type         EventType
	Args         *EventArgs
}

func (e Event) String() string {
	return fmt.Sprintf("<Event %s: %s>", e.Name, e.Args)
}

func (e Event) Error() error {
	if e.Args == nil {
		return nil
	}

	errIntf, ok := (*e.Args)["err"]
	if !ok {
		return nil
	}
	err, ok := errIntf.(error)
	if ok {
		return err
	}
	return nil
}

func NewEvent(name string) *Event {
	return &Event{Name: name, Args: &EventArgs{}}
}

func NewErrorEvent(name string, err error) *Event {
	if err == nil {
		return NewEvent(name)
	}
	return &Event{
		Name: name,
		Args: &EventArgs{"err": err},
	}
}

type EventArgs map[string]interface{}

func (ea EventArgs) String() string {
	aggr := []string{}
	for k, v := range ea {
		aggr = append(aggr, fmt.Sprintf("%s: %+q", k, v))
	}
	return strings.Join(aggr, ", ")
}