BINARY=sveltekit
PUBLIC_DIR=public

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
	if [ -d ${PUBLIC_DIR} ] ; then rm -rf ${PUBLIC_DIR} ; fi

static: clean
	cd frontend;\
	npm install;\
	npm run build

build: static
	docker build . -t sveltekit:latest

run:
	docker run -d -p 8080:8080 --name sveltekit sveltekit:latest

.PHONY: clean static
