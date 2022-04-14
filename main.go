package main

import (
	"github.com/samdevbr/gandalf/queue"
	"github.com/samdevbr/gandalf/spec"
	"github.com/samdevbr/gandalf/worker"
	"log"
	"runtime"
	"sync"
)

var NumWorkers = runtime.NumCPU() * 2

func main() {
	specification, err := spec.New("gandalf.toml")

	if err != nil {
		log.Fatalln(err.Error())
	}

	q := new(queue.Queue)

	for _, scope := range specification.Scopes {
		scopeFiles, err := scope.GetFiles()

		if err != nil {
			log.Println(err.Error())
			continue
		}

		for _, file := range scopeFiles {
			q.Push(file)
		}
	}

	var group sync.WaitGroup

	for i := 0; i < NumWorkers; i++ {
		w := worker.New(q)

		group.Add(1)
		go w.Start(&group)
	}

	group.Wait()
}
