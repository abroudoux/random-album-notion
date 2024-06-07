# random-album-notion

![Repo Size](https://img.shields.io/github/repo-size/abroudoux/random-album-notion)
![License](https://img.shields.io/badge/license-MIT-red)
![Go Version](https://img.shields.io/github/go-mod/go-version/abroudoux/random-album-notion)

## 💻・About

From a Notion todo list, retrieve the albums not listened to, choose randomly and play it with a single command.

## 🎯・Setup

To run the project execute `main.go`

```bash
go run cmd/main.go
```

## 📚・Ressources

- [Notion API](https://developers.notion.com/)
- [Notion API SDK](https://github.com/jomei/notionapi/tree/74e249c47bb5634b670b8b45cad2be7fd56ddc98/)
- [Shpotify](https://github.com/jomei/notionapi/tree/74e249c47bb5634b670b8b45cad2be7fd56ddc98)

## 🧑‍🤝‍🧑・Contributing

Feel free to contribute to the project !

## 🎯・Roadmap

- [ ] Improve perfs
- [ ] Keep in memory last modifications to not retrieve same data
- [ ] Rewrite shpotify library to start album at the first song and improve perfs
- [ ] Improve search accuracy

## 📑・Licence

This project is under MIT license. For more information, please see the file [LICENSE](./LICENSE).
