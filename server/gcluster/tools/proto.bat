rem ����ͬ�����ݵ�proto
protogen.exe --proto_path protogen --tmpl_path protogen --tmpl_file proto.tmpl --output ../conf/protobuf/*.proto playerdata.proto
rem ����ͬ�����ݵ�go����
protogen.exe --proto_path protogen --tmpl_path protogen --tmpl_file go.tmpl --output ../baseapp/entity/*.go playerdata.proto
rem ����ͬ�����ݵ�protobuf����
protoc.exe --plugin=protoc-gen-go.exe --go_out=../proto ../conf/protobuf/logic.proto ../conf/protobuf/battle.proto ../conf/protobuf/cluster.proto ../conf/protobuf/playerdata.proto -I ../conf/protobuf/