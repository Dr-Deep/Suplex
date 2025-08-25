package command

/*
* als Suplex mit "typing..."
```
2. Multi-User-Session-Management
Um ein LLM so zu benutzen, dass es verschiedene User „getrennt“ behandeln kann, braucht man Sitzungen:
Pro Discord-User ein eigener Kontext (History, Persona, Memory).
Speicherung in DB
Bot erkennt den User über Message.Author.ID und holt den zugehörigen Kontext.
So können mehrere Leute gleichzeitig mit dem gleichen LLM über denselben Bot reden,
ohne dass die Kontexte sich vermischen.

In Thread oder command
```
https://github.com/go-deepseek/deepseek
*/
