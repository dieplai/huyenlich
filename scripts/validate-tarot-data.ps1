$ErrorActionPreference = "Stop"

$root = (Get-Location).Path
$errors = New-Object System.Collections.Generic.List[string]

function Add-Error {
    param([string]$Message)
    $errors.Add($Message) | Out-Null
}

function Read-Json {
    param([string]$Path)

    try {
        return Get-Content -Raw -Encoding UTF8 -LiteralPath $Path | ConvertFrom-Json
    }
    catch {
        Add-Error "Invalid JSON: $Path :: $($_.Exception.Message)"
        return $null
    }
}

function Assert-Equal {
    param(
        [object]$Actual,
        [object]$Expected,
        [string]$Message
    )

    if ($Actual -ne $Expected) {
        Add-Error "$Message Expected=$Expected Actual=$Actual"
    }
}

function Assert-True {
    param(
        [bool]$Condition,
        [string]$Message
    )

    if (-not $Condition) {
        Add-Error $Message
    }
}

$tarotDir = Join-Path $root "data\tarot"
$jsonFiles = Get-ChildItem -LiteralPath $tarotDir -Filter "*.json" -File

foreach ($file in $jsonFiles) {
    [void](Read-Json $file.FullName)
}

$deck = Read-Json (Join-Path $tarotDir "rider-waite-smith.json")
$sourceMap = Read-Json (Join-Path $tarotDir "source-map.json")
$elements = Read-Json (Join-Path $tarotDir "elemental-dignities.json")
$spreads = Read-Json (Join-Path $tarotDir "spreads.json")
$reversals = Read-Json (Join-Path $tarotDir "reversal-rules.json")
$numbers = Read-Json (Join-Path $tarotDir "numerology-patterns.json")
$combinations = Read-Json (Join-Path $tarotDir "combination-rules.json")
$questionContexts = Read-Json (Join-Path $tarotDir "question-contexts.json")
$cardV2Sample = Read-Json (Join-Path $tarotDir "rider-waite-smith.v2.sample.json")
$cardV2MajorRest = Read-Json (Join-Path $tarotDir "rider-waite-smith.v2.major-rest.json")
$cardV2Wands = Read-Json (Join-Path $tarotDir "rider-waite-smith.v2.wands.json")
$cardV2Cups = Read-Json (Join-Path $tarotDir "rider-waite-smith.v2.cups.json")
$cardV2Swords = Read-Json (Join-Path $tarotDir "rider-waite-smith.v2.swords.json")
$cardV2Pentacles = Read-Json (Join-Path $tarotDir "rider-waite-smith.v2.pentacles.json")
$testQuestions = Read-Json (Join-Path $tarotDir "test-questions.json")

