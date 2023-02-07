package main

import (
	"context"
	"fmt"

	"github.com/GRTheory/ent-explore/edges/ent"
	"github.com/GRTheory/ent-explore/edges/ent/node"
)

// the command to add a new schema 
// is " go run -mod=mod entgo.io/ent/cmd/ent init xxx "

func Do(ctx context.Context, client *ent.Client) error {
	head, err := client.Node.
		Create().
		SetValue(1).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating the head: %w", err)
	}
	curr := head
	// Generate the following linked-list: 1<->2<->3<->4<->5.
	for i := 0; i < 4; i++ {
		curr, err = client.Node.
			Create().
			SetValue(curr.Value + 1).
			SetPrev(curr).
			Save(ctx)
		if err != nil {
			return err
		}
	}

	// loop over the list and print it. `Firstx` panics if an error occur.
	for curr = head; curr != nil; curr.QueryNext().FirstX(ctx) {
		fmt.Printf("%d ", curr.Value)
	}

	// Make the linked-list circular:
	// The tail of the list, has no "next",
	tail, err := client.Node.
		Query().
		Where(node.Not(node.HasNext())).
		Only(ctx)
	// Check that the change actually applied:
	prev, err := head.QueryPrev().Only(ctx)
	if err != nil {
		return fmt.Errorf("gettting head's prev: %w", err)
	}
	fmt.Printf("\n%v", prev.Value == tail.Value)
	
	return nil
}
