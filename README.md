
# Goline

Goline is a very simple notification CLI app writng with Go.
## Feature
- CRUD operation to Create, Delete, Update and Delete Task from embedded database (SQlite)
- Set time to send notification to line chat (Get token from https://notify-bot.line.me/th/)
## Installation
1. Clone Repository with this command
```
git clone github.com/tigertony2536/goline-v2.git
```
2. enter to project's root directory and install project
```
go install .\app\goline
go install .\app\notiapp
```
## Command
###  create 
```
goline get [-h |--help] [-a |--all] [-i |--id <taskID>] [-n |--name <taskname>] [-d |--date <date>] [-n |--name <>]
```
###  get 
```
goline get [-h |--help] [-a |--all] [-i |--id <taskID>] [-n |--name <taskname>] [-d |--date <date>] [-n |--name <>]
```
###  update 
```
goline get [-h |--help] [-a |--all] [-i |--id <taskID>] [-n |--name <taskname>] [-d |--date <date>] [-n |--name <>]
```
###  delete 
```
goline get [-h |--help] [-a |--all] [-i |--id <taskID>] [-n |--name <taskname>] [-d |--date <date>] [-n |--name <>]
```
### notify 
Start notificaition app in the background
```
goline notify start 
```
Stop notificaition app in the background

```
goline notify stop 
```
Check status of notificaition app

```
goline notify status 
```
