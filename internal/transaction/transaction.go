package transaction

import (
	"sync"

	"github.com/aurelius15/lf/internal/storage"
	"github.com/aurelius15/lf/internal/verification"
)

const (
	numOfWorkers = 30 // move to config pkg and make it configurable
	pendingS     = "pending"
	processedS   = "processed"
)

type transaction struct {
	senderId            string
	receiverId          string
	amount              int
	statusOfTransaction string
}

var transactionChannel = make(chan *transaction, 1_000)

func CreateTransaction(senderId, receiverId string, amount int) {
	go func() {
		transactionChannel <- &transaction{
			senderId:            senderId,
			receiverId:          receiverId,
			amount:              amount,
			statusOfTransaction: pendingS,
		}
	}()
}

func BaseCheck(userId, senderId string, amount int) bool {
	receiver, err := storage.GetUser(userId)
	if err != nil {
		return false
	}

	sender, err := storage.GetUser(senderId)
	if err != nil {
		return false
	}

	isVerified := true

	if !receiver.VerificationStatus {
		isVerified = false
		verification.VerifyUser(receiver)
	}

	if !sender.VerificationStatus {
		isVerified = false
		verification.VerifyUser(sender)
	}

	if !isVerified {
		return false
	} else if sender.Balance < amount {
		return false
	}

	return true
}

func updateBalanceProcess(receiverId, senderId string, amount int) {
	// atomic ?
	r, _ := storage.GetUser(receiverId)
	s, _ := storage.GetUser(senderId)

	r.Balance += amount
	s.Balance -= amount

	_, _ = storage.SaveUser(r)
	_, _ = storage.SaveUser(s)
}

// TransJob (think about errors)
func TransJob() {
	var wg sync.WaitGroup

	for i := 0; i < numOfWorkers; i++ {
		wg.Add(1)
		go func() {
			for {
				if t, ok := <-transactionChannel; ok {
					if BaseCheck(t.receiverId, t.senderId, t.amount) {
						// load our user and update balance for both of them
						updateBalanceProcess(t.receiverId, t.senderId, t.amount)
					} else {

					}
					wg.Done()
				} else {
					break
				}
			}
		}()
	}

	wg.Wait()
}
