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

**Offer**:
An Exchange one Owner has proposed and the partner has not yet accepted or declined.
_Avoid_: Proposal, request, invitation

**Broken**:
An Exchange that ended through one side's fault: it did not publish within the period or did not keep its post up on the agreed terms.
_Avoid_: Failed, violated (for the Exchange as a whole)

**Lapsed**:
An Exchange whose period ended with neither side having published; no one is at fault.
_Avoid_: Expired, failed, cancelled
