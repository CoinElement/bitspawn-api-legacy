package queue

type Message struct {
	Title  string
	Author string
	Body   string
}

type Option func(c *SqsClient)

func WithDelay(seconds int64) Option {
	return func(c *SqsClient) {
		c.delay = &seconds
	}
}
