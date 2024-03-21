.PHONY: ini jwt cli redigo all cache

all:
	make ini && make jwt && make redigo && make cli

ini:
	rm -rf ini && \
	git clone https://github.com/go-ini/ini ini && \
	rm -rf ini/.git && \
	rm -rf ini/.github && \
	rm -rf ini/testdata && \
	rm -rf ini/*_test.go && \
	rm -rf ini/go.* && \
	git add ini && git commit -m 'feat: sync clear github.com/go-ini/ini'

jwt:
	rm -rf jwt && \
	git clone https://github.com/golang-jwt/jwt jwt && \
	rm -rf jwt/.git && \
	rm -rf jwt/.github && \
	rm -rf jwt/cmd && \
	rm -rf jwt/request && \
	rm -rf jwt/test && \
	rm -rf jwt/*_test.go && \
	rm -rf jwt/go.* && \
	git add jwt && git commit -m 'feat: sync clear github.com/golang-jwt/jwt'

redigo:
	rm -rf redigo && \
	git clone https://github.com/gomodule/redigo redigo && \
	rm -rf redigo/.git && \
	rm -rf redigo/.github && \
	rm -rf redigo/redisx && \
	rm -rf redigo/go.* && \
	cp redigo/redis/* redigo && \
	rm -rf redigo/redis && \
	rm -rf redigo/*_test.go && \
	sed -i '' 's/package redis/package redigo/' redigo/*.go && \
	git add redigo && git commit -m 'feat: sync clear github.com/gomodule/redigo'

cli:
	rm -rf cli && \
	git clone https://github.com/urfave/cli cli && \
	rm -rf cli/.git && \
	rm -rf cli/.github && \
	rm -rf cli/autocomplete && \
	rm -rf cli/docs && \
	rm -rf cli/internal && \
	rm -rf cli/testdata && \
	rm -rf cli/*_test.go && \
	rm -rf cli/go.* && \
	rm -rf cli/suggestions.go && \
	git add cli && git commit -m 'feat: sync clear github.com/urfave/cli'

resize:
	rm -rf resize && \
	git clone https://github.com/nfnt/resize resize && \
	rm -rf resize/.git && \
	rm -rf resize/*_test.go && \
	git add resize && git commit -m 'feat: sync clear github.com/nfnt/resize'

cache:
	rm -rf cache && \
	git clone https://github.com/patrickmn/go-cache cache && \
	rm -rf cache/.git && \
	rm -rf cache/*_test.go && \
	git add cache && git commit -m 'feat: sync clear github.com/patrickmn/go-cache'

msgpack:
	rm -rf msgpack && \
	git clone https://github.com/shamaton/msgpack msgpack && \
	rm -rf msgpack/.git && \
	rm -rf msgpack/.github && \
	rm -rf msgpack/*_test.go && \
	rm -rf msgpack/testdata && \
	rm -rf msgpack/go.* && \
	sed -i '' 's/github.com\/shamaton\/msgpack/github.com\/opentoys\/pkg\/msgpack/' msgpack/*.go && \
	sed -i '' 's/github.com\/shamaton\/msgpack/github.com\/opentoys\/pkg\/msgpack/' msgpack/time/*.go && \
	sed -i '' 's/github.com\/shamaton\/msgpack/github.com\/opentoys\/pkg\/msgpack/' msgpack/time/*.go && \
	sed -i '' 's/github.com\/shamaton\/msgpack/github.com\/opentoys\/pkg\/msgpack/' msgpack/internal/decoding/*.go && \
	sed -i '' 's/github.com\/shamaton\/msgpack/github.com\/opentoys\/pkg\/msgpack/' msgpack/internal/encoding/*.go && \
	sed -i '' 's/github.com\/shamaton\/msgpack/github.com\/opentoys\/pkg\/msgpack/' msgpack/internal/stream/decoding/*.go && \
	sed -i '' 's/github.com\/shamaton\/msgpack/github.com\/opentoys\/pkg\/msgpack/' msgpack/internal/stream/encoding/*.go && \
	git add cache && git commit -m 'feat: sync clear github.com/shamaton/msgpack'

mod:
	go mod tidy