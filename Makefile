SHELL := /bin/bash

test:
	curl "http://127.0.0.1:8888/microservice/name";
	@echo ' ';
	curl "http://127.0.0.1:8888/user/profile" --header "Username: Uncle Bob";
	@echo ' ';
	curl "http://127.0.0.1:8888/user/profile" --header "Username: Dolly";
	@echo ' ';
	