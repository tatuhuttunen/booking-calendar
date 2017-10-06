.PHONY: pb vet run dependencies update gateway clean

pb:
	for f in pb/**/*.proto; do \
		protoc -I/usr/local/include -I. \
		  -I $(GOPATH)/src \
		  -I $(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		  --go_out=plugins=grpc:. \
		  $$f; \
		echo compiled: $$f; \
	done

gateway:
	for f in pb/**/*.proto; do \
		protoc -I/usr/local/include -I. \
			-I$(GOPATH)/src \
			-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
			--grpc-gateway_out=logtostderr=true:. \
			$$f; \
		protoc -I/usr/local/include -I. \
			-I$(GOPATH)/src \
			-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
			--swagger_out=logtostderr=true:. \
			$$f; \
	done

	cat pb/**/*.swagger.json | \
	jq '.info.title = "Booking Calendar"' | \
	jq '.info.version = "1.0.0"' | \
	jq '.server.url = "http://localhost:8080"' | \
	jq --slurp 'reduce .[] as $$item ({}; . * $$item)' > cmd/api/swagger.json

vet:
	./bin/lint.sh

run:
	docker-compose build
	docker-compose up

clean:
	docker-compose down

dependencies:
	glide install

update:
	glide update
