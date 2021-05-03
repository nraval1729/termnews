# termnews

`termnews` is a small utility program that lets you read news in your terminal. It uses [NewsAPI](https://newsapi.org/) to fetch the news, and as a
 result, provides options for you to customize your news reading experience.
 For example, you can tell `termnews` to only fetch (say) `sports`, `business` and `science` related news from (say) `India`.
 It also keeps the news data updated by periodically fetching the latest news in the background to ensure that you always get the latest news.
 Finally, it lets you open the articles in your browser by using keyboard shortcuts (defined below)
 
 ![termnews.gif](https://imgur.com/7mjzvoM.gif)

# Features
 - Customizable news sources, categories and region
 - Fully keyboard based navigation
 - In built pagination
 - Ability to open articles in your browser for a detailed read
 - Periodic background updates, so you always see the latest and greatest news
 - Works on multiple terminal emulators (tested on iterm, alacritty, Mac terminal)
  
 
# Getting started
- Get your API key from [NewsAPI](https://newsapi.org/)
- Create `config.yml` at `~/.termnews/config.yml` (read below to understand more, and to see examples of a few `config.yml` files)
- Install [Go](https://golang.org/) (You can skip this if you already have it)
- `go get github.com/nraval1729/termnews`
- Run `termnews` in your terminal and enjoy!


# `config.yml`
`termnews` expects `config.yml` present at `~/.termnews/config.yml` in order for it to function. This is where you can customize the behaviour of
 `termnews`. These are the configuration options currently supported:
 - `apiKey` (**required**) - The API key you'll fetch from [NewsAPI](https://newsapi.org).
 They have multiple plan offerings, and while I use the free plan, you can go ahead and choose whichever plan you'd like.
 If you _are_ using the free plan, please ensure that you set the `refreshFrequency` correctly as a courtesy to the NewsAPI team.
 You can read more about the API rate limits [here](https://newsapi.org/pricing).
 - `sources` - This is a list of sources that `termnews` will use to fetch the news.
 You can read more about the supported sources [here](https://newsapi.org/docs/endpoints/sources).
 - `categories` - This is a list of categories that `termnews` will use to fetch the news.
 The currently supported categories are: `business`, `entertainment`, `general`, `health`, `science`, `sports`, and `technology`.
 - `country` - This is a country code that will ensure `termnews` only fetches news associated with that country.
 The currently supported country codes are: `ae`, `ar`, `at`, `au`, `be`, `bg`, `br`, `ca`, `ch`, `cn`, `co`, `cu`, `cz`, `de`, `eg`, `fr`, `gb
 `, `gr`, `hk`, `hu`, `id`, `ie`, `il`, `in`, `it`, `jp`, `kr`, `lt`, `lv`, `ma`, `mx`, `my`, `ng`, `nl`, `no`, `nz`, `ph`, `pl`, `pt`, `ro`, `rs`, `ru`, `sa`, `se`, `sg`, `si`, `sk`, `th`, `tr`, `tw`, `ua`, `us`, `ve`, `za`
 - `refreshFrequency` (**optional**) - This tells `termnews` about the interval to use for fetching the news. Defaults to `15` minutes.
 For example, a `refreshFrequency` of `10` tells `termnews` to fetch news every `10` minutes.
 
 *Note*: NewsAPI requires that `sources` not be used when either `country` or `categories` are used.
 
# **Sample `config.yml` files**
 
```
apiKey: <your_API_key>
sources: [hacker-news, wired, ars-technica, bbc-sport, bloomberg, google-news, google-news-in, new-scientist, techcrunch, the-next-web, the-wall-street-journal, wired]
```
This will tell `termnews` to fetch news from the above sources every `15` minutes. 

```
apiKey: <your_API_key>
categories: [technology, business, science]
country: in
```
This will fetch tell `termnews` to fetch `technology`, `business` and `science` related news specific to `India` (`in`) every `15` minutes

```
apiKey: <your_API_key>
categories: [technology]
country: in
refreshFrequency: 5
```
This will tell `termnews` to fetch `technology` news in `India` (`in`) every `5` minutes


# Keyboard shortcuts
- `CtrlA` - previous page
- `CtrlD` - next page
- `Up Arrow` - move to the previous article
- `Down Arrow` - move to the next article
- `Right Arrow` - open selected article in the browser