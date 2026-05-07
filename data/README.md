# Tarot Data Pack

Thu muc nay chua seed data Tarot dung chung cho backend, frontend va asset pipeline.

## Tarot

File chinh: `tarot/rider-waite-smith.json`

- Deck: Rider-Waite-Smith, 78 la.
- Anh: bo scan "Rider-Waite-Smith tarot deck (TaionWC)" tren Wikimedia Commons.
- Ban quyen: Rider-Waite goc duoc Wikimedia Commons ghi la public domain o US/UK; van can tranh dung cac ban to mau/retouch khong ro license.
- Truong `image.download_url` dung `Special:Redirect/file/...` de tai file goc tu Commons.
- Truong `image.local_path` la noi frontend nen doc sau khi tai asset ve local.

Tai anh:

```powershell
.\scripts\download-tarot-assets.ps1
```

Mac dinh script ghi vao `public/textures/cards/rws`.

## Nguyen tac dung data

- Backend nen load JSON nay thanh embedded data hoac import vao DB luc migration/seed.
- Frontend khong nen hotlink anh Commons khi production; hay tai ve local/CDN cua minh.
- Cac doan y nghia la noi dung san pham, khong phai khang dinh khoa hoc.
