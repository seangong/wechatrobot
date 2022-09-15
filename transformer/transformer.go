package transformer

import (
	"bytes"
	"fmt"

	"wechatrobot/model"
)

// TransformToMarkdown transform alertmanager notification to wechat markdow message
func TransformToMarkdown(notification model.Notification) (markdown *model.WeChatMarkdown, robotURL string, err error) {

	status := notification.Status

	annotations := notification.CommonAnnotations
	robotURL = annotations["wechatRobot"]

	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("### 告警状态 : <font color=\"warning\">%s</font>\n", status))

	for _, alert := range notification.Alerts {
		labels := alert.Labels
		annotations := alert.Annotations

		buffer.WriteString(fmt.Sprintf("\n> 告警名称: %s", labels["alertname"]))
		buffer.WriteString(fmt.Sprintf("\n> 告警级别: %s", labels["severity"]))
		buffer.WriteString(fmt.Sprintf("\n> 告警主题: %s", annotations["summary"]))
		buffer.WriteString(fmt.Sprintf("\n> 告警详情: %s", annotations["description"]))
		buffer.WriteString(fmt.Sprintf("\n> 故障主机: %s", labels["instance"]))
		buffer.WriteString(fmt.Sprintf("\n> 触发时间: %s", alert.StartsAt.Format("2006-01-02 15:04:05")))
	}

	markdown = &model.WeChatMarkdown{
		MsgType: "markdown",
		Markdown: &model.Markdown{
			Content: buffer.String(),
		},
	}

	return
}
