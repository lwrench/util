### 基本使用
brew 安装 kafka
> https://github.com/orgs/Homebrew/discussions/2530
```
HOMEBREW_BOTTLE_DOMAIN= brew install kafka
```

创建topic:

```
kafka-topics --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --topic test
```

生产者：
```
kafka-console-producer --broker-list localhost:9092 --topic test
```

消费者
```
kafka-console-consumer --bootstrap-server localhost:9092 --topic test --from-beginning
```

查询topic
```
kafka-topics --list --bootstrap-server localhost:9092
```

单播
```
kafka-console-consumer --bootstrap-server localhost:9092 --consumer-property group.id=testGroup --topic test
```

多播
```
kafka-console-consumer --bootstrap-server localhost:9092 --consumer-property group.id=testGroup1 --topic test
kafka-console-consumer --bootstrap-server localhost:9092 --consumer-property group.id=testGroup2 --topic test
```

查看消费组及信息
```
kafka-consumer-groups --bootstrap-server localhost:9092 --list
kafka-consumer-groups --bootstrap-server localhost:9092 --describe --group testGroup
```

### 基础概念
- topic主题
一个业务类型的别名
- partition分区
主题中消息非常大，可以通过分区分布式存储，同时可以并行写
```
kafka-topics --create --bootstrap-server localhost:9092 --partitions 2 --topic test1
kafka-topics --describe --bootstrap-server localhost:9092 --topic test1
```