package context

import (
	"github.com/stretchr/testify/assert"
	"github.com/zclconf/go-cty/cty"
	"testing"
)

func TestNewFeatureContextAndStop(t *testing.T) {
	ctx:= New(map[string]cty.Value{
		"firstname": cty.StringVal("Sally"),
	}).(*orionContext)
	assert.NotNil(t,ctx.Variables())
	assert.NotNil(t,ctx.EvalContext())
	assert.Equal(t,ctx.EvalContext().Variables["firstname"],cty.StringVal("Sally"))
	ctx.StartScenario()
	ctx.StopScenario()
	assert.NotNil(t,ctx.metrics.endTime)
}

func TestNewFeatureContextAndFail(t *testing.T) {
	ctx:= New(map[string]cty.Value{
		"firstname": cty.StringVal("Sally"),
	}).(*orionContext)
	assert.NotNil(t,ctx.Variables())
	assert.NotNil(t,ctx.EvalContext())
	assert.Equal(t,ctx.EvalContext().Variables["firstname"],cty.StringVal("Sally"))
	ctx.StartScenario()
	ctx.FailScenario()
	assert.NotNil(t,ctx.metrics.endTime)
}