package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-co-op/gocron/v2"
)

const httpPort = ":8000"

// https://api.open-meteo.com/v1/forecast?latitude=55.75222&longitude=37.61556&current=temperature_2m
// https://geocoding-api.open-meteo.com/v1/search?name=Moscow&count=1&language=ru&dformat=json
func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/{city}", func(w http.ResponseWriter, r *http.Request) {
		// можно так еще c 1.24
		// city := r.PathValue("city")
		// fmt.Printf("Requasted city: %s\n", city)
		city := chi.URLParam(r, "city")

		fmt.Printf("Requasted city: %s\n", city)
		_, err := w.Write([]byte("welcome"))
		if err != nil {
			log.Println(err)
		}
	})

	s, err := gocron.NewScheduler()
	if err != nil {
		panic(err)
	}

	jobs, err := initJobs(s)
	if err != nil {
		panic(err)
	}

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		fmt.Println("Start server", httpPort)
		err := http.ListenAndServe(httpPort, r) //Блокирующий
		if err != nil {
			panic(err)
		}
	}()

	go func() {
		defer wg.Done()

		fmt.Println("Start cron", jobs[0].ID())
		s.Start()
	}()

	wg.Wait()
}

// func runCon() {
// 	// each job has a unique id
// 	fmt.Println(j.ID())

// 	// start the scheduler
// 	s.Start()

// 	// block until you are ready to shut down
// 	select {
// 	case <-time.After(time.Minute):
// 	}

// 	// when you're done, shut it down
// 	err = s.Shutdown()
// 	if err != nil {
// 		// handle error
// 	}

// }

func initJobs(scheduler gocron.Scheduler) ([]gocron.Job, error) {
	j, err := scheduler.NewJob(
		gocron.DurationJob(
			10*time.Second,
		),
		gocron.NewTask(
			func() {
				fmt.Println("Hello cron")
			},
		),
	)

	if err != nil {
		return nil, err
	}

	return []gocron.Job{j}, nil
}
