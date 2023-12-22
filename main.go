package main

import (
	cmdTalker "talker/manager/api/src/cmd"
	"talker/manager/api/src/services"
	talkerutils "talker/manager/api/src/utils/io"
	ioadapters "talker/manager/api/src/utils/io/adapters"
)

func main() {
	reader := talkerutils.New(new(ioadapters.ReaderAdapter), new(ioadapters.WriterAdapter))
	crud := services.New(reader)

	cmdTalker.BuildCommands(crud)
}
