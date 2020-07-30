env GOOS=darwin  GOARCH=amd64 go build -o target/osx64/DarkAdventures
env GOOS=linux   GOARCH=amd64 go build -o target/linux64/DarkAdventures
env GOOS=windows GOARCH=amd64 go build -o target/win64/DarkAdventures.exe