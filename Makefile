BINARY_NAME=dpswitch
INSTALL_DIR=$(HOME)/.local/bin
SERVICE_DIR=$(HOME)/.config/systemd/user
SERVICE_NAME=dpswitch.service

.PHONY: all build install uninstall clean

all: build

build:
	go build -o $(BINARY_NAME) main.go

install: build
	mkdir -p $(INSTALL_DIR)
	cp $(BINARY_NAME) $(INSTALL_DIR)/$(BINARY_NAME)
	mkdir -p $(SERVICE_DIR)
	cp $(SERVICE_NAME) $(SERVICE_DIR)/$(SERVICE_NAME)
	systemctl --user daemon-reload
	systemctl --user enable $(SERVICE_NAME)
	@echo "Installation complete. Start the service with: systemctl --user start $(SERVICE_NAME)"

uninstall:
	systemctl --user disable $(SERVICE_NAME) || true
	rm -f $(SERVICE_DIR)/$(SERVICE_NAME)
	systemctl --user daemon-reload
	rm -f $(INSTALL_DIR)/$(BINARY_NAME)
	@echo "Uninstallation complete."

clean:
	rm -f $(BINARY_NAME)
