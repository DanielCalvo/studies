VERSION = $(shell date +%Y%m%d_%H%M)

build:
	docker build -t dani-salt-master .

run:
	docker run -p 4505:4505 -p 4506:4506 -p 8000:8000 -d --name dani-salt-master dani-salt-master

shell:
	docker container exec -i -t dani-salt-master /bin/bash

top:
	docker container top dani-salt-master

del:
	docker container rm dani-salt-master --force

minion:
	#vagrant up
	echo "todo"