if ($deck) {
    $cards = @($deck.cards)
    Assert-Equal $cards.Count 78 "Deck must contain 78 cards."

    $ids = $cards | Group-Object id | Where-Object Count -gt 1
    Assert-Equal @($ids).Count 0 "Card ids must be unique."

    $indices = $cards | Group-Object deck_index | Where-Object Count -gt 1
    Assert-Equal @($indices).Count 0 "Card deck_index values must be unique."

    $missingIndices = @()
    for ($i = 0; $i -lt 78; $i++) {
        if (-not ($cards | Where-Object deck_index -eq $i)) {
            $missingIndices += $i
        }
    }
    Assert-Equal $missingIndices.Count 0 "Deck indices must cover 0..77."

    Assert-Equal @($cards | Where-Object arcana -eq "major").Count 22 "Major Arcana count must be 22."
    Assert-Equal @($cards | Where-Object arcana -eq "minor").Count 56 "Minor Arcana count must be 56."

    foreach ($suit in @("wands", "cups", "swords", "pentacles")) {
        Assert-Equal @($cards | Where-Object suit -eq $suit).Count 14 "Suit '$suit' count must be 14."
    }

    foreach ($card in $cards) {
        Assert-True (-not [string]::IsNullOrWhiteSpace($card.id)) "Card missing id."
        Assert-True (-not [string]::IsNullOrWhiteSpace($card.name_en)) "Card $($card.id) missing name_en."
        Assert-True (-not [string]::IsNullOrWhiteSpace($card.name_vi)) "Card $($card.id) missing name_vi."
        Assert-True (-not [string]::IsNullOrWhiteSpace($card.upright_vi)) "Card $($card.id) missing upright_vi."
        Assert-True (-not [string]::IsNullOrWhiteSpace($card.reversed_vi)) "Card $($card.id) missing reversed_vi."
        Assert-True (@($card.keywords_vi).Count -gt 0) "Card $($card.id) must have keywords_vi."

        if (-not [string]::IsNullOrWhiteSpace($card.image.local_path)) {
            $relative = $card.image.local_path -replace '/', [System.IO.Path]::DirectorySeparatorChar
            $imagePath = Join-Path $root $relative
            Assert-True (Test-Path -LiteralPath $imagePath) "Card $($card.id) image missing: $($card.image.local_path)"
        }
        else {
            Add-Error "Card $($card.id) missing image.local_path."
        }

        if ($card.arcana -eq "major") {
            Assert-True ($null -eq $card.suit) "Major card $($card.id) must not have suit."
        }

        if ($card.arcana -eq "minor") {
            Assert-True ($card.suit -in @("wands", "cups", "swords", "pentacles")) "Minor card $($card.id) has invalid suit."
            Assert-True ($card.rank -in @("ace", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "page", "knight", "queen", "king")) "Minor card $($card.id) has invalid rank."
        }
    }
}

if ($sourceMap) {
    $sourceIds = @($sourceMap.sources | ForEach-Object id)
    $duplicates = $sourceIds | Group-Object | Where-Object Count -gt 1
    Assert-Equal @($duplicates).Count 0 "Source ids must be unique."
    foreach ($required in @("waite", "golden_dawn_book_t", "mary_greer_reversals", "product_editorial")) {
        Assert-True ($sourceIds -contains $required) "source-map.json missing source id '$required'."
    }
}

if ($elements) {
    foreach ($suit in @("wands", "cups", "swords", "pentacles")) {
        Assert-True ($elements.suit_to_element.$suit -in @("fire", "water", "air", "earth")) "Invalid element for suit '$suit'."
    }

    foreach ($a in @("fire", "water", "air", "earth")) {
        foreach ($b in @("fire", "water", "air", "earth")) {
            Assert-True ($null -ne $elements.matrix.$a.$b) "Missing elemental matrix relation $a.$b."
        }
    }

    Assert-True (@($elements.triad_rules).Count -ge 3) "elemental-dignities.json should define triad rules."
}

if ($spreads) {
    $spreadItems = @($spreads.spreads)
    Assert-True ($spreadItems.Count -ge 5) "spreads.json should include at least five spreads."

    $spreadIds = @($spreadItems | ForEach-Object id)
    $spreadDuplicates = $spreadIds | Group-Object | Where-Object Count -gt 1
    Assert-Equal @($spreadDuplicates).Count 0 "Spread ids must be unique."
    Assert-True ($spreadIds -contains $spreads.default_spread_id) "default_spread_id must exist in spreads."

    foreach ($spread in $spreadItems) {
        $positions = @($spread.positions)
        Assert-Equal $positions.Count $spread.card_count "Spread $($spread.id) position count must equal card_count."
        $positionIndexes = @($positions | ForEach-Object index)
        for ($i = 1; $i -le $spread.card_count; $i++) {
            Assert-True ($positionIndexes -contains $i) "Spread $($spread.id) missing position index $i."
        }
    }
}

if ($reversals) {
    Assert-True ($reversals.orientation_policy.reversal_probability -ge 0 -and $reversals.orientation_policy.reversal_probability -le 1) "reversal_probability must be between 0 and 1."
    $modeIds = @($reversals.modes | ForEach-Object id)
    Assert-True ($modeIds -contains "blocked") "reversal-rules.json missing blocked mode."
    Assert-True ($modeIds -contains "internalized") "reversal-rules.json missing internalized mode."
    Assert-True ($modeIds -contains "shadow") "reversal-rules.json missing shadow mode."
}

