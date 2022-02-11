
   
EXE  := sense-docker

$(EXE): go.mod *.go
	docker build -t $(EXE) .


