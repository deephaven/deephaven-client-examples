package main

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v8/arrow"

	dhclient "github.com/deephaven/deephaven-core/go-client/client"
)

func main() {
	ctx := context.Background()

	client, err := dhclient.NewClient(ctx, "localhost", "10000", "python")
	if err != nil {
		fmt.Println("got an error:", err)
		return
	}
	defer client.Close()

	// This is the equivalent of the sync_fields argument in python
	err = client.FetchTablesOnce(ctx)
	if err != nil {
		fmt.Println("got an error:", err)
		return
	}

	// Creation of someRecord omitted,
	// because it depends on Apache Arrow's Go API.
	var someRecord arrow.Record

	handle1, err := client.ImportTable(ctx, someRecord)
	if err != nil {
		fmt.Println("got an error:", err)
		return
	}
	defer handle1.Release(ctx)

	handle2, err := client.EmptyTable(ctx, 10)
	if err != nil {
		fmt.Println("got an error:", err)
		return
	}
	defer handle2.Release(ctx)

	handle3, err := handle1.Where(ctx, "a > 10", "b % 2 == 0")
	if err != nil {
		fmt.Println("got an error:", err)
		return
	}
	defer handle3.Release(ctx)

	handle4, err := dhclient.Merge(ctx, "a", handle1, handle3)
	if err != nil {
		fmt.Println("got an error:", err)
		return
	}
	defer handle4.Release(ctx)

	comboAgg := dhclient.NewAggBuilder().
		Sum("foo = a", "bar = c").
		Avg("avgA = a", "avgB = b")
	handle5, err := handle4.AggBy(ctx, comboAgg, "g")
	if err != nil {
		fmt.Println("got an error:", err)
		return
	}
	defer handle5.Release(ctx)

	snapshotRecord, err := handle5.Snapshot(ctx)
	if err != nil {
		fmt.Println("got an error:", err)
		return
	}
	defer snapshotRecord.Release()

	fmt.Println(snapshotRecord)
}
