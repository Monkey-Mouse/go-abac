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

// Value will return value for the key in context
func (c DefaultContext) Value(key interface{}) interface{} {
	if keyAsString, ok := key.(string); ok {
		val, _ := c[keyAsString]
		return val
	}
	return nil
}

// Deadline always returns that there is no deadline (ok==false),
func (c *DefaultContext) Deadline() (deadline time.Time, ok bool) {
	return
}

// Done always returns nil (chan which will wait forever),
func (c *DefaultContext) Done() <-chan struct{} {
	return nil
}

// Err always returns nil, maybe you want to use Request.Context().Err() instead.
func (c *DefaultContext) Err() error {
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

// andProcessRule for and logic process
func andProcessRule(ctx ContextType, rules RulesType) (bool, error) {
	for _, rule := range rules {
		rule.ProcessContext(ctx)
		if res, err := rule.JudgeRule(); err != nil || !res {
			return false, err
		}
	}
	return true, nil
}

// orProcessRule for or logic process
func orProcessRule(ctx ContextType, rules RulesType) (res bool, err error) {
	for _, rule := range rules {
		rule.ProcessContext(ctx)
		if res, err = rule.JudgeRule(); err == nil && res {
			return true, nil
		}
	}
	return false, err
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
