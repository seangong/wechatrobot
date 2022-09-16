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

func TransformToMarkdown(notification model.Notification) (markdown *model.WeChatMarkdown, robotURL string, err error) {

	annotations := notification.CommonAnnotations
	robotURL = annotations["wechatRobot"]

	var (
		buffer         bytes.Buffer
		bufferFiring   bytes.Buffer
		bufferResolved bytes.Buffer
		firingFlag     bool
		resolvedFlag   bool
	)

	bufferFiring.WriteString(fmt.Sprintf("\n### <font color=\"warning\">告警触发 : %s</font>\n", "firing"))
	bufferResolved.WriteString(fmt.Sprintf("\n### <font color=\"info\">告警恢复 : %s</font>\n", "resolved"))

	for _, alert := range notification.Alerts {
		labels := alert.Labels
		annotations := alert.Annotations

		if alert.Status == "firing" {
			firingFlag = true
			bufferFiring.WriteString(fmt.Sprintf("\n> **alertname: <font color=\"warning\">%s</font>**", labels["alertname"]))
			bufferFiring.WriteString(fmt.Sprintf("\n> **severity:** %s", labels["severity"]))
			bufferFiring.WriteString(fmt.Sprintf("\n> **status:** <font color=\"warning\">%s</font>", alert.Status))
			bufferFiring.WriteString(fmt.Sprintf("\n> **instance:** %s", labels["instance"]))
			delete(labels, "alertname")
			delete(labels, "severity")
			MapToString(&bufferFiring, "annotations", &annotations)
			MapToString(&bufferFiring, "labels", &labels)
			bufferFiring.WriteString(fmt.Sprintf("\n> **time:** <font color=\"comment\">%s</font>\n", alert.StartsAt.Format("2006-01-02 15:04:05")))
		} else {
			resolvedFlag = true
			bufferResolved.WriteString(fmt.Sprintf("\n> **alertname: <font color=\"info\">%s</font>**", labels["alertname"]))
			bufferResolved.WriteString(fmt.Sprintf("\n> **severity:** %s", labels["severity"]))
			bufferResolved.WriteString(fmt.Sprintf("\n> **status:** <font color=\"info\">%s</font>", alert.Status))
			bufferResolved.WriteString(fmt.Sprintf("\n> **instance:** %s", labels["instance"]))
			delete(labels, "alertname")
			delete(labels, "severity")
			MapToString(&bufferResolved, "annotations", &annotations)
			MapToString(&bufferResolved, "labels", &labels)
			bufferResolved.WriteString(fmt.Sprintf("\n> **time:** <font color=\"comment\">%s</font>\n", alert.StartsAt.Format("2006-01-02 15:04:05")))
		}
	}

	if firingFlag {
		buffer.WriteString(bufferFiring.String())
	}
	if resolvedFlag {
		buffer.WriteString(bufferResolved.String())
	}

	markdown = &model.WeChatMarkdown{
		MsgType: "markdown",
		Markdown: &model.Markdown{
			Content: buffer.String(),
		},
	}

	return
}
