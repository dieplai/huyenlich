# Tarot Flow Prototype

Prototype nay dung de test flow Tarot truoc khi scaffold app chinh.

## File

- `prototype/tarot-flow.html`: UI test flow.
- `prototype/tarot_test_server.py`: static server + proxy den provider AI.

## Flow

```text
User question
-> AI classifier via json_schema
-> local Tarot engine draws cards and analyzes rules
-> AI writer writes reading from structured context
```

## Run

Dung bien moi truong cho API key, khong hardcode vao HTML:

```powershell
$env:CKEY_API_KEY = "sk-..."
python prototype/tarot_test_server.py 5175
```

Mo:

```text
http://127.0.0.1:5175/prototype/tarot-flow.html
```

## Notes

- `Local proxy` la mode nen dung vi tranh CORS va khong expose key cho browser.
- `Direct browser` chi dung khi provider cho phep CORS va ban chap nhan paste key vao UI.
- Neu AI writer loi, UI fallback sang reading local de flow khong bi dung.
- Prototype co load optional cac file `data/tarot/rider-waite-smith.v2.*.json`. Hien tai V2 da cover 78/78 la.
- Day la prototype, chua phai frontend production.
