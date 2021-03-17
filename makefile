examples/basic-vpc/init:
	@cd ./examples/basic-vpc ; \
		terraform init

examples/basic-vpc/plan: examples/basic-vpc/init
	@cd ./examples/basic-vpc ; \
		terraform plan

examples/basic-vpc/apply: examples/basic-vpc/plan
	@cd ./examples/basic-vpc ; \
		terraform apply \
			-auto-approve

examples/basic-vpc/destroy:
	@cd ./examples/basic-vpc ; \
		terraform destroy \
			-auto-approve

test/basic-vpc/integration-full: examples/basic-vpc/apply
	@cd ./test/basic-vpc ; \
		go test -v 
	@make examples/basic-vpc/destroy