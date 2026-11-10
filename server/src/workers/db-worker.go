package workers

import (
	"airun/brain-tracker/server/src/models"
	"airun/brain-tracker/server/src/repository"
	"context"
	"log"
	"time"
)

var dbWriteChannelCap int = 4096
var dbWriteBatchCap int = 256
var DbWriteChannel = make(chan models.SignalData, dbWriteChannelCap)

func InitiateDBWorker(mainContext context.Context) {
	go func() {
		defer func() {
			workerErr := recover()
			if workerErr != nil {
				log.Println("DB Worker Error: ", workerErr)
			}
		}()
		var ticker = time.NewTicker(2000 * time.Millisecond)
		var signalsBatch []models.SignalData = []models.SignalData{}

		flushSingals := func() {
			if len(signalsBatch) == 0 {
				return
			}
			err := repository.CreateBulkSignalData(&signalsBatch)
			if err != nil {
				log.Printf("dropping batch of %d after flush error: %v", len(signalsBatch), err)
			}
			signalsBatch = []models.SignalData{}
		}

	mainLoop:
		for {
			select {
			case signal := <-DbWriteChannel:
				log.Printf("Singal %v recieved", len(signalsBatch))
				if len(signalsBatch) >= dbWriteBatchCap {
					log.Printf("Batch Filled, Flushing")
					flushSingals()
				}
				signalsBatch = append(signalsBatch, signal)
			case <-ticker.C:
				flushSingals()
			case <-mainContext.Done():
				break mainLoop
			}
		}

	}()
}
