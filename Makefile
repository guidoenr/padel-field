
# install
install: go get . && npm install padel-ui/

# compile
compile: go build .

# run
run: compile \
 	 go run padel-field \
 	 npm start padel-ui/