### Микросервис работающий с шарами

На локалке работает с clickhouse кластером (см. infrastructure/clickhouse/clickhouse-docker) из трех узлов

- Хранилище  Clickhouse
- Для реплицируемости важно делать запросы к кластеру (например: CREATE DATABASE IF NOT EXISTS mpmhouse ON CLUSTER clickhouse_cluster;)
- при создании таблиц указывать "ENGINE = ReplicatedMergeTree"
- команды миграций:
  - make migrate-force n=1 (перейти на указанный шаг миграции)
  - make migrate-up (выполнить миграции)
  - make migrate-create name=create_table_coin (создать новую миграцию)

Используем golang пакет clickhouse-go

В настройках пакета указываем все ноды кластера, clickhouse-go сам управляет записью в одну из нод

