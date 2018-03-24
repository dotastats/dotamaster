
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE vp_matches (
	id SERIAL PRIMARY KEY,
	match_id INTEGER,
	series_id INTEGER,
	team_a_id TEXT,
	team_b_id TEXT,
	team_a TEXT,
	team_b TEXT,
	logo_a TEXT,
	logo_b TEXT,
	team_a_short TEXT,
	team_b_short TEXT,
	tournament TEXT,
	tournament_logo TEXT,
	game TEXT,
	best_of TEXT,
	vp_match_id TEXT,
	url TEXT,
	time TIMESTAMP WITH TIME ZONE,
	match_name TEXT,
	mode_name TEXT,
	mode_desc TEXT,
	handicap_amount TEXT,
	handicap_team TEXT,
	ratio_a FLOAT,
	ratio_b FLOAT,
	winner TEXT,
	status TEXT,
	score_a FLOAT,
	score_b FLOAT,
	note TEXT,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE vp_matches;
