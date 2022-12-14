/*
 * @Author: robert zhang <robertzhangwenjie@gmail.com>
 * @Date: 2022-08-05 11:52:41
 * @LastEditTime: 2022-09-25 22:51:43
 * @LastEditors: robert zhang
 * @Description:
 */
package commit

import (
	"bytes"
	"log"
	"strings"
	"text/template"

	"github.com/AlecAivazis/survey/v2"
)

type Form struct {
	Name     string
	Desc     string
	Form     string
	Options  []SelectOption
	Required bool
}

type SelectOption struct {
	Name string
	Desc string
}

// fill out the form
func fillOutForm(qs []*survey.Question, tmpl string) ([]byte, error) {

	answers := map[string]interface{}{}
	if err := survey.Ask(qs, &answers); err != nil {
		return nil, err
	}

	// process answers
	for i, answer := range answers {
		if ans, ok := answer.(survey.OptionAnswer); ok {
			answers[i] = ans.Value
		} else if ans, ok := answer.(string); ok {
			answers[i] = strings.TrimSpace(ans)
		}
	}

	// render template
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, answers); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// loadForm load the from with messgaeConfig
func loadForm(msgConfig *messageConfig) (qs []*survey.Question, template string, err error) {
	// customize template for showing multiline's answer in new line
	survey.MultilineQuestionTemplate = multilineQuestionTemplate
	// customize selectQuestionTemplate
	survey.SelectQuestionTemplate = selectQuestionTemplate

	for _, item := range msgConfig.Items {
		q := &survey.Question{
			Name: item.Name,
		}

		if item.Required {
			q.Validate = survey.Required
		}

		switch item.Form {
		case "input":
			q.Prompt = &survey.Input{
				Message: item.Desc,
			}
			q.Validate = survey.MaxLength(72)
		case "multiline":
			q.Prompt = &survey.Multiline{
				Message: item.Desc,
			}

		case "select":
			prompt := &survey.Select{
				Message:  item.Desc,
				PageSize: 8,
			}

			for _, option := range item.Options {
				prompt.Options = append(prompt.Options, option.Name)
			}

			prompt.Description = func(options []SelectOption) func(value string, index int) string {
				return func(value string, index int) string {
					return options[index].Desc
				}
			}(item.Options)

			q.Prompt = prompt

		default:
			log.Fatalf("question not support form: '%s'", item.Form)
		}
		qs = append(qs, q)
	}
	return qs, msgConfig.Template, nil
}
