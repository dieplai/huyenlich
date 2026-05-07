# Tarot Reading Framework

Tai lieu nay mo ta cach Tarot doc va ket hop cac la Tarot. Muc tieu la de backend, AI prompt builder va frontend cung dung mot logic, thay vi chi ghep nghia tung la.

## Nguyen Tac Cot Loi

Mot reading khong phai danh sach card meanings. Mot reading la cau chuyen sinh ra tu:

1. cau hoi nguoi dung.
2. spread va y nghia tung position.
3. card trong tung position.
4. orientation cua card.
5. quan he giua cac card.
6. pattern tong the cua spread.
7. loi khuyen hanh dong.

Thu tu uu tien khi giai:

```text
question -> spread position -> card core -> orientation -> adjacent relation -> spread pattern -> synthesis
```

Neu co xung dot giua nghia chung cua la va position, position thang. Vi du `The Tower` o vi tri "warning" khac voi `The Tower` o vi tri "what to release".

## Layer 1: Question Context

Loai cau hoi nen duoc classify truoc:

- `general`
- `love`
- `career`
- `money`
- `inner_work`
- `decision`

Question context quyet dinh phan `contexts` nao trong card data duoc uu tien.

Vi du:

- Love question: uu tien `contexts.love`, relationship tone, emotional pattern.
- Career question: uu tien `contexts.career`, role, skill, conflict, opportunity.
- Money question: uu tien `contexts.money`, risk, stability, resources.

Neu user khong nhap cau hoi, mac dinh `general`.

## Layer 2: Spread Position

Moi spread position phai co:

```json
{
  "id": "present",
  "name_vi": "Hien tai",
  "function": "diagnostic|advice|forecast|shadow|outcome",
  "prompt_instruction": "Doc la nay nhu nang luong chinh cua hien tai."
}
```

Position function:

- `diagnostic`: mo ta hien trang.
- `advice`: chuyen nghia thanh hanh dong nen lam.
- `forecast`: chi xu huong neu tiep tuc hien tai.
- `shadow`: doc mat bi chan/noi tam/diem mu.
- `outcome`: tong hop kha nang, khong phan quyet tuyet doi.

## Layer 3: Card Core

Moi card can co:

- core meaning.
- keywords.
- upright meaning.
- reversed meaning.
- context meanings.
- symbols.
- advice/warning.

Card core la "tu vung". Position va combination bien no thanh "cau".

## Layer 4: Orientation / Reversal

Reversal trong Tarot co 5 mode:

- `blocked`: nang luong bi chan.
- `internalized`: nang luong quay vao ben trong.
- `delayed`: tien trinh cham lai.
- `shadow`: mat bong toi/bai hoc chua nhin.
- `overexpressed`: nang luong qua da.

Khong nen mac dinh reversal = nghia xau. AI prompt nen chon 1-2 mode phu hop dua tren question + position.

Vi du:

- `Ace of Cups reversed` trong love: co the la kho mo long, cam xuc bi nen.
- Trong inner work: co the la can tu cham soc truoc khi trao di.

## Layer 5: Adjacent Card Relations

Adjacent card la cac la dung gan nhau trong thu tu spread hoac co lien ket position.

Moi cap card nen duoc gan relation:

```json
{
  "relation": "amplifies|softens|challenges|clarifies|transitions|contrasts",
  "basis": ["keyword", "element", "number", "arcana", "symbol", "position_flow"]
}
```

### Amplifies

Hai la cung suit, cung number, cung chu de, hoac cung symbolism.

Vi du:

- `Three of Pentacles` + `Eight of Pentacles`: nhan manh cong viec, tay nghe, qua trinh xay dung.
- `The Star` + `Temperance`: nhan manh healing, hy vong, dieu hoa.

### Challenges

Hai la tao xung dot:

- Fire vs Water.
- Air vs Earth.
- mot la muon hanh dong, mot la muon dung lai.
- mot la mo long, mot la phong thu.

Vi du:

- `Knight of Wands` + `Four of Swords`: xung dot giua thuc day hanh dong va nhu cau nghi.
- `Two of Cups` + `Seven of Swords`: ket noi tinh cam nhung co van de minh bach.

### Clarifies

La sau lam ro la truoc.

Vi du:

- `The Moon` + `Ace of Swords`: su mo ho can duoc cat bang su that/giao tiep ro.

### Transitions

Thu tu card tao cau chuyen theo giai doan.

Vi du:

- `Two of Cups` -> `Eight of Cups`: tu ket noi den viec roi di vi can chieu sau moi.
- Dao thu tu `Eight of Cups` -> `Two of Cups`: roi bo dieu cu de mo cho ket noi moi.

## Layer 6: Combination Techniques

### 1. Keyword Synthesis

Chon 1-2 keyword chinh tu moi la, ghep theo position.

Format:

```text
Card A keyword + Card B keyword + position flow = combined meaning
```

Dung cho reading ngan/free.

### 2. Numerology Pattern

Ap dung cho numbered Minor va Major co number.

Cycle:

- 1/Ace: khoi dau.
- 2: lua chon, doi ung, can bang.
- 3: phat trien, bieu dat, hop tac.
- 4: cau truc, on dinh.
- 5: xung dot, bien dong.
- 6: dieu hoa, phuc hoi.
- 7: thu thach, chon loc, noi tam.
- 8: suc manh, tien trinh, dieu khien.
- 9: gan hoan tat, chieu sau, ket tinh.
- 10: ket thuc chu ky, chuyen pha.

Pattern:

- Cung number: chu de do duoc nhan manh.
- Number tang: tien trinh dang di toi.
- Number giam: quay lai giai doan truoc, can sua nen.
- Nhieu Ace: nhieu cua moi.
- Nhieu 5: nhieu bien dong.
- Nhieu 10: dong chu ky.

