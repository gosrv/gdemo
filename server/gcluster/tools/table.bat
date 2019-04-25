rem ����go����
excelreader.exe -meta tabtmpl/meta.json -excel ../../../excel -out ../conf/table -data true
rem ����csharp����
excelreader.exe -meta tabtmpl/meta.json -excel ../../../excel -out ../../../client/Assets/Resources/table -data true
rem ����proto
excelreader.exe -meta tabtmpl/meta.json -excel ../../../excel -out ../conf/table/table.proto -tmpl proto
rem ����go����
excelreader.exe -meta tabtmpl/meta.json -excel ../../../excel -out ../table/table.go -tmpl go
go fmt ../table/table.go
rem ����go��protobuf����
protoc --plugin=protoc-gen-go.exe --go_out=../table ../conf/table/table.proto -I ../conf/table
rem ����csharp����
excelreader.exe -meta tabtmpl/meta.json -excel ../../../excel -out ../../../client/Assets/script/table/TableMgr.cs -tmpl csharp