# Build Server Only
version=2.0v

rm -rf build/
rm -rf publish/

go build src/Calculator.go

mkdir build/
mv Calculator build/

cp RunScript/Calculator.sh build/
zip Calculator-$version.zip build/

mkdir publish
mv Calculator-$version.zip publish/