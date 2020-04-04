package emails

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/pkg/errors"
)

type EmailTemplate struct {
	body    string
	subject string
}

func NewCandidateStepOneTemplate(candidateEmail string) (EmailTemplate, error) {
	data := struct {
		Email string
	}{Email: candidateEmail}

	tmplS, err := parseBody(candidateStepOneBody, data)
	if err != nil {
		return EmailTemplate{}, errors.Wrap(err, "while creating candidate step one")
	}

	return EmailTemplate{
		body:    tmplS,
		subject: "Recruitment process",
	}, nil
}

func NewJudgeDecisionTemplate(judgeEmail, acceptEndpoint, denyEndpoint string) (EmailTemplate, error) {
	data := struct {
		Email          string
		AcceptEndpoint string
		DenyEndpoint   string
	}{Email: judgeEmail, AcceptEndpoint: acceptEndpoint, DenyEndpoint: denyEndpoint}

	tmplS, err := parseBody(judgeDecisionBody, data)
	if err != nil {
		return EmailTemplate{}, errors.Wrap(err, "while creating judge decision")
	}

	return EmailTemplate{
		body:    tmplS,
		subject: "Recruitment process",
	}, nil
}

func NewApproveOutcomeTemplate(candidateEmail string) (EmailTemplate, error) {
	data := struct {
		Email string
	}{Email: candidateEmail}

	tmplS, err := parseBody(approveOutcomeBody, data)
	if err != nil {
		return EmailTemplate{}, errors.Wrap(err, "while creating outcome approved")
	}

	return EmailTemplate{
		body:    tmplS,
		subject: "You have been approved",
	}, nil
}

func NewDenyOutcomeTemplate(candidateEmail string) (EmailTemplate, error) {
	data := struct {
		Email string
	}{Email: candidateEmail}

	tmplS, err := parseBody(deniedOutcomeBody, data)
	if err != nil {
		return EmailTemplate{}, errors.Wrap(err, "while creating outcome denied")
	}

	return EmailTemplate{
		body:    tmplS,
		subject: "You have been denied",
	}, nil
}

func parseBody(body string, data interface{}) (string, error) {
	t, err := template.New("template").Parse(body)
	if err != nil {
		return "", errors.Wrap(err, fmt.Sprintf("problem while creating template"))
	}

	var tpl bytes.Buffer
	err = t.Execute(&tpl, data)
	if err != nil {
		return "", errors.Wrap(err, fmt.Sprintf("problem while executing template"))
	}

	return tpl.String(), nil
}
