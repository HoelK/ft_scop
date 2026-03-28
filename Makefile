GO=go
FLAGS=-buildmode=c-archive
NAME=scop

#main
MAIN_SRC=./src/main.rs

#parser
PARSER_SRCS=./src/bridge/bridge.go

#libraries
LIB_DIR=lib
PARSER_LIB=libparser.a
LIBS=$(LIB_DIR)/libparser.a

#headers
HEADER_DIR=header
HEADERS=$(HEADER_DIR)/libparser.h

all: $(PARSER_LIB)
	cd src && cargo build && cd ..

$(PARSER_LIB): $(PARSER_SRCS)
	mkdir -p lib
	mkdir -p header
	go build $(FLAGS) -o libparser.a $(PARSER_SRCS)
	mv libparser.h ./header/.
	mv libparser.a ./lib/.

test:
	go build ./src/bridge

run:
	go run ./src/bridge

clean:
	rm -f $(NAME) $(LIBS) $(HEADERS)
	rm -f main

fclean: clean
	go clean -cache
