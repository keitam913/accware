build:
	go build .
	docker build -t accware/api .

run:
	docker run --network accware --name accware/api -v /private/var/lib/accware:/var/lib/accware -v /private/etc/accware:/etc/accware accware/api

.PHONY: build run
