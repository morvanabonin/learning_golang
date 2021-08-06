package main

import (
        "net/http"
        "fmt"
        "time"

        "github.com/prometheus/client_golang/prometheus"
        "github.com/prometheus/client_golang/prometheus/promhttp"
        "github.com/prometheus/client_golang/prometheus/promauto"
)

type (
        iMetrics interface {
                recordMetrics()

        }

        stExchange struct {
                Name string
                Qtype uint
                Qclass uint
        }

        stMetrics struct {

        }

        ExampleMetrics interface {
                iMetrics
        }

)


var (
        opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
                Name: "myapp_processed_ops_total",
                Help: "The total number of processed events",
        })
        executeExchange chan stExchange
)


func processRoutine(number int) {
        for {
                st := <-executeExchange
                fmt.Printf("%03d) %+v\n", number, st)
        }
}

func channelExecute() {
        executeExchange = make(chan stExchange)

        for i := 1; i <= 10; i++ {
                go processRoutine(i)
        }

        executeExchange<- stExchange{"Test1", 1, 1}
        executeExchange<- stExchange{"Test2", 1, 1}
        executeExchange<- stExchange{"Test3", 1, 1}
        executeExchange<- stExchange{"Test4", 1, 1}
        executeExchange<- stExchange{"Test5", 1, 1}
        executeExchange<- stExchange{"Test6", 1, 1}
        executeExchange<- stExchange{"Test7", 1, 1}
        executeExchange<- stExchange{"Test8", 1, 1}
        executeExchange<- stExchange{"Test9", 1, 1}

        time.Sleep(10 * time.Second)
}

func (m stMetrics) recordMetrics() {
        go func() {
                for {
                        opsProcessed.Inc()
                        time.Sleep(2* time.Second)
                }
        }()
}


func main() {
        ExampleMetrics.iMetrics

        //channelExecute()
        //recordMetrics()

        http.Handle("/metrics", promhttp.Handler())
        http.ListenAndServe(":2112", nil)
}

