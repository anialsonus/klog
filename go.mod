module gitlab.adsw.io/platform/klog

go 1.18

require go.uber.org/zap v1.21.0

require (
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
)

replace go.uber.org/zap v1.21.0 => github.com/arenadata/zap v1.21.1
