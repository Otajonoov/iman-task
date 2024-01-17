package usecase

import (
	"iman-task/internal/collector/domain"
	"strconv"
	"sync"
)

type collectorUsecase struct {
	collector domain.Collector
}

type CollectorUsecase interface {
	CollectPosts() (*domain.Posts, error)
}

func NewCollectorUsecase(collector domain.Collector) CollectorUsecase {
	return &collectorUsecase{
		collector: collector,
	}
}

func (c *collectorUsecase) CollectPosts() (*domain.Posts, error) {

	// Create a wait group to wait for all goroutines to finish
	var wg sync.WaitGroup
	// Create a channel to communicate errors from goroutines
	errorChannel := make(chan error)
	// Create a channel to close when all workers are done
	doneChannel := make(chan bool)

	posts := make(chan []domain.Post)

	// Slice to store posts
	postsSlice := domain.Posts{}
	// Start a goroutine to handle received posts and close the channel
	go func() {
		defer close(doneChannel)
		for {
			select {
			case posts, ok := <-posts:
				if !ok {
					// posts channel closed, all workers are done
					return
				}
				postsSlice.Posts = append(postsSlice.Posts, posts...)
			}
		}
	}()

	// Send page numbers to the workers
	for i := 1; i <= 50; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			// Fetch data for a specific page
			post, err := c.collector.GetPage(strconv.Itoa(i))
			if err != nil {
				errorChannel <- err
				return
			}
			// Send fetched post to the posts channel
			posts <- post
		}(i)
	}

	// Start a goroutine to close the done channel when all workers are done
	go func() {
		wg.Wait()
		close(posts)
	}()

	// Check for any errors reported by workers
	select {
	case err := <-errorChannel:
		return nil, err
	case <-doneChannel:
		return &postsSlice, nil
	}
}
