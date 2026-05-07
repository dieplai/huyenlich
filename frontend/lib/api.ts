import axios from 'axios'

const apiOrigin = process.env.NEXT_PUBLIC_API_URL?.replace(/\/$/, '')

const api = axios.create({
  baseURL: apiOrigin ? `${apiOrigin}/api/v1` : '/api/v1',
  headers: { 'Content-Type': 'application/json' },
})

api.interceptors.request.use((config) => {
  const token = typeof window !== 'undefined' ? localStorage.getItem('tarot_token') : null
  if (token) config.headers.Authorization = `Bearer ${token}`
  return config
})

export interface TarotRequest {
  question: string
  spread_id?: string
  card_ids: string[]
  orientations?: string[]
  analysis?: Record<string, unknown>
}

export interface TarotPrepareRequest {
  question: string
  preferred_card_count?: number
}

export interface TarotSpreadPosition {
  index: number
  id: string
  name_vi: string
  function: string
  prompt_instruction_vi: string
}

export interface TarotSpreadInfo {
  id: string
  name_vi: string
  card_count: number
  layout: string
  positions: TarotSpreadPosition[]
}

export interface TarotPrepareResponse {
  analysis: Record<string, unknown>
  spread: TarotSpreadInfo
}

export interface ReadingResponse {
  reading_id: string
  free: Record<string, unknown>
  paid?: Record<string, unknown>
  ai_content?: string
  ai_status?: 'ready' | 'fallback'
  is_unlocked: boolean
}

export const readingAPI = {
  prepare: (data: TarotPrepareRequest) =>
    api.post<TarotPrepareResponse>('/readings/tarot/prepare', data),

  tarot: (data: TarotRequest) =>
    api.post<ReadingResponse>('/readings/tarot', data),
}

type TarotStreamEvent = 'start' | 'delta' | 'done' | 'error' | 'message'

export interface TarotStreamHandlers {
  onStart?: (data: ReadingResponse) => void
  onDelta?: (text: string) => void
  onDone?: (data: ReadingResponse) => void
  onError?: (message: string) => void
}

function parseSSEBlock(block: string) {
  let event: TarotStreamEvent = 'message'
  const dataLines: string[] = []

  for (const line of block.split(/\r?\n/)) {
    if (line.startsWith('event:')) {
      event = line.slice(6).trim() as TarotStreamEvent
      continue
    }
    if (line.startsWith('data:')) {
      dataLines.push(line.slice(5).trimStart())
    }
  }

  return { event, data: dataLines.join('\n') }
}

export async function streamTarotReading(data: TarotRequest, handlers: TarotStreamHandlers) {
  const token = typeof window !== 'undefined' ? localStorage.getItem('tarot_token') : null
  const baseURL = apiOrigin ? `${apiOrigin}/api/v1` : '/api/v1'
  const response = await fetch(`${baseURL}/readings/tarot/stream`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      ...(token ? { Authorization: `Bearer ${token}` } : {}),
    },
    body: JSON.stringify(data),
  })

  if (!response.ok || !response.body) {
    throw new Error(`Tarot stream failed with status ${response.status}`)
  }

  const reader = response.body.getReader()
  const decoder = new TextDecoder()
  let buffer = ''

  while (true) {
    const { value, done } = await reader.read()
    if (done) break

    buffer += decoder.decode(value, { stream: true })
    const parts = buffer.split(/\r?\n\r?\n/)
    buffer = parts.pop() ?? ''

    for (const part of parts) {
      const { event, data: rawData } = parseSSEBlock(part)
      if (!rawData) continue

      let parsed: unknown = rawData
      try {
        parsed = JSON.parse(rawData)
      } catch {
        // Some SSE payloads may be plain text.
      }

      if (event === 'start') {
        handlers.onStart?.(parsed as ReadingResponse)
      } else if (event === 'delta') {
        const delta = typeof parsed === 'object' && parsed && 'text' in parsed
          ? String((parsed as { text?: string }).text ?? '')
          : String(parsed)
        if (delta) handlers.onDelta?.(delta)
      } else if (event === 'done') {
        handlers.onDone?.(parsed as ReadingResponse)
      } else if (event === 'error') {
        const message = typeof parsed === 'object' && parsed && 'error' in parsed
          ? String((parsed as { error?: string }).error ?? 'reading failed')
          : String(parsed)
        handlers.onError?.(message)
      }
    }
  }
}
