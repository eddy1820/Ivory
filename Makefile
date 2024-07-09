DB_URL=mysql://root:123456@tcp(localhost:7600)/IvoryDb
name=add_users

u:
	docker compose up -d
d:
	docker compose down
b:
	docker compose build

r:
	docker compose down
	docker compose up -d


migrateup:
	migrate -path migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path migration -database "$(DB_URL)" -verbose down 1

migrate_force:
	migrate -path migration -database "$(DB_URL)" -verbose force 1

new_migration:
	migrate create -ext sql -dir migration -seq $(name)