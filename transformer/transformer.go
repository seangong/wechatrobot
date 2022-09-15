package transformer

import (
	"bytes"
	"fmt"

	"wechatrobot/model"
)

// TransformToMarkdown transform alertmanager notification to wechat markdow message
func TransformToMarkdown(notification model.Notification) (markdown *model.WeChatMarkdown, robotURL string, err error) {

	annotations := notification.CommonAnnotations
	robotURL = annotations["wechatRobot"]

	var buffer bytes.Buffer

	for _, alert := range notification.Alerts {
		labels := alert.Labels
		annotations := alert.Annotations

		if alert.Status == "firing" {
			buffer.WriteString(fmt.Sprintf("### <font color=\"warning\">告警触发 : %s</font>\n", alert.Status))
			buffer.WriteString(fmt.Sprintf("\n> **alertname:** <font color=\"warning\">%s</font>", labels["alertname"]))
		} else {
			buffer.WriteString(fmt.Sprintf("### <font color=\"info\">告警恢复 : %s</font>\n", alert.Status))
			buffer.WriteString(fmt.Sprintf("\n> **alertname:** <font color=\"info\">%s</font>", labels["alertname"]))
		}

		buffer.WriteString(fmt.Sprintf("\n> **severity:** %s", labels["severity"]))

		if alert.Status == "firing" {
			buffer.WriteString(fmt.Sprintf("\n> **status:** <font color=\"warning\">%s</font>", alert.Status))
		} else {
			buffer.WriteString(fmt.Sprintf("\n> **status:** <font color=\"info\">%s</font>", alert.Status))
		}

		buffer.WriteString(fmt.Sprintf("\n> **instance:** %s", labels["instance"]))
		buffer.WriteString(fmt.Sprintf("\n> **annotations:**\n **·** summary: <font color=\"comment\">%s</font>\n **·** description: %s", annotations["summary"], annotations["description"]))

		delete(labels, "alertname")
		delete(labels, "severity")

		if len(labels) > 0 {
			buffer.WriteString("\n> **labels:**")
			for label, value := range labels {
				buffer.WriteString(fmt.Sprintf("\n **·** %s: %s", label, value))
				fmt.Println(label, ":", value)
			}
		}

		buffer.WriteString(fmt.Sprintf("\n> **time:** <font color=\"comment\">%s</font>\n\n", alert.StartsAt.Format("2006-01-02 15:04:05")))
	}

	markdown = &model.WeChatMarkdown{
		MsgType: "markdown",
		Markdown: &model.Markdown{
			Content: buffer.String(),
		},
	}

	return
}
