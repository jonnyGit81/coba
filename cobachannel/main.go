package cobachannel

import (
	"fmt"
	"github.com/jonnyGit81/coba/cobachannel/entity"
	"github.com/jonnyGit81/coba/cobachannel/senderworker"
	"github.com/jonnyGit81/coba/cobachannel/worker"
	"math/rand"
	"strconv"
	"time"
)

var request_queue = make(chan *entity.Request, 1000*1000)

func main() {
	rand.Seed(time.Now().UnixNano())

	w := worker.Worker{}
	s := senderworker.NewSenderWorker(request_queue)

	go w.WaitForRequest(request_queue)

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("len", len(request_queue))
			r := &entity.Request{ID: strconv.Itoa(i) + "_MAIN"}
			request_queue <- r
			// time.Sleep(time.Duration(rand.Intn(10 - 2 + 1) + 2) * time.Second)
		}
	}()
	go s.Send()

	time.Sleep(time.Hour * 1)
}
