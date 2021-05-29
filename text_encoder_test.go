package kv_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/13k/kv-go"
)

func TestTextEncoder(t *testing.T) {
	suite.Run(t, &TextEncoderSuite{})
}

type TextEncoderSuite struct {
	Suite
}

func (s *TextEncoderSuite) TestEncode() {
	require := s.Require()

	//nolint:lll
	testCases := []struct {
		Subject  kv.KeyValue
		Expected []byte
		Err      string
	}{
		{
			Subject:  kv.NewKeyValue(kv.TypeEnd, "", "", nil),
			Expected: nil,
			Err:      "kv: cannot encode nodes of type End",
		},
		{
			Subject:  kv.NewKeyValue(kv.TypeInvalid, "", "", nil),
			Expected: nil,
			Err:      "kv: cannot encode nodes of type Invalid",
		},
		{
			Subject:  kv.NewKeyValueString("K", "S", nil),
			Expected: []byte(`"K" "S"`),
		},
		{
			Subject:  kv.NewKeyValueInt32("K", "1", nil),
			Expected: []byte(`"K" "1"`),
		},
		{
			Subject:  kv.NewKeyValuePointer("K", "1", nil),
			Expected: []byte(`"K" "1"`),
		},
		{
			Subject:  kv.NewKeyValueColor("K", "1", nil),
			Expected: []byte(`"K" "1"`),
		},
		{
			Subject:  kv.NewKeyValueInt64("K", "1", nil),
			Expected: []byte(`"K" "1"`),
		},
		{
			Subject:  kv.NewKeyValueUint64("K", "1", nil),
			Expected: []byte(`"K" "1"`),
		},
		{
			Subject:  kv.NewKeyValueFloat32("K", "1.23", nil),
			Expected: []byte(`"K" "1.23"`),
		},
		{
			Subject: kv.NewKeyValueRoot("K").AddString("s", "S"),
			Expected: []byte(
				`"K" {
  "s" "S"
}
`,
			),
		},
		{
			Subject: kv.NewKeyValueRoot("RP").
				AddString("status", "#DOTA_RP_LEAGUE_MATCH_PLAYING_AS").
				AddString("steam_display", "#DOTA_RP_LEAGUE_MATCH_PLAYING_AS").
				AddString("num_params", "3").
				AddString("lobby", "lobby_id: 26628083760328052 lobby_state: RUN password: true game_mode: DOTA_GAMEMODE_CM member_count: 10 max_member_count: 10 name: \"Team Secret vs Vikin.GG \" lobby_type: 1").
				AddString("party", "party_state: IN_MATCH").
				AddString("WatchableGameID", "26628083760328052").
				AddString("param0", "#DOTA_lobby_type_name_lobby").
				AddString("param1", "8").
				AddString("param2", "#npc_dota_hero_grimstroke"),
			Expected: textData1,
		},
		{
			Subject: kv.NewKeyValueRoot("RP").
				AddString("status", "#DOTA_RP_PLAYING_AS").
				AddString("steam_display", "#DOTA_RP_PLAYING_AS").
				AddString("num_params", "3").
				AddString("CustomGameMode", "0").
				AddString("WatchableGameID", "26628083785444387").
				AddString("party", "party_id: 26628083781803523 party_state: IN_MATCH open: false members { steam_id: 76561198054320440 }").
				AddString("param0", "#DOTA_lobby_type_name_ranked").
				AddString("param1", "15").
				AddString("param2", "#npc_dota_hero_zuus"),
			Expected: textData2,
		},
		{
			Subject: kv.NewKeyValueRoot("RP").
				AddString("status", "#DOTA_RP_PLAYING_AS").
				AddString("steam_display", "#DOTA_RP_PLAYING_AS").
				AddString("num_params", "3").
				AddString("watching_server", "[A:1:3441750017:14553]").
				AddString("watching_from_server", "[A:1:1671049217:14554]").
				AddString("party", "party_state: IN_MATCH").
				AddString("WatchableGameID", "26628083799951455").
				AddString("steam_player_group", "26628083752106249").
				AddString("steam_player_group_size", "2").
				AddString("param0", "#DOTA_lobby_type_name_ranked").
				AddString("param1", "2").
				AddString("param2", "#npc_dota_hero_rubick"),
			Expected: textData3,
		},
		{
			Subject: kv.NewKeyValueRoot("RP").
				AddString("status", "#DOTA_RP_HERO_SELECTION").
				AddString("steam_display", "#DOTA_RP_HERO_SELECTION").
				AddString("num_params", "1").
				AddString("WatchableGameID", "26628083824762603").
				AddString("steam_player_group", "26628083765767134").
				AddString("steam_player_group_size", "2").
				AddString("party", "party_id: 26628083765767134 party_state: IN_MATCH open: false members { steam_id: 76561198235766844 } members { steam_id: 76561197978446698 }").
				AddString("watching_server", "[A:1:300033030:14554]").
				AddString("watching_from_server", "[A:1:1361739785:14554]").
				AddString("param0", "#DOTA_lobby_type_name_ranked"),
			Expected: textData4,
		},
	}

	for testCaseIdx, testCase := range testCases {
		b := &bytes.Buffer{}
		enc := kv.NewTextEncoder(b)
		err := enc.Encode(testCase.Subject)
		actual := b.Bytes()

		if testCase.Err == "" {
			require.NoErrorf(err, "test case %d", testCaseIdx)
		} else {
			require.EqualErrorf(err, testCase.Err, "test case %d", testCaseIdx)
		}

		expected := testCase.Expected

		if expected == nil {
			require.Nilf(actual, "test case %d", testCaseIdx)
		} else {
			require.NotNilf(actual, "test case %d", testCaseIdx)
		}

		require.Equalf(expected, actual, "test case %d", testCaseIdx)
	}
}
