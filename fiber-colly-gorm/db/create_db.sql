SELECT 'CREATE DATABASE quotes' 
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'quotes')\gexec