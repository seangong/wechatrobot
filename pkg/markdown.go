package pkg

import (
	"bytes"
	"fmt"

	"wechatrobot/model"
)

func Markdown(notification model.Notification) (markdown *model.WeChatMarkdown, robotURL string, err error) {

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
			MarkdownTemplate(&bufferFiring, &labels, &annotations, &alert, "warning")
		} else {
			resolvedFlag = true
			MarkdownTemplate(&bufferResolved, &labels, &annotations, &alert, "info")
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
