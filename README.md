# Go Todo CLI

A simple command-line todo application written in Go. Easily add, list, complete, delete, and tag your tasks from the terminal.

## Features
- Add tasks with optional descriptions and tags
- List all tasks
- Mark tasks as complete
- Delete tasks
- Tag tasks for organization
- Interactive and editor-based input for descriptions

## Usage

### Build
```
make build
```

### Run
```
./todo [command] [flags]
```

### Commands
- `add [Task Name]` : Add a new task
- `list` : List all tasks
- `complete [Task Name]` : Mark a task as complete
- `delete [Task Name]` : Delete a task
- `tag [Task Name] [Tag]` : Add a tag to a task

### Flags
- `-d, --description` : Description of the task
- `-t, --tags` : Tags to apply to the task
- `-D, --interactive-description` : Enter a description interactively

## Example
```
./todo add "Buy groceries" -d "Milk, eggs, bread" -t shopping,errands
./todo list
./todo complete "Buy groceries"
./todo delete "Buy groceries"
```

## Editor-based Description
If you want to enter a description using your editor (like Vim), use the interactive flag:
```
./todo add "Write blog post" -D
```

## Requirements
- Go 1.18+
- Cobra CLI library

## Project Structure
```
commands/   # CLI command implementations
models/     # Data models
utils/      # Utility functions
main.go     # Entry point
Makefile    # Build automation
```

## License
MIT
