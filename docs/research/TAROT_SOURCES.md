# Tarot Research Sources

Tai lieu nay ghi lai he nguon nen dung khi phat trien Tarot cho Tarot.

## Ket Luan Research

Tarot khong co mot co quan "chinh chu" duy nhat giong tieu chuan ky thuat. Cach lam dung cho san pham la chon mot he quy chieu, cong khai he do, va giu tat ca data/prompt/logic nhat quan.

Voi Tarot, he chuan nen la:

- Deck: Rider-Waite-Smith.
- Nghia goc: A. E. Waite, `The Pictorial Key to the Tarot`.
- Cach doc hien dai: position trong spread, orientation/reversal, adjacent cards, suit/element balance, number patterns, symbolism, court-card roles.
- Tone san pham: self-reflection, guidance, khong phan quyet tuong lai tuyet doi.

## Source Hierarchy

### Tier 1: Primary / Public Domain

Dung de neo nghia goc va symbol RWS.

1. The Pictorial Key to the Tarot - Wikisource
   - URL: https://en.wikisource.org/wiki/The_Pictorial_Key_to_the_Tarot
   - Tac gia: Arthur Edward Waite.
   - Minh hoa: Pamela Colman Smith.
   - Gia tri: nguon goc cho meaning, symbolism, Greater/Lesser Arcana va mot so phuong phap trai bai.
   - Ghi chu: public domain o US; text tren Wikisource co dieu khoan rieng cua Wikisource khi dung ban transcription.

2. The Pictorial Key to the Tarot PDF - Wikisource
   - URL: https://en.wikisource.org/wiki/File:The_Pictorial_Key_to_the_Tarot.pdf
   - Gia tri: scan/pdf de doi chieu khi can can cu nguyen ban.

3. Rider-Waite-Smith tarot deck image set - Wikimedia Commons
   - URL: https://commons.wikimedia.org/wiki/Category:Rider-Waite-Smith_tarot_deck_(TaionWC)
   - Gia tri: bo anh 78 la da dung trong `data/tarot/rider-waite-smith.json`.

### Tier 2: Institutional / History

Dung de kiem chung cau truc deck, lich su va boi canh.

1. Encyclopaedia Britannica - Tarot
   - URL: https://www.britannica.com/topic/tarot
   - Gia tri: cau truc 78 la, 22 Major, 56 Minor, 4 suit, vai tro spread, reversal va adjacent cards.

2. Victoria and Albert Museum - A history of tarot cards
   - URL: https://www.vam.ac.uk/articles/tarot-cards
   - Gia tri: lich su tarot, RWS, suit/element correspondences, boi canh tu game bai den divination.

3. The Metropolitan Museum of Art - It's in the Cards
   - URL: https://www.metmuseum.org/en/perspectives/tarot
   - Gia tri: boi canh nghe thuat va cach tarot duoc tai dien giai qua cac deck.

4. Johns Hopkins Libraries - Rider-Waite-Smith Pam-A deck
   - URL: https://aspace.library.jhu.edu/repositories/3/archival_objects/330032
   - Gia tri: xac nhan deck Pam-A hoan chinh, 78 la, chia 56 Minor va 22 Major.

### Tier 3: Practice / Reading Methods

Dung de thiet ke reading engine va docs ket hop la. Khong copy noi dung dai; chi rut framework.

1. Biddy Tarot - Ultimate Guide to Tarot Card Combinations
   - URL: https://biddytarot.com/blog/ultimate-guide-tarot-card-combinations/
   - Gia tri: phuong phap ket hop bang keywords, numerology, symbolism, elements, Majors/Minors.

2. Labyrinthos - Tarot Elements and Elemental Dignities
   - URL: https://labyrinthos.co/blogs/learn-tarot-with-labyrinthos-academy/tarot-elements-correspondences-and-working-with-elemental-dignities
   - Gia tri: suit -> element, polarity, elemental interaction rules.

