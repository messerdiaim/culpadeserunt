	ctx := context.Background()
	key := datastore.NameKey("User", "alice", nil)
	user := &User{
		Name: "Alice",
		Age:  25,
	}
	if _, err := client.Put(ctx, key, user); err != nil {
		return err
	}  
