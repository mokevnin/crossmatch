# What a bot can learn about a channel and its posts

Research for issue #4. It records facts, not decisions. Sources were read on 2026-11-03:

- Bot API 10.3, dated August 24, 2026, at <https://core.telegram.org/bots/api>
- MTProto API layer 225, at <https://core.telegram.org/api>
- Telegram Terms of Service, the Bot Developer Terms and the API Terms
- the TGStat API documentation at <https://api.tgstat.ru/docs>

A claim the documentation does not settle is marked **unverified**. All such claims are collected in
[Open questions](#open-questions).

## The variants that actually exist

The issue names three variants. The sources change them in two ways.

**There is no "bot as a plain member" of a channel.** `channels.inviteToChannel` lists the error
`USER_BOT`, "Bots can only be admins in channels"
([method page](https://core.telegram.org/method/channels.inviteToChannel)). A bot is either:

- outside the channel. It can then see only what is public about a public channel.
- an administrator. Its rights can be as small as none. MTProto `chatAdminRights` has an `other`
  flag: "Set this flag if none of the other flags are set, but you still want the user to be an
  admin" ([constructor](https://core.telegram.org/constructor/chatAdminRights)).

**A bot token also works over MTProto.** A library such as gotd/td for Go can log a bot in to
MTProto. Every MTProto method page states either "Both users and bots can use this method" or
"Only users can use this method". Some methods the Bot API does not expose are open to bots this
way.

That leaves four variants:

1. Bot API, bot as admin without `can_post_messages`. This is the "admin without posting rights"
   variant.
2. The same bot over MTProto. See [Bot over MTProto](#bot-over-mtproto).
3. MTProto user account (userbot).
4. External services (TGStat API).

## Summary matrix

| Fact | Bot API, admin, no posting rights | Bot over MTProto | Userbot | TGStat API |
| --- | --- | --- | --- | --- |
| User is an admin of the channel | `getChatMember`, `getChatAdministrators`, `my_chat_member`, `KeyboardButtonRequestChat` | `channels.getParticipant` | `channels.getParticipant` | no |
| Subscriber count | `getChatMemberCount` | `channels.getFullChannel` → `participants_count` | same, plus `stats.getBroadcastStats` (admin only) | `channels/get`, `channels/stat` |
| Views of a post | **no** (Bot API `Message` has no views field) | `channels.getMessages` → `message.views` (**unverified** for bots) | `messages.getMessagesViews`, `channels.getMessages`, `stats.getMessageStats` | `posts/stat` → `viewsCount`, hourly or daily growth |
| New posts and edits | `channel_post`, `edited_channel_post` | updates, `updates.getChannelDifference` | updates; public non-member channels by polling | Callback `new_post`, `edit_post` |
| Post deletion | **no update**; existence check through `forwardMessages` | `updateDeleteChannelMessages` (**unverified** for bots); `channels.getMessages` | `updateDeleteChannelMessages` | Callback `remove_post` |
| Per-exchange invite link | `createChatInviteLink` (`can_invite_users`) | `messages.exportChatInvite` | same | no |
| Joins through a link | `chat_member` updates carry `invite_link`; no usage counter | (as Bot API) | `chatInviteExported.usage`, `messages.getChatInviteImporters` (user only) | no |

## Bot API

### Receiving updates

- Updates arrive through long polling with `getUpdates` or through a webhook. The two are mutually
  exclusive. "Incoming updates are stored on the server until the bot receives them either way,
  but they will not be kept longer than 24 hours"
  ([Getting updates](https://core.telegram.org/bots/api#getting-updates)).
- The default `allowed_updates` covers "all update types except `chat_member`, `message_reaction`,
  and `message_reaction_count`". Those three must be named explicitly
  ([getUpdates](https://core.telegram.org/bots/api#getupdates)).
- "All bots, regardless of settings, will receive … All messages from channels where they are a
  member" ([FAQ](https://core.telegram.org/bots/faq#what-messages-will-my-bot-get)). For a channel,
  "member" means admin, as shown above.

### Verifying that a user is an admin of the channel

- `getChatMember(chat_id, user_id)` returns a `ChatMember`. For an owner or admin, `status` is
  `creator` or `administrator`, together with the rights flags. "The method is only guaranteed to
  work for other users if the bot is an administrator in the chat"
  ([getChatMember](https://core.telegram.org/bots/api#getchatmember)).
- `getChatAdministrators(chat_id)` lists the admins
  ([getChatAdministrators](https://core.telegram.org/bots/api#getchatadministrators)).
- When someone adds the bot, the bot receives `my_chat_member`. Its `from` field is the "Performer
  of the action, which resulted in the change"
  ([ChatMemberUpdated](https://core.telegram.org/bots/api#chatmemberupdated)). Bots can only be
  admins in a channel, and adding an admin needs `can_promote_members`. So the performer should
  hold that right. This is an inference from the two documented facts, not a statement in the
  docs.
- `KeyboardButtonRequestChat` lets the client do the filtering. With `chat_is_channel`,
  `chat_is_created` ("a chat owned by the user"), `user_administrator_rights` and
  `bot_administrator_rights`, the Telegram client lists only channels that meet the criteria. "The
  bot will be granted requested rights in the chat if appropriate"
  ([KeyboardButtonRequestChat](https://core.telegram.org/bots/api#keyboardbuttonrequestchat)).

### Subscriber count

- `getChatMemberCount(chat_id)` returns an Integer
  ([getChatMemberCount](https://core.telegram.org/bots/api#getchatmembercount)). `chat_id` may be
  `@username` for a channel. The docs do not say whether the call works on a public channel the bot
  has not joined (**unverified**).
- `ChatFullInfo`, returned by `getChat`, has no subscriber-count field
  ([ChatFullInfo](https://core.telegram.org/bots/api#chatfullinfo)).

### Post views

- The Bot API `Message` object has no views counter
  ([Message](https://core.telegram.org/bots/api#message)). A text search of the whole Bot API page
  finds no views field and no statistics method.
- `ChatAdministratorRights.can_post_messages` reads: "can post messages in the channel, approve
  suggested posts, or access channel statistics"
  ([ChatAdministratorRights](https://core.telegram.org/bots/api#chatadministratorrights)). The
  statistics themselves are reachable only through MTProto user methods (see
  [Userbot](#userbot-mtproto-user-account)).

### New posts and edits

- `channel_post` is a "New incoming channel post of any kind". `edited_channel_post` is a "New
  version of a channel post that is known to the bot and was edited. This update may at times be
  triggered by changes to message fields that are either unavailable or not actively used by your
  bot" ([Update](https://core.telegram.org/bots/api#update)).
- Fields of `Message` that matter here:
  - `edit_date`
  - `forward_origin`
  - `author_signature`
  - `has_protected_content`, "the message can't be forwarded"
  - `media_group_id`, which ties the parts of an album together
- Admin rights are not needed to receive these updates. Being an admin is what puts the bot in the
  channel.

### Detecting that a post was deleted

- The Bot API has no update for a deleted channel post. The only deletion update is
  `deleted_business_messages`, and it covers connected business accounts only
  ([Update](https://core.telegram.org/bots/api#update)).
- The workaround is to try to forward or copy the post:
  - `forwardMessage` returns the sent `Message`. "Service messages and messages with protected
    content can't be forwarded" ([forwardMessage](https://core.telegram.org/bots/api#forwardmessage)).
  - `forwardMessages` takes 1–100 ids in strictly increasing order. "If some of the specified
    messages can't be found or forwarded, they are skipped", and it returns an array of
    `MessageId` ([forwardMessages](https://core.telegram.org/bots/api#forwardmessages)). One call
    therefore checks up to 100 posts. The documentation does not state the exact error text
    `forwardMessage` returns for a missing post.
  - `copyMessage` has the same shape and does not link back to the original. Its documented
    exclusions are "Service messages, paid media messages, giveaway messages, giveaway winners
    messages, and invoice messages" ([copyMessage](https://core.telegram.org/bots/api#copymessage)).
    Its behaviour on protected content is **unverified**.
- Costs and limits of the workaround:
  - every check sends a real message to some target chat;
  - a channel with protected content cannot be checked this way;
  - "In a single chat, avoid sending more than one message per second", and "In a group, bots are
    not be able to send more than 20 messages per minute"
    ([FAQ](https://core.telegram.org/bots/faq#my-bot-is-hitting-limits-how-do-i-avoid-this)).
- Telegram's own paid suggested posts carry a deletion rule. `Message.is_paid_post`: "such posts
  must not be deleted for 24 hours to receive the payment and can't be edited". The refund reason
  `post_deleted` means "the post was deleted within 24 hours of being posted"
  ([SuggestedPostRefunded](https://core.telegram.org/bots/api#suggestedpostrefunded)). The
  mechanism runs through channel direct messages and is paid in Stars or TON
  ([Bot features](https://core.telegram.org/bots/features)).

### Per-exchange invite links and counting joins

- `createChatInviteLink(chat_id, name, expire_date, member_limit, creates_join_request)` returns a
  `ChatInviteLink`. The bot "must be an administrator in the chat for this to work and must have
  the appropriate administrator rights", meaning `can_invite_users`.
  - `name` is 0–32 characters.
  - `member_limit` is "The maximum number of users that can be members of the chat simultaneously
    after joining the chat via this invite link; 1-99999".
  - With `creates_join_request` true, "member_limit can't be specified"
  ([createChatInviteLink](https://core.telegram.org/bots/api#createchatinvitelink)).
- "Each administrator in a chat generates their own invite links. Bots can't use invite links
  generated by other administrators." Links created by another admin show up with the second part
  replaced by "…" ([ChatInviteLink](https://core.telegram.org/bots/api#chatinvitelink)).
- Bot API `ChatInviteLink` has no count of joins. It has only `pending_join_request_count` for
  links that use join requests.
- Joins are counted from `chat_member` updates:
  - "The bot must be an administrator in the chat and must explicitly specify `chat_member` in the
    list of allowed_updates" ([Update](https://core.telegram.org/bots/api#update)).
  - `ChatMemberUpdated.invite_link` is the "Chat invite link, which was used by the user to join
    the chat; for joining by invite link events only"
    ([ChatMemberUpdated](https://core.telegram.org/bots/api#chatmemberupdated)).
  - A later leave arrives as a status change without the link. Attributing leaves therefore needs a
    user-to-link record the bot keeps itself.
  - Updates are kept for 24 hours at most. An update the bot misses cannot be recovered through the
    Bot API.
- With `creates_join_request`, the bot gets a `chat_join_request` update, which carries
  `invite_link` and `from`. It then calls `approveChatJoinRequest` or `declineChatJoinRequest`.
  Both need `can_invite_users` ([ChatJoinRequest](https://core.telegram.org/bots/api#chatjoinrequest)).

### Minimum rights

| Need | Minimum right |
| --- | --- |
| Be in the channel and receive `channel_post` and `edited_channel_post` | admin with any rights, including none |
| `chat_member` updates, `getChatMember` for other users | admin |
| Create, edit or revoke invite links; handle join requests | `can_invite_users` |
| Access channel statistics (MTProto only) | `can_post_messages`, per the Bot API wording. See the conflict noted in [Userbot](#userbot-mtproto-user-account). |

### Rate limits and price

- The Bot API is free. Published limits cover sending only:
  - about 1 message per second in a single chat;
  - 20 messages per minute in a group;
  - about 30 messages per second for broadcasts. Paid broadcasts go up to 1000 per second at
    0.1 Star per message beyond the free 30. They need 100,000 Stars on balance and 100,000
    monthly active users
    ([FAQ](https://core.telegram.org/bots/faq#my-bot-is-hitting-limits-how-do-i-avoid-this)).
- When flood control is exceeded, the error carries `retry_after`, "the number of seconds left to
  wait before the request can be repeated"
  ([ResponseParameters](https://core.telegram.org/bots/api#responseparameters)). No limit is
  published for read methods such as `getChatMember` or `getChatMemberCount`.

## Bot over MTProto

A bot token can log in to MTProto. That takes an `api_id` and `api_hash`
([Obtaining api_id](https://core.telegram.org/api/obtaining_api_id)). Availability per method,
from each method page:

| Method | Who may call it | Relevant content |
| --- | --- | --- |
| `channels.getMessages` | users and bots | returns `message` objects; the schema has `views:flags.10?int` and `forwards` ([method](https://core.telegram.org/method/channels.getMessages)) |
| `channels.getFullChannel` | users and bots | `channelFull.participants_count`, `can_view_stats`, `stats_dc` ([method](https://core.telegram.org/method/channels.getFullChannel)) |
| `channels.getParticipant` | users and bots | errors include `CHAT_ADMIN_REQUIRED` and `USER_NOT_PARTICIPANT` ([method](https://core.telegram.org/method/channels.getParticipant)) |
| `updates.getChannelDifference` | users and bots | the channel's event stream, including deletions ([method](https://core.telegram.org/method/updates.getChannelDifference)) |
| `messages.exportChatInvite` | users and bots | `expire_date`, `usage_limit`, `request_needed`, `title` ([method](https://core.telegram.org/method/messages.exportChatInvite)) |
| `messages.getMessagesViews` | users only | ([method](https://core.telegram.org/method/messages.getMessagesViews)) |
| `messages.getHistory` | users only | ([method](https://core.telegram.org/method/messages.getHistory)) |
| `messages.getChatInviteImporters` | users only | ([method](https://core.telegram.org/method/messages.getChatInviteImporters)) |
| `messages.getExportedChatInvites`, `messages.getExportedChatInvite` | users only | these return `chatInviteExported.usage` |
| `stats.getBroadcastStats`, `stats.getMessageStats` | users only | ([method](https://core.telegram.org/method/stats.getBroadcastStats)) |
| `channels.getAdminLog` | users only | ([method](https://core.telegram.org/method/channels.getAdminLog)) |
| `channels.joinChannel` | users only | ([method](https://core.telegram.org/method/channels.joinChannel)) |

- `updateDeleteChannelMessages` reports "Some messages in a supergroup/channel were deleted". It
  carries the message ids
  ([constructor](https://core.telegram.org/constructor/updateDeleteChannelMessages)).
- "The API will usually send passive updates … for channels/supergroups the user/bot is a member
  of" ([Working with updates](https://core.telegram.org/api/updates#subscribing-to-updates-of-channels-supergroups)).
- Two things are **unverified**: whether the server fills `views` in messages returned to a bot
  session, and whether a bot session receives `updateDeleteChannelMessages`. The Bot API's lack of
  both may come from the Bot API layer alone or from the server.

## Userbot (MTProto user account)

### What it can read

- **Admin status:** `channels.getParticipant`, and `channelFull` flags.
- **Subscriber count:** `channels.getFullChannel` returns `participants_count`.
- **Views:**
  - `messages.getMessagesViews(peer, id, increment)`: "Get and increase the view counter of a
    message sent or forwarded from a channel". `increment` decides "Whether to mark the message as
    viewed and increment the view counter".
  - Messages from `channels.getMessages` or `messages.getHistory` carry `views` and `forwards`.
  - Errors include `CHANNEL_PRIVATE`, "You haven't joined this channel/supergroup".
- **Channel statistics:** "Administrators of channels of a certain size (the exact limit is a
  server-side config, returned in the `can_view_stats` flag of `channelFull`) can invoke
  `stats.getBroadcastStats`". The result includes:
  - `followers`
  - `views_per_post` and `shares_per_post`
  - `recent_posts_interactions`, which has views, forwards and reactions per post
  - per-message statistics through `stats.getMessageStats`

  Source: [Channel statistics](https://core.telegram.org/api/stats#channel-statistics).
- **The sources disagree on the right needed for statistics.** MTProto `chatAdminRights.other`
  says that if this or any other flag is set, "the admin can … get chat statistics, get message
  statistics in channels". The Bot API attaches "access channel statistics" to
  `can_post_messages`. The live server behaviour is **unverified**.
- **Joins through a link:** `chatInviteExported` has `usage` and `usage_limit`. With the `link`
  parameter, `messages.getChatInviteImporters` lists the users who joined through that link. The
  caller must be an admin; the error is `CHAT_ADMIN_REQUIRED`
  ([Invite links](https://core.telegram.org/api/invites)).
- **New posts, edits and deletions:** `updateNewChannelMessage`, `updateEditChannelMessage` and
  `updateDeleteChannelMessages` arrive for channels the account is a member of
  ([Working with updates](https://core.telegram.org/api/updates)).

### Public channels the account has not joined

- The account can poll `updates.getChannelDifference`. The server then sends passive updates for
  as long as polling continues.
- The docs ask clients to "limit to 10 the maximum number of channels/supergroups short-polled
  using the above mechanism"
  ([Working with updates](https://core.telegram.org/api/updates#subscribing-to-updates-of-channels-supergroups)).
- Joining a channel instead has its own cost. The joined account is counted as a subscriber of that
  channel. An account can join at most 500 channels and supergroups, or 1000 with Premium
  (`channels_limit_default` and `channels_limit_premium` in
  [client configuration](https://core.telegram.org/api/config)).

### Terms and ban risk

The [api_id page](https://core.telegram.org/api/obtaining_api_id#using-the-api-id) states:

> all API client libraries are strictly monitored to prevent abuse. If you use the Telegram API for
> flooding, spamming, faking subscriber and view counters of channels, you will be banned forever.
> Due to excessive abuse of the Telegram API, all accounts that log in using unofficial Telegram
> API clients are automatically put under observation to avoid violations of the Terms of Service.

A banned developer can write to `recover@telegram.org`. The same page says "emails are checked by
a human". Each phone number can have only one `api_id`.

The [API Terms of Service](https://core.telegram.org/api/terms) add:

- 1.4: it is forbidden to make "actions on behalf of the user without the user's knowledge and
  consent".
- 1.5: data obtained through the API may not be used for AI or ML training.
- 3.3: "If your app allows accessing content from Telegram channels, you must include support for
  official sponsored messages".
- 4.1–4.2: Telegram notifies the account first. Access is cut "within 10 days" if the breach is not
  fixed.

Flood control appears as `420 FLOOD_WAIT_X`, "A wait of X seconds is required"
([Error handling](https://core.telegram.org/api/errors)). Concrete per-method limits are not
published.

The `increment` flag of `messages.getMessagesViews` raises the view counter. Raising it from a
service account touches the "faking … view counters" clause quoted above.

## Telegram terms that apply to every variant

- **Bot Developer Terms 5.2(d)** ([Operation](https://telegram.org/tos/bot-developers#5-2-operation)).
  A bot "must not misrepresent the services or functions it provides, including but not limited to
  misleading users into performing certain actions in the pursuit of an unachievable outcome". The
  forbidden examples include "(ii) Social growth manipulation". The terms do not define the phrase.
- **Bot Developer Terms 5.2(f).** A bot "must not attempt to circumvent or otherwise undermine
  Telegram rate limits and moderation … must not operate by proxy (i.e., using Bot API credentials
  supplied by other users)".
- **Bot Developer Terms 4.3**
  ([Data scraping](https://telegram.org/tos/bot-developers#4-3-data-scraping)). A bot may not
  "collect, store, aggregate or process data beyond what is essential for the operation of your
  services", and "scraping public group or channel contents" is named as always prohibited.
- **Bot Developer Terms 4.2.** User data must be deleted on request, when no longer needed, and when
  the bot stops operating. 4.1 requires a privacy policy, set in `@BotFather` if the standard one
  does not fit.
- **The [Content Licensing and AI Scraping](https://telegram.org/tos/content-licensing) terms**
  apply to everyone. Access to user content is allowed "solely to the extent strictly required to
  operate the relevant service".
- **[Terms of Service](https://telegram.org/tos)**, Termination. Non-compliance "may result in a
  temporary or a permanent ban from Telegram or some of its services".

## TGStat API

TGStat is a third-party index. Its data comes from TGStat's own collection, not from the channel
admin.

### Products and pricing

The API is sold as three products ([Introduction](https://api.tgstat.ru/docs/ru/start/intro.html)):

- **API Stat.** "Можно попробовать бесплатно на тарифе FREE для каналов (не более 2), владельцем
  которых вы являетесь. Платно для остальных каналов." In English: a free tier for at most two
  channels you own; every other channel is paid. Quotas count monthly requests and unique channels.
- **API Search.**
- **API Callback.** It is billed by the number of tracked objects.

Prices are **unverified**. The tariff pages `tgstat.ru/api/stat` and `tgstat.ru/api/callback`
returned HTTP 403 to an automated fetch. Quota use is visible through `usage/stat`. Its example
shows `spentChannels: "1989/2500"` and `spentRequests: "89152/400000"` for "тариф L"
([usage/stat](https://api.tgstat.ru/docs/ru/usage/stat.html)).

### What it returns

- **`channels/get`** returns `participants_count`, `tg_id`, `category`, `language` and
  `rkn_verification`, the channel's status in the Russian register of bloggers
  ([channels/get](https://api.tgstat.ru/docs/ru/channels/get.html)).
- **`channels/stat`** returns:
  - `participants_count`
  - `avg_post_reach`
  - `adv_post_reach_12h`, `adv_post_reach_24h` and `adv_post_reach_48h`
  - `err_percent` and `err24_percent`
  - `er_percent`
  - `daily_reach`
  - `ci_index`
  - mention and forward counts
  - `posts_count`, the "общее количество неудаленных публикаций" (non-deleted posts)

  It is available on all API Stat tariffs ([channels/stat](https://api.tgstat.ru/docs/ru/channels/stat.html)).
- **`posts/stat`** takes `postId` as a `t.me/username/123` link, a `t.me/c/…` link or a TGStat
  id. It returns:
  - `viewsCount`, `sharesCount`, `commentsCount`, `reactionsCount`, `forwardsCount` and
    `mentionsCount`
  - a `views` array with the growth "за первые 15 дней после публикации" (over the first 15 days
    after publication), by hour or by day

  Source: [posts/stat](https://api.tgstat.ru/docs/ru/posts/stat.html).
- **Callback `callback/subscribe-channel`** takes `event_types` from `new_post`, `edit_post` and
  `remove_post` ([subscribe-channel](https://api.tgstat.ru/docs/ru/callback/subscribe-channel.html)).
  - Delivery is a POST to a callback URL that must be verified as the subscriber's own.
  - A failed delivery is retried every 3 seconds until it gets a 2xx. The response timeout is
    3 seconds.
  - The average delay is "примерно 2-3 секунды" (about 2–3 seconds)
    ([Callback intro](https://api.tgstat.ru/docs/ru/callback/intro.html)).
- **Not available:** admin verification and invite links. TGStat has no access to a channel's admin
  data.
- **Rate limiting** shows up as `flood_control_10` and `flood_control_60`. "На тарифах выше S
  данные ограничения более дружелюбные" (the limits are looser on tariffs above S). Quota errors
  include `quota_requests_reached` and `quota_callback_objects_reached`
  ([Errors](https://api.tgstat.ru/docs/ru/errors.html)).

### Gaps

- A channel's data exists only if TGStat indexes that channel.
- No terms of use for the API were found among the documentation pages. They are **unverified**.

## Open questions

These need a live test or another source.

1. Do `getChatMemberCount` and `getChat` work on a public channel the bot has not joined, given by
   `@username`?
2. Does `forwardMessage` or `forwardMessages` work from a public channel the bot has not joined?
3. What exact error does `forwardMessage` return for a deleted post? Does `copyMessage` refuse
   protected content?
4. Does a bot session over MTProto get `views` in messages from `channels.getMessages`?
5. Does a bot session over MTProto receive `updateDeleteChannelMessages` for a channel where it is
   an admin?
6. Are `chat_member` updates delivered reliably for every join in large channels? The docs set no
   size limit.
7. Which admin right does `stats.getBroadcastStats` require: any admin right, or
   `post_messages`? What is the current subscriber threshold behind `can_view_stats`?
8. What are TGStat's current prices and tariff quotas, and its terms of use?
9. Other external services, such as Telemetr, were not researched.
