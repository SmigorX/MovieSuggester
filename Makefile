all: run

run:	
	. ./db/.env
	go build
	./MovieSuggester


.PHONY: all run
