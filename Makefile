PROJECT="shares-timeseries"

default:
	echo ${PROJECT}

.PHONY: protogen
protogen:
	# здесь /home/dmitry/include/googleapis - путь к официальному репозиторию googleapis
	# который нужно предварительно склонировать командой: git clone https://github.com/googleapis/googleapis.git в эту (или другую) директорию
	# подробности читать тут: https://laradrom.ru/tag/proto/
	protoc --go_out=. --go-grpc_out=. -I.  -I/home/dmitry/include/googleapis proto/shares.proto


####################### Миграции
# Используем бинарник migrate
# можно установить готовый бинарный файл с поддержкой всех драйверов:
# curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz
# получаем бинарник и копируем его в /home/dmitry/go/pkg/mod/bin/migrate

# В базе будет создана таблица scheme_migrationa с полями: version и  dirty
# version - номер миграции (индекс файла миграции вида 000001, 000002 и т.д.)
# dirty - status выполнения миграйии, false - успешно, true - не выполнена (грязное выполнение)

# Запуск: make migrate-up DATABASE_URL="postgres://user:password@localhost:5432/mydb?sslmode=disable"
# ну или задаем DATABASE_URL напрямую здесь (это база с таблицей миграций)
DATABASE_URL="clickhouse://mpmhouse:mpmhouse@localhost:9000/default"

# Путь к миграциям
MIGRATIONS_DIR = ./migration

# Команда для выполнения миграций
MIGRATE = migrate

# Если DATABASE_URL не задан, выводим ошибку
#ifndef DATABASE_URL
#$(error DATABASE_URL is not set. Pass it as an argument: make <target> DATABASE_URL=<connection_string>)
#endif

# Создание миграции
# Вызов:  make migrate-create name=create_table_coin
# тут name - корень имени файла миграции
migrate-create:
	$(MIGRATE) create -ext sql -dir $(MIGRATIONS_DIR) -seq $(name)

# Вызов:  make migrate-up
migrate-up:
	$(MIGRATE) -path $(MIGRATIONS_DIR) -database "clickhouse://mpmhouse:mpmhouse@localhost:9000/mpmhouse" up

# восстановить состояние
# n — это номер миграции, на который вы хотите вернуться. Это может быть версия,
# которая была завершена перед возникновением ошибки, либо последняя успешная версия.
migrate-force:
	$(MIGRATE) -path $(MIGRATIONS_DIR) -database "clickhouse://mpmhouse:mpmhouse@localhost:9000/mpmhouse" force $(n)

# Откатить одну миграцию
migrate-down:
	$(MIGRATE) -path $(MIGRATIONS_DIR) -database "$(DATABASE_URL)" down 1

# Откатить все миграции
migrate-reset:
	$(MIGRATE) -path $(MIGRATIONS_DIR) -database "$(DATABASE_URL)" down -all

# Применить определённое количество миграций
migrate-step:
	$(MIGRATE) -path $(MIGRATIONS_DIR) -database "$(DATABASE_URL)" up $(n)

# Откатить определённое количество миграций
migrate-down-step:
	$(MIGRATE) -path $(MIGRATIONS_DIR) -database "$(DATABASE_URL)" down $(n)

