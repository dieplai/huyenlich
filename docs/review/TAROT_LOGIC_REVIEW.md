# Tarot Logic Review

Ngay review: 2026-04-28

## Ket Luan Ngan

Co the chuyen sang lam giao dien prototype/MVP sau buoc nay.

Chua nen coi logic hien tai la production-ready. Phan data va rule da du tot de tao bai Tarot co can cu, nhung can them engine tach rieng, seed/replay, golden tests va eval AI truoc khi release that.

## Tieu Chuan "Dung" Cho Tarot

Tarot khong dung theo nghia du doan tuyet doi. Do dung cua san pham nen duoc danh gia bang cac tieu chuan sau:

- Khong doi nghia goc cua la Rider-Waite-Smith.
- La bai phai duoc doc theo dung orientation: upright/reversed.
- La bai phai duoc doc theo dung position trong spread.
- Cau hoi user phai anh huong den context, spread, tone va goc doc.
- Bai doc phai ket noi cac la, khong chi giai tung la roi rac.
- Combination phai dua tren rule: arcana weight, suit/element, reversal density, number pattern, court cards, adjacent relation.
- Cau hoi nhay cam phai di qua safety boundary, khong dua chan doan y te/phap ly/tai chinh.
- AI writer khong duoc rut bai, doi spread, doi ten la hoac them nghia trai data.

## Nhung Gi Da Du

### Data la bai

- Base deck co 78/78 la Rider-Waite-Smith.
- Anh local co du 78/78 la.
- V2 card meaning da cover 78/78 la.
- Major Arcana duoc viet rieng theo archetype.
- Minor Arcana duoc sinh tu rank + suit framework de giu nhat quan nghia goc, giam rui ro lech semantic.
- Moi la V2 co:
  - `core`
  - `orientation.upright`
  - `orientation.reversed`
  - context meanings: general, love, career, money, inner_work, advice, warning
  - `position_overrides`
  - `relationship_signals`
  - `symbols`
  - `reader_notes_vi`

Danh gia: du de AI writer viet bot chung chung va bam dung nghia la.

### Rule va spread

- Co `spreads.json` cho 1 la, 3 la, 5 la love/career, Celtic Cross later.
- Co `elemental-dignities.json` cho suit -> element va ma tran tuong tac.
- Co `reversal-rules.json` de doc la nguoc theo nhieu mode, khong may moc thanh nghia xau.
- Co `numerology-patterns.json` cho Minor rank, court rank va repeated-number pattern.
- Co `combination-rules.json` de dat thu tu tong hop reading.
- Co `question-contexts.json` cho classifier free-form question.
- Co `test-questions.json` gom 20 cau hoi vang de test classifier/writer.

Danh gia: rule nen tang da dung huong va du tot cho MVP.

### Prototype flow

Flow hien tai:

```text
User question
-> AI classifier json_schema
-> choose spread
-> draw cards
-> local rule analysis
-> AI writer json_schema
-> fallback local reading neu AI loi
```

Danh gia: flow dung kien truc san pham. AI phan tich cau hoi va viet bai, nhung engine van giu quyen rut bai/phan tich rule.

## Nhung Gi Chua Du

### 1. Engine van nam trong HTML

Hien tai reading logic nam trong `prototype/tarot-flow.html`.

Rui ro:

- Kho test tu dong.
- Kho reuse cho app that.
- UI co the vo tinh lam thay doi logic.

Can lam truoc production:

- Tach thanh module rieng, vi du `src/domain/tarot`.
- API can nhan input va tra output theo contract co schema.

### 2. Chua replay duoc reading

Hien tai draw card dung random trong browser. Reversal dung `Math.random()`.

Rui ro:

- User refresh se mat ket qua.
- Khong co `reading_id` de xem lai.
- Kho debug bai doc bi sai.

Can lam:

- Moi reading can co `reading_id`, `seed`, `spread_id`, `cards`, `orientation`, `classification`, `created_at`.
- Auto draw phai deterministic theo seed.
- Neu user tu chon bai, backend luu card ids va orientation.

### 3. Combination rules moi duoc dung mot phan

Prototype da dung:

- major/minor count
- reversal count
- court count
- suit dominance
- adjacent elemental relation
- repeated minor numbers

Prototype chua dung day du:

- missing suit.
- sequence rules tang/giam so.
- triad elemental dignity voi la trung tam.
- court card in advice position.
- card-pair special cases.
- weighting/priority khi nhieu pattern xung dot.

Danh gia: du cho MVP 3 la, nhung can nang cap neu muon bai doc sau va it lap.

### 4. Chua co eval AI tu dong

Da co golden questions, nhung chua co script goi classifier/writer va cham diem.

Can lam:

- Chay 20 cau hoi trong `data/tarot/test-questions.json`.
- So sanh output classifier voi expected routing.
- Luu output writer de review:
  - co bam cau hoi khong?
  - co dung card/position khong?
  - co noi chung chung khong?
  - co vi pham safety khong?

### 5. Safety moi la prompt/routing, chua thanh policy layer rieng

Da co safety flags va prompt rule, nhung production nen co policy layer rieng truoc writer.

Can lam:

- Neu `self_harm`: khong viet nhu boi bai binh thuong; chuyen sang support/safety response.
- Neu `medical`, `pregnancy`, `legal`, `financial_investment`: chi soi chieu tam the/rui ro/cau hoi can hoi chuyen gia.
- Khong cho writer tao cam giac dam bao tuyet doi.

### 6. Asset chua toi uu

Anh JPG goc du de prototype. UI san pham nen co:

- WebP.
- Thumbnail.
- Card back.
- Asset manifest co width/height/hash.

## So La Nen Dung Cho MVP

Khuyen nghi:

- Free quick reading: 1 la.
- Main MVP reading: 3 la `situation_action_outcome`.
- Love paid/richer: 5 la `love_five`.
- Career paid/richer: 5 la `career_five`.
- Celtic Cross: de later/VIP, chua nen dua vao UI dau.

Ly do:

- 1 la de nhanh, de giu user moi.
- 3 la du de co story va action, khong qua dai.
- 5 la chi dung khi user da co intent ro hoac paid.
- 10 la rat de lam user moi ngop va ton token.

## Readiness Matrix

| Hang muc | Trang thai | Ghi chu |
|---|---:|---|
| 78 card data | Ready | V1 + V2 du 78/78 |
| Card images | Ready for prototype | Can WebP/thumb cho production |
| Source-backed docs | Ready | Da co source map va bibliography |
| Question classification | Ready for prototype | Can eval 20 golden questions |
| Spread selection | Ready for prototype | Nen mac dinh 3 la |
| Rule analysis | MVP-ready | Chua full data-driven |
| AI writer prompt | MVP-ready | Can eval de giam chung chung |
| Safety | Partial | Can policy layer rieng |
| Replay/persistence | Missing | Bat buoc truoc production |
| Engine tests | Missing | Bat buoc truoc production |
| UI readiness | Yes, for prototype | Nen build UI tren contract on dinh |

## Recommend Buoc Tiep Theo

Thu tu nen lam:

1. Lam giao dien prototype tren flow hien tai de test cam giac user.
2. Khi UI shape on, tach reading engine ra module/app service.
3. Them seed/replay va `reading_id`.
4. Viet eval script cho 20 golden questions.
5. Nang combination logic: missing suit, triad, sequence, card-pair special cases.
6. Toi uu asset sang WebP/thumb.

Quyet dinh cuoi: du dieu kien de chuyen qua UI, mien la coi UI tiep theo la prototype/MVP validation, khong phai ban production release.
