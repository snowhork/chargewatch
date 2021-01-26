package main

import "chargewatch/repo/dynamo"

type ChargeWatchTable struct {
	HashKey      string `dynamo:",hash"`
	RangeKey     int    `dynamo:",range"`
	UserID       string `index:"UserItemIndex,hash"`
	UserItemType string `index:"UserItemIndex,range"`
}

func main() {
	c := dynamo.NewLocalDynamoClient()
	err := c.DB.CreateTable("ChargeWatchTable", ChargeWatchTable{}).
		Provision(4, 2).
		ProvisionIndex("UserItemIndex", 1, 2).
		Run()

	if err != nil {
		panic(err)
	}

	println("Create Table Success.")
}
