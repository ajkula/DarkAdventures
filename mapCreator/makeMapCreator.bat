env GOOS=darwin  GOARCH=amd64 go build -o ../target/osx64/mapCreator
env GOOS=linux   GOARCH=amd64 go build -o ../target/linux64/mapCreator
env GOOS=windows GOARCH=amd64 go build -o ../target/win64/mapCreator.exe