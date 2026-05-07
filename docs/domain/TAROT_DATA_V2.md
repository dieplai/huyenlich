# Tarot Data V2

Tai lieu nay mo ta cach nang data nghia la Tarot tu seed data len ban du sau cho AI writer.

## Muc Tieu

Data V1 cho biet la bai la gi va y nghia ngan. Data V2 phai giup bai doc bot chung chung bang cach cung cap:

- nghia theo context: general, love, career, money, inner_work.
- nghia theo upright/reversed.
- override theo position: situation, action, outcome, challenge.
- relationship signals theo giai doan user suy ra tu cau hoi.
- symbols va reader notes.
- source_basis cho moi la.

## Files

- `data/tarot/rider-waite-smith.v2.schema.json`: schema chuan.
- `data/tarot/rider-waite-smith.v2.sample.json`: 5 Major dau tien.
- `data/tarot/rider-waite-smith.v2.major-rest.json`: 17 Major con lai.
- `data/tarot/rider-waite-smith.v2.wands.json`: 14 la Wands.
- `data/tarot/rider-waite-smith.v2.cups.json`: 14 la Cups.
- `data/tarot/rider-waite-smith.v2.swords.json`: 14 la Swords.
- `data/tarot/rider-waite-smith.v2.pentacles.json`: 14 la Pentacles.

Coverage hien co:

- 78/78 la.
- Major duoc viet rieng theo tung archetype.
- Minor duoc sinh tu framework rank + suit de giu nhat quan nghia goc va giam rui ro lech semantic.

## Cach Prototype Dung V2

`prototype/tarot-flow.html` load optional files:

```text
data/tarot/rider-waite-smith.v2.sample.json
data/tarot/rider-waite-smith.v2.major-rest.json
data/tarot/rider-waite-smith.v2.wands.json
data/tarot/rider-waite-smith.v2.cups.json
data/tarot/rider-waite-smith.v2.swords.json
data/tarot/rider-waite-smith.v2.pentacles.json
```

Neu la duoc rut co V2 context, prompt writer se nhan them:

- `essence_vi`
- `upright_context_vi`
- `reversed_context_vi`
- `position_overrides`
- `relationship_signal_vi`
- `symbols`
- `reader_notes_vi`

Neu la chua co V2, prototype fallback ve V1 meaning.

## Chien Luoc Mo Rong

Khong nen viet 78 la cung luc. Thu tu:

1. Test 20 cau hoi that voi 78 la V2.
2. Sua schema/prompt neu bai doc con chung chung.
3. Them card-pair special cases neu can.
4. Viet reading engine tach khoi prototype HTML.

## Tieu Chuan Chat Luong Moi La

Moi la V2 phai tra loi duoc:

- La nay noi gi khi user hoi tinh yeu?
- La nay noi gi khi user hoi cong viec?
- La nay canh bao dieu gi?
- Neu la nay nam o action, user nen lam gi?
- Neu la nay nam o outcome, dau hieu gan toi la gi?
- Neu user hoi co nen tiep tuc moi quan he, la nay doc ra sao?

Neu mot la chi co nghia tong quat, AI writer se lai viet chung chung.
