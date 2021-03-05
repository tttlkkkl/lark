t?=latest
run:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -tags=jsoniter -ldflags="-s -w" -o $(shell pwd)/dev/app -work $(shell pwd)/dev
	docker build -t h.bitkinetic.com/public/lark:$(t) .
	docker push h.bitkinetic.com/public/lark:$(t)
	helm3 upgrade lark -n dev -f .value.yml --set image.tag=$(t) dn/micro