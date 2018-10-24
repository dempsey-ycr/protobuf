 protoc --go_out=./go/protocol -I=./models  ./models/asset/*.proto
 protoc --go_out=./go/protocol -I=./models ./models/action/*.proto
 protoc --go_out=./go/protocol -I=./models  ./models/basic/*.proto