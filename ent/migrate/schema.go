// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AdventuresColumns holds the columns for the "adventures" table.
	AdventuresColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString, Default: "No description provided"},
		{Name: "price", Type: field.TypeFloat64, Default: 0},
		{Name: "public", Type: field.TypeBool, Default: false},
	}
	// AdventuresTable holds the schema information for the "adventures" table.
	AdventuresTable = &schema.Table{
		Name:       "adventures",
		Columns:    AdventuresColumns,
		PrimaryKey: []*schema.Column{AdventuresColumns[0]},
	}
	// GamesColumns holds the columns for the "games" table.
	GamesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"Paid", "Unpaid", "Expired"}, Default: "Unpaid"},
		{Name: "code", Type: field.TypeString, Unique: true, Size: 40},
		{Name: "game_adventure", Type: field.TypeInt},
		{Name: "game_current_puzzle", Type: field.TypeInt},
		{Name: "user_game", Type: field.TypeInt},
	}
	// GamesTable holds the schema information for the "games" table.
	GamesTable = &schema.Table{
		Name:       "games",
		Columns:    GamesColumns,
		PrimaryKey: []*schema.Column{GamesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "games_adventures_adventure",
				Columns:    []*schema.Column{GamesColumns[3]},
				RefColumns: []*schema.Column{AdventuresColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "games_puzzles_current_puzzle",
				Columns:    []*schema.Column{GamesColumns[4]},
				RefColumns: []*schema.Column{PuzzlesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "games_users_game",
				Columns:    []*schema.Column{GamesColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// GuessesColumns holds the columns for the "guesses" table.
	GuessesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "content", Type: field.TypeString},
	}
	// GuessesTable holds the schema information for the "guesses" table.
	GuessesTable = &schema.Table{
		Name:       "guesses",
		Columns:    GuessesColumns,
		PrimaryKey: []*schema.Column{GuessesColumns[0]},
	}
	// PuzzlesColumns holds the columns for the "puzzles" table.
	PuzzlesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "answer", Type: field.TypeString},
		{Name: "order", Type: field.TypeInt},
		{Name: "adventure_puzzles", Type: field.TypeInt, Nullable: true},
		{Name: "guess_puzzle", Type: field.TypeInt, Nullable: true},
	}
	// PuzzlesTable holds the schema information for the "puzzles" table.
	PuzzlesTable = &schema.Table{
		Name:       "puzzles",
		Columns:    PuzzlesColumns,
		PrimaryKey: []*schema.Column{PuzzlesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "puzzles_adventures_puzzles",
				Columns:    []*schema.Column{PuzzlesColumns[4]},
				RefColumns: []*schema.Column{AdventuresColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "puzzles_guesses_puzzle",
				Columns:    []*schema.Column{PuzzlesColumns[5]},
				RefColumns: []*schema.Column{GuessesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "verify_code", Type: field.TypeString, Unique: true},
		{Name: "verify_expiry", Type: field.TypeTime},
		{Name: "reset_code", Type: field.TypeString, Default: ""},
		{Name: "reset_expiry", Type: field.TypeTime},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "passhash", Type: field.TypeString, Default: ""},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"Unverified", "Verified", "Disabled"}, Default: "Unverified"},
		{Name: "admin", Type: field.TypeBool, Default: false},
		{Name: "guess_team", Type: field.TypeInt, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_guesses_team",
				Columns:    []*schema.Column{UsersColumns[10]},
				RefColumns: []*schema.Column{GuessesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AdventuresTable,
		GamesTable,
		GuessesTable,
		PuzzlesTable,
		UsersTable,
	}
)

func init() {
	GamesTable.ForeignKeys[0].RefTable = AdventuresTable
	GamesTable.ForeignKeys[1].RefTable = PuzzlesTable
	GamesTable.ForeignKeys[2].RefTable = UsersTable
	PuzzlesTable.ForeignKeys[0].RefTable = AdventuresTable
	PuzzlesTable.ForeignKeys[1].RefTable = GuessesTable
	UsersTable.ForeignKeys[0].RefTable = GuessesTable
}
