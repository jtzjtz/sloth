test:
	go test ./...
	go vet ./...

#PROJECT_PATH = gitlab-ci-k8s-demo
#CI_COMMIT_SHA =dsfaslk333
#CI_COMMIT_REF_NAME =master
#VERSION =399

#	go build \
#	  -race \
#	  -ldflags "-X version.Version=$(shell cat VERSION) \
#	  -X version.Revision=${CI_COMMIT_SHA} \
#	  -X version.Branch=${CI_COMMIT_REF_NAME} \
#	  -X version.BuildUser=$(shell whoami)@$(shell hostname) \
#	  -X version.BuildDate=$(shell date +%Y%m%d %H:%M:%S) \
#	  -extldflags '-static'" \
#	  -o app

build:
	go build \
	  -race \
	  -ldflags "-X main.VERSION=${CI_COMMIT_SHA} \
	  -X 'main.BuildTime=`date  +%Y年%m月%d日%H:%M:%S`'  \
	  -X 'main.GoVersion=$(shell go version)' \
	-X main.Branch=${CI_COMMIT_REF_NAME} \
	-X main.BuildUser=$(shell whoami)@$(shell hostname) "\
	  -o output/app

#.PHONY: test build

test2:
	echo "-X ${PROJECT_PATH}/vendor/github.com/prometheus/common/version.Version=$(cat VERSION) \
         	  -ldflags -X ${PROJECT_PATH}/vendor/github.com/prometheus/common/version.Version=$(shell cat VERSION) \
         	  -X ${PROJECT_PATH}/vendor/github.com/prometheus/common/version.Revision=${CI_COMMIT_SHA} \
         	  -X ${PROJECT_PATH}/vendor/github.com/prometheus/common/version.Branch=${CI_COMMIT_REF_NAME} \
         	  -X ${PROJECT_PATH}/vendor/github.com/prometheus/common/version.BuildUser=$(shell whoami)@$(shell hostname) \
         	  -X ${PROJECT_PATH}/${CI_PROJECT_PATH}/vendor/github.com/prometheus/common/version.BuildDate=$(shell date +%Y%m%d-%H:%M:%S) \
         	  -extldflags '-static'"

test3:
	go build \
	-ldflags "-X main.VERSION=${VERSION} -X 'main.BUILD_TIME=$(shell date +%Y-%m-%d %H:%M:%S)'  -X 'main.GO_VERSION=$(shell go version)'"\
	 -o app