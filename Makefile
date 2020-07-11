RELEASE_DIR = out/gin-app-start
PROJECT_NAME = gin-app-start
RELEASE_COPY = app config
VERSION = 1.1.0

LDFLAGS = -ldflags "-X 'main.Version=${VERSION}'"

run:
	@SERVER_ENV=local go run main.go

build:	
	@go build ${LDFLAGS} -a -v -o gin-app-start .

docker:
	@docker build -t gin-app-start:latest .

tar:
	@echo 'Clean out...'
	@rm -rf ./out
	@echo 'Copy files'
	@mkdir -p ${RELEASE_DIR}
	@if [ `echo $$OSTYPE | grep -c 'darwin'` -eq 1 ]; then \
		cp -r ${RELEASE_COPY} ${RELEASE_DIR}; \
	else \
		cp -rL ${RELEASE_COPY} ${RELEASE_DIR}; \
	fi

	@cp main.go ${RELEASE_DIR}
	@cp Dockerfile ${RELEASE_DIR}
	@cp Makefile ${RELEASE_DIR}
	@cp go.mod ${RELEASE_DIR}
	@cp go.sum ${RELEASE_DIR}
	@cp config.yaml ${RELEASE_DIR}
	@echo "all codes are in ${RELEASE_DIR}"
	@cd out && tar zcf ${PROJECT_NAME}.tgz ${PROJECT_NAME}

clean:
	@echo 'Clean out...'
	@rm -rf ./out