package worker

import (
	"fmt"
	"github.com/jonnyGit81/coba/cobachannel/entity"
	"math/rand"
	"time"
)

type Worker struct{}

func (Worker) WaitForRequest(ch <-chan *entity.Request) {
	rand.Seed(time.Now().UnixNano())
	for req := range ch {
		fmt.Println("len", len(ch))
		fmt.Println(req.ID)
		time.Sleep(time.Duration(rand.Intn(10-2+1)+2) * time.Second)
	}
}
