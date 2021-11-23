REMOTE=117.50.36.69
APPNAME=goblog

deploy:
	@echo "\n--- 开始构建可执行文件 ---"
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -v -o tmp/$(APPNAME)_tmp
	
	@echo "\n--- 上传可执行文件 ---"
	scp tmp/$(APPNAME)_tmp root@$(REMOTE):/data/www/goblog.com/

	@echo "\n--- 停止服务 ---"
	ssh root@$(REMOTE) "supervisorctl stop $(APPNAME)"

	@echo "\n--- 替换新文件 ---"
	ssh root@$(REMOTE) "cd /data/www/goblog.com/ \
							&& rm $(APPNAME) \
							&& mv $(APPNAME)_tmp $(APPNAME) \
							&& chown www-data:www-data $(APPNAME)"

	@echo "\n--- 开始服务 ---"
	ssh root@$(REMOTE) "supervisorctl start $(APPNAME)"

	@echo "\n--- 部署完毕 ---\n"

.PHONY: clean
clean:
	rm -rf tmp/*

nginx-test:
	ssh root@$(REMOTE) "nginx -t"

task1: task2 task3
	@echo "\n--- 执行命令 ---\n"
	@$(MAKE) task3
	@echo "\n--- 执行更多命令 ---\n"
	@$(MAKE) task4

task2:  
	@echo "\n--- task2 ---\n"

task3:  
	@echo "\n--- task3 ---\n"

task4:  
	@echo "\n--- task4 ---\n"
