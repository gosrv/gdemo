rem ����go����
excelreader.exe -meta tabtmpl/meta.json -excel ../../../excel -out ../conf/table -data true
rem ����csharp����
excelreader.exe -meta tabtmpl/meta.json -excel ../../../excel -out ../../../client/Assets/Resources/table -data true
rem ����proto
excelreader.exe -meta tabtmpl/meta.json -excel ../../../excel -out ../conf/table/table.proto -tmpl proto
rem ����go����
excelreader.exe -meta tabtmpl/meta.json -excel ../../../excel -out ../proto/table.go -tmpl go
rem ����go��protobuf����
protoc --plugin=protoc-gen-go.exe --go_out=../proto ../conf/table/table.proto -I ../conf/table
rem ����csharp����
excelreader.exe -meta tabtmpl/meta.json -excel ../../../excel -out ../../../client/Assets/script/table/TableMgr.cs -tmpl csharp