all: run

run:
	go run cmd/main/main.go

push:
	git push git@github.com:RB-PRO/infocolor.git

pull:
	git pull git@github.com:RB-PRO/infocolor.git

pushW:
	git push https://github.com/RB-PRO/infocolor.git

pullW:
	git pull https://github.com/RB-PRO/infocolor.git