3. Mary K. Greer - Elemental Dignities
   - URL: https://marykgreer.com/2008/02/05/24/
   - Gia tri: goc Golden Dawn/Book T cua elemental dignities, cach dung de xem la nao manh/yeu trong spread.

4. Joan Bunning - Learning the Tarot
   - URL: https://books.google.com/books/about/Learning_the_Tarot.html?id=he3J6UYYPakC
   - Gia tri: phuong phap hoc card meanings, card pairs, story of reading. Nen tham khao nhu sach thuc hanh uy tin, khong copy noi dung.

5. Aeclectic Tarot - Minor Arcana meanings
   - URL: https://www.aeclectic.net/tarot/tarot-card-meanings/minor-arcana/
   - Gia tri: tong quan suit meanings va cac ten suit khac nhau; huu ich de doi chieu.

## He Chuan Da Chon Cho Tarot

### Deck

- Rider-Waite-Smith.
- Khong tron voi Thoth hoac Marseille trong V1.
- Neu sau nay ho tro deck khac, moi deck phai co data rieng.

### Suit/Element

- Wands = Fire.
- Cups = Water.
- Swords = Air.
- Pentacles = Earth.

Day la mapping pho bien cho RWS hien dai va duoc V&A, Labyrinthos, Biddy Tarot cung nhac theo huong nay.

### Major/Minor

- Major Arcana: chu de lon, archetype, turning point, bai hoc mang tinh doi song/noi tam.
- Minor Arcana: su kien doi thuong, cam xuc, suy nghi, tien bac, cong viec, hanh dong cu the.

### Reversal

Reversal khong nen chi la "nghia nguoc". Trong product, reversal nen duoc doc theo cac mode:

- blocked: nang luong bi chan.
- internalized: nang luong dang dien ra ben trong.
- delayed: nang luong cham/tri hoan.
- shadow: mat bong toi cua la.
- overexpressed: nang luong qua da.

Moi card V2 nen co `reversal_modes` de AI khong viet may moc.

### Adjacent Cards

Adjacent cards co the:

- amplify: tang manh cung chu de.
- soften: lam mem/y neu co element ho tro.
- conflict: tao xung dot neu element/meaning trai nhau.
- clarify: la sau lam ro la truoc.
- sequence: tao cau chuyen theo thu tu trai -> phai hoac position order.

Britannica ghi nhan y nghia la bi sua doi boi orientation, position trong spread va adjacent cards. Day la nen tang de chung ta xay reading engine, khong chi hien tung la doc lap.

## Source Usage Rules

- Duoc paraphrase va tong hop.
- Khong copy dai noi dung cua website hien dai.
- Voi Waite/Wikisource public domain, van nen uu tien dien giai lai bang tieng Viet cua Tarot.
- Moi field trong data V2 nen co `source_basis`: `waite`, `rws_symbolism`, `modern_practice`, `product_editorial`.
- Neu noi dung la sang tao san pham, danh dau `product_editorial` de khong nham la nghia goc.

## Bibliography Cho Rule Data

Chi tiet sach/nguon nen dung de tao "data quy luat that" nam o:

- `docs/research/TAROT_RULES_BIBLIOGRAPHY.md`

## Data Files Nen Tao Sau Research

- `data/tarot/sources.json`
- `data/tarot/spreads.json`
- `data/tarot/combination-rules.json`
- `data/tarot/elemental-dignities.json`
- `data/tarot/numerology-patterns.json`
- `data/tarot/rider-waite-smith.v2.json`

## Nhung Diem Can Canh Giac

- Tarot co nhieu tradition; "dung" nghia la dung voi tradition da chon.
- Justice/Strength order co khac giua deck; Tarot theo RWS: Strength = 8, Justice = 11.
- Mot so deck doi Wands/Swords element; Tarot khong lam vay trong V1.
- Court cards co the doc la nguoi, vai tro, nang luong, hoac cach hanh dong; khong mac dinh la gioi tinh sinh hoc.
- AI khong duoc bien bai thanh loi khuyen y te, phap ly, tai chinh chac chan.

