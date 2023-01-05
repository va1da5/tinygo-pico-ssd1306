.PHONY: help
help:
	@echo "Usage: "
	@echo " bitmap 			- build and flash bitmap example"
	@echo " animation 		- build and flash animation example"
	@echo " help 			- show this message"

.PHONY: bitmap
bitmap:
	cd ./bitmap && make all

.PHONY: animation
animation:
	cd ./animation && make all
