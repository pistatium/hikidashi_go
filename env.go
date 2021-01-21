package main

type Env struct {
	Backend string `default:"dynamodb" envconfig:"BACKEND"`
	DynamodbEndpointUrl string `envconfig:"DYNAMODB_ENDPOINT_URL"`
	DynamodbItemTableName string `default:"items" envconfig:"DYNAMODB_ITEM_TABLE_NAME"`
	RequireInitializeRepository bool `default:"false" envconfig:"REQUIRE_INITIALIZE_REPOSITORY"`
}
