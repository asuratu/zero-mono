air := $(shell which air)
sql2pb := $(shell which sql2pb)
goctl := $(shell which goctl)
date = $(shell date "+%Y-%m-%d-%H:%M:%S")

dir = $(shell pwd)

api:
	@cd $(dir)/app/$(m)/api && $(air) $(m).go -f etc/$(m).yaml

rpc:
	@cd $(dir)/app/$(m)/rpc && $(air) $(m).go -f etc/$(m).yaml

job:
	@cd $(dir)/app/mqueue/cmd/job && $(air) mqueue.go -f etc/mqueue.yaml

scheduler:
	@cd $(dir)/app/mqueue/cmd/scheduler && $(air) scheduler.go -f etc/scheduler.yaml

sql2pb:
	@$(sql2pb) -go_package ./pb -host localhost -package user -password tu4211241992 -port 3306 -schema go-zero-looklook -table user,user_auth -service_name user -user root > user1.proto

model:
	@cd deploy/script/mysql && sh genModel.sh $(m) $(t) $(c)

model.origin:
	@cd deploy/script/mysql && sh genModelOrigin.sh $(m) $(t) $(c)

pb:
	@cd $(dir)/app/$(m)/rpc && $(goctl) rpc protoc $(m).proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=./ --style=goZero -m


############################################################
# 数据库相关:
# - docker exec -it mysql mysqldump -u用户名 -p密码 数据库 > /mnt/vdb/data/mysql/test_db.sql
############################################################

db.init:
	docker exec -i mysql_stage mysql --default-character-set=utf8mb4 -uroot -pPXDN93VRKUm8TeE7 ms_user < ./deploy/sql/migrations.sql

db.dump:
	docker exec -it mysql_stage mysqldump --no-data --default-character-set=utf8mb4 -R -E --hex-blob -uroot -pPXDN93VRKUm8TeE7 ms_user > ./deploy/sql/backup_schema_$(date).sql

db.dump.with.data:
	docker exec -it mysql_stage mysqldump --default-character-set=utf8mb4 -R -E --hex-blob -uroot -pPXDN93VRKUm8TeE7 ms_user > ./deploy/sql/backup_schema_data_$(date).sql
