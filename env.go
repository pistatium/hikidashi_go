package main

type Env struct {
	Backend string `default:"dynamodb"`
	DynamodbEndpointUrl string
	RequireInitializeRepository bool `default:"false"`
}
