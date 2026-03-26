GO=go
PARSER_SRCS=srcs/parser.go \
			srcs/file.go \
			srcs/cmds.go
NAME=parser

all:
	go build $(PARSER_SRCS)

run:
	go run $(PARSER_SRCS)

clean:
	rm -f $(NAME)
