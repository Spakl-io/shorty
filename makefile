RELEASE_VERSION ?= v0.0.1
RELEASE_COMMIT_COUNT ?= 1
RELEASE_MESSAGE ?= $(shell git log -n $(RELEASE_COMMIT_COUNT) --pretty=format:"%h - %ad | %s [%an]" --graph --date=short)
# %h - Abbreviated commit hash
# %ad - Author date (you can format the date using --date=option, here I used short for a concise date format)
# %s - Commit subject (message)
# %d - Ref names
# %an - Author name
# --graph - Shows a text-based graph of the commits on the left-hand side
# --date=short - Formats the date to a shorter form
clear-local-tags:
	git tag -d $(shell git tag --list)

release:
	@echo "Creating release $(RELEASE_VERSION)"
	@echo "Changelog:"
	@echo "$(RELEASE_MESSAGE)"
	git tag -a $(RELEASE_VERSION) -m "$(RELEASE_MESSAGE)"
	git push origin master --tag $(RELEASE_VERSION)
build:
	go build -o ./bin/shorty.so -buildmode=plugin .


vendor:
	go mod vendor
