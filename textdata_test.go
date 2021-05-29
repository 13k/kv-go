package kv_test

var (
	//nolint:lll
	textData1 = []byte(
		`"RP" {
  "status" "#DOTA_RP_LEAGUE_MATCH_PLAYING_AS"
  "steam_display" "#DOTA_RP_LEAGUE_MATCH_PLAYING_AS"
  "num_params" "3"
  "lobby" "lobby_id: 26628083760328052 lobby_state: RUN password: true game_mode: DOTA_GAMEMODE_CM member_count: 10 max_member_count: 10 name: \"Team Secret vs Vikin.GG \" lobby_type: 1"
  "party" "party_state: IN_MATCH"
  "WatchableGameID" "26628083760328052"
  "param0" "#DOTA_lobby_type_name_lobby"
  "param1" "8"
  "param2" "#npc_dota_hero_grimstroke"
}
`,
	)

	//nolint:lll
	textData2 = []byte(
		`"RP" {
  "status" "#DOTA_RP_PLAYING_AS"
  "steam_display" "#DOTA_RP_PLAYING_AS"
  "num_params" "3"
  "CustomGameMode" "0"
  "WatchableGameID" "26628083785444387"
  "party" "party_id: 26628083781803523 party_state: IN_MATCH open: false members { steam_id: 76561198054320440 }"
  "param0" "#DOTA_lobby_type_name_ranked"
  "param1" "15"
  "param2" "#npc_dota_hero_zuus"
}
`,
	)

	//nolint:lll
	textData3 = []byte(
		`"RP" {
  "status" "#DOTA_RP_PLAYING_AS"
  "steam_display" "#DOTA_RP_PLAYING_AS"
  "num_params" "3"
  "watching_server" "[A:1:3441750017:14553]"
  "watching_from_server" "[A:1:1671049217:14554]"
  "party" "party_state: IN_MATCH"
  "WatchableGameID" "26628083799951455"
  "steam_player_group" "26628083752106249"
  "steam_player_group_size" "2"
  "param0" "#DOTA_lobby_type_name_ranked"
  "param1" "2"
  "param2" "#npc_dota_hero_rubick"
}
`,
	)

	//nolint:lll
	textData4 = []byte(
		`"RP" {
  "status" "#DOTA_RP_HERO_SELECTION"
  "steam_display" "#DOTA_RP_HERO_SELECTION"
  "num_params" "1"
  "WatchableGameID" "26628083824762603"
  "steam_player_group" "26628083765767134"
  "steam_player_group_size" "2"
  "party" "party_id: 26628083765767134 party_state: IN_MATCH open: false members { steam_id: 76561198235766844 } members { steam_id: 76561197978446698 }"
  "watching_server" "[A:1:300033030:14554]"
  "watching_from_server" "[A:1:1361739785:14554]"
  "param0" "#DOTA_lobby_type_name_ranked"
}
`,
	)
)
