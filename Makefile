MAXIPOWER=7

deps:
	go mod vendor

build:
	go build -o ./bin/iterator ./cmd/iterator

run: deps build
	./bin/iterator -maxIPower $(MAXIPOWER)

run_profiled: deps build
	./bin/iterator -maxIPower $(MAXIPOWER) -cpuprofile cpu.prof -memprofile mem.prof

profile_cpu:
	go tool pprof -http=localhost:8080 cpu.prof

profile_mem:
	go tool pprof -http=localhost:8080 mem.prof



# ----
## LINTER stuff start

linter_include_check:
	@[ -f linter.mk ] && echo "linter.mk include exists" || (echo "getting linter.mk from github.com" && curl -sO https://raw.githubusercontent.com/spacetab-io/makefiles/master/golang/linter.mk)

.PHONY: lint
lint: linter_include_check
	@make -f linter.mk go_lint

## LINTER stuff end
# ----