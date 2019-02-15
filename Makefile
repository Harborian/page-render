ifndef VERBOSE
	MAKEFLAGS += --silent
endif

.PHONY: run
run:
	go build && ./page-render