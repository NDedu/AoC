# Advent of Code runner
# Usage:
#   make go    2015/day1
#   make c     2015/day1.c       (opt=debug for sanitizers)
#   make zig   2015/day1.zig     (opt=fast for speed)
#   make zigt  2015/day1.zig     (run tests)
#   make odin  2015/day1         (opt=speed for speed)
#   make odint 2015/day1         (run tests)
#   make new   y=2015 d=10 l=go
#   make input y=2015 d=10
#   make clean

.DEFAULT_GOAL := help

RUN_TARGETS = go c zig zigt odin odint
TARGETS     = $(RUN_TARGETS) new input clean help

# Grab the file path from the first non-target argument
FILE = $(wordlist 2,2,$(MAKECMDGOALS))
YEAR = $(dir $(FILE))

# Run targets need a path: fail early instead of running in the wrong folder
RUN_GOAL     = $(firstword $(filter $(RUN_TARGETS),$(MAKECMDGOALS)))
example_path = 2015/day1$(if $(filter c,$1),.c)$(if $(filter zig zigt,$1),.zig)
ifneq ($(RUN_GOAL),)
ifeq ($(FILE),)
$(error Missing path, e.g. make $(RUN_GOAL) $(call example_path,$(RUN_GOAL)))
endif
endif

# Go and Odin days are folders (one package per day), e.g. 2015/day1/day1.go.
# Their targets take the folder (2015/day1, 2015/day1/ or a file inside it)
# and run from the year folder, so input/dayN.txt paths work as for the others.
DAYDIR = $(patsubst %/,%,$(if $(suffix $(FILE)),$(dir $(FILE)),$(FILE)))

# Puzzle input of a day, for make new / make input
INPUT_FILE = $(y)/input/day$(d).txt

# --- Go ---
go:
	@cd $(dir $(DAYDIR)) && go run ./$(notdir $(DAYDIR))

# --- C ---
# opt: fast (-O2) | debug (-O0 -g with AddressSanitizer and UBSan)
c: opt ?= fast
C_OPT_fast  = -O2
C_OPT_debug = -O0 -g -fsanitize=address,undefined

c:
	@$(if $(C_OPT_$(opt)),,$(error opt for C must be fast or debug, got '$(opt)'))
	@gcc $(C_OPT_$(opt)) -Wall -Wextra -Ic_headers -o /tmp/aoc_c $(FILE) -lm && cd $(YEAR) && /tmp/aoc_c

# --- Zig ---
# opt: debug | fast | safe | small
zig zigt: opt ?= debug

zig:
	@cd $(YEAR) && zig run -O $(opt) $(notdir $(FILE))

zigt:
	@cd $(YEAR) && zig test -O $(opt) $(notdir $(FILE))

# --- Odin ---
# opt: none | minimal | size | speed | aggressive
# -out keeps binaries in /tmp: odin skips its cleanup after a Ctrl-C or failure.
odin odint: opt ?= minimal

odin:
	@cd $(dir $(DAYDIR)) && odin run $(notdir $(DAYDIR)) -o:$(opt) -out:/tmp/aoc_odin

odint:
	@cd $(dir $(DAYDIR)) && odin test $(notdir $(DAYDIR)) -o:$(opt) -out:/tmp/aoc_odin_test

# --- Download a day's puzzle input ---
# Usage: make input y=2015 d=10
# Needs your adventofcode.com session cookie, either exported in your shell
# or as a line AOC_SESSION=<cookie> in .env (gitignored, no quotes).
# AoC asks scripts to identify themselves: set AOC_UA the same way to put
# your own contact info (e.g. repo URL or email) in the User-Agent.
# The cookie only reaches this recipe and goes to curl via stdin, not argv.
-include .env
AOC_UA ?= aoc-makefile-input
input: export AOC_SESSION := $(AOC_SESSION)

input:
	@if [ -z "$(y)" ] || [ -z "$(d)" ]; then \
		echo "Usage: make input y=2015 d=10"; exit 1; \
	fi
	@if [ -s $(INPUT_FILE) ]; then \
		echo "$(INPUT_FILE) already downloaded"; exit 0; \
	fi; \
	if [ -z "$$AOC_SESSION" ]; then \
		echo "Set AOC_SESSION to your adventofcode.com session cookie (env var or .env)"; exit 1; \
	fi; \
	mkdir -p $(y)/input && \
	printf 'Cookie: session=%s\n' "$$AOC_SESSION" | \
		curl -fsS -A "$(AOC_UA)" -H @- -o $(INPUT_FILE).part \
			https://adventofcode.com/$(y)/day/$(d)/input && \
	mv $(INPUT_FILE).part $(INPUT_FILE) && \
	echo "Saved $(INPUT_FILE)" || \
	{ rm -f $(INPUT_FILE).part; exit 1; }

# --- Clean build artifacts ---
clean:
	@rm -f /tmp/aoc_c
	@rm -f /tmp/aoc_odin*
	@find . -name "*.o" -delete
	@find . -name "zig-cache" -type d -exec rm -rf {} + 2>/dev/null || true
	@find . -name ".zig-cache" -type d -exec rm -rf {} + 2>/dev/null || true
	@echo "Cleaned."

# --- Create a new day from template ---
# Usage: make new y=2015 d=10 l=go
# Templates live in templates/<lang>.tmpl
# Use {{DAY}} in templates as a placeholder for the day number
# Go and Odin days get their own folder: 2015/day10/day10.go
FOLDER_LANGS = go odin
NEW_FILE = $(if $(filter $(l),$(FOLDER_LANGS)),$(y)/day$(d)/day$(d).$(l),$(y)/day$(d).$(l))

new:
	@if [ -z "$(y)" ] || [ -z "$(d)" ] || [ -z "$(l)" ]; then \
		echo "Usage: make new y=2015 d=10 l=go"; exit 1; \
	fi
	@mkdir -p $(y)/input
	@touch $(INPUT_FILE)
	@if [ -f $(NEW_FILE) ]; then \
		echo "$(NEW_FILE) already exists"; \
	elif [ ! -f templates/$(l).tmpl ]; then \
		echo "No template found at templates/$(l).tmpl"; \
	else \
		mkdir -p $(dir $(NEW_FILE)); \
		sed 's/{{DAY}}/$(d)/g' templates/$(l).tmpl > $(NEW_FILE); \
		echo "Created $(NEW_FILE)"; \
	fi

# --- Help ---
help:
	@echo "Usage:"
	@echo "  make go    2015/day1         Run Go solution"
	@echo "  make c     2015/day1.c       Run C solution (opt=debug for sanitizers)"
	@echo "  make zig   2015/day1.zig     Run Zig solution (opt=fast for speed)"
	@echo "  make zigt  2015/day1.zig     Run Zig tests"
	@echo "  make odin  2015/day1         Run Odin solution (opt=speed for speed)"
	@echo "  make odint 2015/day1         Run Odin tests"
	@echo "  make new   y=2015 d=10 l=go  Create new day from template"
	@echo "  make input y=2015 d=10       Download puzzle input (needs AOC_SESSION)"
	@echo "  make clean                   Remove build artifacts"

# Swallow the path argument so make doesn't try to build it.
# Only that one word: a mistyped target name still fails.
ifneq ($(filter-out $(TARGETS),$(FILE)),)
.PHONY: $(FILE)
$(FILE):
	@true
endif

.PHONY: $(TARGETS)
