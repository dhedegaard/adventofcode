all:
	for F in day*.go; do \
		echo $$F; \
		go run $$F; \
	done
