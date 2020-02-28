export GO111MODULE = on

build: go.sum
ifeq ($(OS),Windows_NT)
	go build -mod=readonly $(BUILD_FLAGS) -o build/gosdk.exe ./...
else
	go build -mod=readonly $(BUILD_FLAGS) -o build/gosdk.exe ./...
endif