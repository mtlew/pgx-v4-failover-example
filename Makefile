# Step 1. Run db
.PHONY: run
run:
	docker compose up
# Run APP (main.go)
# Open http://localhost:8080/list
# You can see data from db1/postgres.public.test

# Step 2. Shutdown db1
.PHONY: down-db-1
down-db-1:
	docker compose stop db1
# F5
# You can see the message: Error, FATAL: terminating connection due to administrator command (SQLSTATE 57P01)!
# F5
# You can see data from db2/postgres.public.test

# Step 3. Toggle db: down db2, up db1
.PHONY: toggle-db1-db2
toggle-db1-db2:
	docker compose start db1 && docker compose stop db2
# F5
# You can see the message: Error, FATAL: terminating connection due to administrator command (SQLSTATE 57P01)!
# F5
# You can see data from db1/postgres.public.test