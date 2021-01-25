package api

import (
	"net/http"
	"strconv"

	"github.com/magna5/lm2cw/util"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// InfoMetric contains system information
	InfoMetric = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "api_info",
		Help: "Information about the service",
	}, []string{"service", "version"})
	opsProcessing = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "api_processing_ops_total",
		Help: "The number of events processing",
	})
	responseMetric = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "api_responses_total",
		Help: "The number of responses by endpoint and status",
	},
		[]string{"code", "method", "endpoint"})
)

// ErrorCounter exposes error count
var (
	ErrorCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "error_counter",
		Help: "Total error count.",
	},
	)
)

//DevicesSynchronizedGauge exposes devices synchronized
var (
	DevicesSynchronizedGauge = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "devices_synchronized_gauge",
			Help: "Number of devices synchronized.",
		},
		[]string{"company"},
	)
)

// LastSyncTime exposes the last time when the devices where synched
var (
	LastSyncTime = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "last_sync_time",
		Help: "Last time when devices where synched.",
	},
		[]string{"status"},
	)
)

// LastSyncDuration exposes the duration of the last sync
var (
	LastSyncDuration = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "last_sync_duration",
		Help: "Time taken when last time the devices were synched.",
	},
	)
)

// NextSyncTime exposes the next time for synching devices
var (
	NextSyncTime = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "next_sync_time",
		Help: "Next time the devices would be synched.",
	},
	)
)

//Add the devices which do not have company name set
var (
	CompanyNotSet = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "not_synced_no_company",
			Help: "Number of devices which do not have company name set.",
		},
		[]string{"device"},
	)
)

//Add the devices for which company name was not found in connectwise
var (
	CompanyNotFound = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "not_synced_company_not_found",
			Help: "Number of devices for which company was not found in connectwise.",
		},
		[]string{"company"},
	)
)

// mwMetrics is simple middleware to count ongoing requests.
func mwMetrics(next http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		defer opsProcessing.Dec()
		opsProcessing.Inc()
		lw := util.WrapWriter(w)
		next.ServeHTTP(lw, r)
		responseMetric.WithLabelValues(strconv.Itoa(lw.Status()), r.Method, r.URL.Path).Inc()
	}
	return http.HandlerFunc(f)
}
