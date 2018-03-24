dotastats master
===

# Set up project

1. Copy config from config/config.yaml.example to config/config.yaml

2. Create database `dotastats`

3. Copy goose config from migration/db/dbconf.yml.example to migration/db/dbconf.yml

4. Remember to change db username and password

5. Create new schema: 
 - `cd dotamaster/`
 - `goose -path migration/db/ create open_dota_match_create_table sql`



