all: clean
	go run reader.go

clean:
	$(RM) *.gch *.test *.out
