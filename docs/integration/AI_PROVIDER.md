# AI Provider Integration

Tai lieu nay ghi chu cach tich hop provider AI dang dung cho Tarot.

## Endpoint Da Test

Provider hien tai:

```text
POST https://ckey.vn/v1/chat/completions
model: gpt-5.5
style: OpenAI-compatible Chat Completions
```

Ket qua test ngay 2026-04-28:

- Simple Chat Completions request: OK.
- `response_format.type = json_schema`: OK voi schema classifier rut gon.
- Response shape co `choices[0].message.content`, `finish_reason`.

## Model Compatibility Tests

### `gpt-5.5`

Trang thai: dung duoc cho prototype Tarot.

Da test:

- Simple request: OK.
- JSON schema classifier: OK.
- Phu hop de lam AI Question Classifier va AI Reading Writer.

### `claude-opus-4-6`

Trang thai: khong nen dung cho Tarot flow hien tai.

Da test ngay 2026-04-28:

- Simple request `Reply with exactly: OK`: OK.
- JSON schema classifier cho cau hoi Tarot/tinh cam: fail ve hanh vi.
- Model tra loi theo kieu "toi la tro ly lap trinh" va tu choi chu de Tarot/tinh cam.
- Co luc tra markdown JSON khong dung schema yeu cau.

Ket luan: endpoint co model nay, nhung model dang bi system prompt/provider scope gioi han cho coding/technical assistance. Khong phu hop de classify cau hoi Tarot hoac viet bai doc Tarot.

Khuyen nghi:

- Giu `gpt-5.5` lam default cho prototype.
- Chi dung `claude-opus-4-6` neu provider thay doi system scope hoac cap mot route/model khong bi gioi han coding.

Khong luu API key vao repo. Dung bien moi truong:

```powershell
$env:CKEY_API_KEY = "..."
```

## Khuyen Nghi Bao Mat

Neu API key da bi paste vao chat/log, nen rotate key.

Repo khong duoc commit:

- `.env`
- API key hard-coded
- request logs co Authorization header
- response logs chua thong tin ca nhan user

## Classifier Flow

Dung endpoint nay cho AI Question Classifier:

```text
user question
-> Chat Completions + response_format json_schema
-> structured classification JSON
-> Tarot rule engine
```

Schema tham chieu:

- `data/tarot/question-contexts.json`

## Vi Sao Khong De AI Tu Boi Het

AI provider chi nen lam:

- classify cau hoi.
- viet lai reading thanh tieng Viet hay.

AI provider khong nen:

- tu rut bai.
- tu doi spread.
- tu them rule Tarot.
- tu phan quyet ket qua y te/phap ly/tai chinh.

Rule phai nam o:

- `data/tarot/elemental-dignities.json`
- `data/tarot/spreads.json`
- `data/tarot/reversal-rules.json`
- `data/tarot/numerology-patterns.json`
- `data/tarot/combination-rules.json`

## OpenAI Compatibility Notes

Endpoint dang dung theo API style Chat Completions. Neu sau nay doi sang OpenAI truc tiep:

- Structured Outputs chinh thuc duoc OpenAI khuyen dung khi can output theo JSON schema.
- Responses API la API moi hon cho workflow agentic/stateful.
- Agents SDK huu ich khi can tools, handoff, streaming va tracing.

Nguon:

- https://platform.openai.com/docs/guides/structured-outputs
- https://platform.openai.com/docs/api-reference/responses
- https://platform.openai.com/docs/guides/agents-sdk/

