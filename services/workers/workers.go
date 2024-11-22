package workers

import (
	"fmt"

	"github.com/kaellybot/kaelly-notifier/models/constants"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func New() *Impl {
	service := &Impl{
		taskChan: make(chan Task, viper.GetInt(constants.TaskBuffer)),
	}

	workerPoolCount := viper.GetInt(constants.WorkerPool)
	service.wg.Add(workerPoolCount)
	for range workerPoolCount {
		go service.worker()
	}
	return service
}

// Work submits a task to the service, blocking if the buffer is full.
func (service *Impl) Work(task Task) {
	service.taskChan <- task
}

func (service *Impl) Shutdown() {
	close(service.taskChan)
	service.wg.Wait()
}

func (service *Impl) worker() {
	defer service.wg.Done()

	for task := range service.taskChan {
		// Panic recovery wrapper to make it resilient.
		func() {
			defer func() {
				if r := recover(); r != nil {
					log.Error().
						Str(constants.LogPanic, fmt.Sprintf("%v", r)).
						Msg("Panicking while performing task, ignoring it")
				}
			}()
			task()
		}()
	}
}
