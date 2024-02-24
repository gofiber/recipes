SELECT 'CREATE DATABASE colly' 
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'colly')\gexec