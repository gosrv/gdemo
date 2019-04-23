rem 生成同步数据的proto
protogen.exe --proto_path protogen --tmpl_path protogen --tmpl_file proto.tmpl --output ../conf/protobuf/*.proto playerdata.proto
rem 生成同步数据的go代码
protogen.exe --proto_path protogen --tmpl_path protogen --tmpl_file go.tmpl --output ../baseapp/entity/*.go playerdata.proto
rem 生成同步数据的protobuf代码
protoc.exe --plugin=protoc-gen-go.exe --go_out=../proto ../conf/protobuf/logic.proto ../conf/protobuf/battle.proto ../conf/protobuf/cluster.proto ../conf/protobuf/playerdata.proto -I ../conf/protobuf/
rem 生成同步数据的cs代码
protogen.exe --proto_path protogen --tmpl_path protogen --tmpl_file cs.tmpl --output ../../../client/Assets/script/proto/*.cs playerdata.proto

cd protocs

del logic.proto
copy ..\..\conf\protobuf\logic.proto .\logic.proto
protogen.exe -i:logic.proto -o:../../../../client/Assets/script/proto/logic.cs -p:detectMissing
del logic.proto

del battle.proto
copy ..\..\conf\protobuf\battle.proto .\battle.proto
protogen.exe -i:battle.proto -o:../../../../client/Assets/script/proto/battle.cs -p:detectMissing
del battle.proto

del cluster.proto
copy ..\..\conf\protobuf\cluster.proto .\cluster.proto
protogen.exe -i:cluster.proto -o:../../../../client/Assets/script/proto/cluster.cs -p:detectMissing
del cluster.proto

del playerdata.proto
copy ..\..\conf\protobuf\playerdata.proto .\playerdata.proto
protogen.exe -i:playerdata.proto -o:../../../../client/Assets/script/proto/pbplayer.cs -p:detectMissing
del playerdata.proto

cd ..