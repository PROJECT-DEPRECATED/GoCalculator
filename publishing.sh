# Build Server Only
version=1.0v

rm -rf build/
rm -rf publishing/

go build src/Calculator.go

mkdir build/
mv Calculator build/

mv build/Calculator Calculator-$version
cp build/Calculator-$version Calculator-$version.exe

cp RunScript/Calculator.sh build/
zip Calculator-1.0v.zip build/

mkdir publishing
mv Calculator-$version.zip publishing/