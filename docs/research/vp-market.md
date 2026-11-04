# How the Telegram mutual-promotion market works

Research for [#5](https://github.com/mokevnin/crossmatch/issues/5). Question: which terms of a 1-to-1 mutual promotion (ВП, взаимный пиар) are customary among Russian-speaking Telegram channels, so that the terms of an Exchange feel familiar to an Owner. Collected 2026-11-03.

What the bot can observe about channels and posts (Bot API) is out of scope here; see [#4](https://github.com/mokevnin/crossmatch/issues/4).

## Sources and how far to trust them

Every claim below carries a tier:

- **[P] Primary.** Rules or documentation published by a service about its own product:
  - Telega.in [service rules](https://telega.in/page/rules), revision of 09.12.2025, and its [FAQ](https://telega.in/faq). Telega.in is the largest Russian Telegram ad exchange; its rules are the closest thing the market has to a written standard for placement terms.
  - TGStat Academy articles (TGStat is the main Russian Telegram analytics service), each with its date, and the TGStat API [`channels/stat` reference](https://api.tgstat.ru/docs/ru/channels/stat.html).
  - The ARIR and Telega.in [market study](https://interactivead.ru/wp-content/uploads/2025/10/issledovanie_rynka_telegram_reklamy_arir_i_telega_in_2025_1-1.pdf), October 2025, a survey of more than 300 advertisers and agencies.
- **[S] Secondary.** Marketing publications that describe practice: eLama, Каллибри, Netpeak, Bidfox.
- **[A] Anecdotal.** Blog posts, product launch write-ups and user comments: vc.ru, DTF, ProductRadar, catalog pages. They show what practitioners say, not what is standard.
- **[snippet]** Seen only in a search-engine summary; the page itself could not be opened. Treat it as a lead, not a fact.

There is no primary source for ВП as such. Paid exchanges publish rules; mutual promotion is a handshake between two admins, and its conventions are inherited from paid placements. Most findings below are therefore about paid placements, and the text says so where it matters.

## Vocabulary an Owner already knows

| Term | Meaning | Source |
|---|---|---|
| ВП, взаимопиар | Cross-promotion of channels with links to each other. TGStat: a free swap of ad posts between channels close in topic, "своего рода бартер". | [S] [Netpeak](https://netpeak.net/ru/blog/sleng-v-telegram-kanaly-oborotni-prikhody-i-zakrepy/) 2024; [P] [TGStat](https://tgstat.ru/academy/1290-mutual-pr) 2025-09-04 |
| Мега | Several channels grouped into one list post that every participant publishes at an agreed time. TGStat calls such list posts "подборки" and warns that with 5 or more channels their effect is "крайне низкая" compared with a 1-to-1 swap. | [S] Netpeak; [P] TGStat 2025-09-04 |
| Топ, «N часов в топе» | The time a post stays the newest one in the channel. Telega.in defines it as "без обновлений и вытеснения новыми постами": no new posts above it. | [S] Netpeak; [P] Telega.in rules п. 3.7 |
| Лента | The channel feed; the total time the post stays published. | [P] Telega.in FAQ |
| 1/24, 2/48, 3/72 | Hours in the top / hours in the feed. See [Hold formats](#hold-formats). | [P] Telega.in rules п. 3.7 |
| Холд | The period after an ad is published during which the host channel publishes nothing. | [S] Netpeak |
| Закреп | A pinned message. In groups, Telega.in pins the ad for the whole term instead of holding it in the top. | [S] Netpeak; [P] Telega.in rules п. 3.7 |
| Нативный | A post written by the host channel's author and approved by the advertiser. | [P] Telega.in rules п. 3.7 |
| Охват, средний охват | Views of a post. TGStat reports the average reach of a post and, separately, the average reach of ad posts after 12, 24 and 48 hours. | [P] TGStat API |
| ERR | Share of subscribers who view posts: average views divided by subscribers. TGStat also reports ERR24, views in the first 24 hours. | [P] TGStat API; [P] [TGStat](https://tgstat.ru/academy/1212-advertising-channel) 2025-03-28 |
| ER | Interactions (reposts, comments, reactions) divided by reach. | [P] TGStat 2025-03-28 |
| ИЦ, индекс цитирования | TGStat's measure of how much other channels mention a channel; it ranges from 0 to several thousand. | [P] [TGStat](https://tgstat.ru/academy/1018-advertisers-guide) 2023-07-07 |
| CPM | Cost per thousand views. TGStat: cost = CPM × ad-post reach / 1000. | [P] TGStat 2025-03-28 |
| ОП, обязательная подписка | A bot that unlocks its features only after the user subscribes to listed channels; cheap, unstable traffic. Netpeak uses «ОП» for "exchange of posts between two channels"; TGStat's meaning is the common one in fraud checks. | [P] TGStat 2023-07-07; [S] Netpeak |
| Канал-трансформер, канал-оборотень | A channel that gathered an audience on one topic and switched to another (for example memes to business). | [P] [TGStat](https://tgstat.ru/academy/1277-channels-transformers) 2025-08-18; [S] Netpeak |
| Красная и чёрная метка | TGStat labels: red for inflated metrics, black for fraud or breaking an ad agreement. | [P] [TGStat](https://tgstat.ru/academy/1161-red-black-label) 2025-07-21 |
| Сетка | Several channels run by one author. | [S] Netpeak |

## Hold formats

The only written definitions found come from Telega.in (rules п. 3.7, FAQ) [P]:

| Format | In the feed at least | In the top at least | After the term |
|---|---|---|---|
| 1/24 | 24 h | 1 h | the post "должен быть удален" |
| 2/48 | 48 h | 2 h | deleted |
| 3/72 | 72 h | 3 h | deleted |
| 30 дней | 30 days | 3 h | deleted |
| Нативный | 7 days | 1 h | deleted |
| Репост | 1/24 of a repost from the advertised channel | 1 h | deleted |

- In groups the post is pinned for the whole term instead, and pinning other posts is forbidden for that term [P].
- The default when an advertiser specifies nothing is 1/24: "1 час в топе, 24 часа в ленте" [P] Telega.in FAQ. TGStat also calls 24 hours "наиболее распространенный формат" and describes it as "на срок 24 часа, затем реклама удаляется" [P] TGStat 2023-07-07. TGStat's ad-reach metrics stop at 12, 24 and 48 hours, the terms it says an ad usually stays up [P] TGStat API and 2025-03-28.
- Deletion after the term is required, not tolerated. Telega.in ties it to ad marking: a marked post must be removed when its term ends, otherwise the host must obtain a new marking token (п. 3.8) [P].
- Publishing without a notification ("в беззвучном режиме") is forbidden (п. 3.9) [P].
- TGStat advises agreeing that the ad will not be pushed down by other posts for "от 1 до 12 часов", because a post gets most of its reach in its first hours [P] [TGStat](https://tgstat.ru/academy/1243-best-advertising-time) 2025-07-07. In a normal view curve the first hour brings about 20 % of the first day's views [P] [TGStat](https://tgstat.ru/academy/1002-tools-views-analyze) 2022-06-09.

What the ticket calls "1/48" does not appear as a format anywhere; the documented two-day format is 2/48. "Без удаления" (keep forever) is not described by any primary source either; the nearest documented terms are 30 days and the 7-day native post. The TGStat article on ВП itself names no hold format at all. A search summary claims 1/24 is the usual ВП term [snippet], which fits the paid-market default but was not verified.

## Publication time and how dates are agreed

- **Same time on both sides.** For ВП, TGStat says the posts should go out "одновременно": it puts both sides on equal terms and lowers the risk that one side backs out after seeing a weak result from the partner's post [P] TGStat 2025-09-04.
- **Pick the slot from data.** TGStat recommends looking at both channels' activity by hour [P] 2025-09-04. Across Telegram, activity is lowest early in the morning and highest after 13:00, and lower on weekends, so advertisers "обычно избегают" Saturday and Sunday [P] TGStat 2025-07-07.
- **Moscow time.** TGStat shows dates and times in MSK [P] 2025-07-07; Telega.in deadlines are in UTC+3 (п. 7.9) [P].
- **How a paid exchange schedules.** An order states the date and time of publication among its requirements (п. 3.2) [P]. The host has 24 hours to accept or decline and then 36 hours to publish, unless the order says otherwise (п. 3.4, 7.4) [P]. With an exact date, the host can submit the link to the published post until 23:59 UTC+3 of that date; with "в ближайшее время", within 36 hours of accepting (п. 7.9) [P]. A discounted "hot" slot gives the host a fixed window of three hours either side of the announced time (п. 9.4) [P].
- **What practitioners write into agreements.** The exact date and time or an acceptable window, the time in the top, the total term, and when the post may be deleted [S] [eLama](https://elama.ru/blog/dogovor-na-reklamu-v-telegram-kanale/) 2026-08-28.
- **Pace.** "1–2 ВП в неделю" is called optimal [A] [DTF](https://dtf.ru/pro-smm/4762864-vzaimnyj-piar-telegram-kanalov); ВП Сервис caps a channel at 3 swaps a day [A] [ProductRadar](https://productradar.ru/product/vp-servis/).

## Text and creative

- **In ВП each side writes its own ad.** TGStat: "Вы составляете пост для размещения в канале-партнёра, а партнёр — для размещения в вашем", then the posts are exchanged and published [P] 2025-09-04. Its sample opening message to a partner states one's own average reach, which shows reach is the first thing compared.
- **The host may not change it.** Telega.in forbids changing the text unilaterally (п. 7.6) and lets the advertiser reject an order whose text was altered without agreement, whose media is missing, or whose link is dead or wrong (п. 3.12.2–3.12.4) [P]. Formatting is done by hand to the advertiser's requirements (п. 7.5) [P].
- **Native is the exception.** In the native format the host writes the post and the advertiser approves it (п. 3.7) [P].
- eLama adds: agree where and by when the final creative is approved, and forbid changing "текст, изображение или ссылку" without approval [S] 2026-08-28.

## How channels are compared

Thresholds TGStat publishes [P]:

| Signal | Threshold | Source |
|---|---|---|
| Reach | Partner's average 24-hour reach within ±10 % of yours, for ВП | TGStat 2025-09-04 |
| ERR | 10 % or more | TGStat 2025-09-04 and 2023-07-07 |
| ИЦ | 30 or more, for ad buying | TGStat 2023-07-07 |
| Red label | None | TGStat 2025-09-04 |
| Topic | Same or adjacent | TGStat 2025-09-04 |

- **Reach, not subscribers.** TGStat tells advertisers to rely on average *ad* reach for 12, 24 or 48 hours rather than average reach, since that is what an ad actually gets [P] 2025-03-28. Average reach per post is the headline figure in its catalog search [S] [Каллибри](https://callibri.ru/blog/kak-analizirovat-telegram-kanaly-v-tgstat).
- **Price.** A paid placement is priced as CPM × ad reach / 1000; TGStat suggests deriving a CPM from 10 or more similar channels [P] 2025-03-28. Advertisers say they judge Telegram ads by conversions (54 %), ER (41 %), CPM (39 %) and reach (36 %), and see the best return in micro- and nano-channels (46 %) [P] ARIR and Telega.in 2025.
- **Entry bar of an exchange.** Telega.in lists a channel only with at least 1,500 subscribers, an age of more than 3 months, at least 14 days of data in analytics services, Russian-language content and no suspicious audience (п. 5.6) [P].
- **Practitioner rules of thumb, weaker:** reach within 30 % and ER above 5 % [A] DTF; ERR at least 15 % and reach at least 40 % of subscribers [A] [vc.ru](https://vc.ru/id4642799/1842395-vzaimopiar-v-telegram-kak-delat-kollaboracii-kotorye-realno-rabotayut) 2025-03-02. When reach differs, a surcharge or an extra post is reported as the usual fix [snippet]; money is out of scope for the MVP anyway.

## Common cheats and how they are caught

### Breaking the placement terms

Telega.in's grounds for rejecting a placement (п. 3.12) are the most complete list of what goes wrong [P]:

1. Not published in the agreed channel.
2. Published without the image or video.
3. Link dead or different from the agreed one.
4. Text changed without agreement.
5. Date, time or formatting ignored.
6. Channel blocked or inactive.
7. Less time in the feed (or pinned) than the format requires.
8. Less time in the top than the format requires.
9. Suspicious audience: inflated subscribers or views.
10. Ad-marking token or advertiser data missing or changed.

Silent posting (п. 3.9) belongs on the same list. eLama's practitioner examples match: posted on the wrong day, removed early (after a week instead of 30 days), and "еще четыре публикации" within 20 minutes of the ad [S] 2026-08-28.

How they are caught:

- **The buyer checks, the exchange holds the money.** The host submits a link to the post; the buyer has 24 to 72 hours, by format, to complain; silence counts as acceptance (п. 7.9, 10.2) [P]. Money is reserved until then [P] Telega.in FAQ. Claims are reviewed by moderators, and repeated substantiated claims can get a channel removed from the catalog [P] [Telega.io FAQ](https://telega.io/faq).
- **Analytics keep deleted posts.** TGStat shows a channel's deleted posts unless "Скрывать удаленные посты" is set, which is how an advertiser sees every ad ever run [P] 2025-03-28. The Каллибри blog notes that deleted ads stay viewable on TGStat [S].
- **Simultaneous publication** removes the incentive to defect after seeing the partner's result [P] TGStat 2025-09-04.

### Inflated subscribers and views (накрутка)

TGStat's published signals [P]:

- Many subscriptions within a single hour (except when a private channel approves a backlog of join requests at once) [2023-07-07].
- Unsubscribes below 1 % of subscribes: "Не бывает живых подписчиков, которые не отписываются" [2023-07-07].
- Growth above 1 % of the ad reach the channel bought: bots added while ads were running, to hide among real joins [2023-07-07].
- ИЦ between 0 and 20 with a large audience: traffic from outside Telegram, bots, ОП or inflation [2023-07-07].
- ОП traffic: more than 80 % of those who joined in 24 hours have already left [2023-07-07].
- Hourly views: a normal post peaks in the first hour, falls hour by hour, dips at night and rises in the morning. Flat equal bars, an even plateau followed by a drop, or views arriving almost 20 hours late without any repost indicate inflation. A 10-minute breakdown of the first hour exposes careful inflation. Judge on at least 10 posts, not one [2022-06-09].

Who flags it:

- **TGStat red label**, given after an analyst's manual review and only with "100% уверенности"; TGStat itself warns that not every inflated channel carries one [P] 2025-07-21 and 2023-07-07. Its catalog hides labeled channels by default [S] Каллибри.
- **Telega.in** rejects or removes such channels (п. 5.6.7, 5.7), withholds an offender's earnings as a fine (п. 5.8) and bars re-moderation (п. 5.11) [P]. It also refuses channels whose audience came largely from grey methods, naming "взаимного пиара (меги)" among them (п. 5.6.13) [P].
- **eLama** lists the same checks for manual use: spikes, low ER, synchronised reactions, equal hourly view portions [S] [eLama](https://elama.ru/blog/kak-proverit-telegram-kanal-na-nakrutku-besplatno-bez-registracii-i-sms/) 2026-02-03.

### Bait-and-switch

- **Topic switch.** A channel grows on a cheap broad topic and then switches to an expensive narrow one; its audience does not engage. TGStat's "История канала" shows every change of username, title and description with dates; "Привлечение подписчиков" and the topic mix of ИЦ show whether the audience came before or after the switch [P] 2025-08-18.
- **Fake admin, substituted channel, fake statistics screenshots, delete-after-payment.** Described as common schemes in direct deals; the advice is to check the channel in TGStat or Telemetr and not to trust screenshots [S] [Bidfox](https://blog.bidfox.ru/polezno/1575) 2026.
- **Fraud and breach of an ad agreement.** TGStat gives the black label after a complaint with evidence is published by a partner channel that handles such complaints; TGStat does not examine complaints itself, and the label is lifted only after the conflict is settled [P] 2025-07-21.

## Existing ВП services

Evidence here is thin and mostly old; all of it is [A].

- **ВП Сервис** (Telegram bot; the page gives no date). Free; channels are moderated by hand and pick up to 10 of 150 categories; the bot shows matching channels, and once both accept, the parties arrange the placement themselves, because the bot has no right to post. Users complained about onboarding and multi-day moderation [A] [ProductRadar](https://productradar.ru/product/vp-servis/).
- **Erzy** (platform, 2022). Explicitly 1-to-1 ("ВП один-к-одному"); the parties set the start and end of an exchange; the bot watches that the partner does not delete the promo post and, if they do, the post is published again and both are notified; invite links give exact click counts [A] [vc.ru](https://vc.ru/tribuna/527933-erzy-platforma-vzaimnogo-piara-telegram-kanalov) 2022-10-28. Re-publishing implies the bot had posting rights in the channels.
- **@MutualTrafficBot** (2017). "Сам контролирует проведение рекламной акции" and removes a violator's links from posts if they delete the ad from their channel [A] [tgreef](https://tgreef.com/2754-bot-vzaimopiar.html) 2017-12-14.
- **ВП chats** (topic chats where admins look for partners) exist, but their rules could not be read; the handles listed in a DTF article were not verified [A].
- **TGStat** offers no ВП mechanics; it supplies the partner search, the checks above and, for measuring the result, its analytics bot and per-partner invite links ("Пригласительные ссылки") [P] 2025-09-04.
- **Telega.in** is a paid exchange, not a ВП service, but its rules are the template that ВП terms borrow from.

## Gaps

- No written source for "без удаления" or "1/48" as ВП terms.
- No primary description of how ВП chats police their members (ban lists, reputation).
- No current, verified source on how ВП bots verify publication; the only description of automatic checks is Erzy's 2022 launch post.
- The ARIR and Telega.in study has no CPM or format-share figures.

## Implications for CrossMatch

Inferences from the findings above, not findings themselves. They are inputs for [#9](https://github.com/mokevnin/crossmatch/issues/9) (Exchange lifecycle) and [#10](https://github.com/mokevnin/crossmatch/issues/10) (bot as admin).

- **Format vocabulary.** Express the hold as "N/M": hours in the top / hours in the feed, with Telega.in's meaning of top ("no new posts above it"). 1/24 is the familiar default; 2/48 and 3/72 are the familiar longer options. Offering "без удаления" would be a CrossMatch invention, since nothing documents it.
- **One shared slot.** A single agreed time for both posts, in MSK, matches TGStat's advice and closes the "publish second, then back out" gap.
- **Each Owner supplies their own post**, and the publishing Owner may not edit it. This matches TGStat's ВП procedure and Telega.in п. 7.6.
- **Violations.** Telega.in п. 3.12 plus п. 3.9 is a ready-made list: not published, wrong time, text or link changed, media missing, too short in the top, too short in the feed, silent post, inflated audience. Which of them the bot can detect depends on #4 and #10.
- **Partner fit.** "24-hour reach within ±10 %" and "ERR at least 10 %" are TGStat's published numbers and will look familiar. Owners also expect channels marked for inflation to be excluded.
- **Result.** A separate invite link per partner is the established way to count joins, which supports one invite link per Exchange.
- **1-to-1 only** is supported by the market: TGStat calls multi-channel lists far less effective, and Telega.in counts an audience from "меги" as grey.