if ($numbers) {
    foreach ($key in @("ace", "2", "3", "4", "5", "6", "7", "8", "9", "10")) {
        Assert-True ($null -ne $numbers.minor_numbers.$key) "numerology-patterns.json missing minor number '$key'."
    }
    foreach ($rank in @("page", "knight", "queen", "king")) {
        Assert-True ($null -ne $numbers.court_rank_patterns.$rank) "numerology-patterns.json missing court rank '$rank'."
    }
}

if ($combinations) {
    Assert-True (@($combinations.evaluation_order).Count -gt 0) "combination-rules.json missing evaluation_order."
    Assert-True (@($combinations.rule_categories).Count -ge 5) "combination-rules.json should include at least five rule categories."
    foreach ($category in $combinations.rule_categories) {
        Assert-True (@($category.rules).Count -gt 0) "Combination category $($category.id) must include rules."
    }
}

if ($questionContexts) {
    $contextIds = @($questionContexts.contexts | ForEach-Object id)
    foreach ($required in @("general", "love", "career", "money", "inner_work", "decision", "wellbeing_sensitive")) {
        Assert-True ($contextIds -contains $required) "question-contexts.json missing context '$required'."
    }

    $spreadIds = @()
    if ($spreads) {
        $spreadIds = @($spreads.spreads | ForEach-Object id)
    }
    foreach ($context in $questionContexts.contexts) {
        foreach ($spreadId in @($context.recommended_spreads)) {
            Assert-True ($spreadIds -contains $spreadId) "Question context $($context.id) references unknown spread '$spreadId'."
        }
    }
}

if ($testQuestions) {
    $items = @($testQuestions.items)
    Assert-True ($items.Count -ge 20) "test-questions.json should contain at least 20 golden questions."

    $itemIds = @($items | ForEach-Object id)
    $itemDuplicates = $itemIds | Group-Object | Where-Object Count -gt 1
    Assert-Equal @($itemDuplicates).Count 0 "Test question ids must be unique."

    $contextIds = @()
    if ($questionContexts) {
        $contextIds = @($questionContexts.contexts | ForEach-Object id)
    }

    $spreadIds = @()
    if ($spreads) {
        $spreadIds = @($spreads.spreads | ForEach-Object id)
    }

    foreach ($item in $items) {
        Assert-True (-not [string]::IsNullOrWhiteSpace($item.question_vi)) "Test question $($item.id) missing question_vi."
        Assert-True ($contextIds -contains $item.expected.primary_context) "Test question $($item.id) references unknown primary_context '$($item.expected.primary_context)'."
        foreach ($spreadId in @($item.expected.acceptable_spread_ids)) {
            Assert-True ($spreadIds -contains $spreadId) "Test question $($item.id) references unknown spread '$spreadId'."
        }
        Assert-True (@($item.expected.safety_flags).Count -gt 0) "Test question $($item.id) must define safety_flags."
        Assert-True (@($item.quality_checks_vi).Count -gt 0) "Test question $($item.id) must define quality checks."
    }
}

