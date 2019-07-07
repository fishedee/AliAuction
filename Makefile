dev:stop 
	fishcmd run aliauction --watch
start:stop install
	./manager &
	tail -f ./log/access.log
install:
	go install aliauction
	rm -rf aliauction
	cp ${GOPATH}/bin/aliauction .
stop:
	-pkill -9 aliauction
release:stop install
	RUNMODE=prod nohup ./aliauction &

