# Golang Diagnostic Task Runner

I've got a task from my company for making a set list of diagnostic job and this thing is created. This whole project is inspired by : https://github.com/docker/go-healthcheck . I simplify those project by removing several methods and defining each of identifier with my own glosarium :p.

## Getting started

Here is the step you need to do

### Fetch it into your libraries

You could install this by using golang command
```
go get -u github.com/dewanggasurya/godtr
```

### Example

```
package main

import (
	"github.com/dewanggasurya/godiag"
	"github.com/dewanggasurya/godiag/tasks"
	"fmt"
	"errors"
)

func main() {
	d := godiag.NewDiagnostic()

	if e := d.RegisterFunc("SuccessfulTask", func() error {
		return nil
	}); e != nil {
		panic(e)
	}
	t.Log("SuccessfulTask is registered perfectly")

	if e := d.RegisterFunc("FailedTask", func() error {
		return errors.New("Wew, something happend...")
	}); e != nil {
		panic(e)
	}
	t.Log("FailedTask is also registered perfectly")

	//=== Adding pre-defined task and check if nginx is running
	if e := d.Register("Nginx", tasks.IsProcessRunning("nginx")); e != nil {
		panic(e)
	}
	
	fmt.Println("Here is the result :", d.Run())
}
```

### Create your own pre-defined task

This is an example of expanding re-usable task function you define your own re-usable task and call it on diagnostic check This example is a task to check some process is currently running or not.

```
func TaskIsProcessRunning(process string) diagnostic.Task {
	var out, err bytes.Buffer
	var e error
	var cmd *exec.Cmd

	return diagnostic.TaskFunc(func() error {
		os := runtime.GOOS
		switch os {
		case "windows":
			cmd = exec.Command("cmd", "/C", "tasklist", "/fo", "csv", "/nh")
		case "linux":
			cmd = exec.Command("ps", "-def")
		default:
			cmd = nil
		}

		if cmd != nil {
			cmd.Stdout = &out
			cmd.Stderr = &err
			e = cmd.Run()

			if e != nil {
				return errors.New(err.String())
			}

			if strings.Contains(out.String(), process) {
				return nil
			}

			return errors.New(fmt.Sprint("No process named '", process, "' is running"))
		}

		return errors.New(fmt.Sprint("No OS '", os, "' command handler defined"))
	})
}
```

## Contributing

Feel free to contribute. You know da wey .. :P

## License

MIT License

## Author

Surya Dewangga - [@dewanggasurya](https://twitter.com/dewanggasurya)