package transaction

import (
	"log"
	"sync"

	"github.com/aurelius15/lf/internal/utils"

	"github.com/aurelius15/lf/internal/storage"
	"github.com/aurelius15/lf/internal/verification"
)

const (
	numOfWorkers = 30 // move to config pkg and make it configurable
	pendingS     = "pending"
	processedS   = "processed"
)

type transaction struct {
	id                  string
	receiverId          string
	senderId            string
	amount              int
	statusOfTransaction string
}

var transactionChannel = make(chan *transaction, 1_000)

func CreateTransaction(receiverId, senderId string, amount int) {
	go func() {
		transactionChannel <- &transaction{
			id:                  utils.GenerateUUID(),
			receiverId:          receiverId,
			senderId:            senderId,
			amount:              amount,
			statusOfTransaction: pendingS,
		}
	}()
}

func BaseCheck(receiverId, senderId string, amount int) bool {
	receiver, err := storage.Instance().GetUser(receiverId)
	if err != nil {
		return false
	}

	sender, err := storage.Instance().GetUser(senderId)
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

	if sender.Balance < amount {
		isVerified = false
	}

	if !isVerified {
		return false
	}

	return true
}

// TransJob (think about errors)
func TransJob() {
	var wg sync.WaitGroup

	for i := 0; i < numOfWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
		Loop:
			for {
				select {
				case t, ok := <-transactionChannel:
					if ok {
						if BaseCheck(t.receiverId, t.senderId, t.amount) {
							err := storage.Instance().TransferMoney(t.receiverId, t.senderId, t.amount)
							log.Println(err)
							log.Printf("Transaction %s(%s->%s %d): processed \n", t.id, t.senderId, t.receiverId, t.amount)
						} else {
							log.Printf("Transaction %s(%s->%s %d): did not processed \n", t.id, t.senderId, t.receiverId, t.amount)
						}
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
