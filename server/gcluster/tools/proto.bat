rem 生成同步数据的proto
protogen.exe --proto_path protogen --tmpl_path protogen --tmpl_file proto.tmpl --output ../conf/protobuf/*.proto playerdata.proto
rem 生成同步数据的go代码
protogen.exe --proto_path protogen --tmpl_path protogen --tmpl_file go.tmpl --output ../baseapp/entity/*.go playerdata.proto
rem 生成同步数据的protobuf代码
protoc.exe --plugin=protoc-gen-go.exe --go_out=../proto ../conf/protobuf/logic.proto ../conf/protobuf/battle.proto ../conf/protobuf/cluster.proto ../conf/protobuf/playerdata.proto -I ../conf/protobuf/