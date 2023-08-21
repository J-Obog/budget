package workers

import (
	"github.com/J-Obog/paidoff/manager"
	"github.com/J-Obog/paidoff/queue"
)

type CategoryDeletionWorker struct {
	queue              queue.Queue
	categoryManager    *manager.CategoryManager
	transactionManager *manager.TransactionManager
}

func (worker *CategoryDeletionWorker) Work() error {
	message, err := worker.queue.Pop(queue.QueueName_CategoryDeleted)
	task := queue.CategoryDeletedMessage{}
	queue.FromMessage(*message, &task)

	if err != nil {
		return err
	}

	if message != nil {
		ok, err := worker.categoryManager.Exists(task.CategoryId, task.AccountId)

		if err != nil {
			return err
		}

		if !ok {
			transactions, err := worker.transactionManager.GetByCategory(task.CategoryId, task.AccountId)

			if err != nil {
				return err
			}

			for _, transaction := range transactions {
				_, err := worker.transactionManager.NullCategory(transaction.Id, task.AccountId)
				if err != nil {
					return err
				}
			}
		}

		return nil
	}

	return nil
}
