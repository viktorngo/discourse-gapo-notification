deploy-vm:
	@echo "Deploying to production..."
	go build -v -o server
	systemctl stop discourse-noti.service
	systemctl start discourse-noti.service
	@echo "Deployed to production"