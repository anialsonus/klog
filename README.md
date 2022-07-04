# Klog

KLog is a thin wrapper around go.uber.org/zap logger.
It provide structured json logger with two additional fields - source and date.

## Quick Start

```go
klog.InitLogger("New Source")
defer klog.Sync()
klog.Info("First message")
```