package main

import (
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
)

func main() {
	knowledgeBase := ast.NewKnowledgeBase("Tutorial", "0.1.1")
	knowledgeBuilder := builder.NewRuleBuilder(knowledgeBase)

	rule := `rule TutorialRule "Test the tutorial rule"  salience 10 {
		when
			MF.IntAttribute == 123 && MF.StringAttribute == "Hello World" && "Hello, World!"`
}

type MyFact struct {
	InAttribute      int
	StringAttribute  string
	BooleanAttribute bool
}

func (mf *MyFact) GetObjectType() string {
	return "MyFact"
}
