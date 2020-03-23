protoc --go_out=plugins=grpc:. *.proto && mv github.com/thethan/fdr_proto/* . && rm -r github.com


protoc  --js_out=import_style=commonjs:../MDB-FDR/fdr_proto_js --grpc-web_out=import_style=commonjs,mode=grpcwebtext:../MDB-FDR/fdr_proto_js  *.proto