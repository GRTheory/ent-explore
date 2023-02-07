package main

// the command to add a new schema
// is " go run -mod=mod entgo.io/ent/cmd/ent init xxx "

// func Do(ctx context.Context, client *ent.Client) error {
// 	head, err := client.Node.
// 		Create().
// 		SetValue(1).
// 		Save(ctx)
// 	if err != nil {
// 		return fmt.Errorf("creating the head: %w", err)
// 	}
// 	curr := head
// 	// Generate the following linked-list: 1<->2<->3<->4<->5.
// 	for i := 0; i < 4; i++ {
// 		curr, err = client.Node.
// 			Create().
// 			SetValue(curr.Value + 1).
// 			SetPrev(curr).
// 			Save(ctx)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	// loop over the list and print it. `Firstx` panics if an error occur.
// 	for curr = head; curr != nil; curr.QueryNext().FirstX(ctx) {
// 		fmt.Printf("%d ", curr.Value)
// 	}

// 	// Make the linked-list circular:
// 	// The tail of the list, has no "next",
// 	tail, err := client.Node.
// 		Query().
// 		Where(node.Not(node.HasNext())).
// 		Only(ctx)
// 	// Check that the change actually applied:
// 	prev, err := head.QueryPrev().Only(ctx)
// 	if err != nil {
// 		return fmt.Errorf("gettting head's prev: %w", err)
// 	}
// 	fmt.Printf("\n%v", prev.Value == tail.Value)

// 	return nil
// }

// func DoO2OBidirectional(ctx context.Context, client *ent.Client) error {
// 	a8m, err := client.User.
// 		Create().
// 		SetAge(30).
// 		SetName("a8m").
// 		Save(ctx)
// 	if err != nil {
// 		return fmt.Errorf("creating user: %w", err)
// 	}
// 	nati, err := client.User.
// 		Create().
// 		SetAge(28).
// 		SetName("nati").
// 		SetSpouse(a8m).
// 		Save(ctx)
// 	if err != nil {
// 		return fmt.Errorf("creating user: %w", err)
// 	}

// 	// Query the spouse edge.
// 	// Unlike `Only`, `OnlyX` panics if an error occurs.
// 	spouse := nati.QuerySpouse().OnlyX(ctx)
// 	fmt.Println(spouse.Name)

// 	// Query how many users have a spouse.
// 	// Unlike `Count`, `CountX` panics if an error occurs.
// 	count := client.User.
// 		Query().
// 		Where(user.HasSpouse()).
// 		CountX(ctx)
// 	fmt.Println(count)

// 	// Get the user, that has a spouse with name="a8m"
// 	spouse = client.User.
// 		Query().
// 		Where(user.HasSpouseWith(user.Name("a8m"))).
// 		OnlyX(ctx)
// 	fmt.Println(spouse.Name)

// 	return nil
// }

// func DoO2MSameType(ctx context.Context, client *ent.Client) error {
// 	root, err := client.Node.
// 		Create().
// 		SetValue(2).
// 		Save(ctx)
// 	if err != nil {
// 		return fmt.Errorf("creating the root: %w", err)
// 	}

// 	// Add additional nodes to the tree:
// 	//
// 	//       2
// 	//     /   \
// 	//    1     4
// 	//        /   \
// 	//       3     5
// 	//
// 	// Unlike `Save`, `SaveX` panics if an error occurs.
// 	n1 := client.Node.
// 		Create().
// 		SetValue(1).
// 		SetParent(root).
// 		SaveX(ctx)
// 	n4 := client.Node.
// 		Create().
// 		SetValue(4).
// 		SetParent(root).
// 		SaveX(ctx)
// 	n3 := client.Node.
// 		Create().
// 		SetValue(3).
// 		SetParent(n4).
// 		SaveX(ctx)
// 	n5 := client.Node.
// 		Create().
// 		SetValue(5).
// 		SetParent(n4).
// 		SaveX(ctx)

// 	fmt.Println("Tree leafs", []int{n1.Value, n3.Value, n5.Value})

// 	ints := client.Node.
// 		Query().
// 		Where(node.Not(node.HasChildren())).
// 		Order(ent.Asc(node.FieldValue)).
// 		GroupBy(node.FieldValue).
// 		IntsX(ctx)
// 	fmt.Println(ints)

// 	orphan := client.Node.
// 		Query().
// 		Where(node.Not(node.HasParent())).
// 		OnlyX(ctx)
// 	fmt.Println(orphan)

// 	return nil
// }

// func DoM2MBidirectional(ctx context.Context, client *ent.Client) error {
// 	// Unlike `Save`, `SaveX` panics if an error accurs.
// 	a8m := client.User.
// 		Create().
// 		SetAge(30).
// 		SetName("a8m").
// 		SaveX(ctx)
// 	nati := client.User.
// 		Create().
// 		SetAge(28).
// 		SetName("nati").
// 		AddFriends(a8m).
// 		SaveX(ctx)

// 	// Query friends.
// 	friends := nati.
// 		QueryFriends().
// 		AllX(ctx)
// 	fmt.Println(friends)

// 	friends = a8m.
// 		QueryFriends().
// 		AllX(ctx)
// 	fmt.Println(friends)

// 	friends = client.User.Query().
// 		Where(user.HasFriends()).
// 		AllX(ctx)
// 	fmt.Println(friends)

// 	return nil
// }
