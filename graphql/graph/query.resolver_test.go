package graph

import (
	"testing"

	"github.com/99designs/gqlgen/graphql/handler"

	"github.com/99designs/gqlgen/client"
)

func TestQueryResolver_Account(t *testing.T) {
	t.Run("should query account correctly", func(t *testing.T) {

		c := client.New(handler.NewDefaultServer(NewExecutableSchema(Config{})))

		query := `query {
      accounts {
        id
        name
      }
    }`

		var response struct {
			Account struct {
				ID   string
				Name string
			}
		}

		err := c.Post(query, &response)

		// t.Logf(response.Account.ID, err)
		t.Logf(response.Account.Name, err)
	})
}
