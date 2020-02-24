build:
	npm run build
	docker build -t accware-ui .

run:
	docker run --network accware -p 443:443 -v /private/etc/pki/nginx:/etc/pki/nginx --name accware-ui accware-ui

.PHONY: build run
