version=1.0v

if [ "$1" == "build" ]; then
  go build src/Calculator.go

  mkdir build/
  mv Calculator build/

elif [ "$1" == "clean" ]; then
  rm -rf build/

  if [ "$2" == "build" ]; then
    go build src/Calculator.go

    mkdir build/
    mv Calculator build/
  fi

# Build Server Only
elif [ "$1" == "publishing" ]; then
  go build src/Calculator.go

  mkdir build/publishing

  mv Calculator build/publishing

  mv Calculator Calculator-$version
  cp Calculator Calculator-$version.exe

  cp RunScript/Calculator.sh build/publishing

  zip Calculator-1.0v.zip build/publishing

else
  echo "Invalid Command: $1"
fi