if ($cardV2Sample -or $cardV2MajorRest -or $cardV2Wands -or $cardV2Cups -or $cardV2Swords -or $cardV2Pentacles) {
    $deckIds = @()
    $majorIds = @()
    $minorIds = @()
    if ($deck) {
        $deckIds = @($deck.cards | ForEach-Object id)
        $majorIds = @($deck.cards | Where-Object arcana -eq "major" | ForEach-Object id)
        $minorIds = @($deck.cards | Where-Object arcana -eq "minor" | ForEach-Object id)
    }

    $v2DataSets = @()
    if ($cardV2Sample) {
        $v2DataSets += $cardV2Sample
    }
    if ($cardV2MajorRest) {
        $v2DataSets += $cardV2MajorRest
    }
    foreach ($minorDataSet in @($cardV2Wands, $cardV2Cups, $cardV2Swords, $cardV2Pentacles)) {
        if ($minorDataSet) {
            $v2DataSets += $minorDataSet
        }
    }

    $allV2Cards = @()
    foreach ($v2DataSet in $v2DataSets) {
        $v2Cards = @($v2DataSet.cards)
        Assert-True ($v2Cards.Count -gt 0) "V2 dataset $($v2DataSet.schema_version) must contain at least one card."
        Assert-Equal $v2DataSet.coverage.card_count $v2Cards.Count "V2 dataset $($v2DataSet.schema_version) coverage.card_count must match cards count."
        $allV2Cards += $v2Cards
    }

    $v2DuplicateIds = $allV2Cards | Group-Object id | Where-Object Count -gt 1
    Assert-Equal @($v2DuplicateIds).Count 0 "V2 card ids must be unique across all V2 datasets."
    Assert-Equal @($allV2Cards).Count 78 "V2 coverage should contain all 78 cards."

    foreach ($majorId in $majorIds) {
        Assert-True (@($allV2Cards | Where-Object id -eq $majorId).Count -eq 1) "V2 Major coverage missing '$majorId'."
    }
    foreach ($minorId in $minorIds) {
        Assert-True (@($allV2Cards | Where-Object id -eq $minorId).Count -eq 1) "V2 Minor coverage missing '$minorId'."
    }

    foreach ($card in $allV2Cards) {
        Assert-True ($deckIds -contains $card.id) "V2 sample card id '$($card.id)' does not exist in base deck."
        Assert-True (-not [string]::IsNullOrWhiteSpace($card.core.essence_vi)) "V2 card $($card.id) missing core.essence_vi."
        Assert-True (@($card.core.keywords_vi).Count -gt 0) "V2 card $($card.id) missing core.keywords_vi."
        Assert-True (-not [string]::IsNullOrWhiteSpace($card.orientation.upright.general_vi)) "V2 card $($card.id) missing upright.general_vi."
        Assert-True (-not [string]::IsNullOrWhiteSpace($card.orientation.reversed.general_vi)) "V2 card $($card.id) missing reversed.general_vi."
        Assert-True (-not [string]::IsNullOrWhiteSpace($card.orientation.upright.love_vi)) "V2 card $($card.id) missing upright.love_vi."
        Assert-True (-not [string]::IsNullOrWhiteSpace($card.orientation.reversed.love_vi)) "V2 card $($card.id) missing reversed.love_vi."
        Assert-True ($null -ne $card.relationship_signals.default) "V2 card $($card.id) missing relationship_signals.default."
        Assert-True (@($card.symbols).Count -gt 0) "V2 card $($card.id) should include symbols."
        Assert-True (@($card.reader_notes_vi).Count -gt 0) "V2 card $($card.id) should include reader notes."
    }
}

if ($errors.Count -gt 0) {
    Write-Host "Tarot data validation FAILED"
    foreach ($errorItem in $errors) {
        Write-Host "- $errorItem"
    }
    exit 1
}

Write-Host "Tarot data validation OK"
Write-Host "JSON files: $($jsonFiles.Count)"
if ($deck) {
    Write-Host "Cards: $(@($deck.cards).Count)"
}
if ($spreads) {
    Write-Host "Spreads: $(@($spreads.spreads).Count)"
}
if ($questionContexts) {
    Write-Host "Question contexts: $(@($questionContexts.contexts).Count)"
}
if ($testQuestions) {
    Write-Host "Golden questions: $(@($testQuestions.items).Count)"
}
if ($cardV2Sample) {
    $v2Total = @($cardV2Sample.cards).Count
    if ($cardV2MajorRest) {
        $v2Total += @($cardV2MajorRest.cards).Count
    }
    foreach ($minorDataSet in @($cardV2Wands, $cardV2Cups, $cardV2Swords, $cardV2Pentacles)) {
        if ($minorDataSet) {
            $v2Total += @($minorDataSet.cards).Count
        }
    }
    Write-Host "V2 cards: $v2Total"
}
