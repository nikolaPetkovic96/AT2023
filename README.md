# AT2023
Agentski/aktorski sistem za upravljenje lancem snabdevanja<br>
Projekat iz predmeta Agentske tehonolgije, skolsa godina 2022/2023<br>
Clanovi tima:<br>
RA 139/2019 Aleksandar Hornjak,<br>
RA 173/2019 Stanimirovic Dusan,<br>
RA 232/2015 Nikola Petkovic<br>
<br>
====================================<br>
POKRETANJE PROJEKTA (iz at23 foldera)<br>
STORAGE:(za sad portovi i nazivSkladista moraju biti jedinstveni)<br>
cd storage<br>
go run main.go nazivSkladista adresa, port, portClustera (example: go run main.go ryzen 127.0.0.1 8081 6331)<br>
<br>
COORDINATOR:<br>
cd coordinator<br>
go run main.go<br>
<br>
Generisanje iz proto fajla<br>
cd messages/proto<br>
protoc --go_out=../ --go_opt=paths=source_relative --proto_path=. --go-grpc_out=../ --go-grpc_opt=paths=source_relative messages.proto<br>
