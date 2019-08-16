.PHONY: create-pipeline
create-pipeline:
	aws --profile adam-buglass-default cloudformation create-stack \
		--stack-name demonstration \
		--template-body file://./code-pipeline-stack.yaml \
		--capabilities CAPABILITY_IAM