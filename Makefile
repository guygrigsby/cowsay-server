.PHONY: test
test:
	go test --cowsayPath `which cowsay`
