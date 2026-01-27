# Agent Manager (agent-manager)

Employee agent orchestration system for managing AI agents in tmux sessions. Use when working with agents/ directory employee agents - starting, stopping, monitoring, or assigning tasks to Dev/QA agents running in tmux sessions. Completely independent of CAO, uses only tmux + Python.

## Overview

Simple, installation-agnostic agent lifecycle management.

## Installation

```bash
# Clone from oh-my-code workspace
cd /home/elliot245/workspace/elliot245/oh-my-code
git clone https://github.com/fractalmind-ai/agent-manager-skill.git

# Go to fractalbot directory
cd /home/elliot245/workspace/elliot245/oh-my-code/fractalbot
```

## Usage

After installation:

```bash
# List all agents
python3 .claude/skills/agent-manager/scripts/main.py list
```

## License

MIT
