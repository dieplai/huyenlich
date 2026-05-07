# Tarot Domain Context

Tai lieu nay la ngu canh chuan cho phan Tarot cua Tarot. Khi phat trien backend, frontend, prompt AI, SEO hoac noi dung paid/free, uu tien bam file nay truoc khi mo rong.

## Muc Tieu Chat Luong

Tarot trong Tarot can dat 3 muc tieu:

- Dung he quy chieu: bam Rider-Waite-Smith, khong tron nghia lung tung tu nhieu deck.
- Hay ve dien dat: tieng Viet am, sau, co tinh goi mo nhung khong noi vo can cu.
- Dung trong san pham: moi ket qua phai co cau truc ro, ho tro free preview, paid reading va AI personalization.

Tarot khong duoc trinh bay nhu mot ket luan tuyet doi ve tuong lai. Cach viet nen la "nang luong/hinh mau/tinh huong co the dang dien ra", tap trung vao phan chieu, lua chon va hanh dong.

## Chuan Deck

Deck mac dinh:

- Ten: Rider-Waite-Smith Tarot.
- Artist: Pamela Colman Smith.
- He nghia: Rider-Waite-Smith / A. E. Waite.
- So la: 78.
- Major Arcana: 22 la.
- Minor Arcana: 56 la, gom Wands, Cups, Swords, Pentacles.

Nguon anh hien tai:

- `data/tarot/rider-waite-smith.json`
- `public/textures/cards/rws/*.jpg`
- Source category: https://commons.wikimedia.org/wiki/Category:Rider-Waite-Smith_tarot_deck_(TaionWC)

Nguon nghia nen tham chieu:

- The Pictorial Key to the Tarot, A. E. Waite.
- Wikisource: https://en.wikisource.org/wiki/The_Pictorial_Key_to_the_Tarot
- Sacred Texts mirror: https://archive.sacred-texts.com/tarot/pkt/

Ghi chu ban quyen: Rider-Waite goc duoc Wikimedia Commons ghi nhan public domain o US/UK, nhung mot so ban to mau, retouch hoac deck phai sinh co the van con ban quyen. Production nen dung asset local cua minh, khong hotlink.

## Data Hien Co

File JSON hien tai co:

- `schema_version`
- deck metadata
- 3 spread co ban
- 78 card records
- moi la co:
  - `id`
  - `deck_index`
  - `name_en`
  - `name_vi`
  - `arcana`
  - `number`
  - `suit`
  - `rank`
  - `element` cho Minor Arcana
  - `keywords_vi`
  - `upright_vi`
  - `reversed_vi`
  - `image`

Day la seed data du de lam MVP rut bai, hien anh, hien nghia ngan va tao prompt AI co can cu. Chua du de lam paid reading chat luong cao.

## Research Va Reading Framework

Tai lieu lien quan:

- `docs/research/TAROT_SOURCES.md`: he nguon da research, source hierarchy va chuan Tarot da chon.
- `docs/domain/TAROT_READING_FRAMEWORK.md`: cach ket hop la, adjacent cards, elemental dignities, numerology patterns, spread-level synthesis.
- `docs/research/TAROT_RULES_BIBLIOGRAPHY.md`: sach/nguon dung de tao rule data.
- `docs/domain/TAROT_QUESTION_AGENT.md`: cach dung AI de classify cau hoi tu do cua user.
- `docs/domain/TAROT_DATA_V2.md`: schema va chien luoc mo rong data nghia la sau.

Khi co xung dot giua y nghia rieng tung la va logic ket hop, uu tien thu tu:

```text
question -> spread position -> card core -> orientation -> adjacent relation -> spread pattern -> synthesis
```

Rule data hien co:

- `data/tarot/source-map.json`
- `data/tarot/elemental-dignities.json`
- `data/tarot/spreads.json`
- `data/tarot/reversal-rules.json`
- `data/tarot/numerology-patterns.json`
- `data/tarot/combination-rules.json`
- `data/tarot/question-contexts.json`
- `data/tarot/rider-waite-smith.v2.schema.json`
- `data/tarot/rider-waite-smith.v2.sample.json`
- `data/tarot/rider-waite-smith.v2.major-rest.json`
- `data/tarot/rider-waite-smith.v2.wands.json`
- `data/tarot/rider-waite-smith.v2.cups.json`
- `data/tarot/rider-waite-smith.v2.swords.json`
- `data/tarot/rider-waite-smith.v2.pentacles.json`

## Data Can Nang Cap Len V2

Moi la nen co them cac field sau:

