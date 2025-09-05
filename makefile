PROJ_DIR := $(CURDIR)

.PHONY: test sayhi

test:
	cd ${PROJ_DIR}/internal/aedificium && go test
	cd ${PROJ_DIR}/internal/hi && go test

sayhi:
	cd ${PROJ_DIR}/cmd/sayhi && go build . && mv sayhi ${PROJ_DIR}/
