package main

import "chargewatch/repo/dynamo"

func main() {
	c := dynamo.NewLocalDynamoClient()
	err := c.Table.DeleteTable().Run()

	if err != nil {
		panic(err)
	}

	println("Delete Table Success.")
}
