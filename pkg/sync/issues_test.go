package sync

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/gnoverse/gh-sql/ent"
	"github.com/gnoverse/gh-sql/pkg/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_fetchIssueEventType(t *testing.T) {
	f, err := os.ReadFile("testdata/events_gnolang_gno_3597.json")
	require.NoError(t, err)
	var fi []fetchIssueEventType
	err = json.Unmarshal(f, &fi)
	assert.NoError(t, err)
	assert.Equal(t, fetchIssueEventType{
		TimelineEvent: ent.TimelineEvent{
			ID: "RRE_lADOE-u6Jc6nWXfDzwAAAAO8joIJ",
		},
		Actor: &model.SimpleUser{
			Login: "thehowl",
			ID:    4681308,
		},
	}, fetchIssueEventType{
		TimelineEvent: ent.TimelineEvent{
			ID: fi[0].ID,
		},
		Actor: fi[0].Actor,
		User:  fi[0].User,
	}, "ID should be filled, actor should be filled")
}
