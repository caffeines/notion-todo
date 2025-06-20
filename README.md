# Notion Todo CLI

A command-line interface for managing todo items using Notion database integration. This Go-based CLI tool allows you to add and manage todo items directly from your terminal to your Notion workspace.

## Installation

### Linux (Recommended)

```bash
# Quick install script
curl -sSL https://raw.githubusercontent.com/caffeines/notion-todo/main/install.sh | bash
```

Or download manually from the [releases page](https://github.com/caffeines/notion-todo/releases).

### macOS (Homebrew)

```bash
# Add the tap
brew tap caffeines/tap

# Install notion-todo (available as 'todo' command)
brew install notion-todo

# Now you can use it with:
todo guide
```

### Using Go Install

```bash
go install github.com/caffeines/notion-todo@latest
```

### Download Binary

Download the latest release for your platform from the [releases page](https://github.com/caffeines/notion-todo/releases).

### Build from Source

1. Clone the repository:

   ```bash
   git clone https://github.com/caffeines/notion-todo.git
   cd notion-todo
   ```

2. Build the project:

   ```bash
   make build
   # or
   go build -o todo main.go
   ```

3. Move the binary to your PATH:

   ```bash
   sudo mv todo /usr/local/bin/
   ```

## Quick Start

1. **Build the project:**

   ```bash
   make build
   # or
   go build -o todo main.go
   ```

2. **Run the interactive setup guide:**

   ```bash
   todo guide
   # Or use the short alias
   todo g
   ```

3. **Configure the CLI:**

   ```bash
   todo config
   # Or use the short alias
   todo c
   ```

4. **Start adding todos:**

   ```bash
   todo add "Buy groceries"
   # Or use the short alias
   todo a "Complete project" --date 25-06-2025
   ```

5. **View and manage your todos:**

   ```bash
   todo list
   # Or use short aliases
   todo l
   # or
   todo ls
   ```

That's it! The interactive guide will walk you through everything else.

## Features

- 💡 **Interactive Setup Guide** - Step-by-step walkthrough for first-time users
- 🔧 Easy configuration setup with Notion API token and database ID
- ➕ Add todo items to your Notion database with optional due dates
- 📋 Interactive list view for managing todos
- 📝 Simple and intuitive command-line interface with short aliases for faster usage
- 🔒 Secure credential storage
- 🎯 Direct integration with Notion API
- 📊 Status tracking (Not started, In progress, Done)
- 📅 Due date support for better task management
- ⚡ Quick commands with short aliases (`todo v`, `todo a`, `todo l`, etc.)

## Prerequisites

- Go 1.19 or higher
- A Notion account with API access
- A Notion database set up for todo items

## Configuration

Before using the CLI, you need to configure it with your Notion API credentials. You have two options:

### Option 1: Interactive Setup Guide (Recommended)

Run the interactive setup guide for a step-by-step walkthrough:

```bash
todo guide
```

This interactive guide will walk you through:

- Setting up your todo database (with template option)
- Getting your database ID
- Creating a Notion integration
- Getting your API credentials
- Connecting the integration to your database
- Configuring the CLI
- Testing your setup

**Navigation in the Guide:**

- **Next step**: `Enter`, `Space`, `n`, `l`, or `→` (right arrow)
- **Previous step**: `p`, `h`, or `←` (left arrow)
- **Restart**: `r`
- **Quit**: `q` or `Ctrl+C`

### Option 2: Manual Configuration

If you prefer manual setup or already have your credentials:

```bash
todo config
```

This will prompt you to enter:

- **Notion API Token**: Your Notion integration token
- **Database ID**: The ID of your Notion database

### Getting Notion Credentials

#### Quick Database Setup (Recommended)

Use the pre-configured template for the fastest setup:

1. **Get the Template**:
   - Go to [Notion Todo Template](https://www.notion.so/templates/cli-todo)
   - Click "Get template" and select your workspace
   - The template already has the correct database structure

2. **Get Database ID**:
   - Open your database in Notion
   - Copy the database ID from the URL: `https://www.notion.so/{workspace}/{database_id}?v=...`
   - The database ID is the string after the last `/` and before `?`
   - Example: `217e31359630803999d6ecaabdf4e11f`

#### Manual Database Setup

If you prefer to create the database manually:

**Required Properties:**

- **Title** (Title) - The main todo text
- **Status** (Select) - Todo completion status with options:
  - `Todo` (default)
  - `In progress`
  - `Done`
- **Due Date** (Date) - Optional due date for todos
- **Tags** (Multi-select) - Optional categorizing tags

**Setup Instructions:**

1. **Create Integration**:
   - Go to [Notion Integrations](https://www.notion.so/profile/integrations)
   - Click "New integration"
   - Name it "Todo CLI" or similar
   - Select your workspace and choose "Internal" type
   - Click "Save"

2. **Grant Database Access**:
   - In the integration configuration popup, click "Access" tab
   - Select "pages" to see your database
   - Select your todo database from the list
   - Click "Update Access"
   - Go back to "Configuration" tab

3. **Copy Integration Token**:
   - Click "Show" next to "Internal Integration Token"
   - Copy the token and keep it secure

## Usage

### Add a Todo Item

```bash
# Basic todo
todo add "Buy groceries"

# Todo with due date
todo add "Complete project documentation" --date 25-06-2025

# Todo with specific date format
todo add "Schedule dentist appointment" --date 2025-06-30
```

The todo items will be created in your Notion database with:

- **Title**: Your todo text
- **Status**: Set to "Todo" by default  
- **Due Date**: Due date if specified

### List and Manage Todos

```bash
# View all todos in an interactive interface
todo list
```

The list command provides an interactive interface where you can:

- View todos by status (Todo, In progress, Done)
- Navigate through your todos
- See due dates and completion status
- Manage your todo items efficiently

### Available Commands

- `todo guide` (or `todo g`) - **Interactive setup guide** for first-time users (recommended)
- `todo config` (or `todo c`) - Configure Notion API credentials manually
- `todo add <todo-text>` (or `todo a`) - Add a new todo item
- `todo add <todo-text> --date YYYY-MM-DD` - Add todo with due date
- `todo list` (or `todo l`, `todo ls`) - View and manage existing todos in interactive mode
- `todo version` (or `todo v`) - Show version information
- `todo help` - Show help information

#### Short Command Aliases

For faster usage, you can use these short aliases:

- `todo v` → `todo version`
- `todo a` → `todo add`
- `todo l` → `todo list`
- `todo ls` → `todo list`
- `todo g` → `todo guide`
- `todo c` → `todo config`

### Command Examples

```bash
# First-time setup
todo guide
# Or use the short alias
todo g

# Add todos
todo add "Buy milk"
todo a "Finish presentation" --date 2025-06-25

# View and manage
todo list
# Or use short aliases
todo l
todo ls

# Check version
todo version
todo v

# Get help
todo help
```

## Project Structure

```
├── cmd/                    # CLI commands
│   ├── add.go             # Add todo command
│   ├── config.go          # Configuration command  
│   ├── root.go            # Root command
│   └── version.go         # Version command
├── models/                # Data models
│   ├── config.go          # Configuration model
│   ├── createTodoPayload.go # Notion API payload
│   └── todoItem.go        # Todo item structure
├── notion/                # Notion API integration
│   ├── notion.go          # Interface
│   └── notionImpl.go      # Implementation
├── service/               # Business logic services
│   ├── credential.go      # Credential interface
│   ├── credentialImpl.go  # Credential implementation
│   ├── file.go           # File interface
│   ├── FileImpl.go       # File implementation
│   ├── todo.go           # Todo interface
│   └── todoImpl.go       # Todo implementation
└── main.go               # Application entry point
```

## Configuration Storage

The application stores your configuration in a hidden directory:
- Path: `~/.notion-todo/config.json`
- Contains encrypted credentials for secure storage

## Troubleshooting

### Common Issues

**"Failed to create todo" or "Database not found"**

- Make sure your integration is connected to the database (step 5 in the guide)
- Verify your Database ID is correct
- Check that your API token is valid

**"Property not found" errors**

- Ensure your database has the exact properties: "Title", "Status", "Due Date"
- Property names are case-sensitive  
- The Status property must be of type "Select" with options: "Todo", "In progress", "Done"
- Consider using the template for correct setup

**Configuration issues**

- Run `todo config` to reconfigure your credentials
- Check if `~/.notion-todo/config.json` exists and has valid JSON
- Re-run the setup guide: `todo guide`

**Need help?**

- Run `todo guide` for the interactive setup
- Use `todo help` for command information
- Check that your Notion integration has the correct permissions

## Architecture

The project follows a clean architecture pattern with:

- **Commands** ([`cmd/`](cmd/)): CLI command handlers using Cobra
- **Services** ([`service/`](service/)): Business logic with interfaces and implementations
- **Models** ([`models/`](models/)): Data structures for API communication
- **Notion Integration** ([`notion/`](notion/)): Notion API client implementation

Key interfaces:

- [`service.Todo`](service/todo.go): Todo operations
- [`service.Credential`](service/credential.go): Credential management
- [`service.File`](service/file.go): File operations
- [`notion.Notion`](notion/notion.go): Notion API operations

## Development

### Building

```bash
# Build the project
go build -o todo main.go

# Or use the Makefile
make build
```

### Testing

```bash
go test ./...
```

## Dependencies

- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - Terminal UI framework for interactive guide
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) - Style and layout for terminal interfaces
- [PromptUI](https://github.com/manifoldco/promptui) - Interactive prompts

## License

See [LICENSE](LICENSE) file for details.

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request