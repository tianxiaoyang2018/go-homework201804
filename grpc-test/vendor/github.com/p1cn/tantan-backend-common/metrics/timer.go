package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// NewTimer ...
// new一个timer且注册到普罗米修斯，请注意，metricName在一个进程内不可以重复，否则panic
// namespace 统一用"tantan"(name.go中定义了常量) 或者其他产品
// metricName是指标名字，确保一个进程内唯一性
// help是描述指标用途
// labels 是维度
func NewTimer(namespace, metricName, help string, labels []string) *Timer {
	summary := prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace:  namespace,
			Name:       metricName + "_s",
			Help:       help + " (summary)",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
		labels)

	prometheus.MustRegister(summary)

	histogram := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Name:      metricName + "_h",
			Help:      help + " (histogram)",
		},
		labels)

	prometheus.MustRegister(histogram)
	return &Timer{
		name:      metricName,
		summary:   summary,
		histogram: histogram,
	}
}

type Timer struct {
	name      string
	summary   *prometheus.SummaryVec
	histogram *prometheus.HistogramVec
}

// Timer 返回一个函数，并且开始计时，结束计时则调用返回的函数
// 请参考timer_test.go 的demo
func (t *Timer) Timer() func(values ...string) {
	if t == nil {
		return func(values ...string) {}
	}

	now := time.Now()

	return func(values ...string) {
		seconds := float64(time.Since(now)) / float64(time.Second)
		t.summary.WithLabelValues(values...).Observe(seconds)
		t.histogram.WithLabelValues(values...).Observe(seconds)
	}
}

// Observe ：传入duration和labels，
func (t *Timer) Observe(duration time.Duration, label ...string) {
	if t == nil {
		return
	}

	seconds := float64(duration) / float64(time.Second)
	t.summary.WithLabelValues(label...).Observe(seconds)
	t.histogram.WithLabelValues(label...).Observe(seconds)
}
