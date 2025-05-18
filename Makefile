.SILENT:
db_migrate: 
	atlas schema apply -u $(DATABASE_URL) --to file://schema.sql --dev-url "docker://postgres"
