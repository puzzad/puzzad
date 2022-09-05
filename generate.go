package puzzad

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --target ./ent ./schema

//go:generate go run -mod=mod entgo.io/ent/cmd/ent describe ./schema
