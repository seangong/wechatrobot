package pkg

import (
	"bytes"
	"fmt"

	"wechatrobot/model"
)

func MapToString(buf *bytes.Buffer, name string, data *map[string]string) {
	if len(*data) > 0 {
		buf.WriteString(fmt.Sprintf("\n> **%s:**", name))
		for key, value := range *data {
			buf.WriteString(fmt.Sprintf("\n **·** %s: %s", key, value))
		}
	}
}

func MarkdownTemplate(buf *bytes.Buffer, labels *map[string]string, annotations *map[string]string, alert *model.Alert, color string) {

	buf.WriteString(fmt.Sprintf("\n> **alertname: <font color=\"%s\">%s</font>**", color, (*labels)["alertname"]))
	buf.WriteString(fmt.Sprintf("\n> **severity:** %s", (*labels)["severity"]))
	buf.WriteString(fmt.Sprintf("\n> **status: <font color=\"%s\">%s</font>**", color, alert.Status))

	if value, ok := (*labels)["instance"]; ok {
		buf.WriteString(fmt.Sprintf("\n> **instance:** %s", value))
	}

	delete((*labels), "alertname")
	delete((*labels), "severity")

	MapToString(buf, "labels", labels)
	MapToString(buf, "annotations", annotations)
	buf.WriteString(fmt.Sprintf("\n> **time:** <font color=\"comment\">%s</font>\n", alert.StartsAt.Format("2006-01-02 15:04:05")))
}
