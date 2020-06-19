package descriptors

import (
	"myproject/pkg/handlers"
	//v1 "myproject/pkg/apis/v1"
	def "github.com/caicloud/nirvana/definition"
)

func init() {
	register([]def.Descriptor{{
		Path:        "/allUsers",
		Definitions: []def.Definition{listUsers},
	}, {
		Path:        "/oneUser",
		Definitions: []def.Definition{getUser},
	},
		{
			Path:        "/addUser",
			Definitions: []def.Definition{addUser},
		},
		{
			Path:        "/deleteUser",
			Definitions: []def.Definition{deleteUser},
		},
		{
			Path:        "/changeInfo",
			Definitions: []def.Definition{changeInfo},
		},
	}...)

}

var addUser = def.Definition{
	Method:      def.Create,
	Summary:     "Add one user",
	Description: "Add a user with parameters 'identifier','name' and 'age'",
	Function:    handlers.AddUser,
	Parameters: []def.Parameter{
		{
			Source:      def.Query,
			Name:        "Identifier",
			Default:     "Unknown",
			Description: "Identifier of the user",
		},
		{
			Source:      def.Query,
			Name:        "Name",
			Default:     "Anonymity",
			Description: "Name of the user",
		},
		{
			Source:      def.Query,
			Name:        "Age",
			Default:     "20",
			Description: "Age of the user",
		},
	},
	Results: def.DataErrorResults("Add a user"),
}

var listUsers = def.Definition{
	Method:      def.List,
	Summary:     "List all the users",
	Description: "Query a specified number of messages and returns an array",
	Function:    handlers.AllUsers,
	Results:     def.DataErrorResults("A list of users"),
}

var getUser = def.Definition{
	Method:      def.Get,
	Summary:     "Get user",
	Description: "Get a user by id",
	Function:    handlers.OneUser,
	Parameters: []def.Parameter{
		{
			Source:      def.Query,
			Name:        "Identifier",
			Default:     "Undefined",
			Description: "The id to find the specific user",
		},
	},
	Results: def.DataErrorResults("A specific user"),
}

var deleteUser = def.Definition{
	Method:      def.Delete,
	Summary:     "Delete user",
	Description: "Delete a user by id",
	Function:    handlers.DeleteUser,
	Parameters: []def.Parameter{
		{
			Source:      def.Query,
			Name:        "Identifier",
			Default:     "Undefined",
			Description: "The id to find the specific user",
		},
	},
	Results: def.DataErrorResults("Delete a specific user"),
}

var changeInfo = def.Definition{
	Method:      def.Patch,
	Summary:     "Change information ",
	Description: "Change information of a user by id",
	Function:    handlers.ChangeInfo,
	Parameters: []def.Parameter{
		{
			Source:      def.Query,
			Name:        "Identifier",
			Default:     "Undefined",
			Description: "The id to find the specific user",
		},
		{
			Source:      def.Query,
			Name:        "Name",
			Description: "Name of the user",
		},
		{
			Source: def.Query,
			Name:   "Age",
			//Default:      -100,
			Description: "Age of the user",
		},
	},
	Results: def.DataErrorResults("Delete a specific user"),
}
