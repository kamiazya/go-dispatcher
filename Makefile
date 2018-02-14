.PHONY: lifecycle gen


lifecycle:
	cat ./lifecycle.dot | dot -Kdot -Grankdir=TD -Tpng -o lifecycle.png

gen:
	go generate
