package verification

import (
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
			for {
				if user, ok := <-verificationChannel; ok {
					user.VerificationStatus = true
					_, _ = storage.SetAsVerified(user.ID)
					wg.Done()
				} else {
					break
				}
			}
		}()
	}

	wg.Wait()
}
