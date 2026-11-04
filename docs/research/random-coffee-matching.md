# Что происходит после мэтча: Random Coffee, Donut, Tinder, Hinge, Bumble

Вопрос: как сервисы в духе Random Coffee (и дейтинг-приложения как эталонная модель) устроены после того, как пара составлена:

1. есть ли шаг подтверждения перед раскрытием контакта;
2. можно ли отказаться от конкретной пары и узнаёт ли об этом партнёр;
3. что делают с гостингом (напоминания, перемэтч, санкции, вопрос «встреча состоялась?»);
4. что делает анматч или блок: скрывает ли контакт, удаляет ли переписку, влияет ли на будущий подбор;
5. как поставить участие на паузу.

Все источники прочитаны **2026-11-03**. Здесь только факты из первоисточников (справка, FAQ, страницы сервиса, статьи и посты самих команд), без выводов для CrossMatch. Цитаты даны дословно на языке оригинала. Статьи справочных центров Tinder, Hinge и Bumble (Zendesk) отдают 403 на обычный запрос страницы, поэтому их текст взят из JSON Zendesk Help Center API (`/api/v2/help_center/en-us/articles/<id>.json`), то есть из той же статьи; рядом указана дата последней правки статьи (`edited_at`).

## Сводка

| | Подтверждение до раскрытия контакта | Отказ от конкретной пары | Гостинг | Анматч / блок | Пауза |
|---|---|---|---|---|---|
| **Random Coffee** (боты сообществ) | Нет подтверждения пары. Есть только еженедельное «участвуешь на следующей неделе?», затем в понедельник обоим сразу приходит @username и анкета партнёра | Кнопки отказа не описаны. Правило: «лучше не сливаться», написать партнёру, лучше до среды | «Рейтинг ответственности» падает за неявку, вопрос «Как прошла встреча?» в конце недели, ответы обоих сверяются | Отдельной функции не описано. По жалобам других участников доступ могут прекратить | Пауза на неделю или месяц, «время сгорает» |
| **@randomcoffee_official_bot** (групповые чаты) | Нет. Пятничный опрос «кто пойдёт», в понедельник пары публикуются в общем чате | Не описано | Не описано | Не описано | Не идти в пятничном опросе |
| **Donut** (Slack) | Нет. Бот сразу создаёт групповой DM с парой | «🚫 I'm unavailable» или «♻️ I need a new match». Партнёру сообщают, причина конфиденциальна | Чек-ины в DM, статус «no response». После трёх подряд «призрак» исключается из раундов, ему об этом сообщают | «Don't pair me with...», выбор конфиденциален | Snooze на выбранную дату (до 4 циклов), opt-out, выход из канала |
| **Tinder** | Мэтч = взаимный лайк, сразу открывается чат в приложении. Телефон или ник не раскрывается | Unmatch, необратимо | В справке механизмов не найдено | Unmatch: исчезаете из списков друг друга, без уведомления. Block Profile: перестаёте видеть друг друга | Discovery off |
| **Hinge** | То же: мэтч открывает чат в приложении | Unmatch с выбором причины; Hide вместо анматча | «Your Turn»-напоминания; «Your Turn Limits» блокирует новые лайки; «We Met» спрашивает о свидании | Unmatch: переписка исчезает у обоих. Block List по контактам, без уведомления | Pause |
| **Bumble** | Мэтч, затем 72 часа на первое сообщение | Unmatch, необратимо | Мэтч истекает за 72 часа без сообщения. Extend / Rematch | Unmatch, без отдельного уведомления (но см. противоречие ниже). Блок приватный, необратимый | Snooze Mode со статусом «Away» |

## Random Coffee (random-coffee.ru), боты сообществ

