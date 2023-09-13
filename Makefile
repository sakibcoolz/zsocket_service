include compose/.env
include variables.mk

run: 
	DATABASE=$(database) \
	HOST=$(localhost) \
	PORT=$(dbport) \
	USER=$(dbuser) \
	PASSWORD=$(dbpassword) \
	HIVEMQ=$(hivemq) \
	HUSERNAME=$(husername) \
	HPASSWORD=$(hpassword) \
	HCLIENT=$(clientid) \
	HTOPIC=$(htopic) \
	SERVICEHOST=$(localhost) \
	SERVICEPORT=$(port) \
	ENROLL=${topicenroll} \
	SECRETKEY=${secretkey} \
	$(GORUN) cmd/main.go

docker-image:
	make -C docker docker-image

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $(build) ./cmd/

docker-start:
	make -C compose start

docker-stop:
	make -C compose stop

kick-start: docker-image docker-start

clean:
	@docker rmi -f $(service)
	@docker rmi -f $(database)
	@docker image prune -f
	

.PHONY: run

