# install
install: go get . && npm install padel-ui/

# compile
compile: go build -o .

# run
run: compile \
 	 go run padel-field \
 	 npm start padel-ui/