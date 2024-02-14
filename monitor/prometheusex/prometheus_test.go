package prometheusex

import (
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"
)

func TestCounter(t *testing.T) {
	const metadata = `
		# HELP ns_subsys_test_counter A value that represents a counter
		# TYPE ns_subsys_test_counter counter
	`

	counter := prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "ns",
		Subsystem: "subsys",
		Name:      "test_counter",
		Help:      "A value that represents a counter",
	})
	counter.Inc()
	want := `
	ns_subsys_test_counter 1
	`

	err := testutil.CollectAndCompare(counter, strings.NewReader(metadata+want), "ns_subsys_test_counter")
	assert.NoError(t, err)
}

func TestGauge(t *testing.T) {
	const metadata = `
		# HELP ns_subsys_test_gauge A value that represents a gauge
		# TYPE ns_subsys_test_gauge gauge
	`

	gauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "ns",
		Subsystem: "subsys",
		Name:      "test_gauge",
		Help:      "A value that represents a gauge",
	})
	gauge.Set(3.14)
	want := `
	ns_subsys_test_gauge 3.14
	`

	err := testutil.CollectAndCompare(gauge, strings.NewReader(metadata+want), "ns_subsys_test_gauge")
	assert.NoError(t, err)
}

func TestHistogram(t *testing.T) {
	const metadata = `
		# HELP ns_subsys_test_histogram An example of a histogram
		# TYPE ns_subsys_test_histogram histogram
	`

	histogram := prometheus.NewHistogram(prometheus.HistogramOpts{
		Namespace: "ns",
		Subsystem: "subsys",
		Name:      "test_histogram",
		Help:      "An example of a histogram",
		Buckets:   []float64{1, 2, 3},
	})
	histogram.Observe(2.5)
	want := `
		ns_subsys_test_histogram{le="1"} 0
		ns_subsys_test_histogram{le="2"} 0
		ns_subsys_test_histogram{le="3"} 1
		ns_subsys_test_histogram_bucket{le="+Inf"} 1
		ns_subsys_test_histogram_sum 2.5
		ns_subsys_test_histogram_count 1
	`
	err := testutil.CollectAndCompare(histogram, strings.NewReader(metadata+want), "ns_subsys_test_histogram")
	assert.NoError(t, err)
}

func TestSummary(t *testing.T) {
	const metadata = `
		# HELP ns_subsys_test_summary An example of a summary
		# TYPE ns_subsys_test_summary summary
	`
	summary := prometheus.NewSummary(prometheus.SummaryOpts{
		Namespace: "ns",
		Subsystem: "subsys",
		Name:      "test_summary",
		Help:      "An example of a summary",
	})

	for i := 1; i <= 100; i++ {
		summary.Observe(float64(i))
	}

	want := `
		ns_subsys_test_summary_sum 5050
		ns_subsys_test_summary_count 100
	`
	err := testutil.CollectAndCompare(summary, strings.NewReader(metadata+want), "ns_subsys_test_summary")
	assert.NoError(t, err)
}

// TODO: test vector metrics
