protoc --go_out=plugins=grpc:. *.proto && mv github.com/thethan/fdr_proto/* . && rm -r github.com


protoc --js_out=import_style=commonjs:. --grpc-web_out=import_style=commonjs,mode=grpcwebtext:.  *.proto