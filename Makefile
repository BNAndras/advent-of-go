day = $(shell date +'%-d')

new:
	@echo "Creating new file structure for day" $(day)"..."

	@if [ $(day) -lt 10 ] ; then \
  		mkdir calendar/day-0$(day); \
  		cp template calendar/day-0$(day)/day0$(day).go; \
  	else \
  		mkdir calendar/day-$(day); \
		cp template calendar/day-$(day)/day$(day).go; \
    fi
	@echo "Files successfully created.. happy hacking :)"
	@git add calendar/


