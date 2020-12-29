package vlp16

// listenOptions are the configuration options for a VLP-16 listener.
type listenOptions struct {
	bufferSizeBytes int
	batchSize       int
}

// defaultListenOptions returns receiver options with sensible default values.
func defaultListenOptions() *listenOptions {
	return &listenOptions{
		batchSize:       10,
		bufferSizeBytes: 2097152, // 2MB
	}
}

// ListenOption configures a VLP-16 listener.
type ListenOption func(*listenOptions)

// WithReadBufferSize configures the kernel read buffer size (in bytes).
func WithReadBufferSize(n int) ListenOption {
	return func(o *listenOptions) {
		o.bufferSizeBytes = n
	}
}

// WithReadBatchSize configures the max number of messages to read from the kernel in a single batch.
func WithReadBatchSize(n int) ListenOption {
	return func(o *listenOptions) {
		o.batchSize = n
	}
}
