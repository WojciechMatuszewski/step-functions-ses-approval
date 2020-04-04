package emails

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCandidateStepOneTemplate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		tmpl, err := NewCandidateStepOneTemplate("foo@fo.com")

		assert.NoError(t, err)
		assert.True(t, strings.Contains(tmpl.body, "foo@fo.com"))
	})
}

func TestNewJudgeDecisionTemplate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		tmpl, err := NewJudgeDecisionTemplate("cand@cand.com", "https://accept.com", "https://deny.com")

		assert.NoError(t, err)
		assert.True(t, strings.Contains(tmpl.body, "cand@cand.com"))
		assert.True(t, strings.Contains(tmpl.body, "https://accept.com"))
		assert.True(t, strings.Contains(tmpl.body, "https://deny.com"))
	})
}

func TestNewDenyOutcomeTemplate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		tmpl, err := NewDenyOutcomeTemplate("cand@cand.com")

		assert.NoError(t, err)
		assert.True(t, strings.Contains(tmpl.body, "cand@cand.com"))
		assert.True(t, strings.Contains(tmpl.subject, "denied"))
	})
}

func TestNewApproveOutcomeTemplate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		tmpl, err := NewApproveOutcomeTemplate("cand@cand.com")

		assert.NoError(t, err)
		assert.True(t, strings.Contains(tmpl.body, "cand@cand.com"))
		assert.True(t, strings.Contains(tmpl.subject, "approved"))
	})
}
