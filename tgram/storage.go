package tgram

type Storage interface {
	Bucket(key []byte) (bucket Bucket, found bool, err error)
	GetOrCreateBucket(key []byte) (Bucket, error)
}

type Bucket interface {
	Tx() (BucketTx, error)
	Get(key []byte) (value []byte, found bool, err error)
	Set(key, value []byte) error
	ForEach(f func(key, value []byte) error) error
}

type BucketTx interface {
	Bucket
	Commit() error
	Rollback() error
}
