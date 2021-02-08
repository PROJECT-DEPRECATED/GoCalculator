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

else
  echo "Invalid Command: $1"
fi