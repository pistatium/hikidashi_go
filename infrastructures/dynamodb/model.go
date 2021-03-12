package dynamodb

import (
	"time"
)

type itemTable struct {
	Path string  `dynamo:",hash"` // Primary key
	Group string  // Path の先頭部分を DynamoDB の Hash として利用する for Secondary Index
	GroupPath string // Path から Group 部分を除いたもの for Secondary Index
	Value string
	ContentType string
	UpdatedAt time.Time
}
