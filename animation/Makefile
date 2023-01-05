.PHONY: all
all: build flash

.PHONY: build
build:
	tinygo build -target=pico -o main.bin .

.PHONY: flash
flash:
	tinygo flash -target=pico
