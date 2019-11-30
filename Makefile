run:
	go run main.go

test:	
	docker-compose -f docker-compose.test.yml down
	docker-compose -f docker-compose.test.yml up -d
	go test -v -p 1 ./service/...
	docker-compose -f docker-compose.test.yml down

