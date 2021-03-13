START=${pwd}

protoc: 
	PATH="${PATH}:${HOME}/go/bin" protoc -I protobufs \
	 --go_out=plugins=grpc:protobufs --gogrpcmock_out=:protobufs \
	 protobufs/*.proto

protoc-js:
	mkdir -p protobufs/js
	PATH="${PATH}:${HOME}/go/bin" protoc -I protobufs/ \
		--js_out=import_style=commonjs:protobufs/js \
		--grpc-web_out=import_style=commonjs,mode=grpcwebtext:protobufs/js \
		protobufs/*.proto

proxy:
	grpcwebproxy \
	--backend_addr=localhost:9090 \
	--run_tls_server=false \
	--allow_all_origins
