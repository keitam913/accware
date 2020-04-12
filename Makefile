build:
	cd ./react && npm run build && rsync -av --del build/ ../assets

.PHONY: build
