# 编译生成各个语言版本
MODE_PATH     = models
SOURCE_PATH   = `ls ${MODE_PATH}`
PROTO_PATH    = ./go/protocol

SOURCES_PROTO = $(wildcard ${MODE_PATH}/*.proto)

all:
	@(for dir in ${SOURCE_PATH};do \
		protoc --go_out=${PROTO_PATH} -I=${MODE_PATH} ${MODE_PATH}/$$dir/*.proto; \
	done)

clean:
	rm -rf ${PROTO_PATH}/* > /dev/null 2>&1