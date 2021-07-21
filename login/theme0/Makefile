all: remove-no-use cpjs compile

check:
	ifeq (, $(shell which adm))
		$(error "adm not found, consider doing: go install github.com/GoAdminGroup/go-admin/adm")
	endif

remove-no-use:
	find ./ -name ".DS_Store" -depth -exec rm {} \;

combine:
	adm combine js --src=./assets/src/js/combine/ --dist=./assets/login/dist/all.min.js
	adm combine css --src=./assets/src/css/ --dist=./assets/login/dist/all.min.css

cpjs:
	cp ./assets/src/js/combine/* ./assets/login/dist/

compile:
	adm compile asset --src=./assets/login/dist/ --dist=./ --package=theme0
	GO111MODULE=off go fmt ./...

.PHONY: all check remove-no-use cpjs compile