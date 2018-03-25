DGRAPH_RUN_DIR := ./dgraph-var

dgraph-start:
	-mkdir ${DGRAPH_RUN_DIR}
	(cd ${DGRAPH_RUN_DIR} && \
		nohup dgraph zero 0<&- &> dgraph-zero.log &)
	(cd ${DGRAPH_RUN_DIR} && \
		nohup dgraph server --memory_mb 1024 0<&- &> dgraph-server.log &)
	(cd ${DGRAPH_RUN_DIR} && \
		nohup dgraph-ratel 0<&- dgraph-ratel.log &)

dgraph-stop:
	-killall dgraph
	-killall dgraph-ratel

dgraph-clean:
	rm -rf ${DGRAPH_RUN_DIR}
