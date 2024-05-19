refact:
	go mod edit -module github.com/noolingo/deck-service
	-- rename all imported module
	find . -type f -name '*.go' \
  	-exec sed -i -e 's,github.com/noolingo/card-service,github.com/noolingo/deck-service,g' {} \;