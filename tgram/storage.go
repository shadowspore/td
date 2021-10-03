package tgram

type Storage interface {
	Bucket(key string) (bucket Bucket, found bool, err error)
	GetOrCreateBucket(key string) (Bucket, error)
}

type Bucket interface {
	Tx() (BucketTx, error)
	Get(key string) (value []byte, found bool, err error)
	Set(key string, value []byte) error
	ForEach(f func(key string, value []byte) error) error
}

type BucketTx interface {
	Bucket
	Commit() error
	Rollback() error
}
