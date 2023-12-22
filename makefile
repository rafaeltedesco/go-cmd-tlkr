ID := 1

build: 
	go build -o ./bin/talker
run:
	go run .
get-all:
	go run . get-all
get-by-id:
	go run . get-by-id -id=$(ID)
create:
	go run . create -json '{"name":"Rafael", "age": 32, "talk": {"watchedAt": "21/12/2023", "rate": 4}}'
delete:
	go run . delete -id=$(ID)
restore:
	cp ./backup/people.json ./src/files/people.json