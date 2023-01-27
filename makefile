
cloudIP.sqlite3.db:
		$(eval ROOTDIR=$(shell pwd))
		rm -rf cloudIPtoDB 
		git clone git@github.com:stclaird/cloudIPtoDB.git; 
		cd cloudIPtoDB/cmd/main && \
		go build -o cloudIPtoDB -v && \
		chmod +x cloudIPtoDB && \
		ls -lasi && \
		mv output/cloudIP.sqlite3.db $(ROOTDIR)/cmd/api/cloudIP.sqlite3.db;
build:
		$(eval GIT_HASH = $(shell git log --format="%h" -n 1))
		ls -lasi 
		docker build --tag ${REPOPATH}/${APPLICATION_NAME}:${GIT_HASH} .
push:	
		docker push ${REPOPATH}/${APPLICATION_NAME}

