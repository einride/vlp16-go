package vlp16

// receiverOptions are the configuration options for a VLP-16 receiver.
type receiverOptions struct {
	bufferSizeBytes int
	batchSize       int
}

// defaultReceiverOptions returns receiver options with sensible default values.
func defaultReceiverOptions() *receiverOptions {
	return &receiverOptions{
		batchSize:       10,
		bufferSizeBytes: 2097152, // 2MB
	}
}

// ReceiverOption configures an LCM receiver.
type ReceiverOption func(*receiverOptions)

// WithReceiveBufferSize configures the kernel read buffer size (in bytes).
func WithBufferSize(n int) ReceiverOption {
	return func(o *receiverOptions) {
		o.bufferSizeBytes = n
	}
}

// WithReceiveBatchSize configures the max number of messages to receive from the kernel in a single batch.
func WithBatchSize(n int) ReceiverOption {
	return func(o *receiverOptions) {
		o.batchSize = n
	}
}
