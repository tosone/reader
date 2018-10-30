all: clean
	$(CC) -c reader.h
	go run -a reader.go

clean:
	$(RM) *.gch *.test *.out
