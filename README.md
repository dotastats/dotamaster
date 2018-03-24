dotastats master
===

Copy config from config/config.yaml.example to config/config.yaml

Create database `dotastats`

Copy goose config from migration/db/dbconf.yml.example to migration/db/dbconf.yml

Remember to change db username and password

Create new schema : 
 - cd dotamaster/
 - goose -path migration/db/dbconf.yml create open_dota_match_create_table sql



