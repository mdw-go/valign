package valign

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestOn(t *testing.T) {
	should.So(t,
		On("FROM",
			"SELECT 'a' FROM table;",
			"SELECT 'hello' FROM table;",
			"I'm a special snowflake",
			"SELECT 'goodbye' FROM table;",
			"SELECT 'really-super-long-name' FROM table;",
		),
		should.Equal, []string{
			"SELECT 'a'                      FROM table;",
			"SELECT 'hello'                  FROM table;",
			"I'm a special snowflake",
			"SELECT 'goodbye'                FROM table;",
			"SELECT 'really-super-long-name' FROM table;",
		},
	)
}

func TestSplitBlocks(t *testing.T) {
	should.So(t, Blocks("hi"), should.BeNil)
	should.So(t,
		Blocks("FROM",
			"SELECT 'a' FROM table;",
			"SELECT 'hello' FROM table;",
			"I'm a special snowflake",
			"SELECT 'goodbye' FROM table;",
			"SELECT 'really-super-long-name' FROM table;",
		),
		should.Equal, [][]string{
			{
				"SELECT 'a' FROM table;",
				"SELECT 'hello' FROM table;",
			},
			{
				"I'm a special snowflake",
			},
			{
				"SELECT 'goodbye' FROM table;",
				"SELECT 'really-super-long-name' FROM table;",
			},
		},
	)
}
