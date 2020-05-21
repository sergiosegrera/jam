run:
	@go build -o game.o ./main.go && ./game.o
build:
	@go build -o game.o ./main.go
windows:
	@GOOS=windows GOARCH=386 go build -o game.exe main.go
	@zip windows.zip assets/* game.exe
clean:
	@rm -f *.{exe,o,zip}
todo:
	@grep TODO -RIn --exclude={README.md,Makefile} *