### 3. Elemental Dignities

Mapping:

- Wands = Fire.
- Cups = Water.
- Swords = Air.
- Pentacles = Earth.

Relation:

- Same element: strengthens.
- Fire + Air: supports.
- Water + Earth: supports.
- Fire + Water: weakens/challenges.
- Air + Earth: weakens/challenges.
- Fire + Earth: neutral, can manifest or weigh down.
- Water + Air: neutral, can balance heart and mind or create emotional overthinking.

Use:

- Khong bo qua la "yeu".
- Dung de xac dinh la nao dang manh/y trong spread.
- Dung de viet cau tong hop: hanh dong bi cam xuc lam cham, ly tri gap thuc te, cam xuc duoc nen tang ho tro...

### 4. Arcana Weight

- Major + Major: chu de lon, bai hoc archetype, kho xu ly bang meo ngan han.
- Major + Minor: Major la theme, Minor la cach no dien ra hang ngay.
- Minor + Minor: tinh huong thuc te, co the thay doi bang hanh dong.
- Nhieu Major trong spread: reading nen co tone nghiem tuc, chuyen hoa, turning point.
- It Major/nhieu Minor: reading nen thuc te, huong hanh dong.

### 5. Court Card Roles

Court cards co the la:

- a person.
- a role.
- a behavioral mode.
- a maturity level of suit energy.
- a message about how to act.

Rank lens:

- Page: hoc, tin moi, bat dau, to mo.
- Knight: hanh dong, theo duoi, dong luc, cuc doan cua suit.
- Queen: tiep nhan, lam chu noi tam, nuoi duong suit.
- King: quan tri, the hien ra ngoai, trach nhiem, authority.

Khong gan court card cung gioi tinh sinh hoc. Giai theo nang luong/vai tro.

### 6. Symbolic Echo

Neu hai la co cung motif, motif do la theme:

- water: cam xuc, dong chay, vo thuc.
- mountain: thu thach, tam nhin dai.
- sun/light: ro rang, sinh luc.
- moon/night: mo ho, tiem thuc, truc giac.
- child: tinh moi, su trong treo, qua khu.
- angel: healing, calling, protection, dieu hoa.
- crown/throne: quyen luc, lam chu, trach nhiem.
- road/path: hanh trinh, lua chon, tien trinh.

V2 card data can co `symbols` de engine co the detect echo.

### 7. Visual Direction

Neu UI/data co direction tags:

- nhan vat nhin trai: qua khu/noi tam.
- nhan vat nhin phai: tuong lai/huong ra ngoai.
- nhan vat quay lung: roi di, tim kiem, tu tach.
- nhan vat doi dien nhau: doi thoai, ket noi, xung dot.
- nhan vat khong nhin nhau: lech pha, mat ket noi.

Khong bat buoc V1, nhung huu ich cho reading hay.

## Layer 7: Spread-Level Pattern

Sau khi giai tung la, phai quet pattern toan spread:

```json
{
  "major_count": 0,
  "minor_count": 0,
  "suit_distribution": {},
  "element_distribution": {},
  "reversal_count": 0,
  "court_count": 0,
  "number_patterns": [],
  "dominant_themes": []
}
```

Interpretation:

- Dominant suit = main field of life.
- Missing suit = nang luong thieu.
- Many reversals = block/internal processing.
- Many courts = con nguoi/vai tro xa hoi quan trong.
- Many majors = turning point.
- Number cluster = chu ky dang noi bat.

## Reading Output Structure

Free output:

```json
{
  "summary": "",
  "visible_cards": [],
  "key_message": "",
  "advice": "",
  "paywall_teaser": ""
}
```

Paid output:

```json
{
  "opening": "",
  "spread_overview": "",
  "cards": [
    {
      "position": "",
      "card": "",
      "orientation": "",
      "interpretation": "",
      "advice": ""
    }
  ],
  "combinations": [],
  "patterns": [],
  "action_steps": [],
  "closing": ""
}
```

## Prompt Builder Rules

AI prompt phai nhan:

- question.
- spread positions.
- selected cards.
- card data snippets.
- combination analysis tu engine.
- tone guide.
- safety boundaries.

AI prompt khong duoc yeu cau:

- tu rut bai.
- tu tinh combination tu dau neu engine da tinh.
- dua ra ket luan y te/phap ly/tai chinh chac chan.
- dung ngon ngu doa so.

## Example: Three-Card Spread

Question: "Moi quan he nay nen tiep tuc khong?"

Cards:

- Past: Two of Cups upright.
- Present: Seven of Swords upright.
- Future: Eight of Cups upright.

Engine notes:

- Suit mix: Cups + Swords + Cups.
- Flow: connection -> secrecy/avoidance -> departure/search.
- Combination:
  - Two of Cups + Seven of Swords: ket noi co van de minh bach.
  - Seven of Swords + Eight of Cups: neu tiep tuc ne tranh su that, xu huong la roi di.

Good synthesis:

- "Spread nay khong phu nhan rang da tung co ket noi that, nhung hien tai van de nam o minh bach va long tin. Tuong lai khong nhat thiet la chia tay bat buoc, nhung neu hai ben khong noi ro su that, nang luong Eight of Cups cho thay mot nguoi se chon roi di de tim su binh yen hon."

Bad synthesis:

- "Hai ban chac chan se chia tay."

## Files Nen Tao Tiep

- `data/tarot/spreads.json`
- `data/tarot/combination-rules.json`
- `data/tarot/elemental-dignities.json`
- `data/tarot/numerology-patterns.json`
- `data/tarot/symbols.json`