```json
{
  "core_meaning_vi": "",
  "symbolism_vi": [
    {
      "symbol": "",
      "meaning": ""
    }
  ],
  "contexts": {
    "general": {
      "upright": "",
      "reversed": ""
    },
    "love": {
      "upright": "",
      "reversed": ""
    },
    "career": {
      "upright": "",
      "reversed": ""
    },
    "money": {
      "upright": "",
      "reversed": ""
    },
    "inner_work": {
      "upright": "",
      "reversed": ""
    },
    "advice": {
      "upright": "",
      "reversed": ""
    },
    "warning": {
      "upright": "",
      "reversed": ""
    }
  },
  "yes_no": {
    "upright": "yes|no|maybe",
    "reversed": "yes|no|maybe",
    "note_vi": ""
  },
  "timing": {
    "season": "",
    "pace": "",
    "note_vi": ""
  },
  "as_person_vi": "",
  "prompt_notes_vi": []
}
```

## Spread Chuan

V1 nen ho tro cac spread sau:

- `one_card`: mot thong diep chinh.
- `past_present_future`: qua khu, hien tai, tuong lai.
- `situation_action_outcome`: tinh huong, hanh dong nen chon, ket qua tiem nang.
- `love_five`: ban, nguoi kia, dong luc ket noi, thu thach, huong phat trien.
- `career_five`: hien trang, the manh, diem can sua, co hoi, buoc tiep theo.

Moi position can co prompt context rieng. Cung mot la nhung o vi tri "qua khu" khac voi "hanh dong nen chon".

## Combination Rules

Khong can viet het 78 x 78 combinations o giai doan dau. Nen co rule tong hop:

- Nhieu Major Arcana: chu de co tinh buoc ngoat, bai hoc lon, kho kiem soat bang y chi ngan han.
- Nhieu Wands: hanh dong, dam me, toc do, sang tao, xung luc.
- Nhieu Cups: cam xuc, tinh yeu, ket noi, chua lanh.
- Nhieu Swords: suy nghi, xung dot, quyet dinh, su that, cang thang tinh than.
- Nhieu Pentacles: cong viec, tien bac, than the, nen tang, thoi gian dai.
- Nhieu court cards: con nguoi, vai tro xa hoi, mau tinh cach dang tac dong.
- Nhieu la reversed: nang luong bi chan, noi tam chua xu ly, hoac can dieu chinh cach tiep can.
- Nhieu so 5: bien dong va xung dot.
- Nhieu so 10: ket thuc chu ky.
- Nhieu Ace: khoi dau moi.

AI reading phai dua ra tong hop spread, khong chi ghep nghia tung la roi ket thuc.

## Chat Luong Noi Dung

Mot bai Tarot tot nen co:

- Mo dau gan voi cau hoi nguoi dung.
- Tom tat nang luong tong the cua spread.
- Giai tung la theo position.
- Ket noi cac la voi nhau.
- Neu co reversed, giai thich la nang luong bi chan/lech, khong chi viet nguoc nghia.
- Loi khuyen hanh dong cu the.
- Mot cau ket co tinh goi mo, khong phan quyet.

Can tranh:

- Noi chac chan ve tai nan, chet choc, benh tat, mang thai, ly hon, dau tu loi lo.
- Dung ngon ngu doa so de ep mua paid.
- Viet chung chung kieu "ban dang co nhieu cam xuc".
- De AI tu sang tac nghia la trai voi data.

## Free Vs Paid

Free:

- 1 la hoac tom tat ngan cua spread.
- Keyword, nghia ngan, mot loi khuyen.
- Khong qua 20-25% tong gia tri noi dung.

Paid:

- Day du tat ca la.
- Dien giai theo cau hoi.
- Tong hop pattern cua spread.
- Loi khuyen hanh dong.
- Giai thich symbol noi bat.
- Co the xuat PDF.

## Asset Pipeline

Da co anh `.jpg` goc local. Can them:

- resize ban hien thi: 400 x 700 hoac 420 x 720.
- `.webp` cho web.
- thumbnail: 160 x 280.
- card back texture.
- blur placeholder hoac dominant color.
- manifest hash/size de QA.

Khong nen de frontend tai file JPG goc dung luong lon neu render 78 la cung luc.

## QA Bat Buoc

Can co script/test kiem tra:

- du 78 la.
- du 22 Major, 56 Minor.
- moi suit du 14 la.
- moi la co anh local.
- moi la co `upright` va `reversed`.
- khong trung `deck_index`.
- khong trung `id`.
- source image ton tai trong manifest.

## Uu Tien Phat Trien Tarot

1. Tao Data V2 cho 78 la, them context theo tinh yeu, su nghiep, tien bac, loi khuyen.
2. Tao `tarot-spreads.json` rieng cho spread positions.
3. Tao `tarot-combination-rules.json`.
4. Tao script resize/convert anh sang webp/thumbnail.
5. Scaffold frontend Tarot page va scene don gian.
6. Tao backend/service rut bai deterministic theo seed/session.
7. Tao AI prompt builder dua tren data, khong de AI tu bịa he nghia.

