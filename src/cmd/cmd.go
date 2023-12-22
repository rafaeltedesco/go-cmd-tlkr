package cmdTalker

import (
	"flag"
	"fmt"
	"os"
	"talker/manager/api/src/models"
	"talker/manager/api/src/services"
)

func BuildCommands(crudService services.Crud) {
	args := os.Args
	if len(args) < 2 {
		panic("expect at least one subcommand")
	}
	getAll := flag.NewFlagSet("get-all", flag.ExitOnError)
	getById := flag.NewFlagSet("get-by-id", flag.ExitOnError)
	id := getById.Int("id", 1, "talker id")
	create := flag.NewFlagSet("create", flag.ExitOnError)
	jsonData := create.String("json", `{"name": "John Doe", "age": 32, "talk": {"watchedAt": "23/10/2022", "rate": 5 }}`, "Send data to create a Talker in Json format")
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteId := deleteCmd.Int("id", 1, "talker id")

	switch args[1] {
	case "get-all":
		getAll.Parse(args[2:])
		talkers := crudService.GetTalkers()
		fmt.Println(talkers)
	case "get-by-id":
		getById.Parse(args[2:])
		if talker, err := crudService.GetTalkerById(*id); err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(talker)
		}
	case "create":
		create.Parse(args[2:])
		var talker models.Talker
		crudService.CreateTalker(*jsonData, &talker)
		fmt.Printf("New talker created %v", talker)
	case "delete":
		deleteCmd.Parse(args[2:])
		if success := crudService.DeleteTalker(*deleteId); success {
			fmt.Printf("Talker with id %d delete successfully", *deleteId)
		}
	}
}
