all: run

run:	
	go build -v
	./MovieSuggester


.PHONY: all run
