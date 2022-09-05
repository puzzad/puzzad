package puzzad

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate ./ent/schema

// segfaults, so... no. go:generate go run -mod=mod entgo.io/ent/cmd/ent describe ./ent/schema
