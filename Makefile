PHONY: run-docker-%
run-docker-%:
	@echo "Building docker image..."
	@docker build -q $(@:run-docker-%=%) -t herrnan/aoc-go-$(@:run-docker-%=%) > /dev/null
	@echo "Part1:"
	@docker run -e part=part1 herrnan/aoc-go-$(@:run-docker-%=%)
	@echo "Part2:"
	@docker run -e part=part2 herrnan/aoc-go-$(@:run-docker-%=%)
