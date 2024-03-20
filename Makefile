ini:
	rm -rf github.com/ini && \
	git clone https://github.com/go-ini/ini github.com/ini && \
	rm -rf github.com/ini/.git && \
	rm -rf github.com/ini/.github && \
	rm -rf github.com/ini/testdata && \
	rm -rf github.com/ini/*_test.go && \
	rm -rf github.com/ini/go.* && \
	git add github.com/ini && git commit -m 'feat: sync clear github.com/go-ini/ini'

jwt:
	rm -rf github.com/jwt && \
	git clone https://github.com/golang-jwt/jwt github.com/jwt && \
	rm -rf github.com/jwt/.git && \
	rm -rf github.com/jwt/.github && \
	rm -rf github.com/jwt/cmd && \
	rm -rf github.com/jwt/request && \
	rm -rf github.com/jwt/test && \
	rm -rf github.com/jwt/*_test.go && \
	rm -rf github.com/jwt/go.* && \
	git add github.com/jwt && git commit -m 'feat: sync clear github.com/golang-jwt/jwt'

redigo:
	rm -rf github.com/redigo && \
	git clone https://github.com/gomodule/redigo github.com/redigo && \
	rm -rf github.com/redigo/.git && \
	rm -rf github.com/redigo/.github && \
	rm -rf github.com/redigo/redisx && \
	rm -rf github.com/redigo/go.* && \
	cp github.com/redigo/redis/* github.com/redigo && \
	rm -rf github.com/redigo/redis && \
	rm -rf github.com/redigo/*_test.go && \
	sed -i '' 's/package redis/package redigo/' github.com/redigo/*.go && \
	git add github.com/redigo && git commit -m 'feat: sync clear github.com/gomodule/redigo'

cli:
	rm -rf github.com/cli && \
	git clone https://github.com/urfave/cli github.com/cli && \
	rm -rf github.com/cli/.git && \
	rm -rf github.com/cli/.github && \
	rm -rf github.com/cli/autocomplete && \
	rm -rf github.com/cli/docs && \
	rm -rf github.com/cli/internal && \
	rm -rf github.com/cli/testdata && \
	rm -rf github.com/cli/*_test.go && \
	rm -rf github.com/cli/go.* && \
	rm -rf github.com/cli/suggestions.go && \
	git add github.com/cli && git commit -m 'feat: sync clear github.com/urfave/cli'
	
github:
	make ini && make jwt && make redigo && make cli

mod:
	go mod tidy