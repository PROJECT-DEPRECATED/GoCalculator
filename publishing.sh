# Build Server Only
version=1.0v

rm -rf build/
rm -rf publishing/

go build src/Calculator.go

mv Calculator Calculator-$version
cp Calculator-$version Calculator-$version.exe

mkdir build/
mv Calculator-$version build/

cp RunScript/Calculator.sh build/
zip Calculator-1.0v.zip build/

mkdir publishing/
mv Calculator-$version.zip publishing/