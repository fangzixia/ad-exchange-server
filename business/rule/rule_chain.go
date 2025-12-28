package rule

import (
	_interface "ad-exchange-server/core/interface"
	"ad-exchange-server/core/model"
	"errors"
	"log"
)

// RuleChain 规则责任链
type RuleChain struct {
	rules []_interface.MediaRequestRule
}

// NewRuleChain 创建规则责任链
func NewRuleChain(rules []_interface.MediaRequestRule) *RuleChain {
	return &RuleChain{
		rules: rules,
	}
}

// Execute 执行规则链过滤
func (rc *RuleChain) Execute(c *model.AdMediaContent) (bool, error) {
	for _, rule := range rc.rules {
		log.Printf("执行规则: %s", rule.GetRuleName())
		pass, err := rule.Filter(c)
		if err != nil {
			return false, errors.New("规则[" + rule.GetRuleName() + "]执行失败: " + err.Error())
		}
		if !pass {
			return false, errors.New("请求被规则[" + rule.GetRuleName() + "]过滤")
		}
	}
	return true, nil
}
