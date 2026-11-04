# CrossMatch

A Telegram bot that pairs channel admins for mutual promotion: it proposes fitting channels, opens contact on mutual interest, and sees the agreed post swap through.

## Language

### Participants

**Channel**:
A Telegram channel registered in the bot by one of its admins.
_Avoid_: Group, chat, account

**Owner**:
The admin who registered a Channel and acts for it in the bot.
_Avoid_: User, client, admin (when the registered admin is meant)

**Detached**:
A Channel the bot can no longer act for, because the bot was removed from it or lost a right that Channel requires; it is neither shown as a Candidate nor offered any, and returns to normal once the bot is back.
_Avoid_: Disabled, paused, inactive

**Reach**:
The typical number of views a post in a Channel gets, as the bot measures it from the Channel's recent posts. A Channel whose Reach is not known is neither shown as a Candidate nor offered any.
_Avoid_: Subscribers, coverage, ERR, average reach

### Matching

**Candidate**:
A Channel the bot proposes to an Owner as a possible partner for one of their Channels.
_Avoid_: Suggestion, recommendation, offer

**Decision**:
An Owner's answer on a Candidate: interested or skip.
_Avoid_: Vote, like, swipe

**Match**:
Two Channels whose Owners each decided "interested" on the other; it lets them start an Exchange, and shows each Owner the other's contact unless that Owner chose to hide it.
_Avoid_: Pair, connection, deal

### Exchange

**Exchange**:
One agreed swap of posts within a Match: within an agreed period, each Channel publishes a post its Owner wrote about the other Channel. A Match may hold several Exchanges over time, one at a time.
_Avoid_: Swap, deal, mutual PR (as a term for a single agreement), ВП

**Exchange Post**:
The post a Channel publishes about its partner within an Exchange, written by that Channel's Owner in their own words.
_Avoid_: Ad, promo post, placement

**Exchange Link**:
An invite link the bot creates in a Channel for one Exchange, offered to the partner to put in their Exchange Post so the Channel's Owner can see how many people came through it.
_Avoid_: Invite link (as the general term), referral link, tracking link

**Pitch**:
A few lines an Owner writes about their own Channel for partners to draw on when writing their Exchange Post.
_Avoid_: Brief, ad text, description (Telegram's own channel description)

**Offer**:
An Exchange one Owner has proposed and the partner has not yet accepted or declined.
_Avoid_: Proposal, request, invitation

**Broken**:
An Exchange that ended through one side's fault: it did not publish within the period, or removed its post or the bot before the post's 24 hours were up.
_Avoid_: Failed, violated (for the Exchange as a whole)

**Lapsed**:
An Exchange whose period ended with neither side having published; no one is at fault.
_Avoid_: Expired, failed, cancelled

**Review**:
A participant's thumbs up or thumbs down on a completed Exchange, given to the partner's Channel.
_Avoid_: Rating, score, feedback

### Trust

**Violation**:
A breach of the Exchange terms recorded against a Channel: either the bot observed it (the Channel broke an Exchange) or Complaints about the Channel came from two different partners. It stays with the Telegram channel whoever its Owner is.
_Avoid_: Strike, penalty, complaint

**Complaint**:
A partner's report that a Channel's Exchange Post broke a specific rule the bot cannot check itself, such as not staying on top for an hour or being published silently.
_Avoid_: Report, flag, Review
