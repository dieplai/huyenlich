# Tarot Rules Bibliography

Tai lieu nay tra loi cau hoi: neu Tarot muon co "data quy luat that" cho Tarot thi nen dua vao sach/nguon nao, va moi nguon dung de tao loai rule nao.

## Ket Luan Thang Than

Tarot khong co mot tieu chuan duy nhat nhu ISO hay cong thuc toan. "Dung" trong Tarot nghia la:

- Chon mot tradition ro rang.
- Trung thanh voi tradition do.
- Co nguon cho moi rule chinh.
- Khong de AI tu bia them rule khong co trong source map.

Voi Tarot, nen chon:

- Deck/tradition: Rider-Waite-Smith.
- Nghia la goc: A. E. Waite.
- He quy tac element/correspondence: Golden Dawn / Book T, dung co chon loc vi RWS co lien he voi Golden Dawn nhung Waite khong cong khai toan bo he GD.
- Reading modern: Joan Bunning, Mary K. Greer, Rachel Pollack, Benebell Wen.
- History/symbolism: Robert M. Place, V&A, Britannica.

## Core Sources

### 1. A. E. Waite - The Pictorial Key to the Tarot

Loai nguon: primary/public domain cho Rider-Waite-Smith.

Dung de tao:

- card meanings goc.
- upright/reversed divinatory meanings.
- Greater vs Lesser Arcana.
- Celtic Cross style method.
- mot so rule ve recurrence/card dealing.
- symbolism cua Major Arcana theo Waite.

Links:

- Wikisource: https://en.wikisource.org/wiki/The_Pictorial_Key_to_the_Tarot
- Sacred Texts table of contents: https://sacred-texts.com/tarot/pkt/pktcont.htm

Can luu y:

- Van phong Waite co tinh co dien, kho hieu, doi khi mau thuan.
- Khong nen copy nguyen van dai. Nen paraphrase thanh tieng Viet hien dai.
- Waite la goc RWS nhung khong du de tao reading engine hay; can them framework doc spread/combinations.

Rule/data nen lay:

```text
card_meanings.source_basis = "waite"
major_symbolism.source_basis = "waite"
celtic_cross.source_basis = "waite"
```

### 2. Golden Dawn - Book T / Mathers-Felkin Tarot Manuscripts

Loai nguon: esoteric source cho correspondences va elemental dignities.

Dung de tao:

- elemental dignities.
- well-dignified / ill-dignified logic.
- court card elemental structure.
- astrological/qabalistic correspondences neu sau nay can.
- rule adjacent cards trong triad/pair.

Links:

- Book T online text mirror: https://www.mr-kaplan.com/Library/Mathers/Book-T.html
- Book T reference page: https://mytarot.org/2022/01/24/book-t-the-tarot-by-samuel-liddell-macgregor-mathers-and-harriet-felkin-1888-members-of-the-golden-dawn/
- Mary K. Greer explanation: https://marykgreer.com/2008/02/05/24/

Can luu y:

- Book T khong phai "user-friendly"; no la tai lieu he thong Huyen hoc.
- RWS chiu anh huong Golden Dawn, nhung Waite co dieu chinh rieng.
- Nen dung GD cho engine weighting, khong bat buoc phoi het astrology/qabalah vao MVP.

Rule/data nen lay:

```text
elemental_dignities.source_basis = "golden_dawn_book_t"
court_card_elements.source_basis = "golden_dawn_book_t"
adjacent_weighting.source_basis = "golden_dawn_book_t"
```

Canonical elemental matrix cho Tarot:

```text
same element -> strong / amplify
Fire + Air -> friendly / support
Water + Earth -> friendly / support
Fire + Earth -> somewhat friendly / neutral
Air + Water -> somewhat friendly / neutral
Fire + Water -> contrary / weaken
Air + Earth -> contrary / weaken
```

Triad rule nen dung:

```text
left card + center card + right card
center = main voice
left/right = modifiers
if modifiers support center -> center amplified
if one modifier contrary and one supportive/neutral -> center modified but still fairly strong
if both modifiers contrary -> center weakened; warning/conflict emphasis
if left and right contrary to each other -> read center more strongly
```

Khong nen "bo qua" la yeu trong UX. Nen dien giai la nang luong yeu/bi nhieu/bi chan.

