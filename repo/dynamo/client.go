package dynamo

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type client struct {
	DB    *dynamo.DB
	Table dynamo.Table
}

func NewLocalDynamoClient() *client {
	s, _ := session.NewSession()

	_ = os.Setenv("AWS_ACCESS_KEY_ID", "dummy")
	_ = os.Setenv("AWS_SECRET_ACCESS_KEY", "dummy")

	db := dynamo.New(s, &aws.Config{
		Region:      aws.String("ap-northeast-1"),
		Endpoint:    aws.String("localhost:8123"),
		DisableSSL:  aws.Bool(true),
		Credentials: credentials.NewEnvCredentials(),
	})
	table := db.Table("ChargeWatchTable")

	return &client{Table: table, DB: db}
}

var hashKeyName = "HashKey"

type HashKey string

var rangeKeyName = "RangeKey"

type RangeKey int64

var userItemIndexName = "UserItemIndex"
var userItemIndexHashKeyName = "UserID"
var userItemIndexRangeKeyName = "UserItemType"

type UserID string
type UserItemType string
