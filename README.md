# wechatrobot

通过 webhook 方法发送消息到企业微信群机器人。

## 使用方法

```bash
./wechatrobot --RobotKey="899220cd-5ed6-44ad-b053-f3785033da7f"
```

or

```bash
go run main.go --RobotKey="899220cd-5ed6-44ad-b053-f3785033da7f"
```


## 配置

alertmanager.yml

```yml
receivers:
  - name: webhook-test                                                                                                           
    webhook_configs:                                                                                                             
      - url: 'http://127.0.0.1:8989/webhook?key=899220cd-5ed6-44ad-b053-f3785033da7f'
```

prometheus rules configure

```yml
groups:
- name: ansible managed alert rules
  rules:
  - alert: CriticalCPULoad
    expr: (100 * (1 - avg(irate(node_cpu_seconds_total{mode="idle"}[5m])) BY (instance)))
      > 96
    for: 2m
    labels:
      severity: critical
    annotations:
      description: '{{ $labels.instance }} of mountpoint {{ $labels.mountpoint }} has
        Critical CPU load for more than 2 minutes.'
      summary: Instance {{ $labels.instance }} - Critical CPU load
      wechatRobot: "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=899220cd-5ed6-44ad-b053-f3785033da7f"

```

## 测试

### 1、默认使用启动程序时微信机器人 token 发送消息

```json
curl 'http://127.0.0.1:8989/webhook' -H 'Content-Type: application/json' -d '
{
    "receiver":"webhook-test",
    "status":"firing",
    "alerts":[
        {
            "status":"firing",
            "labels":{
                "alertname":"altername test one",
                "instance":"1.1.1.1",
                "severity":"critical"
            },
            "annotations":{
                "info":"Test message, ignore",
                "description":"This is test summary, ignore",
                "summary":"This is test message, ignore"
            },
            "startsAt":"2022-09-15T02:38:30.763785079Z",
            "endsAt":"0001-01-01T00:00:00Z",
            "generatorURL":"critical"
        }
    ],
    "groupLabels":{
        "alertname":"altername test one"
    },
    "commonLabels":{
        "alertname":"altername test one",
        "instance":"1.1.1.1"
    },
    "commonAnnotations":{
        "info":"Test message, ignore",
        "summary":"This is test summary, ignore"
    },
    "externalURL":"http://localhost:9093",
    "version":"4",
    "groupKey":"{}/{alertname=~\"^(?:test.*)$\"}:{alertname=\"altername test one\"}"
}'
```

![default_robot](https://github.com/SeanGong/wechatrobot/blob/master/docs/images/default_robot.png)

### 2、指定微信机器人 tocken 发送消息

```json
curl 'http://127.0.0.1:8989/webhook?key=xxxxxx-xxxxx-xxxxx-xxxxxx-xxxxxxx' -H 'Content-Type: application/json' -d '
{
    "receiver":"webhook-test",
    "status":"firing",
    "alerts":[
        {
            "status":"firing",
            "labels":{
                "alertname":"altername test two",
                "instance":"2.2.2.2",
                "severity":"critical"
            },
            "annotations":{
                "info":"Test message, ignore",
                "description":"This is test summary, ignore",
                "summary":"This is test message, ignore"
            },
            "startsAt":"2022-09-15T02:38:30.763785079Z",
            "endsAt":"0001-01-01T00:00:00Z",
            "generatorURL":"critical"
        }
    ],
    "groupLabels":{
        "alertname":"altername test two"
    },
    "commonLabels":{
        "alertname":"altername test two",
        "instance":"2.2.2.2"
    },
    "commonAnnotations":{
        "info":"Test message, ignore",
        "summary":"This is test summary, ignore"
    },
    "externalURL":"http://localhost:9093",
    "version":"4",
    "groupKey":"{}/{alertname=~\"^(?:test.*)$\"}:{alertname=\"altername test two\"}"
}'
```

![default_robot](https://github.com/SeanGong/wechatrobot/blob/master/docs/images/token_robot.png)


### 3、同时发送多条告警消息

```json
curl 'http://127.0.0.1:8989/webhook' -H 'Content-Type: application/json' -d '
{
    "receiver":"webhook-test",
    "status":"firing",
    "alerts":[
        {
            "status":"firing",
            "labels":{
                "alertname":"altername test two",
                "instance":"2.2.2.2",
                "severity":"critical"
            },
            "annotations":{
                "info":"Test message, ignore",
                "description":"This is test summary, ignore",
                "summary":"This is test message, ignore"
            },
            "startsAt":"2022-09-15T02:38:30.763785079Z",
            "endsAt":"0001-01-01T00:00:00Z",
            "generatorURL":"critical"
        },
        {
            "status":"resolved",
            "labels":{
                "alertname":"NodeIOWaitOvercommit",
                "instance":"192.168.0.4:9100",
                "prometheus":"monitoring/k8s",
                "resource_type":"node",
                "severity":"critical"
            },
            "annotations":{
                "description":"High Node CPU IO Wait",
                "message":"",
                "summary":""
            },
            "startsAt":"2022-09-15T02:38:34.637265174Z",
            "endsAt":"2018-11-22T06:28:34.637265174Z",
            "generatorURL":"http://0.0.0.0:9090/graph?g0.expr=instance:node_cpu_iowait:sum+>+1&g0.tab=1"
        }
    ],
    "groupLabels":{
        "alertname":"altername test two"
    },
    "commonLabels":{
        "alertname":"altername test two",
        "instance":"2.2.2.2"
    },
    "commonAnnotations":{
        "info":"Test message, ignore",
        "summary":"This is test summary, ignore"
    },
    "externalURL":"http://localhost:9093",
    "version":"4",
    "groupKey":"{}/{alertname=~\"^(?:test.*)$\"}:{alertname=\"altername test two\"}"
}'
```

![alerts_robot](https://github.com/SeanGong/wechatrobot/blob/master/docs/images/alerts_robot.png)


### 4、alert 内容指定微信机器人 token

```json
curl 'http://127.0.0.1:8989/webhook' -H 'Content-Type: application/json' -d '
{
    "receiver":"webhook-test",
    "status":"firing",
    "alerts":[
        {
            "status":"firing",
            "labels":{
                "alertname":"altername test three",
                "instance":"3.3.3.3",
                "severity":"critical"
            },
            "annotations":{
                "info":"Test message,ignore",
                "description":"This is test summary,ignore",
                "summary":"This is test message,ignore"
            },
            "startsAt":"2022-09-15T02:38:30.763785079Z",
            "endsAt":"0001-01-01T00:00:00Z",
            "generatorURL":"critical"
        }
    ],
    "groupLabels":{
        "alertname":"altername test three"
    },
    "commonLabels":{
        "alertname":"altername test three",
        "instance":"3.3.3.3"
    },
    "commonAnnotations":{
        "info":"Test message,ignore",
        "summary":"This is test summary,ignore",
        "wechatRobot":"https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=xxxxxx-xxxxx-xxxxx-xxxxxx-xxxxxxx"
    },
    "externalURL":"http://localhost:9093",
    "version":"4",
    "groupKey":"{}/{alertname=~\"^(?:test.*)$\"}:{alertname=\"altername test three\"}"
}'
```
