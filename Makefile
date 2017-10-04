.PHONY: pb vet run dependencies update

pb:
	for f in pb/**/*.proto; do \
		protoc --go_out=plugins=grpc:. $$f; \
		echo compiled: $$f; \
	done

vet:
	./bin/lint.sh

run:
	docker-compose build
	docker-compose up

dependencies:
	glide install

update:
	glide update
