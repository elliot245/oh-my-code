# Workspace Repositories Recovery Checklist

This file records all Git repositories under this workspace to help restore the same setup on another VM.

## Root
- `/home/elliot245/workspace/elliot245/oh-my-code/workspace`

## Repositories

| Path | Current Branch | Origin | Dirty Working Tree |
|---|---|---|---|
| `/home/elliot245/workspace/elliot245/oh-my-code/workspace/agent-manager-skill` | `chore/remove-pyyaml-dep` | `git@github.com:fractalmind-ai/agent-manager-skill.git` | `no` |
| `/home/elliot245/workspace/elliot245/oh-my-code/workspace/clawdbot-feishu` | `main` | `https://github.com/m1heng/clawdbot-feishu.git` | `no` |
| `/home/elliot245/workspace/elliot245/oh-my-code/workspace/clawdbot` | `main` | `git@github.com:elliot245/clawdbot.git` | `no` |
| `/home/elliot245/workspace/elliot245/oh-my-code/workspace/CloudBank/contracts` | `detached HEAD` | `N/A` | `yes` |
| `/home/elliot245/workspace/elliot245/oh-my-code/workspace/cloudbank_main` | `feat/issue-607-reserved-commission` | `git@github.com:SuLabsOrg/CloudBank.git` | `no` |
| `/home/elliot245/workspace/elliot245/oh-my-code/workspace/fractalbot` | `feat/issue-1-telegram-routing` | `git@github.com:fractalmind-ai/fractalbot.git` | `yes` |
| `/home/elliot245/workspace/elliot245/oh-my-code/workspace/fractalbot_main` | `issue-268-slack-app-mention` | `git@github.com:fractalmind-ai/fractalbot.git` | `no` |
| `/home/elliot245/workspace/elliot245/oh-my-code/workspace/openclaw` | `main` | `https://github.com/openclaw/openclaw.git` | `no` |
| `/home/elliot245/workspace/elliot245/oh-my-code/workspace/polymarket-quant` | `main` | `git@github.com:elliot245/polymarket-quant.git` | `no` |
| `/home/elliot245/workspace/elliot245/oh-my-code/workspace/typemind-android` | `main` | `git@github.com:fractalmind-ai/typemind-android.git` | `no` |

## Notes
- `CloudBank/contracts` is currently in detached HEAD state and has local changes.
- `fractalbot` has local uncommitted changes.
- Before migration, consider exporting patches for dirty repos:
  - `git -C <repo> diff > <repo-name>.patch`
  - `git -C <repo> status --short`