### 3. Mary K. Greer - The Complete Book of Tarot Reversals

Loai nguon: sach uy tin ve reversed cards.

Dung de tao:

- reversal modes.
- reversed interpretation cho 78 la.
- cach doc upside-down khong may moc la "nghia xau".
- lien he reversal voi learning opportunity, inner support, blocked energy.

Links:

- Google Books: https://books.google.com/books/about/The_Complete_Book_of_Tarot_Reversals.html?id=0QnoqX0TOS0C
- Llewellyn page: https://www.llewellyn.com/product.php?ean=9781567182859

Can luu y:

- Noi dung co ban quyen. Khong copy card interpretations.
- Dung de xay schema reversal modes va tu viet lai meaning.

Rule/data nen lay:

```text
reversal_modes.source_basis = "mary_greer_reversals"
```

Canonical reversal modes cho Tarot:

```text
blocked
internalized
delayed
shadow
overexpressed
weak_or_undeveloped
```

### 4. Joan Bunning - Learning the Tarot

Loai nguon: giao trinh thuc hanh RWS cho nguoi hoc.

Dung de tao:

- card pairs.
- story of reading.
- keyword/action phrase style.
- similar/opposite card relations.
- beginner-friendly reading procedure.

Links:

- Google Play Books: https://play.google.com/store/books/details/Learning_the_Tarot_A_Tarot_Book_for_Beginners?id=GtLPNvengDoC
- Google Books: https://books.google.com/books/about/Learning_the_Tarot.html?id=he3J6UYYPakC

Can luu y:

- Bunning noi ro day la guidelines hon la rigid rules.
- Rat phu hop voi UX san pham vi de hieu, it qua nang occult.

Rule/data nen lay:

```text
keyword_synthesis.source_basis = "joan_bunning_learning_tarot"
similar_opposite_cards.source_basis = "joan_bunning_learning_tarot"
story_reading.source_basis = "joan_bunning_learning_tarot"
```

### 5. Joan Bunning - Learning Tarot Spreads

Loai nguon: spread design va position meanings.

Dung de tao:

- spread catalog.
- position meanings.
- spread shapes.
- subject-specific spreads: relationship, love, money, work, health, time period.
- quick three-card draws.

Links:

- Google Books: https://books.google.com/books/about/Learning_Tarot_Spreads.html?id=nXyZnU6MA_wC
- Red Wheel/Weiser: https://redwheelweiser.com/book/learning-tarot-spreads-9781578632701/

Rule/data nen lay:

```text
spreads.source_basis = "joan_bunning_learning_tarot_spreads"
position_meanings.source_basis = "joan_bunning_learning_tarot_spreads"
```

### 6. Rachel Pollack - Seventy-Eight Degrees of Wisdom

Loai nguon: modern classic ve symbolism/psychological Tarot.

Dung de tao:

- deeper symbolic interpretation.
- psychological/archetypal lens.
- Major Arcana journey.
- modern tone for paid reading.

Links:

- Google Books: https://books.google.com/books/about/Seventy_Eight_Degrees_of_Wisdom.html?id=8Js0NMhA6B0C
- Publisher/bookstore summary: https://www.labyrinthbooks.com/seventy-eight-degrees-of-wisdom/

Can luu y:

- Rat tot cho "hay", khong phai rule engine thuan.
- Khong copy wording; dung de guide editorial tone.

Rule/data nen lay:

```text
archetype_notes.source_basis = "rachel_pollack"
psychological_lens.source_basis = "rachel_pollack"
```

### 7. Benebell Wen - Holistic Tarot

Loai nguon: compendium hien dai, rat thuc dung cho system design.

Dung de tao:

- RWS-specific study framework.
- beginner/intermediate/advanced learning path.
- reading procedure.
- spread practice.
- court cards.
- correspondences, numerology, elements, journal templates.

Links:

- Author page: https://benebellwen.com/about-the-book/
- Study guides: https://benebellwen.com/about-the-book/holistic-tarot-supplements/
- Google Books: https://books.google.com/books?id=ZwcDBAAAQBAJ

Can luu y:

- Book rat lon, nhieu schema phu hop voi product data.
- Tac gia ghi ro study guide keyed to Rider-Waite-Smith.
- Khong phai tradition co xua duy nhat; la modern integrative system.

