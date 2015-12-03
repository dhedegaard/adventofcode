all:
	for F in day*.go; do \
		echo "---- BEGIN $$F ----"; \
		go run $$F; \
	done
	echo "--- ALL DONE! ----"
