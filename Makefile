PHONY: run-docker-%
run-docker-%: build-docker-%
	@echo "Part1:"
	@docker run -e part=part1 herrnan/aoc-go-$(@:run-docker-%=%)
	@echo "Part2:"
	@docker run -e part=part2 herrnan/aoc-go-$(@:run-docker-%=%)

PHONY: build-docker-%
build-docker-%:
	@echo "Building docker image..."
	@docker build -q $(@:build-docker-%=%) -t herrnan/aoc-go-$(@:build-docker-%=%) > /dev/null
