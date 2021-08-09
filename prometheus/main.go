package main

/*
 * Prometheus examples
 *
 * References by:
 * https://github.com/prometheus/client_golang/blob/master/examples/random/main.go
 */

// A simple example exposing fictional RPC latencies with different types of
// random distributions (uniform, normal, and exponential) as Prometheus
// metrics.

// Um exemplo simples expondo uma latência fictícia de RPC com diferentes tipos de distribuições
// (regular, normal, e exponencial) como métricas do Prometheus
import (
	"flag"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	addr = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")
	uniformDomain = flag.Float64("uniform.domain", 0.0002, "The domain for the uniform distribution.")
	normalDomain = flag.Float64("normal.domain", 0.0002, "The domain for the normal distribution.")
	normalMean = flag.Float64("normal.main", 0.0002, "The mean for the normal distribution.")
	oscillationPeriod = flag.Duration("oscillation-period", 10*time.Minute, "The duration of the rate oscillation period.")
)

var (
	// Create a summary to track fictional interservice RPC latencies for three
	// distinct services with different latency distributions. These services are
	// differentiated via a "service" label.

	// Cria um sumário de uma trilha de latência fictícia de RPC para três
	// serviços distintos com diferente distribuições de latências. Estes serviços
	// são diferenciado por um camada de "serviço".
	rpcDurations = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "rpc_durations_seconds",
			Help: "RPC latency distributions.",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
		[]string{"service"},
	)

	// The same as above, but now as a histogram, and only for the normal
	// distribution. The buckets are targeted to the parameters of the
	// normal distribution, with 20 buckets centered on the mean, each
	// half-sigma wide.

	// Mesmo do exemplo de cima, mas agora como um histograma, e só para o tipo normal de distribuição.
	// Os buckets são marcados para o parâmetro normal de distribuição, com 20 buckets centralizados no mesmo tópico,
	// cada largura do meio sigma.
	rpcDurationsHistogram = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name: "rpc_durations_histogram_seconds",
			Help: "RPC latency distributions.",
			Buckets: prometheus.LinearBuckets(*normalMean-5**normalDomain, .5**normalDomain, 20),
		},
	)
)

func init() {
	// Register the summary and the histogram with Prometheus's default registry.

	// Registra o sumário e o histograma com registro default do Prometheus.
	prometheus.MustRegister(rpcDurations)
	prometheus.MustRegister(rpcDurationsHistogram)

	// Add Go module build info.
	// Obs.: the function called in the example, using the Prometheus package, was deprecated.
	prometheus.MustRegister(collectors.NewBuildInfoCollector())
}

func main() {
	flag.Parse()

	start := time.Now()

	oscillationFactor := func() float64 {
		return 2 + math.Sin(2*math.Pi*float64(time.Since(start))/float64(*oscillationPeriod))
	}

	// Periodically record some sample latencies for the three services.

	// Gravação periódica de algumas simples latências dos 3 serviços
	go func() {
		for {
			var v = rand.Float64() * *uniformDomain
			rpcDurations.WithLabelValues("uniform").Observe(v)
			time.Sleep(time.Duration(100*oscillationFactor()) * time.Millisecond)
		}
	}()

	go func() {
		for {
			v := (rand.NormFloat64() * *normalDomain) + *normalMean
			rpcDurations.WithLabelValues("normal").Observe(v)
			// Demonstrate exemplar support with a dummy ID. This would be something like a trace ID in a real
			// application. Note the necessary type assertion. We already know that rpcDurationsHistogram implements
			// the ExemplarObserver interface and thus don't need to check the outcome of the type assertion.
			rpcDurationsHistogram.(prometheus.ExemplarObserver).ObserveWithExemplar(
				v, prometheus.Labels{
					"dummyID": fmt.Sprint(rand.Intn(100000)),
				},
			)
			time.Sleep(time.Duration(50*oscillationFactor())* time.Millisecond)
		}
	}()

	go func() {
		for {
			v := rand.ExpFloat64() / 1e6
			rpcDurations.WithLabelValues("exponential").Observe(v)
			time.Sleep(time.Duration(50*oscillationFactor()) * time.Millisecond)
		}
	}()

	http.Handle("/metrics", promhttp.HandlerFor(
		prometheus.DefaultGatherer,
		promhttp.HandlerOpts{
			EnableOpenMetrics: true,
		},
	))

	log.Fatal(http.ListenAndServe(*addr, nil))
}