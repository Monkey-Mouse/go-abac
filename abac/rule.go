package abac

import (
	"context"
	"log"
	"time"
)

type ContextType interface {
	Value(key interface{}) interface{}
}

type DefaultContext map[string]interface{}

func (c DefaultContext) Value(key interface{}) interface{} {
	if key == 0 {
		return c
	}
	if keyAsString, ok := key.(string); ok {
		val, _ := c[keyAsString]
		return val
	}
	return nil
}

func processRule(ctx context.Context, rules RulesType) (pass bool) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	doneChan := make(chan bool)
	for _, rule := range rules {
		go func(rule RuleType, ctx context.Context) {
			var res bool
			var err error
			if res, err = rule.JudgeRule(); err != nil {
				log.Println(err)
				res = false
			}
			select {
			case <-ctx.Done():
				return
			case doneChan <- res:
			}
		}(rule, ctx)
	}
	for i := 0; i < len(rules); i++ {
		if d := <-doneChan; d {
			cancel()
			pass = true
			return
		}
	}
	pass = false
	return
}

func andProcessRule(ctx ContextType, rules RulesType) bool {
	for _, rule := range rules {
		rule.ProcessContext(ctx)
		if res, err := rule.JudgeRule(); err != nil || !res {
			return false
		}
	}
	return true
}

func testCtx(ctx context.Context) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	select {
	case <-ctx.Done():
		return false, ctx.Err()
	case <-time.After(1 * time.Minute):

	}
	print("here")
	time.Sleep(time.Minute * 1)
	return true, nil
}
