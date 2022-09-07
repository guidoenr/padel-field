# padel-field

## **Setup** - Windows 10
1. Download [Docker](https://runnable.com/docker/install-docker-on-windows-10) for windows
2. Once installed, open a `cmd` and type `docker --version`
3. The output will be something like `Docker version 20.10.17, build 100c701`, that means you have installed docker succesfully


## Run the api
1. `git clone https://github.com/guidoenr/padel-field`
2. `cd /padel-field`
3. `make start-env`
4. `make build-api`
5. `make start-api`
6. enter `http://localhost:8080` or `http://127.0.0.1:8080`
7. you will see the main page of the api, that means the api is working
8. DB NOT WORKING
