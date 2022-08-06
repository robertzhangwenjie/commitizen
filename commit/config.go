/*
 * @Author: robert zhang <robertzhangwenjie@gmail.com>
 * @Date: 2022-08-04 17:12:31
 * @LastEditTime: 2022-08-05 17:07:27
 * @LastEditors: robert zhang
 * @Description:
 */
package commit

var selectQuestionTemplate = `
{{- define "option"}}
    {{- if eq .SelectedIndex .CurrentIndex }}{{color .Config.Icons.SelectFocus.Format }}{{ .Config.Icons.SelectFocus.Text }} {{else}}{{color "default"}}  {{end}}
    {{- .CurrentOpt.Value}}{{ if ne ($.GetDescription .CurrentOpt) "" }}: {{ $.GetDescription .CurrentOpt }}{{end}}
{{end}}
{{- if .ShowHelp }}{{- color .Config.Icons.Help.Format }}{{ .Config.Icons.Help.Text }} {{ .Help }}{{color "reset"}}{{"\n"}}{{end}}
{{- color .Config.Icons.Question.Format }}{{ .Config.Icons.Question.Text }} {{color "reset"}}
{{- color "default+hb"}}{{ .Message }}{{ .FilterMessage }}{{color "reset"}}
{{- if .ShowAnswer}}{{color "cyan"}} {{.Answer}}{{color "reset"}}{{"\n"}}
{{- else}}
  {{- "  "}}{{- color "cyan"}}[Use arrows to move, type to filter{{- if and .Help (not .ShowHelp)}}, {{ .Config.HelpInput }} for more help{{end}}]{{color "reset"}}
  {{- "\n"}}
  {{- range $ix, $option := .PageEntries}}
    {{- template "option" $.IterateOption $ix $option}}
  {{- end}}
{{- end}}`

var multilineQuestionTemplate = `
{{- if .ShowHelp }}{{- color .Config.Icons.Help.Format }}{{ .Config.Icons.Help.Text }} {{ .Help }}{{color "reset"}}{{"\n"}}{{end}}
{{- color .Config.Icons.Question.Format }}{{ .Config.Icons.Question.Text }} {{color "reset"}}
{{- color "default+hb"}}{{ .Message }} {{color "reset"}}
{{- if .ShowAnswer}}
  {{- "\n"}}{{color "cyan"}}{{.Answer}}{{color "reset"}}
  {{- if .Answer }}{{ "\n" }}{{ end }}
{{- else }}
  {{- if .Default}}{{color "white"}}({{.Default}}) {{color "reset"}}{{end}}
  {{- color "cyan"}}[Enter 2 empty lines to finish]{{color "reset"}}{{"\n"}}
{{- end}}`

var defaultConfig = `{
	"items": [
			{
				"name": "type",
				"desc": "Select the type of change that you're committing:",
				"form": "select",
				"options": [
					{ "name": "feat", "desc": "A new feature" },
					{ "name": "fix", "desc": "A bug fix" },
					{ "name": "docs", "desc": "Documentation only changes" },
					{
					  "name": "style",
					  "desc":
						"Changes that do not affect the meaning of the code(white-space, formatting, missing semi-colons, etc)"
					},
					{
					  "name": "refactor",
					  "desc": "A code change that neither fixes a bug nor adds a feature"
					},
					{
					  "name": "perf",
					  "desc": "A code change that improves performance"
					},
					{ "name": "test", "desc": "Adding missing tests" },
					{
					  "name": "chore",
					  "desc":
						"Changes to the build process or auxiliary tools and libraries such as documentation generation"
					},
					{ "name": "revert", "desc": "Revert to a commit" },
					{ "name": "WIP", "desc": "Work in progress" }
				],
				"required": true
			},
			{
				"name": "scope",
				"desc": "Scope. Could be anything specifying place of the commit change (users, db, poll):",
				"form": "input"
			},
			{
				"name": "subject",
				"desc": "Subject. Concise description of the changes. Imperative, lower case and no final dot:",
				"form": "input",
				"required": true
			},
			{
				"name": "body",
				"desc": "Body. Motivation for the change and contrast this with previous behavior:",
				"form": "multiline"
			},
			{
				"name": "footer",
				"desc": "Footer. Information about Breaking Changes and reference issues that this commit closes:",
				"form": "multiline"
			}
		],
	"template": "{{.type}}{{with .scope}}({{.}}){{end}}: {{.subject}}{{with .body}}\n\n{{.}}{{end}}{{with .footer}}\n\n{{.}}{{end}}"
}
`
