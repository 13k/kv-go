package kv_test

var (
	//nolint:lll
	binData1 = []byte(
		"\x00RP\x00" +
			"\x01status\x00#DOTA_RP_LEAGUE_MATCH_PLAYING_AS\x00" +
			"\x01steam_display\x00#DOTA_RP_LEAGUE_MATCH_PLAYING_AS\x00" +
			"\x01num_params\x003\x00" +
			"\x01lobby\x00lobby_id: 26628083760328052 lobby_state: RUN password: true game_mode: DOTA_GAMEMODE_CM member_count: 10 max_member_count: 10 name: \"Team Secret vs Vikin.GG \" lobby_type: 1\x00" +
			"\x01party\x00party_state: IN_MATCH\x00" +
			"\x01WatchableGameID\x0026628083760328052\x00" +
			"\x01param0\x00#DOTA_lobby_type_name_lobby\x00" +
			"\x01param1\x008\x00" +
			"\x01param2\x00#npc_dota_hero_grimstroke\x00" +
			"\x08",
	)

	//nolint:lll
	binData2 = []byte(
		"\x00RP\x00" +
			"\x01status\x00#DOTA_RP_PLAYING_AS\x00" +
			"\x01steam_display\x00#DOTA_RP_PLAYING_AS\x00" +
			"\x01num_params\x003\x00" +
			"\x01CustomGameMode\x000\x00" +
			"\x01WatchableGameID\x0026628083785444387\x00" +
			"\x01party\x00party_id: 26628083781803523 party_state: IN_MATCH open: false members { steam_id: 76561198054320440 }\x00" +
			"\x01param0\x00#DOTA_lobby_type_name_ranked\x00" +
			"\x01param1\x0015\x00" +
			"\x01param2\x00#npc_dota_hero_zuus\x00" +
			"\x08",
	)

	//nolint:lll
	binData3 = []byte(
		"\x00RP\x00" +
			"\x01status\x00#DOTA_RP_PLAYING_AS\x00" +
			"\x01steam_display\x00#DOTA_RP_PLAYING_AS\x00" +
			"\x01num_params\x003\x00" +
			"\x01watching_server\x00[A:1:3441750017:14553]\x00" +
			"\x01watching_from_server\x00[A:1:1671049217:14554]\x00" +
			"\x01party\x00party_state: IN_MATCH\x00" +
			"\x01WatchableGameID\x0026628083799951455\x00" +
			"\x01steam_player_group\x0026628083752106249\x00" +
			"\x01steam_player_group_size\x002\x00" +
			"\x01param0\x00#DOTA_lobby_type_name_ranked\x00" +
			"\x01param1\x002\x00" +
			"\x01param2\x00#npc_dota_hero_rubick\x00" +
			"\x08",
	)

	//nolint:lll
	binData4 = []byte(
		"\x00RP\x00" +
			"\x01status\x00#DOTA_RP_HERO_SELECTION\x00" +
			"\x01steam_display\x00#DOTA_RP_HERO_SELECTION\x00" +
			"\x01num_params\x001\x00" +
			"\x01WatchableGameID\x0026628083824762603\x00" +
			"\x01steam_player_group\x0026628083765767134\x00" +
			"\x01steam_player_group_size\x002\x00" +
			"\x01party\x00party_id: 26628083765767134 party_state: IN_MATCH open: false members { steam_id: 76561198235766844 } members { steam_id: 76561197978446698 }\x00" +
			"\x01watching_server\x00[A:1:300033030:14554]\x00" +
			"\x01watching_from_server\x00[A:1:1361739785:14554]\x00" +
			"\x01param0\x00#DOTA_lobby_type_name_ranked\x00" +
			"\x08",
	)
)