Rule/data nen lay:

```text
rws_study_path.source_basis = "benebell_holistic_tarot"
advanced_spread_practice.source_basis = "benebell_holistic_tarot"
```

### 8. Robert M. Place - The Tarot: History, Symbolism, and Divination

Loai nguon: history/symbolism, giam rui ro bia lich su.

Dung de tao:

- historical notes.
- symbol explanation co can cu nghe thuat/lich su.
- three-card reading approach.
- myth-busting: khong noi tarot bat nguon tu Ai Cap neu khong co co so.

Links:

- Google Books: https://books.google.com/books/about/The_Tarot.html?id=KCiQEAAAQBAJ
- Review/reference: https://www.aeclectic.net/tarot/books/history-symbolism-divination/

Rule/data nen lay:

```text
history_notes.source_basis = "robert_place"
symbol_history.source_basis = "robert_place"
```

## Data Rule Files Nen Tao

### 1. `data/tarot/source-map.json`

Muc dich: moi rule biet den tu dau.

```json
{
  "waite": {
    "title": "The Pictorial Key to the Tarot",
    "author": "A. E. Waite",
    "use_for": ["card_meanings", "rws_symbolism", "celtic_cross"]
  },
  "golden_dawn_book_t": {
    "title": "Book T - The Tarot",
    "use_for": ["elemental_dignities", "court_card_elements", "adjacent_weighting"]
  }
}
```

### 2. `data/tarot/elemental-dignities.json`

Muc dich: rule element tu Golden Dawn/Book T.

```json
{
  "elements": {
    "wands": "fire",
    "cups": "water",
    "swords": "air",
    "pentacles": "earth"
  },
  "matrix": {
    "fire.fire": "strong",
    "fire.air": "friendly",
    "fire.earth": "neutral",
    "fire.water": "contrary"
  }
}
```

### 3. `data/tarot/spreads.json`

Muc dich: spread va position meanings.

V1 nen co:

- one_card.
- three_card_past_present_future.
- three_card_situation_action_outcome.
- five_card_love.
- five_card_career.
- celtic_cross.

### 4. `data/tarot/combination-rules.json`

Muc dich: rule ket hop khong phai 3003 cap card thu cong.

Categories:

- keyword synthesis.
- number pattern.
- element interaction.
- arcana weight.
- court-card role.
- suit dominance.
- reversal density.
- adjacent flow.

### 5. `data/tarot/reversal-rules.json`

Muc dich: reversal khong bi doc may moc.

Modes:

- blocked.
- internalized.
- delayed.
- shadow.
- overexpressed.
- weak_or_undeveloped.

### 6. `data/tarot/numerology-patterns.json`

Muc dich: pattern theo so trong Minor/Major.

```text
Ace/1 = beginning
2 = polarity/choice/relationship
3 = growth/expression/collaboration
4 = structure/stability
5 = disruption/conflict
6 = harmony/recovery
7 = test/assessment/inner challenge
8 = power/movement/mastery
9 = culmination/solitude/depth
10 = completion/transition
```

## Nen Mua/Doc Sach Nao Truoc

Neu chi chon 5 cuon/nguon de xay product:

1. A. E. Waite - The Pictorial Key to the Tarot.
2. Book T / Golden Dawn material, kem Mary K. Greer giai thich elemental dignities.
3. Mary K. Greer - The Complete Book of Tarot Reversals.
4. Joan Bunning - Learning the Tarot + Learning Tarot Spreads.
5. Benebell Wen - Holistic Tarot.

Neu muon tang do sau/noi dung paid:

6. Rachel Pollack - Seventy-Eight Degrees of Wisdom.
7. Robert M. Place - The Tarot: History, Symbolism, and Divination.

## Chien Luoc San Pham

V1 khong nen co gang lam "AI tarot reader tu do". Nen lam "rules-first reader":

1. Backend rut bai va orientation.
2. Engine tinh spread pattern.
3. Engine tinh elemental dignity/adjacent relation.
4. Engine chon meaning snippets theo question context.
5. AI chi viet thanh bai doc tieng Viet hay dua tren output co cau truc.

Neu lam nhu vay, ta khong tu bia. Ta dang bien cac rule duoc cong nhan thanh structured data va prompt context.

