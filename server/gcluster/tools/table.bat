rem 生成go数据
excelreader.exe -meta tabtmpl/meta.json -excel ../../../excel -out ../conf/table -data true
rem 生成csharp数据
excelreader.exe -meta tabtmpl/meta.json -excel ../../../excel -out ../../../client/Assets/Resources/table -data true
rem 生成proto
excelreader.exe -meta tabtmpl/meta.json -excel ../../../excel -out ../conf/table/table.proto -tmpl proto
rem 生成go代码
excelreader.exe -meta tabtmpl/meta.json -excel ../../../excel -out ../proto/table.go -tmpl go
rem 生成go的protobuf代码
protoc --plugin=protoc-gen-go.exe --go_out=../proto ../conf/table/table.proto -I ../conf/table
rem 生成csharp代码
excelreader.exe -meta tabtmpl/meta.json -excel ../../../excel -out ../../../client/Assets/script/table/TableMgr.cs -tmpl csharp