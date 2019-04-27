rem 生成表数据
excelreader.exe -meta tabtmpl/meta.json -excel ../../../excel -out ../conf/table -data true
rem 生成表cs代码
excelreader.exe -meta tabtmpl/meta.json -excel ../../../excel -out ../../../client/Assets/Resources/table -data true
rem 生成表proto
excelreader.exe -meta tabtmpl/meta.json -excel ../../../excel -out ../conf/table/table.proto -tmpl proto
rem 生成表go代码
excelreader.exe -meta tabtmpl/meta.json -excel ../../../excel -out ../table/table.go -tmpl go
go fmt ../table/table.go
rem 生成表proto代码
protoc --plugin=protoc-gen-go.exe --go_out=../table ../conf/table/table.proto -I ../conf/table
rem 生成表cs代码
excelreader.exe -meta tabtmpl/meta.json -excel ../../../excel -out ../../../client/Assets/script/table/TableMgr.cs -tmpl csharp