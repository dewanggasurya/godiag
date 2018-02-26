package godiag

import (
	"errors"
	"time"
)

type Task interface {
	Check() error
}

type TaskFunc func() error

func (tf TaskFunc) Check() error {
	return tf()
}

type Diagnostic struct {
	TaskList map[string]Task
}

type ResultStatus string

const (
	StatusOK  ResultStatus = "OK"
	StatusNOK ResultStatus = "NOK"
)

type Result struct {
	Name           string
	Message        string
	DurationString string
	Duration       time.Duration
	Timestamp      time.Time
	Status         ResultStatus
}

var _diagnostic *Diagnostic

func NewDiagnostic() *Diagnostic {
	diagnostic := new(Diagnostic)
	diagnostic.TaskList = make(map[string]Task)

	return diagnostic
}

func init() {
	_diagnostic = NewDiagnostic()
}

func Register(name string, task Task) error {
	return _diagnostic.Register(name, task)
}

func (diagnostic *Diagnostic) Register(name string, task Task) error {
	if diagnostic == nil {
		diagnostic = _diagnostic
	}

	if _, ok := diagnostic.TaskList[name]; ok {
		return errors.New("Task name '" + name + "' is already exists")
	}

	diagnostic.TaskList[name] = task

	return nil
}

func RegisterFunc(name string, task func() error) error {
	return _diagnostic.Register(name, TaskFunc(task))
}

func (diagnostic *Diagnostic) RegisterFunc(name string, task func() error) error {
	return diagnostic.Register(name, TaskFunc(task))
}

func Run() []*Result {
	return _diagnostic.Run()
}

func (diagnostic *Diagnostic) Run() []*Result {
	if diagnostic == nil {
		diagnostic = _diagnostic
	}

	results := []*Result{}
	if len(diagnostic.TaskList) > 0 {
		for name, task := range diagnostic.TaskList {
			result := diagnostic.NewResult(name)
			e := task.Check()

			if e != nil {
				result.SetError(e)
			} else {
				result.SetSuccessMessage("")
			}

			results = append(results, result)
		}
	}

	return results
}

func (diagnostic *Diagnostic) NewResult(name string) *Result {
	result := new(Result)
	result.Name = name
	result.Status = StatusOK
	result.Timestamp = time.Now()
	return result
}

func (result *Result) setDuration() *Result {
	result.Duration = time.Since(result.Timestamp)
	result.DurationString = result.Duration.String()
	return result
}

func (result *Result) SetError(e error) *Result {
	result.Status = StatusNOK
	result.Message = e.Error()
	result.setDuration()
	return result
}

func (result *Result) SetErrorString(e string) *Result {
	result.Status = StatusNOK
	result.Message = e
	result.setDuration()
	return result
}

func (result *Result) SetSuccessMessage(message string) *Result {
	result.Status = StatusOK
	result.Message = message
	result.setDuration()
	return result
}

func (result *Result) SetMessage(status ResultStatus, message string) *Result {
	result.Status = status
	result.Message = message
	result.setDuration()
	return result
}
