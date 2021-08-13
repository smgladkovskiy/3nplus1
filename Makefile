MAXIPOWER=7

deps:
	go mod vendor

build:
	go build -o ./bin/collatz main.go

run_iterator: deps build
	./bin/collatz iterator --max-power $(MAXIPOWER)

run_iterator_profiled: deps build
	./bin/collatz iterator --profile-cpu --profile-memory --max-power $(MAXIPOWER)

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