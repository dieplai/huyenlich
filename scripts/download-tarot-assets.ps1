param(
    [string]$DeckDataPath = "data/tarot/rider-waite-smith.json",
    [switch]$Force,
    [int]$ThrottleMilliseconds = 1500,
    [int]$MaxRetries = 5
)

$ErrorActionPreference = "Stop"

if (-not (Test-Path -LiteralPath $DeckDataPath)) {
    throw "Deck data file not found: $DeckDataPath"
}

$deck = Get-Content -Raw -Encoding UTF8 -LiteralPath $DeckDataPath | ConvertFrom-Json
$root = (Get-Location).Path
$cards = $deck.cards
$downloaded = 0
$skipped = 0
$headers = @{
    "User-Agent" = "TarotDataPrep/1.0 (local asset downloader; Wikimedia Commons)"
}

function Invoke-DownloadWithRetry {
    param(
        [string]$Uri,
        [string]$OutFile
    )

    for ($attempt = 1; $attempt -le $MaxRetries; $attempt++) {
        try {
            Invoke-WebRequest -Uri $Uri -OutFile $OutFile -Headers $headers

            $item = Get-Item -LiteralPath $OutFile
            if ($item.Length -lt 1024) {
                throw "Downloaded file is unexpectedly small: $($item.Length) bytes"
            }

            return
        }
        catch {
            if ($attempt -eq $MaxRetries) {
                throw
            }

            $delaySeconds = [Math]::Min(60, 5 * $attempt)
            Write-Warning "Download failed (attempt $attempt/$MaxRetries). Waiting $delaySeconds seconds. $($_.Exception.Message)"
            Start-Sleep -Seconds $delaySeconds
        }
    }
}

foreach ($card in $cards) {
    $relativePath = $card.image.local_path -replace '/', [System.IO.Path]::DirectorySeparatorChar
    $targetPath = Join-Path $root $relativePath
    $targetDir = Split-Path -Parent $targetPath

    if (-not (Test-Path -LiteralPath $targetDir)) {
        New-Item -ItemType Directory -Force -Path $targetDir | Out-Null
    }

    if ((Test-Path -LiteralPath $targetPath) -and -not $Force) {
        $existing = Get-Item -LiteralPath $targetPath
        if ($existing.Length -ge 1024) {
            $skipped++
            continue
        }
    }

    Write-Host "Downloading $($card.id) -> $relativePath"
    Invoke-DownloadWithRetry -Uri $card.image.download_url -OutFile $targetPath
    $downloaded++

    if ($ThrottleMilliseconds -gt 0) {
        Start-Sleep -Milliseconds $ThrottleMilliseconds
    }
}

Write-Host "Done. Downloaded=$downloaded Skipped=$skipped Total=$($cards.Count)"
