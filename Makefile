run:
	docker-machine ip dev
	docker-compose up

build-image: build-all-binaly
	docker-compose build

build-all-binaly:
	$(MAKE) echo
	$(MAKE) goji
	$(MAKE) gin


%:
	cd docker/${@}/src; \
	GOOS=linux GOARCH=amd64 go build -o ../bin/sample -ldflags="-w -s"
