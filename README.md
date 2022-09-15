# wechatrobot-webhook

Convert alertmanger webhook messages to qiye WeChat robots that can receive messages

## How use

```
./wechatbot --RobotKey="899220cd-5ed6-44ad-b053-f3785033da7f"
```

or

```
go run main.go --RobotKey="899220cd-5ed6-44ad-b053-f3785033da7f"
```


## configure

alertmanager.yml

```
receivers:
- name: webhook-test                                                                                                           
  webhook_configs:                                                                                                             
  - url: 'http://127.0.0.1:8999/webhook?key=xxxxxx-xxxxx-xxxxx-xxxxxx-xxxxxxx'
```

prometheus rules configure

```
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
      wechatRobot: "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=xxxxxxxx-xxxxxxxx-xxxxx-xxxxxx"

```

## test

### 1、使用默认微信机器人发送消息

```
curl 'http://127.0.0.1:8999/webhook' -H 'Content-Type: application/json' -d '
  {
    "receiver": "webhook-test",
    "status": "firing",
    "alerts": [
      {
        "status": "firing",
        "labels": {
          "alertname": "altername test one",
          "instance": "1.1.1.1",
          "severity": "critical"
        },
        "annotations": {
          "info": "Test message, ignore",
          "description": "This is test summary, ignore",
          "summary": "This is test message, ignore"
        },
        "startsAt": "2022-09-15T02:38:30.763785079Z",
        "endsAt": "0001-01-01T00:00:00Z",
        "generatorURL": "critical"
      }
    ],
    "groupLabels": {
      "alertname": "altername test one"
    },
    "commonLabels": {
      "alertname": "altername test one",
      "instance": "1.1.1.1"
    },
    "commonAnnotations": {
      "info": "Test message, ignore",
      "summary": "This is test summary, ignore"
    },
    "externalURL": "http://localhost:9093",
    "version": "4",
    "groupKey": "{}/{alertname=~\"^(?:test.*)$\"}:{alertname=\"altername test one\"}"
  }'
```



### 2、使用指定微信机器人发送消息

```
curl 'http://127.0.0.1:8999/webhook?key=xxxxxx-xxxxx-xxxxx-xxxxxx-xxxxxxx' -H 'Content-Type: application/json' -d '
  {
    "receiver": "webhook-test",
    "status": "firing",
    "alerts": [
      {
        "status": "firing",
        "labels": {
          "alertname": "altername test two",
          "instance": "2.2.2.2",
          "severity": "critical"
        },
        "annotations": {
          "info": "Test message, ignore",
          "description": "This is test summary, ignore",
          "summary": "This is test message, ignore"
        },
        "startsAt": "2022-09-15T02:38:30.763785079Z",
        "endsAt": "0001-01-01T00:00:00Z",
        "generatorURL": "critical"
      }
    ],
    "groupLabels": {
      "alertname": "altername test two"
    },
    "commonLabels": {
      "alertname": "altername test two",
      "instance": "2.2.2.2"
    },
    "commonAnnotations": {
      "info": "Test message, ignore",
      "summary": "This is test summary, ignore"
    },
    "externalURL": "http://localhost:9093",
    "version": "4",
    "groupKey": "{}/{alertname=~\"^(?:test.*)$\"}:{alertname=\"altername test two\"}"
  }'
```

### 3、使用 alert 内容指定微信机器人发送消息

```
  curl 'http://127.0.0.1:8999/webhook' -H 'Content-Type: application/json' -d '
  {
    "receiver": "webhook-test",
    "status": "firing",
    "alerts": [
      {
        "status": "firing",
        "labels": {
          "alertname": "altername test three",
          "instance": "3.3.3.3",
          "severity": "critical"
        },
        "annotations": {
          "info": "Test message,ignore",
          "description": "This is test summary,ignore",
          "summary": "This is test message,ignore"
        },
        "startsAt": "2022-09-15T02:38:30.763785079Z",
        "endsAt": "0001-01-01T00:00:00Z",
        "generatorURL": "critical"
      }
    ],
    "groupLabels": {
      "alertname": "altername test three"
    },
    "commonLabels": {
      "alertname": "altername test three",
      "instance": "3.3.3.3"
    },
    "commonAnnotations": {
      "info": "Test message,ignore",
      "summary": "This is test summary,ignore",
      "wechatRobot": "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=xxxxxx-xxxxx-xxxxx-xxxxxx-xxxxxxx"
    },
    "externalURL": "http://localhost:9093",
    "version": "4",
    "groupKey": "{}/{alertname=~\"^(?:test.*)$\"}:{alertname=\"altername test three\"}"
  }'
```
