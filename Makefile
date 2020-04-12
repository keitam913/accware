assets:
	cd ./react && npm run build && rsync -av --del build/ ../assets

image: assets
	docker build -t accware .

.PHONY: assets image
