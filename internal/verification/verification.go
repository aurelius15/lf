package verification

import (
	"log"
	"sync"

	"github.com/aurelius15/lf/internal/model"
	"github.com/aurelius15/lf/internal/storage"
)

const (
	numOfWorkers = 30 // move to config pkg and make it configurable
)

var verificationChannel = make(chan *model.User, 1_000)

// VerifyUser put user in channel for feature verification
func VerifyUser(user *model.User) {
	go func() {
		verificationChannel <- user
	}()
}

// VerifyJob (think about errors)
func VerifyJob() {
	var wg sync.WaitGroup

	for i := 0; i < numOfWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
		Loop:
			for {
				select {
				case user, ok := <-verificationChannel:
					if ok {
						_ = storage.Instance().SetAsVerified(user.ID)
						log.Printf("User %s(%s): verified \n", user.ID, user.Name)
					} else {
						break Loop
					}
				default:
					break Loop
				}
			}
		}()
	}

	wg.Wait()
}
