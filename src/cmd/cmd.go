package cmdTalker

import (
	"flag"
	"os"
	"talker/manager/api/src/presenter"
	"talker/manager/api/src/services"
)

var (
	getAll, getById, create, deleteCmd *flag.FlagSet
	id, deleteId                       *int
	jsonData                           *string
)

func checkCommands(args []string) {
	if len(args) < 2 {
		panic("expect at least one subcommand")
	}
}

func createFlags() {
	getAll = flag.NewFlagSet("get-all", flag.ExitOnError)
	getById = flag.NewFlagSet("get-by-id", flag.ExitOnError)
	id = getById.Int("id", 1, "talker id")
	create = flag.NewFlagSet("create", flag.ExitOnError)
	jsonData = create.String("json", `{"name": "John Doe", "age": 32, "talk": {"watchedAt": "23/10/2022", "rate": 5 }}`, "Send data to create a Talker in Json format")
	deleteCmd = flag.NewFlagSet("delete", flag.ExitOnError)
	deleteId = deleteCmd.Int("id", 1, "talker id")
}

func BuildCommands(crudService services.Crud) {
	args := os.Args
	checkCommands(args)
	createFlags()

	presenter := presenter.New(crudService)

	switch args[1] {
	case "get-all":
		getAll.Parse(args[2:])
		presenter.DisplayTalkers()
	case "get-by-id":
		getById.Parse(args[2:])
		presenter.DisplayTalkerById(*id)
	case "create":
		create.Parse(args[2:])
		presenter.CreateAndDisplay(*jsonData)
	case "delete":
		deleteCmd.Parse(args[2:])
		presenter.DeleteAndDisplay(*deleteId)
	}
}