Основной продукт: встречи один на один внутри тематических сообществ, каждое сообщество работает как отдельный Telegram-бот, подписка платная. Источники: страница «Встречи» [random-coffee.ru/meetings](https://random-coffee.ru/meetings), страница с FAQ [random-coffee.ru/page19036178.html](https://random-coffee.ru/page19036178.html), лендинг [more.random-coffee.ru/meetings](https://more.random-coffee.ru/meetings), «Наши принципы» [random-coffee.ru/nashi-principy](https://random-coffee.ru/nashi-principy), оферта сообщества [random-coffee.ru/pravovaya-informaciya/oferta-soobschestva](https://random-coffee.ru/pravovaya-informaciya/oferta-soobschestva), статья команды на Хабре [«От рандома до модели»](https://habr.com/ru/articles/727142/) (5 апреля 2023).

### Подтверждение и раскрытие контакта

- Подтверждение есть только на уровне недели, а не пары. Сценарий на лендинге: в воскресенье бот спрашивает «Хочешь встретиться на предстоящей неделе?» → «да!», в понедельник приходит «Вот твой партнёр Random Coffee: @pavelkozlov. Договоритесь про место и время.» ([more.random-coffee.ru/meetings](https://more.random-coffee.ru/meetings)).
- «В конце недели бот спрашивает, готовы ли вы участвовать во встречах на следующей.» ([random-coffee.ru/meetings](https://random-coffee.ru/meetings))
- Контакт раскрывается обоим одновременно и без дополнительного согласия на конкретного человека: «Ты получаешь анкету и контакты интересного человека, который станет собеседником, а он твои.» ([random-coffee.ru/soobschestva/dlya-obscheniya](https://random-coffee.ru/soobschestva/dlya-obscheniya)); «В понедельник ты получишь контакты и анкету с основной информацией о человеке» ([page19036178](https://random-coffee.ru/page19036178.html)).
- Согласие на неделю трактуется как договорённость: «Когда вы соглашаетесь на встречу в конце недели, считайте, что вы уже договорились встретиться с партнером где-то на неделе, ведь он тоже получил твои контакты и настроен на встречу.» ([random-coffee.ru/meetings](https://random-coffee.ru/meetings))
- Связь участники налаживают сами: «Бот присылает вам контакты собеседника, а вы списываетесь и назначаете созвон или живую встречу.» ([random-coffee.ru/meetings](https://random-coffee.ru/meetings))

### Механики взаимного интереса

- **Предвыбор**: можно просмотреть анкеты участников сообщества и отметить, с кем хочется встретиться. «В понедельник бот мэтчит тех, у кого случился взаимный лайк в Предвыборе, потом по отдельным лайкам, а дальше включается обычное распределение.» ([page19036178](https://random-coffee.ru/page19036178.html)). О том, видит ли человек, что его лайкнули или не лайкнули в ответ, на странице ничего нет.
- **True Random** (сообщество с опцией дейтинга): статус «в активном поиске» раскрывается только при взаимной положительной оценке после встречи. «Если оба человека в «активном поиске» и оценили встречу на «хорошо» и «отлично», бот сообщит о вашем статусе друг другу.» ([page19036178](https://random-coffee.ru/page19036178.html))

### Отказ от пары после распределения

- До распределения: «Перед распределением пар можно отказаться — если на предстоящей неделе не складывается или нет настроения.» ([page19036178](https://random-coffee.ru/page19036178.html))
- После распределения кнопки отказа или перемэтча в публичных материалах не описаны. Описано только правило поведения: «Когда пара уже назначена, лучше не сливаться. Если не получается встретиться, напиши сообщение своему партнёру. Небольшое извинение всегда лучше, чем игнорирование.» ([page19036178](https://random-coffee.ru/page19036178.html), [more.random-coffee.ru/meetings](https://more.random-coffee.ru/meetings))
- «Наши принципы»: «если так вышло, что ты не сможешь встретиться с партнером на неделе, обязательно напиши ему об этом. […] Будет здорово, если ты сообщишь паре о форс-мажоре до среды, ведь тогда у него будет возможность получить нового партнера.» ([nashi-principy](https://random-coffee.ru/nashi-principy)). Как именно партнёр получает нового собеседника в середине недели, не описано.
- Фраза «Бывает так, что «не зашло». Напиши нам в поддержку…» в FAQ относится к сервису или сообществу в целом (вопрос «А что если мне не понравится?», дальше речь о возврате денег), а не к конкретному собеседнику ([page19036178](https://random-coffee.ru/page19036178.html), [more.random-coffee.ru/meetings](https://more.random-coffee.ru/meetings)).

### Гостинг и неявка

- В конце недели бот спрашивает: «Как прошла встреча? Пойдёшь на новую встречу на следующей неделе?» ([more.random-coffee.ru/meetings](https://more.random-coffee.ru/meetings)). В веб-версии можно «посмотреть историю встреч и оценить прошедшие (а заодно выбрать, кого стоит еще раз позвать на кофе)» ([page19036178](https://random-coffee.ru/page19036178.html)).
- «Наши принципы»: «мы не приветствуем игнора в адрес друг друга.» И далее: «У нас существует рейтинговая система. Рейтинг растет каждый раз, когда встреча успешно состоялась, и падает каждый раз, когда ты согласился на встречу, но не пошел на нее. […] чем выше твой рейтинг ответственности, тем выше этот рейтинг у твоих потенциальных партнеров». Ответы обоих сверяются: «если ты ответишь, что встреча не состоялась, а твой партнер ответит обратное, скорее всего, мы спросим кого-то из вас, где именно случилась ошибка.» ([nashi-principy](https://random-coffee.ru/nashi-principy))
- Хабр, 2023: в раннем, «раундовом» подборе рейтинг ответственности был одним из критериев («насколько ты держишь обещания перед своими собеседниками и ходишь на встречи»). В более поздней ML-модели история посещений — признак, который снижает предсказанную вероятность встречи: «если он пропустил четыре встречи, то модель предскажет, что скорее всего, он пропустит и следующую» ([habr.com/ru/articles/727142](https://habr.com/ru/articles/727142/)). Явного штрафа (исключения из подбора за неявки) в этих источниках нет.

### Блокировка и жалобы

- Функция вида «не ставить меня в пару с этим человеком» в публичных материалах не описана.
- Оферта сообщества: «В случае получения жалоб на Покупателя от других участников чат-бота, Продавец вправе прекратить оказание услуги.» и «В случае противоречащего общепринятым нормам морали и нравственности поведении, получения, в связи с этим жалоб на Покупателя от других участников чат-бота, доступ в чат-бот прекращается, а плата данному Покупателю не возвращается.» ([oferta-soobschestva](https://random-coffee.ru/pravovaya-informaciya/oferta-soobschestva))

### Пауза

- «Если нет времени или желания, встречи можно поставить на паузу на неделю или месяц, но это время сгорает. Когда вы не пользуетесь абонементом в спортзал, он все равно продолжает действовать — так же и у нас.» ([random-coffee.ru/meetings](https://random-coffee.ru/meetings))

## @randomcoffee_official_bot (бесплатный бот для групповых чатов)

Отдельный продукт той же компании. Источник: пост аккаунта randomcoffeeru на Пикабу [«Мы запустили бесплатный Random Coffee бот для ваших групповых чатов в Телеграме»](https://pikabu.ru/story/myi_zapustili_besplatnyiy_random_coffee_bot_dlya_vashikh_gruppovyikh_chatov_v_telegrame_10017458). Точная дата на странице не показана («3 года назад»), судя по ссылкам, март 2023. Условия тарифов: оферта [oferta-chat-bot](https://random-coffee.ru/pravovaya-informaciya/oferta-chat-bot).

- «По пятницам бот спрашивает, кто пойдет на встречу — пока пары не распределены, можно поменять свой выбор в опросе.»
- «В понедельник бот перемешивает желающих рандомно […] Потом он выдает в чате сообщение, кто с кем встречается на этой неделе.» То есть пары объявляются публично в общем чате, а не в личных сообщениях, как в ботах сообществ.
- Отказ после распределения, гостинг, блокировка: в посте не описаны.

## Donut (Slack)

Источник: справочный центр [help.donut.ai](https://help.donut.ai/).

### Подтверждение и раскрытие

- Подтверждения перед знакомством нет: «On a regular basis Donut takes all the members of a specified Slack channel, puts them together into groups, and then sends a DM to each of member encouraging them to meet» ([What does Intros mean?](https://help.donut.ai/en/articles/648344-what-does-intros-mean-how-does-donut-introduce-people)).
- Пара общается в групповом DM, который создаёт бот: Donut делает чек-ины «in the multi-party direct message» ([How do I tell Donut that we met?](https://help.donut.ai/en/articles/4268294-how-do-i-tell-donut-that-we-met)).
- Шаг подтверждения участия перед каждым раундом (opt-in) в справке не найден.

### Отказ от конкретной пары

Источник: [What happens when one person isn't available to meet?](https://help.donut.ai/en/articles/4102261-what-happens-when-one-person-isn-t-available-to-meet)

- В выпадающем меню сообщения-знакомства есть «🚫 I'm unavailable». Нужно выбрать причину, и «the reasons you give are totally confidential and will not be visible to your partner». Затем «Donut will let the other person know that you're not able to meet this time around, and if possible, will try to re-match the other person with someone else.»
- Там же есть «♻️ I need a new match» («This option can only be used once per Intros round and only works with group sizes of 2»). После этого «Donut will mark the meeting as incomplete and let the other participant know». Перемэтч возможен только при встречном запросе: «You can only get matched with someone else if someone else also clicks that option and needs a new intro, so otherwise, it will be marked as unable to meet that week.»
- Точный текст уведомления, которое получает партнёр, в справке не приведён.
- После пропуска знакомства Donut предлагает больше не ставить вас в пару с этим человеком: «If you have recently skipped an Intro with someone, Donut will ask you within the Intro DM if you'd like to prevent future pairings.» ([Can I prevent intros to specific people?](https://help.donut.ai/en/articles/4260148-can-i-prevent-intros-to-specific-people))

### Гостинг

- Чек-ины: «After you're introduced, Donut will check in periodically (in the multi-party direct message) to ask whether or not you've met. At the end of the round, if you haven't responded, Donut will mark the meeting as "no response."» Отметить встречу раньше чек-ина нельзя ([How do I tell Donut that we met?](https://help.donut.ai/en/articles/4268294-how-do-i-tell-donut-that-we-met)).
- «We will check back in with each pair in a few days» ([How Donut works](https://help.donut.ai/en/articles/597029-how-donut-works)).
- Автоисключение «призраков», источник [What happens if someone is repeatedly unresponsive to Intros?](https://help.donut.ai/en/articles/3724232-what-happens-if-someone-is-repeatedly-unresponsive-to-intros):
  - «If someone is unresponsive to Donut Intros three times in a row, we'll automatically mark them as a ghost 👻 and stop including them in future Intros for that channel.»
  - Неотзывчивым считается тот, кто в трёх последних знакомствах не отметил встречу как «met», не участвовал в переписке и не нажимал кнопки в DM знакомства. В группах правило мягче: «If anyone in the group confirms the meeting or shows engagement, no one will be marked as unresponsive.»
  - Исключённому сообщают сразу, в сообщении есть кнопки «remain active in the channel, ask again later, or to exclude and stop asking». Вернуться можно кнопкой «I'm not a ghost» или «Resume Intros».
  - На платных тарифах администратор видит исключённых со статусом «Ghost busted».

### Блокировка

- «Don't pair me with...» в Home-вкладке: список действует во всех каналах, где составляются пары по двое. «Your selections are always confidential.» ([Can I prevent intros to specific people?](https://help.donut.ai/en/articles/4260148-can-i-prevent-intros-to-specific-people))
- Оговорка: «this is only for 1:1 matches - you may be matched with someone on this list as part of a group of 3 or more» ([For participants: setting intros preferences in Slack](https://help.donut.ai/en/articles/10355707-for-participants-setting-intros-preferences-in-slack)).

### Пауза

- Snooze: «Snooze All Intros Until...» или выбор следующей даты знакомства для отдельного канала. «After snoozing, you won't receive any new matches until your Intros resume, but you might get a final check-in message from Donut asking if you've had time to meet just so we can complete reporting.» Snooze личный и не меняет расписание команды ([Can I pause or snooze? (Just for me)](https://help.donut.ai/en/articles/1404870-can-i-pause-or-snooze-just-for-me)).
- В другой статье указан предел: «you can pause your Donut Intros for up to 4 cycles at a time». Кроме Snooze есть «exclude from all Intros for that channel», `/leave` из канала и удаление приложения ([How can I opt out of Donut intros?](https://help.donut.ai/en/articles/4211820-how-can-i-opt-out-of-donut-intros)).

## Tinder

Источник: [help.tinder.com](https://www.help.tinder.com/).

- Мэтч — это взаимный лайк, после него открывается чат в приложении. Контакты вне приложения не раскрываются, обмен ими остаётся делом участников (общая механика; отдельной статьи об этом не читал).
- Unmatch ([Unmatching someone](https://www.help.tinder.com/hc/en-us/articles/115003360106-Unmatching-someone), правка 2024-10-02): «When you unmatch someone, they'll disappear from your match list and you'll disappear from theirs. Unmatching is a permanent action. The person you're unmatching won't be notified of your action, but may notice you're missing from their Matches.» При анматче можно сразу заблокировать: «you both will no longer be able to view each other's profile».
- Block Profile ([Block a Profile](https://www.help.tinder.com/hc/en-us/articles/12425789312013-Block-a-Profile), правка 2026-10-22): «Block Profile ensures you won't see them and they won't see you». Действие необратимо, функция «being tested and isn't available for all users».
- Пауза ([Hide your Tinder profile](https://www.help.tinder.com/hc/en-us/articles/29635636556045-Hide-your-Tinder-profile), правка 2024-08-29): «When Discovery is off, your profile won't be shown as a recommendation to new people […] You can continue to view and chat with your existing matches.» Те, кого вы лайкнули раньше, всё ещё могут с вами сматчиться.
- Напоминаний об ответе, истечения мэтча или санкций за молчание поиск по справке Tinder не нашёл.

## Hinge

Источник: [help.hinge.co](https://help.hinge.co/).

- Unmatch ([Unmatching](https://help.hinge.co/hc/en-us/articles/36311673303571-Unmatching), правка 2026-07-21): «Unmatching is permanent; once you unmatch, the conversation will disappear, and you won't be able to message that person again.» При анматче выбирают причину («Select the Reason for unmatching»). Вместо анматча можно скрыть чат: Hide «removes a conversation from view but keeps the connection intact».
- [How do I remove an existing Match?](https://help.hinge.co/hc/en-us/articles/360011137994-How-do-I-remove-an-existing-Match) (правка 2026-05-01): «You will not see their profile again, nor will they see yours. When you Unmatch a profile from your Matches screen, you immediately disappear from that member's view, and they cannot view or retrieve the Conversation or the Match.» Уведомляют ли второго, ни одна из этих статей прямо не говорит.
- Hide ([How Can I Hide a Chat Conversation?](https://help.hinge.co/hc/en-us/articles/26848748680851-How-Can-I-Hide-a-Chat-Conversation)): «Matches that haven't had any activity in 14 days will be automatically hidden». Если скрытый мэтч напишет, чат возвращается наверх списка.
- Гостинг:
  - [What is Your Turn?](https://help.hinge.co/hc/en-us/articles/26848706496659-What-is-Your-Turn) (правка 2025-12-23): «Your Turn helps keep conversations moving by reminding you when it's your turn to reply. We've found that this helps reduce ghosting».
  - [What is "Your Turn Limits"?](https://help.hinge.co/hc/en-us/articles/31662181659027-What-is-Your-Turn-Limits) (правка 2025-12-23): «When too many people are waiting for a response, daters will not be able to send or receive new Likes until they take action with their current matches.» Выйти из ограничения можно, ответив или закрыв разговор. «Daters will not be notified if their match hits the limit.» Порог в статье не назван, есть только индикатор прогресса над списком мэтчей.
- Вопрос «встреча состоялась?» ([What is "We Met?"](https://help.hinge.co/hc/en-us/articles/360010692913-What-is-We-Met), правка 2026-01-22): «A few days after exchanging phone numbers, we'll ask you if you went on a date and if your date is the type of person you'd like to see again. If you indicate that you don't want a second date, we'll consider that info when recommending new people to you.» И далее: «Any feedback you share through 'We Met' is completely private. We will never share date feedback with existing or future matches.»
- Блок по контактам ([How can I Block People I Know?](https://help.hinge.co/hc/en-us/articles/19659017273235-How-can-I-Block-People-I-Know)): «Blocked contacts won't be notified that you blocked them». Блок не затрагивает существующие мэтчи и переписку.
- Пауза ([Pausing](https://help.hinge.co/hc/en-us/articles/360012595653-Pausing), правка 2025-12-23): «This will prevent you from being shown to new people, and you'll still be able to chat with your existing matches.» Лайк, отправленный до паузы, всё ещё может превратиться в мэтч.

## Bumble

Источник: [support.bumble.com](https://support.bumble.com/).

- Начало разговора ([How conversations start on Bumble](https://support.bumble.com/hc/en-us/articles/36679912947229-How-conversations-start-on-Bumble), правка 2026-08-10): у каждого из двоих есть одно стартовое сообщение, чат открывается после ответа. «You both have 72 hours from matching to send your opening message»; «If neither person sends a message within 72 hours of matching, the match will expire». Эта модель новая: «Opening Moves are no longer part of the experience». Более старые страницы Bumble, например [How to Extend a Match on Bumble](https://bumble.com/the-buzz/bumble-extend-match), по выдаче поиска говорят о 24 часах. Сами страницы я не перечитывал.
- Истёкший мэтч ([Expired matches](https://support.bumble.com/hc/en-us/articles/28424238819357-Expired-matches), правка 2026-08-10): у фото серый круг вместо жёлтого. «Extend: You can give a match an extra 72 hours before it expires. Everyone gets one free Extend per day». Rematch доступен на платных тарифах и перезапускает таймер.
- Unmatch ([Unmatching someone](https://support.bumble.com/hc/en-us/articles/28777122258333-Unmatching-someone), правка 2026-06-04): «Once you unmatch, the conversation will end for both of you, and the person won't be able to contact you again or view the conversation history.» И: «The other person won't get a separate notification, but they may still see in Chats that the conversation has ended».
- **Противоречие в справке Bumble.** [Matches that have disappeared](https://support.bumble.com/hc/en-us/articles/28507915973149-Matches-that-have-disappeared) (правка 2026-08-10) пишет: «If someone unmatches you, we'll let you know they ended the chat. The conversation will remain visible in your Chats, along with the time of the last message, so you can still report them». Это расходится с «won't get a separate notification» и «won't be able to […] view the conversation history» из статьи об анматче. Что на самом деле видит второй участник, по справке однозначно не установить.
- Блок ([Blocking someone](https://support.bumble.com/hc/en-us/articles/28783382700701-Blocking-someone), правка 2026-08-03): «Blocking someone removes them from your match list and prevents future contact.» «Blocking someone is private—they won't be notified. Once you've blocked someone, it's not possible to unblock them.»
- Пауза ([Taking a break from online dating](https://support.bumble.com/hc/en-us/articles/28537106111261-Taking-a-break-from-online-dating), правка 2026-09-01): Snooze Mode на 24 часа, 72 часа, неделю или бессрочно. Можно показать мэтчам статус («I'm traveling», «I'm focused on work», «I'm on a digital detox», «I'm prioritizing me») или «stay invisible». «Any matches you have without active chats will expire while you're in Snooze Mode.»

## Не удалось проверить

- **Интерфейс ботов Random Coffee после выдачи пары.** Есть ли кнопки «не смогу», «нужен новый партнёр» или «собеседник не отвечает», что в этом случае видит партнёр, как он «получает нового партнёра» в середине недели. Публичные страницы описывают только правила поведения. Бота изнутри (подписка платная) я не проверял.
- **Точная формулировка вопроса о встрече в Random Coffee** и варианты ответа. Известны только реплика с лендинга («Как прошла встреча? …») и оценки «хорошо»/«отлично»/«супер».
- **Видна ли в Random Coffee неотвеченная симпатия в Предвыборе.**
- **Посты Telegram-канала @randomcoffeeru.** Веб-версия t.me не показывает календарные даты, а поиск по каналу вернул нерелевантные посты, поэтому канал как источник не использован.
- **Другие русскоязычные «нетворкинг-боты».** Первоисточников (справки, FAQ, постов авторов) с правилами об отказе, гостинге или паузе не нашёл. Находились только кейсы студий-разработчиков и любительские репозитории на GitHub.
- **Donut**: точный текст уведомления партнёру при «I'm unavailable» и «I need a new match»; есть ли тип знакомств с подтверждением участия перед раундом.
- **Tinder**: есть ли напоминания об ответе или истечение мэтча; в справке не найдено.
- **Hinge**: уведомляется ли второй участник об анматче (справка не говорит ни да, ни нет); числовой порог «Your Turn Limits».
- **Bumble**: что именно видит человек, с которым сделали анматч (см. противоречие выше).
