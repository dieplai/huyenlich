# Tarot Question Agent

Tai lieu nay mo ta cach xu ly cau hoi tu do cua user truoc khi boc bai Tarot.

## Quyet Dinh San Pham

Nen de user nhap cau hoi tu do thay vi bat chon dropdown `tinh yeu / su nghiep / tien bac`.

Ly do:

- Trai nghiem tu nhien hon.
- User khong bi ep vao category cung.
- Mot cau hoi co the co nhieu lop, vi du tinh yeu + quyet dinh + noi tam.
- AI co the classify de chon context, spread va safety boundaries.

Nhung khong nen de AI tu doc Tarot tu dau. Kien truc dung la:

```text
user question
-> AI Question Classifier
-> structured context JSON
-> Tarot rules engine
-> selected cards + rule analysis
-> AI Reading Writer
-> final reading
```

AI classify cau hoi. Tarot engine xu ly rule. AI writer chi viet thanh bai doc hay dua tren data/rule.

## Tai Sao Dung Structured Outputs

OpenAI Structured Outputs cho phep model tra ve JSON theo schema minh dinh nghia, tot hon prompt "hay tra ve JSON" thuong. Tai lieu OpenAI khuyen dung Structured Outputs thay vi JSON mode khi co the vi no rang buoc output theo schema.

Nguon:

- https://platform.openai.com/docs/guides/structured-outputs
- https://platform.openai.com/docs/api-reference/responses

Agents SDK co the dung neu flow can tools, handoff, stream va trace. V1 co the bat dau bang Responses API + structured output, chua can agent phuc tap.

Nguon:

- https://platform.openai.com/docs/guides/agents-sdk/

## Classifier Schema

Schema de trong:

- `data/tarot/question-contexts.json`

Output can co:

- `normalized_question`
- `primary_context`
- `secondary_contexts`
- `intent_type`
- `time_horizon`
- `emotional_tone`
- `safety_flags`
- `recommended_spread_id`
- `confidence`
- `reason_vi`

## Prompt Nguyen Tac

System/developer prompt cho classifier:

```text
You classify a Vietnamese Tarot question into structured fields.
Do not answer the Tarot question.
Do not draw cards.
Do not give advice.
Only classify intent, context, time horizon, safety flags, and recommended spread.
Return JSON matching the schema.
```

Input gom:

- user question.
- allowed contexts.
- allowed intent types.
- spread recommendation rules.
- safety categories.

Output chi la JSON.

## Recommended Routing

```text
primary_context = love
-> prefer love_five for paid, situation_action_outcome for MVP

primary_context = career
-> prefer career_five for paid, situation_action_outcome for MVP

primary_context = money
-> prefer situation_action_outcome or career_five
-> activate financial safety wording

primary_context = decision
-> prefer situation_action_outcome

primary_context = inner_work
-> prefer one_card or past_present_future

primary_context = wellbeing_sensitive
-> prefer one_card
-> safety boundary required
```

## Risk Va Guardrails

Rui ro:

- AI classify sai context.
- AI bien cau hoi nhay cam thanh loi khuyen nguy hiem.
- AI tu them nghia Tarot khong co trong data.

Guardrails:

- Neu confidence < 0.6, fallback `general + situation_action_outcome`.
- Neu co `medical`, `legal`, `financial_investment`, `pregnancy`, `self_harm`, route safety.
- Tarot engine khong nhan plain text category; chi nhan enum da validate.
- Reading writer khong duoc rut bai hay doi rule.

## V1 Recommendation

Dung OpenAI Responses API voi Structured Outputs cho classifier.

Chua can full Agents SDK neu flow chi co:

- classify question.
- draw cards.
- write reading.

Dung Agents SDK khi sau nay can:

- trace day du multi-step.
- tool calls phuc tap.
- handoff sang agent khac, vi du Tarot Reader Agent, Payment Agent, Retention Agent.
- streaming nhieu buoc trong UI.
