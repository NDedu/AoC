# Advent of Code runner
# Usage:
#   make go   2015/day1.go       (without utils)
#   make gow  2015/day1.go       (with goUtils.go)
#   make c    2015/day1.c
#   make zig  2015/day1.zig
#   make new  YEAR=2015 DAY=10 LANG=go
#   make clean

# Grab the file path from the first non-target argument
FILE = $(wordlist 2,2,$(MAKECMDGOALS))
YEAR = $(dir $(FILE))

# --- Go ---
go:
	@cd $(YEAR) && go run $(notdir $(FILE))

# --- Go with utils ---
gow:
	@cd $(YEAR) && go run $(notdir $(FILE)) goUtils.go

# --- C ---
c:
	@gcc -O2 -Ic_headers -o /tmp/aoc_c $(FILE) -lm && /tmp/aoc_c

# --- Zig ---
zig:
	@cd $(YEAR) && zig run $(notdir $(FILE))

# --- Clean build artifacts ---
clean:
	@rm -f /tmp/aoc_c
	@find . -name "*.o" -delete
	@find . -name "zig-cache" -type d -exec rm -rf {} + 2>/dev/null || true
	@find . -name ".zig-cache" -type d -exec rm -rf {} + 2>/dev/null || true
	@echo "Cleaned."

# --- Create a new day from template ---
# Usage: make new YEAR=2015 DAY=10 LANG=go
# Templates live in templates/<lang>.tmpl
# Use {{DAY}} in templates as a placeholder for the day number
new:
	@mkdir -p $(YEAR)/input
	@touch $(YEAR)/input/day$(DAY).txt
	@if [ -f $(YEAR)/day$(DAY).$(LANG) ]; then \
		echo "$(YEAR)/day$(DAY).$(LANG) already exists"; \
	elif [ ! -f templates/$(LANG).tmpl ]; then \
		echo "No template found at templates/$(LANG).tmpl"; \
	else \
		sed 's/{{DAY}}/$(DAY)/g' templates/$(LANG).tmpl > $(YEAR)/day$(DAY).$(LANG); \
		echo "Created $(YEAR)/day$(DAY).$(LANG)"; \
	fi

# --- Help ---
help:
	@echo "Usage:"
	@echo "  make go   2015/day1.go       Run Go solution"
	@echo "  make gow  2015/day1.go       Run Go solution with goUtils.go"
	@echo "  make c    2015/day1.c        Run C solution"
	@echo "  make zig  2015/day1.zig      Run Zig solution"
	@echo "  make new  YEAR=2015 DAY=10 LANG=go   Create new day from template"
	@echo "  make clean                   Remove build artifacts"

# Swallow the file path argument silently
%: FORCE
	@true

FORCE:
.PHONY: FORCE

.PHONY: go gow c zig clean new help
