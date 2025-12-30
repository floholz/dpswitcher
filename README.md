# dpswitch

`dpswitch` is a lightweight system tray application written in Go that allows you to quickly toggle your monitors on and off. It is particularly useful for multi-monitor setups where you frequently need to enable or disable specific displays without diving into system settings.

![Logo](assets/logo.svg)

## Features

- **System Tray Integration**: Resides in your system tray for quick access.
- **Real-time Status**: Displays the current state (enabled/disabled) of your connected monitors.
- **Quick Toggle**: Enable or disable displays with a single click.
- **KDE Plasma Support**: Currently supports KDE Plasma environments using `kscreen-doctor`.
- **Automatic Updates**: Periodically polls for display changes to keep the menu up to date.

## Prerequisites

- **Go**: Version 1.16 or higher is recommended.
- **KDE Plasma**: The application currently targets KDE environments.
- **kscreen-doctor**: Ensure `kscreen-doctor` is installed (usually part of KDE Plasma's display management).
- **Fyne Dependencies**: Since it uses the Fyne toolkit, you may need certain system libraries installed (e.g., `libgl1-mesa-dev`, `xorg-dev` on Linux). See [Fyne's Getting Started](https://developer.fyne.io/started/) for details.

## Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/yourusername/dpswitch.git
   cd dpswitch
   ```

2. **Build the application**:
   ```bash
   go build -o dpswitch main.go
   ```

3. **Run the application**:
   ```bash
   ./dpswitch
   ```

## Usage

Once launched, an icon will appear in your system tray. 
- Click the tray icon to see a list of connected monitors.
- A checkmark indicates the monitor is currently enabled.
- Click on a monitor's name to toggle its state.
- Note: The primary monitor is usually protected from being disabled via the menu to prevent accidental loss of all displays.

## Development

The project structure is organized as follows:
- `main.go`: Entry point and tray initialization.
- `cmd/`: Core logic for menu handling and display management.
- `cmd/display-tools/`: Interfaces and implementations for different display configuration tools (currently supporting `kscreen-doctor`).
- `assets/`: Icons and logos used by the application.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details (if available).
