package senderworker

import (
	"github.com/jonnyGit81/coba/cobachannel/entity"
	"math/rand"
	"strconv"
	"time"
)

type SenderWorker struct {
	request_queue chan *entity.Request
}

func NewSenderWorker(request_queue chan *entity.Request) *SenderWorker {
	return &SenderWorker{request_queue: request_queue}
}

func (s *SenderWorker) Send() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		// fmt.Println("len", len(s.request_queue))
		r := &entity.Request{ID: strconv.Itoa(i) + "_FROM_SENDER_WORKER"}
		s.request_queue <- r
		//time.Sleep(time.Duration(rand.Intn(10 - 2 + 1) + 2) * time.Second)
	}
}
