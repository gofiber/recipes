---
title: I18n
keywords: [i18n, go-i18n, internationalization]
description: Internationalization support.
---

# Fiber with i18n

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/i18n) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/i18n)

This is a quick example of how to use [nicksnyder/go-i18n](https://github.com/nicksnyder/go-i18n) package to translate your Fiber application into multiple languages.

## Demo

- Run Fiber application;
- Open `http://127.0.0.1:3000/?unread=1` and see:

```bash
Hello Bob

    I have 1 unread email.
    Bob has 1 unread email.
```

- Next, go to `http://127.0.0.1:3000/?unread=4` and see pluralization of your message:

```bash
Hello Bob

    I have 4 unread emails.
    Bob has 4 unread emails.
```

- OK. Try translation of other languages, just add `&lang=es` (or `&lang=ru`) query to the URL:

```bash
Hola Bob

    Tengo 4 correos electrónicos no leídos
    Bob tiene 4 correos electrónicos no leídos
```

## go-i18n docs

- [Translating a new language](https://github.com/nicksnyder/go-i18n#translating-a-new-language);
- [Translating a new messages (updating)](https://github.com/nicksnyder/go-i18n#translating-new-messages);
