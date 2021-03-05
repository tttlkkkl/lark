package lark

import (
	"fmt"
	"testing"
)

func TestLark_SendMessage(t *testing.T) {
	chat, err := NewLark(
		"xx",
		"xx",
		SetReceiveMessageAPI("xx", "xx"),
	)
	if err != nil {
		t.Error(err)
		return
	}
	msg, err := NewMessageCard(card, false)
	if err != nil {
		t.Error(err)
		return
	}
	msg.ChatID = "oc_299448deb4a1cd8c0e3bf4b8e13af467"
	rs := chat.SendMessage(msg)
	fmt.Println("======>>>>>", rs)
}

var card = `{
	"config": {
	  "wide_screen_mode": true
	},
	"header": {
	  "title": {
		"tag": "plain_text",
		"content": "李华胜申请将xx/xx发布到生产环境"
	  },
	  "template": "wathet"
	},
	"card_link": {
	  "url": "https://www.baidu.com"
	},
	"i18n_elements": {
	  "zh_cn": [
		{
		  "tag": "div",
		  "text": {
			"tag": "lark_md",
			"content": "仓库：xx\n标签：xx\n"
		  }
		},
		{
		  "tag": "action",
		  "layout": "bisected",
		  "actions": [
			{
			  "tag": "button",
			  "text": {
				"tag": "plain_text",
				"content": "批准"
			  },
			  "type": "primary",
			  "value": {
				"key": "value1"
			  },
			  "confirm": {
				"title": {
				  "tag": "plain_text",
				  "content": "版本发布确认"
				},
				"text": {
				  "tag": "plain_text",
				  "content": "已确定完成代码审查，已确认发布风险，已确保发生任何意外情况时有可靠的回滚措施。"
				}
			  }
			},
			{
			  "tag": "button",
			  "text": {
				"tag": "plain_text",
				"content": "驳回"
			  },
			  "type": "default",
			  "value": {
				"key": "value2"
			  }
			}
		  ]
		}
	  ],
	  "en_us": [
		{
		  "tag": "div",
		  "text": {
			"tag": "lark_md",
			"content": "Empowering teams by messenger, video conference, calendar, docs, and emails. It's all in one place."
		  }
		},
		{
		  "tag": "action",
		  "layout": "bisected",
		  "actions": [
			{
			  "tag": "button",
			  "text": {
				"tag": "plain_text",
				"content": "approval"
			  },
			  "type": "primary",
			  "value": {
				"key": "value11"
			  },
			  "confirm": {
				"title": {
				  "tag": "plain_text",
				  "content": "version release confirmation"
				},
				"text": {
				  "tag": "plain_text",
				  "content": "The code review has been completed, the release risk has been confirmed, and reliable rollback measures have been ensured in case of any unexpected situation."
				}
			  }
			},
			{
			  "tag": "button",
			  "text": {
				"tag": "plain_text",
				"content": "reject"
			  },
			  "type": "default",
			  "value": {
				"key": "value22"
			  }
			}
		  ]
		}
	  ]
	}
  }`
var updateCard = `{
	"config": {
	  "wide_screen_mode": true
	},
	"header": {
	  "title": {
		"tag": "plain_text",
		"content": "李华胜申请将xx/xx发布到生产环境"
	  },
	  "template": "wathet"
	},
	"card_link": {
	  "url": "https://www.baidu.com"
	},
	"i18n_elements": {
	  "zh_cn": [
		{
		  "tag": "div",
		  "text": {
			"tag": "lark_md",
			"content": "仓库：xx\n标签：xx\n"
		  }
		}
		  ]
		}
	  ],
	  "en_us": [
		{
		  "tag": "div",
		  "text": {
			"tag": "lark_md",
			"content": "Empowering teams by messenger, video conference, calendar, docs, and emails. It's all in one place."
		  }
		},
		{
		  "tag": "action",
		  "layout": "bisected",
		  "actions": [
			{
			  "tag": "button",
			  "text": {
				"tag": "plain_text",
				"content": "approval"
			  },
			  "type": "primary",
			  "value": {
				"key": "value11"
			  },
			  "confirm": {
				"title": {
				  "tag": "plain_text",
				  "content": "version release confirmation"
				},
				"text": {
				  "tag": "plain_text",
				  "content": "The code review has been completed, the release risk has been confirmed, and reliable rollback measures have been ensured in case of any unexpected situation."
				}
			  }
			},
			{
			  "tag": "button",
			  "text": {
				"tag": "plain_text",
				"content": "reject"
			  },
			  "type": "default",
			  "value": {
				"key": "value22"
			  }
			}
		  ]
		}
	  ]
	}
  }`
