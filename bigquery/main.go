package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/option"
)
type vendor struct {
		BusinessEntityID int
		AccountNumber  string
		Name  string
		CreditRating  int
		PreferredVendorStatus  int
		ActiveFlag  int
		PurchasingWebServiceURL  string
		ModifiedDate  string
}

func main() {
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, "walmart-omni-sre", option.WithCredentialsFile("walmart-omni-sre-2b17b7f14f86.json"))
	if err != nil {
		fmt.Println("client ", err);
		// TODO: Handle error.
	}
	ins := client.Dataset("purchasing").Table("vendor").Inserter()

	type score struct {
		Name string
		Num  int
	}
	vendors := []vendor {
	{BusinessEntityID:10000, AccountNumber:"10000", Name: "vendor_10000", CreditRating:1, PreferredVendorStatus:1, ActiveFlag:1, PurchasingWebServiceURL: "www.bata.com", ModifiedDate:"04/15/2024"},
	{BusinessEntityID:10020, AccountNumber:"10020", Name: "vendor_10020", CreditRating:1, PreferredVendorStatus:1, ActiveFlag:1, PurchasingWebServiceURL: "www.adidas.com", ModifiedDate:"04/15/2024"},
	}
	
	// Schema is inferred from the score type.
	if err := ins.Put(ctx, vendors); err != nil {
		fmt.Println("Put ", err);

		// TODO: Handle error.
	}
}