# AT2023
Agentski/aktorski sistem za upravljenje lancem snabdevanja

Projekat iz predmeta Agentske tehonolgije, skolsa godina 2022/2023

Clanovi tima:

RA 139/2019 Aleksandar Hornjak,

RA 173/2019 Stanimirovic Dusan,

RA 232/2015 Nikola Petkovic


POKRETANJE PROJEKTA (iz at23 foldera)
STORAGE:(za sad portovi i nazivSkladista moraju biti jedinstveni)
cd storage
go run main.go nazivSkladista adresa, port, portClustera (example: go run main.go ryzen 127.0.0.1 8081 6331)

COORDINATOR:
cd coordinator
go run main.go

Generisanje iz proto fajla
cd messages/proto
protoc --go_out=../ --go_opt=paths=source_relative --proto_path=. --go-grpc_out=../ --go-grpc_opt=paths=source_relative messages.proto
