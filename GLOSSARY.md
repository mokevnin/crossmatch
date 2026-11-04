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
One agreed swap of posts within a Match: each Channel publishes a post about the other on agreed terms. A Match may hold several Exchanges over time.
_Avoid_: Swap, deal, mutual PR (as a term for a single agreement), ВП
