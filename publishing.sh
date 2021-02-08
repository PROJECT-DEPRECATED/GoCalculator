# Build Server Only
version=1.0v

rm -rf build/

go build src/Calculator.go
mkdir build/publishing/

mv Calculator build/publishing/

mv Calculator Calculator-$version
cp Calculator Calculator-$version.exe

cp RunScript/Calculator.sh build/publishing
zip Calculator-1.0v.zip build/publishing