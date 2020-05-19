run:
	@go build -o game.o ./main.go && ./game.o
build:
	@go build -o game.o ./main.go
clean:
	@rm ./game.o
todo:
	@grep TODO -RIn --exclude={README.md,Makefile} *
