.PHONY: protos gateway lint vendor-update vendor-install run clean

protos:
	for f in pb/**/*.proto; do \
		docker run --rm -v $$PWD:$$PWD -w $$PWD znly/protoc --go_out=plugins=grpc:. -I. $$f; \
		echo compiled: $$f; \
	done

gateway:
	for f in pb/**/*.proto; do \
		docker run --rm -v $$PWD:$$PWD -w $$PWD znly/protoc --grpc-gateway_out=logtostderr=true:. -I. $$f; \
		echo compiled gateway: $$f; \
		docker run --rm -v $$PWD:$$PWD -w $$PWD znly/protoc --swagger_out=logtostderr=true:. -I. $$f; \
		echo compiled swagger: $$f; \
	done

	docker run --rm -v $$PWD:/go/src/path -w /go/src/path tatuhuttunen/jq-container sh bin/swagger_processing.sh

lint:
	docker run --rm -v $$PWD:/go/src/path -w /go/src/path moogar0880/gometalinter:latest --config=gometalinter.json ./...

run:
	docker-compose build
	docker-compose up

clean:
	docker-compose down

vendor-install:
	docker run --rm -it -v $$PWD:/go/src/github.com/tatuhuttunen/booking-calendar \
	-w /go/src/github.com/tatuhuttunen/booking-calendar instrumentisto/glide install

vendor-update:
	docker run --rm -it -v $$PWD:/go/src/github.com/tatuhuttunen/booking-calendar \
	-w /go/src/github.com/tatuhuttunen/booking-calendar instrumentisto/glide update