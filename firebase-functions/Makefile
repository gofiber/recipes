PROJECT_ID=<ProjectID>
FUNCTION_NAME=<FunctionName>
RUNTIME=go121
ENTRY_POINT=ServerFunction

deploy:
	gcloud config set project $(PROJECT_ID)
	gcloud functions deploy $(FUNCTION_NAME) \
		--runtime $(RUNTIME) \
		--entry-point $(ENTRY_POINT) \
		--trigger-http \
		--no-gen2
