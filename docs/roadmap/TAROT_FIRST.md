# Tarot First Roadmap

## Current Status 2026-04-28

Da hoan thanh:

- Base deck 78/78 la va anh local.
- V2 meaning cover 78/78 la.
- Rule data cho spreads, elemental dignity, reversal, numerology pattern, combination, question contexts.
- Prototype HTML co flow: user question -> AI classifier -> local engine -> AI writer.
- Golden question set: `data/tarot/test-questions.json`.
- Logic review: `docs/review/TAROT_LOGIC_REVIEW.md`.

Ket luan:

- Du dieu kien chuyen qua UI prototype/MVP.
- Chua production-ready cho den khi tach reading engine, them seed/replay, tests va safety policy layer.

Tai lieu nay la ke hoach phat trien Tarot truoc trong Tarot.

## Ly Do Lam Tarot Truoc

Tarot nen lam truoc vi:

- Da co du 78 la va anh local.
- Khong phu thuoc thuat toan am lich/can chi phuc tap.
- De tao trai nghiem 3D an tuong som.
- De test freemium/paywall bang 1 flow gon: hoi cau hoi -> chon bai -> xem free -> unlock paid.
- Co the dat chat luong noi dung cao nhanh hon Tu Vi/BaZi.

## Definition Of Done Cho MVP

MVP Tarot du tot khi co:

- Rut 1 la va 3 la.
- Hien anh bai local.
- Shuffle/chon bai tren UI.
- Ket qua free cho 1 la.
- Ket qua paid preview cho 3 la.
- Backend hoac service layer luu reading input/output.
- Data khong thieu anh, khong thieu nghia la.
- Prompt AI co input tu data, khong tu bịa nghia.

## Phase 1: Data V2

Muc tieu: bien seed data thanh data san pham.

Viec can lam:

- Mo rong 78 la voi contexts:
  - general
  - love
  - career
  - money
  - inner_work
  - advice
  - warning
- Them symbolism cho Major Arcana truoc.
- Them yes/no/timing.
- Tach spreads thanh file rieng.
- Them combination rules.

Output mong muon:

- `data/tarot/rider-waite-smith.v2.json`
- `data/tarot/spreads.json`
- `data/tarot/combination-rules.json`

## Phase 2: Asset Pipeline

Muc tieu: anh nhe va san sang render.

Viec can lam:

- Convert JPG goc sang WebP.
- Tao thumbnail.
- Tao card back texture.
- Tao manifest asset gom width, height, size, hash.

Output mong muon:

- `public/textures/cards/rws-webp/*.webp`
- `public/textures/cards/rws-thumb/*.webp`
- `public/textures/cards/card-back.webp`
- `data/tarot/assets-manifest.json`

## Phase 3: App Scaffold

Muc tieu: co frontend/backend co the chay.

Frontend:

- Next.js app.
- Route `/tarot`.
- Route `/tarot/result/[id]`.
- Component deck/card.
- Tarot scene don gian truoc, 3D nang cap sau.

Backend:

- Go/Gin hoac API mock truoc neu can nhanh.
- Endpoint create reading.
- Endpoint get reading.
- Service rut bai.
- Repository tam bang file/in-memory trong prototype, PostgreSQL khi scaffold backend day du.

## Phase 4: Reading Logic

Muc tieu: reading nhat quan va replay duoc.

Can co:

- spread id.
- question optional.
- selected card ids.
- reversed flag.
- random seed.
- created at.
- free result.
- paid result.

Rule:

- Neu user tu chon bai tren UI, backend nhan card ids da chon.
- Neu auto draw, backend sinh seed va luu lai.
- Ket qua phai replay duoc bang `reading_id`.

## Phase 5: AI Prompt Builder

Muc tieu: AI viet hay nhung co can cu.

Prompt nen nhan:

- cau hoi nguoi dung.
- spread id va position meanings.
- tung la, upright/reversed.
- context snippets tu data.
- combination summary.
- tone guide.
- gioi han nhung noi dung khong duoc khang dinh.

Prompt khong nen nhan:

- raw JSON qua lon.
- yeu cau AI tu rut bai.
- yeu cau AI tu quyet dinh nghia la trai voi data.

## Phase 6: Paywall

Free:

- 1 la dau tien hoac overview 3-5 cau.
- 1 loi khuyen.

Paid:

- day du 3 la.
- lien ket cac la.
- phan tich theo cau hoi.
- loi khuyen hanh dong.
- PDF/report ve sau.

## Viec Nen Lam Ngay Tiep Theo

1. Lam UI prototype dua tren flow hien tai.
2. Test prototype voi 20 cau hoi trong `data/tarot/test-questions.json`, tap trung xem bai doc con chung chung khong.
3. Viet reading engine tach khoi prototype HTML.
4. Them seed/replay va `reading_id`.
5. Convert anh sang WebP/thumbnail.

