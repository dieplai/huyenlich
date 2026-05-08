'use client'

import { FormEvent, useEffect, useMemo, useRef, useState } from 'react'
import { AnimatePresence, motion } from 'framer-motion'
import { readingAPI, streamTarotReading } from '@/lib/api'
import { SparklesCore } from '@/components/ui/sparkles'

type Step = 'question' | 'selecting' | 'loading' | 'result'
type DeckPhase = 'stack' | 'fan'

interface TarotCard {
  id: string
  nameVI: string
  imagePath: string
  keywords: string[]
  uprightVI?: string
  reversedVI?: string
}

interface TarotReadingCard {
  id: string
  name_vi: string
  position: string
  meaning_vi?: string
  orientation?: string
  keywords_vi?: string[]
  upright_vi?: string
  reversed_vi?: string
  image_path?: string
}

interface TarotReadingResult {
  reading_id: string
  free: {
    question: string
    spread_id: string
    spread_name?: string
    cards: TarotReadingCard[]
    summary: string
    synthesis?: string
    advice?: string[]
    cautions?: string[]
  }
  paid?: {
    question?: string
    spread_id?: string
    cards?: TarotReadingCard[]
  }
  ai_content?: string
  ai_status?: 'ready' | 'fallback'
  is_unlocked?: boolean
}

interface TarotSpreadState {
  id: string
  targetCount: number
  desc: string
}

const FALLBACK_CARDS: TarotCard[] = [
  { id: 'major-00-the-fool', nameVI: 'Kẻ Khờ', imagePath: '/textures/cards/rws/major-00-the-fool.jpg', keywords: ['khởi đầu', 'tự do', 'niềm tin'] },
  { id: 'major-01-the-magician', nameVI: 'Nhà Ảo Thuật', imagePath: '/textures/cards/rws/major-01-the-magician.jpg', keywords: ['ý chí', 'kỹ năng', 'chủ động'] },
  { id: 'major-02-the-high-priestess', nameVI: 'Nữ Tư Tế', imagePath: '/textures/cards/rws/major-02-the-high-priestess.jpg', keywords: ['trực giác', 'bí ẩn', 'nội tâm'] },
  { id: 'major-03-the-empress', nameVI: 'Nữ Hoàng', imagePath: '/textures/cards/rws/major-03-the-empress.jpg', keywords: ['phong phú', 'nuôi dưỡng', 'sáng tạo'] },
  { id: 'major-04-the-emperor', nameVI: 'Hoàng Đế', imagePath: '/textures/cards/rws/major-04-the-emperor.jpg', keywords: ['quyền lực', 'ổn định', 'cấu trúc'] },
  { id: 'major-05-the-hierophant', nameVI: 'Giáo Hoàng', imagePath: '/textures/cards/rws/major-05-the-hierophant.jpg', keywords: ['truyền thống', 'hướng dẫn', 'tâm linh'] },
  { id: 'major-06-the-lovers', nameVI: 'Đôi Tình Nhân', imagePath: '/textures/cards/rws/major-06-the-lovers.jpg', keywords: ['tình yêu', 'lựa chọn', 'hài hòa'] },
  { id: 'major-07-the-chariot', nameVI: 'Cỗ Xe Chiến', imagePath: '/textures/cards/rws/major-07-the-chariot.jpg', keywords: ['chiến thắng', 'kiểm soát', 'quyết tâm'] },
  { id: 'major-08-strength', nameVI: 'Sức Mạnh', imagePath: '/textures/cards/rws/major-08-strength.jpg', keywords: ['can đảm', 'kiên nhẫn', 'nội lực'] },
  { id: 'major-09-the-hermit', nameVI: 'Ẩn Sĩ', imagePath: '/textures/cards/rws/major-09-the-hermit.jpg', keywords: ['nội tâm', 'cô độc', 'tìm kiếm'] },
  { id: 'major-10-wheel-of-fortune', nameVI: 'Bánh Xe Vận Mệnh', imagePath: '/textures/cards/rws/major-10-wheel-of-fortune.jpg', keywords: ['vận may', 'chu kỳ', 'thay đổi'] },
  { id: 'major-11-justice', nameVI: 'Công Lý', imagePath: '/textures/cards/rws/major-11-justice.jpg', keywords: ['công bằng', 'sự thật', 'nhân quả'] },
  { id: 'major-12-the-hanged-man', nameVI: 'Kẻ Treo Ngược', imagePath: '/textures/cards/rws/major-12-the-hanged-man.jpg', keywords: ['buông bỏ', 'chờ đợi', 'góc nhìn mới'] },
  { id: 'major-13-death', nameVI: 'Cái Chết', imagePath: '/textures/cards/rws/major-13-death.jpg', keywords: ['chuyển đổi', 'kết thúc', 'tái sinh'] },
  { id: 'major-14-temperance', nameVI: 'Tiết Chế', imagePath: '/textures/cards/rws/major-14-temperance.jpg', keywords: ['cân bằng', 'kiên nhẫn', 'điều độ'] },
  { id: 'major-15-the-devil', nameVI: 'Ác Quỷ', imagePath: '/textures/cards/rws/major-15-the-devil.jpg', keywords: ['ràng buộc', 'vật chất', 'bóng tối'] },
  { id: 'major-16-the-tower', nameVI: 'Tháp', imagePath: '/textures/cards/rws/major-16-the-tower.jpg', keywords: ['vỡ vụn', 'đột biến', 'giải phóng'] },
  { id: 'major-17-the-star', nameVI: 'Ngôi Sao', imagePath: '/textures/cards/rws/major-17-the-star.jpg', keywords: ['hy vọng', 'chữa lành', 'cảm hứng'] },
  { id: 'major-18-the-moon', nameVI: 'Mặt Trăng', imagePath: '/textures/cards/rws/major-18-the-moon.jpg', keywords: ['ảo tưởng', 'tiềm thức', 'sợ hãi'] },
  { id: 'major-19-the-sun', nameVI: 'Mặt Trời', imagePath: '/textures/cards/rws/major-19-the-sun.jpg', keywords: ['vui vẻ', 'thành công', 'sáng suốt'] },
  { id: 'major-20-judgement', nameVI: 'Phán Xét', imagePath: '/textures/cards/rws/major-20-judgement.jpg', keywords: ['thức tỉnh', 'tha thứ', 'phục sinh'] },
  { id: 'major-21-the-world', nameVI: 'Thế Giới', imagePath: '/textures/cards/rws/major-21-the-world.jpg', keywords: ['hoàn thành', 'toàn vẹn', 'thành tựu'] },
]

const DEFAULT_SPREAD = {
  id: 'five_card_cross',
  targetCount: 5,
  desc: 'Trải Bài 5 Lá',
} satisfies TarotSpreadState

const CROSS_POSITION_LABELS = [
  'Hiện trạng',
  'Điều ẩn bên dưới',
  'Tác động bên ngoài',
  'Điểm chuyển hướng',
  'Hướng đi tiếp',
]

const CARD_BACKS = ['/after1.jpg', '/after2.jpg', '/after3.jpg']
const VISUAL_CARD_LIMIT = 32
const DECK_READY_DELAY = 1800
const DECK_FAN_DELAY = 80

function shuffleCards(cards: TarotCard[]) {
  const next = [...cards]
  for (let i = next.length - 1; i > 0; i--) {
    const j = Math.floor(Math.random() * (i + 1))
    ;[next[i], next[j]] = [next[j], next[i]]
  }
  return next
}

function normalizeImagePath(path: string | undefined, id: string) {
  if (!path) return `/textures/cards/rws/${id}.jpg`
  return path.replace(/^public\//, '').startsWith('/')
    ? path.replace(/^public\//, '')
    : `/${path.replace(/^public\//, '')}`
}

function getCardFrame(stageWidth: number) {
  const width = stageWidth < 520 ? 78 : stageWidth < 860 ? 94 : 108
  return { width, height: Math.round(width * 1.75) }
}

function deckTransform(
  index: number,
  total: number,
  selected: boolean,
  selectedOrder: number,
  targetCount: number,
  stageWidth: number,
  step: Step,
  deckPhase: DeckPhase
) {
  if (selected) {
    if ((step === 'loading' || step === 'result') && targetCount === 5) {
      const isMobile = stageWidth < 640
      const isTablet = stageWidth >= 640 && stageWidth < 1024
      const crossX = isMobile
        ? 0
        : isTablet
          ? -Math.min(260, Math.max(190, stageWidth * 0.24))
          : -Math.min(430, Math.max(320, stageWidth * 0.28))
      const crossY = isMobile ? -92 : isTablet ? -4 : 10
      const horizontalGap = isMobile
        ? 96
        : isTablet
          ? Math.min(156, Math.max(136, stageWidth * 0.15))
          : Math.min(204, Math.max(176, stageWidth * 0.12))
      const verticalGap = isMobile
        ? 128
        : isTablet
          ? Math.min(198, Math.max(178, stageWidth * 0.2))
          : Math.min(242, Math.max(218, stageWidth * 0.135))
      const points = [
        { x: 0, y: 0, rotate: 0 },
        { x: 0, y: verticalGap, rotate: 2.5 },
        { x: -horizontalGap, y: 0, rotate: -4 },
        { x: horizontalGap, y: 0, rotate: 4 },
        { x: 0, y: -verticalGap, rotate: -1.5 },
      ]
      const point = points[selectedOrder] ?? points[0]

      return {
        x: crossX + point.x,
        y: crossY + point.y,
        rotate: point.rotate,
        rotateY: 180,
        scale: isMobile ? 0.74 : isTablet ? 0.86 : 0.88,
        delay: 0,
      }
    }

    const spread = Math.min(172, Math.max(86, stageWidth * (step === 'result' ? 0.22 : 0.18)))
    const resultShift = step === 'result' && stageWidth >= 900
      ? -Math.min(280, stageWidth * 0.22)
      : 0
    const x = resultShift + (selectedOrder - (targetCount - 1) / 2) * spread
    const y = step === 'result'
      ? stageWidth < 640 ? 64 : 58
      : stageWidth < 640 ? 138 : 178

    return {
      x,
      y,
      rotate: step === 'result' ? (selectedOrder - (targetCount - 1) / 2) * -5 : 0,
      rotateY: 180,
      scale: step === 'result' ? 1.08 : 1.14,
      delay: 0,
    }
  }

  const isMobile = stageWidth < 640
  const stackCenter = Math.floor(total / 2)
  const stackOffset = index - stackCenter
  const side = Math.sign(stackOffset) || 1
  const absStack = Math.abs(stackOffset)
  const phase = step === 'question' ? 'stack' : deckPhase
  const rowSize = Math.ceil(total / 2)
  const row = index < rowSize ? 0 : 1
  const local = row === 0 ? index : index - rowSize
  const cardsInRow = row === 0 ? rowSize : total - rowSize
  const safeRowSize = Math.max(cardsInRow, 1)
  const localCenter = (safeRowSize - 1) / 2
  const localOffset = local - localCenter
  const fanWidth = isMobile
    ? Math.min(104, Math.max(92, stageWidth * 0.26))
    : Math.min(540, Math.max(124, stageWidth * 0.405))

  if (phase === 'stack') {
    return {
      x: side * Math.min(absStack, 18) * (isMobile ? 0.72 : 0.88),
      y: (isMobile ? -112 : -84) + Math.min(absStack, 18) * 0.08,
      rotate: side * Math.min(absStack, 18) * (isMobile ? 0.14 : 0.18),
      rotateY: 0,
      scale: isMobile ? 0.96 : 1,
      delay: Math.min(absStack, 22) * 0.005,
    }
  }

  const rowGap = isMobile ? 120 : 146
  const rowBase = row === 0 ? -rowGap : rowGap * 0.72
  const p = safeRowSize <= 1 ? 0 : local / (safeRowSize - 1) * 2 - 1
  const angle = p * (isMobile ? 23 : 29)
  const arcDepth = Math.min(isMobile ? 42 : 66, Math.max(34, stageWidth * 0.064))
  const yCurve = p * p * arcDepth

  return {
    x: p * fanWidth,
    y: rowBase + yCurve,
    rotate: angle,
    rotateY: 0,
    scale: 1 + (1 - Math.abs(p)) * 0.014,
    delay: (row * 260 + local * 22) / 1000,
  }
}

function formatKeywords(keywords?: string[]) {
  return keywords?.length ? keywords.join(' · ') : 'Đang cập nhật từ khóa.'
}

function buildFallbackInsight(question: string, cards: TarotReadingCard[]) {
  const names = cards.map(card => `${card.position}: ${card.name_vi}`).join(' · ')
  return `Câu hỏi của bạn đang được phản chiếu qua ${names}. Hãy đọc từng lá như một lớp tín hiệu: điều đang lộ diện, điều cần tỉnh táo nhìn thẳng, và bước nhỏ nên làm sau trải bài này.`
}

export function TarotPageClient() {
  const [step, setStep] = useState<Step>('question')
  const [question, setQuestion] = useState('')
  const [allCards, setAllCards] = useState<TarotCard[]>(FALLBACK_CARDS)
  const [drawPool, setDrawPool] = useState<TarotCard[]>(FALLBACK_CARDS)
  const [selectedIds, setSelectedIds] = useState<string[]>([])
  const [result, setResult] = useState<TarotReadingResult | null>(null)
  const [aiContent, setAiContent] = useState<string | null>(null)
  const [isStreaming, setIsStreaming] = useState(false)
  const [activeCardId, setActiveCardId] = useState<string | null>(null)
  const [deckPhase, setDeckPhase] = useState<DeckPhase>('stack')
  const [deckReady, setDeckReady] = useState(false)
  const stageRef = useRef<HTMLDivElement>(null)
  const questionInputRef = useRef<HTMLInputElement>(null)
  const aiStreamEndRef = useRef<HTMLSpanElement>(null)
  const sequenceTimersRef = useRef<ReturnType<typeof setTimeout>[]>([])
  const autoOpenTimerRef = useRef<ReturnType<typeof setTimeout> | null>(null)
  const typewriterTimerRef = useRef<ReturnType<typeof setTimeout> | null>(null)
  const typewriterQueueRef = useRef('')
  const typedContentRef = useRef('')
  const receivedContentRef = useRef('')
  const streamCompleteRef = useRef(false)
  const readingStartedRef = useRef(false)
  const prepareRequestIdRef = useRef(0)
  const [stageWidth, setStageWidth] = useState(760)
  const [spread, setSpread] = useState<TarotSpreadState>(DEFAULT_SPREAD)
  const [preparedAnalysis, setPreparedAnalysis] = useState<Record<string, unknown> | null>(null)

  const cardFrame = useMemo(() => getCardFrame(stageWidth), [stageWidth])

  const selectedCards = useMemo(
    () => selectedIds
      .map(id => allCards.find(card => card.id === id))
      .filter((card): card is TarotCard => Boolean(card)),
    [allCards, selectedIds]
  )

  const resultCards = useMemo(() => {
    if (!result) return []
    const apiCards = result.paid?.cards?.length ? result.paid.cards : result.free.cards

    return apiCards.map(card => {
      const localCard = allCards.find(item => item.id === card.id)
      return {
        ...card,
        name_vi: card.name_vi || localCard?.nameVI || card.id,
        keywords_vi: card.keywords_vi?.length ? card.keywords_vi : localCard?.keywords,
        upright_vi: card.upright_vi || localCard?.uprightVI,
        reversed_vi: card.reversed_vi || localCard?.reversedVI,
        image_path: card.image_path || localCard?.imagePath,
      }
    })
  }, [allCards, result])

  const displayCards = useMemo<TarotReadingCard[]>(() => {
    if (resultCards.length) return resultCards

    return selectedCards.map((card, index) => ({
      id: card.id,
      name_vi: card.nameVI,
      position: CROSS_POSITION_LABELS[index] ?? `Lá ${index + 1}`,
      keywords_vi: card.keywords,
      upright_vi: card.uprightVI,
      reversed_vi: card.reversedVI,
      image_path: card.imagePath,
      meaning_vi: card.uprightVI,
    }))
  }, [resultCards, selectedCards])

  const streamBlocks = useMemo(
    () => (aiContent ?? '')
      .split(/\n{2,}/)
      .map(block => block.trim())
      .filter(Boolean),
    [aiContent]
  )

  const readingPanelOpen = (step === 'loading' || step === 'result') && selectedCards.length === spread.targetCount
  const activeCard = displayCards.find(card => card.id === activeCardId) ?? displayCards[0]
  const canRead = selectedIds.length === spread.targetCount

  const clearDeckSequence = () => {
    sequenceTimersRef.current.forEach(timer => clearTimeout(timer))
    sequenceTimersRef.current = []
  }

  const queueDeckPhase = (phase: DeckPhase, delay: number) => {
    sequenceTimersRef.current.push(setTimeout(() => setDeckPhase(phase), delay))
  }

  const startDeckDeal = () => {
    clearDeckSequence()
    setDeckReady(false)
    setDeckPhase('stack')
    queueDeckPhase('fan', DECK_FAN_DELAY)
    sequenceTimersRef.current.push(setTimeout(() => setDeckReady(true), DECK_READY_DELAY))
  }

  useEffect(() => {
    return () => {
      sequenceTimersRef.current.forEach(timer => clearTimeout(timer))
      sequenceTimersRef.current = []
      if (autoOpenTimerRef.current) clearTimeout(autoOpenTimerRef.current)
      if (typewriterTimerRef.current) clearTimeout(typewriterTimerRef.current)
    }
  }, [])

  useEffect(() => {
    const previousBodyOverflow = document.body.style.overflow
    const previousHtmlOverflow = document.documentElement.style.overflow

    document.body.style.overflow = 'hidden'
    document.documentElement.style.overflow = 'hidden'

    return () => {
      document.body.style.overflow = previousBodyOverflow
      document.documentElement.style.overflow = previousHtmlOverflow
    }
  }, [])

  useEffect(() => {
    let cancelled = false

    async function loadDeck() {
      try {
        const response = await fetch('/api/v1/tarot/deck')
        if (!response.ok) return
        const deck = await response.json() as {
          cards?: Array<{
            id: string
            name_vi: string
            image_path?: string
            keywords_vi?: string[]
            upright_vi?: string
            reversed_vi?: string
          }>
        }
        const cards = (deck.cards ?? []).map(card => ({
          id: card.id,
          nameVI: card.name_vi,
          imagePath: normalizeImagePath(card.image_path, card.id),
          keywords: card.keywords_vi ?? [],
          uprightVI: card.upright_vi,
          reversedVI: card.reversed_vi,
        }))
        if (!cancelled && cards.length) {
          setAllCards(cards)
          setDrawPool(shuffleCards(cards))
        }
      } catch {
        // Keep local fallback cards when the API is not ready.
      }
    }

    loadDeck()
    return () => {
      cancelled = true
    }
  }, [])

  useEffect(() => {
    const updateStageWidth = () => {
      if (stageRef.current) setStageWidth(stageRef.current.offsetWidth)
    }
    updateStageWidth()
    const observer = new ResizeObserver(updateStageWidth)
    if (stageRef.current) observer.observe(stageRef.current)
    return () => observer.disconnect()
  }, [])

  useEffect(() => {
    if (step !== 'question') {
      window.scrollTo({ top: 0, behavior: 'smooth' })
    }
  }, [step])

  useEffect(() => {
    if (readingPanelOpen && isStreaming) {
      aiStreamEndRef.current?.scrollIntoView({ block: 'nearest' })
    }
  }, [aiContent, isStreaming, readingPanelOpen])

  const resetTypewriter = () => {
    if (typewriterTimerRef.current) {
      clearTimeout(typewriterTimerRef.current)
      typewriterTimerRef.current = null
    }
    typewriterQueueRef.current = ''
    typedContentRef.current = ''
    receivedContentRef.current = ''
    streamCompleteRef.current = false
  }

  const scheduleTypewriter = () => {
    if (typewriterTimerRef.current) return

    const tick = () => {
      const queue = typewriterQueueRef.current
      if (!queue) {
        typewriterTimerRef.current = null
        if (streamCompleteRef.current) setIsStreaming(false)
        return
      }

      const nextChar = queue.slice(0, 1)
      typewriterQueueRef.current = queue.slice(1)
      typedContentRef.current += nextChar
      setAiContent(typedContentRef.current)

      const nextDelay = nextChar === '\n' ? 18 : /[.,;:!?]/.test(nextChar) ? 22 : 7
      typewriterTimerRef.current = setTimeout(tick, nextDelay)
    }

    typewriterTimerRef.current = setTimeout(tick, 7)
  }

  const enqueueAIText = (text: string) => {
    if (!text) return
    receivedContentRef.current += text
    typewriterQueueRef.current += text
    scheduleTypewriter()
  }

  const completeTypewriterWith = (finalText?: string) => {
    const normalized = finalText ?? ''
    if (normalized && normalized.startsWith(receivedContentRef.current)) {
      const missingTail = normalized.slice(receivedContentRef.current.length)
      if (missingTail) enqueueAIText(missingTail)
    } else if (normalized && !receivedContentRef.current) {
      enqueueAIText(normalized)
    }

    streamCompleteRef.current = true
    if (!typewriterQueueRef.current && !typewriterTimerRef.current) {
      setIsStreaming(false)
    }
  }

  const resetReadingState = () => {
    if (autoOpenTimerRef.current) {
      clearTimeout(autoOpenTimerRef.current)
      autoOpenTimerRef.current = null
    }
    resetTypewriter()
    readingStartedRef.current = false
    setSelectedIds([])
    setResult(null)
    setAiContent(null)
    setIsStreaming(false)
    setActiveCardId(null)
  }

  const beginSelection = (event?: FormEvent) => {
    event?.preventDefault()
    const nextQuestion = (questionInputRef.current?.value ?? question).trim()
    if (!nextQuestion) return

    const requestId = prepareRequestIdRef.current + 1
    prepareRequestIdRef.current = requestId

    setQuestion(nextQuestion)
    resetReadingState()
    setPreparedAnalysis(null)
    setSpread(DEFAULT_SPREAD)
    setDrawPool(shuffleCards(allCards))
    setStep('selecting')
    startDeckDeal()

    void readingAPI.prepare({ question: nextQuestion })
      .then(response => {
        if (prepareRequestIdRef.current !== requestId) return
        setPreparedAnalysis(response.data.analysis)
      })
      .catch(() => {
        if (prepareRequestIdRef.current !== requestId) return
        setPreparedAnalysis(null)
      })
  }

  const reshuffle = () => {
    resetReadingState()
    setDrawPool(shuffleCards(allCards))
    setStep('selecting')
    startDeckDeal()
  }

  const toggleCard = (cardId: string) => {
    setSelectedIds(prev => {
      if (prev.includes(cardId)) return prev.filter(id => id !== cardId)
      if (prev.length >= spread.targetCount) return prev
      return [...prev, cardId]
    })
  }

  const submitReading = async (cardIds = selectedIds) => {
    if (cardIds.length !== spread.targetCount || readingStartedRef.current) return
    readingStartedRef.current = true
    resetTypewriter()
    setStep('loading')
    setIsStreaming(true)
    setAiContent(null)

    const payload = {
      question: question.trim(),
      spread_id: spread.id,
      card_ids: cardIds,
      analysis: preparedAnalysis ?? undefined,
    }

    try {
      await streamTarotReading(payload, {
        onStart: (data) => {
          const nextResult = data as unknown as TarotReadingResult
          setResult(nextResult)
          setAiContent('')
          setActiveCardId(cardIds[0] ?? nextResult.free.cards?.[0]?.id ?? null)
          setStep('result')
        },
        onDelta: (text) => {
          enqueueAIText(text)
        },
        onDone: (data) => {
          const finalResult = data as unknown as TarotReadingResult
          setResult(finalResult)
          completeTypewriterWith(finalResult.ai_content ?? buildFallbackInsight(question.trim(), finalResult.free.cards))
        },
        onError: () => {
          readingStartedRef.current = false
          setSelectedIds([])
          resetTypewriter()
          setIsStreaming(false)
          setStep('selecting')
        },
      })
      if (!streamCompleteRef.current) completeTypewriterWith()
    } catch {
      try {
        const response = await readingAPI.tarot(payload)
        const nextResult = response.data as unknown as TarotReadingResult
        const nextContent = nextResult.ai_content ?? buildFallbackInsight(question.trim(), nextResult.free.cards)

        setResult(nextResult)
        setActiveCardId(cardIds[0] ?? null)
        setStep('result')
        enqueueAIText(nextContent)
        completeTypewriterWith(nextContent)
      } catch {
        readingStartedRef.current = false
        setSelectedIds([])
        resetTypewriter()
        setStep('selecting')
      } finally {
        if (!typewriterQueueRef.current && !typewriterTimerRef.current) setIsStreaming(false)
      }
    }
  }

  useEffect(() => {
    if (autoOpenTimerRef.current) {
      clearTimeout(autoOpenTimerRef.current)
      autoOpenTimerRef.current = null
    }

    if (
      step !== 'selecting' ||
      !deckReady ||
      selectedIds.length !== spread.targetCount ||
      readingStartedRef.current
    ) {
      return
    }

    const cardIds = [...selectedIds]
    autoOpenTimerRef.current = setTimeout(() => {
      autoOpenTimerRef.current = null
      void submitReading(cardIds)
    }, 160)

    return () => {
      if (autoOpenTimerRef.current) {
        clearTimeout(autoOpenTimerRef.current)
        autoOpenTimerRef.current = null
      }
    }
  }, [deckReady, selectedIds, spread.targetCount, step])

  const resetAll = () => {
    clearDeckSequence()
    setStep('question')
    setDeckPhase('stack')
    setDeckReady(false)
    setQuestion('')
    prepareRequestIdRef.current += 1
    setSpread(DEFAULT_SPREAD)
    setPreparedAnalysis(null)
    if (questionInputRef.current) questionInputRef.current.value = ''
    resetReadingState()
    setDrawPool(shuffleCards(allCards))
  }

  return (
    <main
      className="relative h-[100svh] overflow-hidden bg-cover bg-center text-[#f7f1df]"
      style={{
        backgroundImage:
          'linear-gradient(180deg, rgba(3,3,10,0.78) 0%, rgba(3,3,10,0.34) 44%, rgba(3,3,10,0.92) 100%), linear-gradient(90deg, rgba(3,3,10,0.8) 0%, rgba(3,3,10,0.18) 52%, rgba(3,3,10,0.75) 100%), url(/background.jpg)',
      }}
    >
      <div className="pointer-events-none absolute inset-x-0 top-0 h-40 bg-gradient-to-b from-black/70 to-transparent" />
      <div className="pointer-events-none absolute inset-x-0 bottom-0 h-56 bg-gradient-to-t from-black/90 to-transparent" />

      <header className="absolute inset-x-0 top-0 z-40 px-4 py-4 md:px-8" data-tarot-header>
        <div className="mx-auto flex w-full max-w-7xl items-center justify-between gap-4">
          <button
            type="button"
            onClick={resetAll}
            className="group text-left"
            aria-label="Về màn bắt đầu"
          >
            <span className="block font-mono text-[10px] uppercase tracking-[0.32em] text-[#d4af37]">Tarot</span>
            <span className="mt-1 block font-serif text-2xl leading-none text-white md:text-3xl">Rider-Waite-Smith</span>
          </button>

          <div className="hidden flex-1 md:block" />

          <div className="flex items-center gap-2">
            {step === 'selecting' && deckReady && (
              <button
                type="button"
                onClick={reshuffle}
                className="rounded-full border border-white/12 bg-black/30 px-4 py-2 text-xs font-bold text-white/72 backdrop-blur-md transition hover:border-[#d4af37]/60 hover:text-white"
              >
                Xáo lại
              </button>
            )}
          </div>
        </div>
      </header>

      <section className="relative z-10 h-[100svh] px-0 pb-0 pt-24 md:pt-28">
        <div className="mx-auto flex h-full w-full max-w-none items-center justify-center">
          <div
            ref={stageRef}
            data-tarot-stage
            className="relative h-full min-h-0 w-full max-w-full overflow-visible bg-transparent"
          >
            <div
              className="absolute left-0 top-0 h-full w-full"
              style={{ perspective: '1200px', perspectiveOrigin: '50% 42%' }}
            >
              <div className="absolute left-1/2 top-[42%] h-0 w-0 md:top-[43%]">
                {drawPool.slice(0, Math.min(drawPool.length, VISUAL_CARD_LIMIT)).map((card, index) => {
                  const selectedOrder = selectedIds.indexOf(card.id)
                  const selected = selectedOrder >= 0
                  const disabled = !selected && selectedIds.length >= spread.targetCount
                  const visualTotal = Math.min(drawPool.length, VISUAL_CARD_LIMIT)
                  const transform = deckTransform(index, visualTotal, selected, selectedOrder, spread.targetCount, stageWidth, step, deckPhase)
                  const active = selected && card.id === activeCard?.id
                  const clickable = (step === 'selecting' && deckReady) || (step === 'result' && selected)
                  const fullSelection = selectedIds.length >= spread.targetCount
                  const readingLayout = step === 'loading' || step === 'result'
                  const clearingDeck = fullSelection && !selected
                  const rowSize = Math.ceil(visualTotal / 2)
                  const row = index < rowSize ? 0 : 1
                  const local = row === 0 ? index : index - rowSize
                  const cardZIndex = deckPhase === 'fan'
                    ? Math.round((row === 0 ? 170 : 118) + local)
                    : index + 1

                  return (
                    <motion.button
                      key={card.id}
                      type="button"
                      onClick={() => {
                        if (step === 'selecting') toggleCard(card.id)
                        if (step === 'result' && selected) setActiveCardId(card.id)
                      }}
                      disabled={!clickable || (step === 'selecting' && disabled)}
                      className="group absolute left-0 top-0 rounded-lg outline-none disabled:cursor-default"
                      style={{
                        width: cardFrame.width,
                        height: cardFrame.height,
                        marginLeft: -cardFrame.width / 2,
                        marginTop: -cardFrame.height / 2,
                        zIndex: selected ? 220 + selectedOrder : cardZIndex,
                        transformStyle: 'preserve-3d',
                      }}
                      initial={false}
                      animate={{
                        x: transform.x,
                        y: transform.y,
                        rotate: transform.rotate,
                        scale: transform.scale,
                        opacity: readingLayout
                          ? selected ? 1 : 0
                          : step === 'question' ? 0
                          : fullSelection && !selected ? 0 : 1,
                        filter: readingLayout && !selected ? 'blur(6px)' : 'blur(0px)',
                      }}
                      whileHover={step === 'selecting' && deckReady && !selected && !disabled ? {
                        y: transform.y - 38,
                        scale: 1.08,
                        zIndex: 240,
                      } : undefined}
                      transition={{
                        duration: clearingDeck ? 0.82 : step === 'selecting' && deckReady ? 0.34 : 1.75,
                        ease: [0.16, 0.86, 0.18, 1],
                        delay: deckReady ? 0 : transform.delay,
                      }}
                      aria-label={step === 'result' ? `Xem ${card.nameVI}` : `Chọn ${card.nameVI}`}
                    >
                      <motion.span
                        className="absolute left-0 top-0 block h-full w-full rounded-lg"
                        animate={{ rotateY: transform.rotateY }}
                        transition={{ duration: 0.62, ease: [0.16, 0.86, 0.18, 1] }}
                        style={{
                          transformStyle: 'preserve-3d',
                          boxShadow: selected
                            ? active
                              ? '0 24px 62px rgba(0,0,0,0.62), 0 0 0 2px rgba(240,208,96,0.8), 0 0 34px rgba(212,175,55,0.28)'
                              : '0 22px 54px rgba(0,0,0,0.58), 0 0 0 2px rgba(212,175,55,0.38)'
                            : deckPhase === 'fan'
                              ? '0 10px 20px rgba(0,0,0,0.24)'
                              : 'none',
                        }}
                      >
                        <span
                          className="absolute left-0 top-0 h-full w-full overflow-hidden rounded-lg bg-black"
                          style={{ backfaceVisibility: 'hidden' }}
                        >
                          <img
                            src={CARD_BACKS[index % CARD_BACKS.length]}
                            alt="Mặt sau lá bài Tarot"
                            className="h-full w-full object-cover"
                          />
                          <span className="pointer-events-none absolute inset-0 rounded-lg shadow-[inset_0_0_0_1px_rgba(12,10,18,0.74),inset_0_0_18px_rgba(0,0,0,0.28)]" />
                        </span>
                        <span
                          className="absolute left-0 top-0 h-full w-full overflow-hidden rounded-lg bg-black"
                          style={{
                            transform: 'rotateY(180deg)',
                            backfaceVisibility: 'hidden',
                          }}
                        >
                          <img
                            src={card.imagePath}
                            alt={card.nameVI}
                            className="h-full w-full object-cover"
                          />
                          <span className="pointer-events-none absolute inset-0 rounded-lg shadow-[inset_0_0_0_1px_rgba(12,10,18,0.72),inset_0_0_18px_rgba(0,0,0,0.28)]" />
                        </span>
                      </motion.span>

                      {selected && (
                        <span className="absolute -right-3 -top-3 z-20 grid h-8 w-8 place-items-center rounded-full bg-[#d4af37] text-sm font-black text-[#130f09] shadow-lg shadow-black/40">
                          {selectedOrder + 1}
                        </span>
                      )}
                      <span className={`absolute left-1/2 top-full z-20 mt-3 w-36 -translate-x-1/2 rounded-md border border-white/10 bg-black/68 px-2 py-1 text-center text-[11px] font-semibold leading-tight text-white backdrop-blur-md transition ${step === 'result' && selected ? 'opacity-100' : 'opacity-0 group-hover:opacity-100'}`}>
                        {selected ? card.nameVI : 'Chọn lá này'}
                      </span>
                    </motion.button>
                  )
                })}
              </div>
            </div>

            <AnimatePresence mode="wait">
              {step === 'question' && (
                <div
                  key="question-line"
                  data-question-panel
                  className="absolute left-1/2 top-1/2 z-30 w-[calc(100vw-2rem)] max-w-[980px] -translate-x-1/2 -translate-y-1/2 text-center md:w-[980px] md:max-w-[calc(100%-3rem)]"
                >
                  <motion.form
                    onSubmit={beginSelection}
                    initial={{ opacity: 0, y: 18 }}
                    animate={{ opacity: 1, y: 0 }}
                    exit={{ opacity: 0, y: -14 }}
                  >
                    <div className="relative mx-auto flex max-w-5xl flex-col items-center justify-center">
                      <div className="pointer-events-none absolute inset-x-[-22%] top-20 h-64 overflow-hidden [mask-image:radial-gradient(700px_280px_at_top,white_16%,transparent_84%)]">
                        <SparklesCore
                          id="tarot-question-sparkles"
                          background="transparent"
                          minSize={0.35}
                          maxSize={1.35}
                          particleDensity={920}
                          className="h-full w-full"
                          particleColor="#FFFFFF"
                          speed={2}
                          fpsLimit={60}
                        />
                      </div>
                      <input
                        id="question"
                        ref={questionInputRef}
                        placeholder="Gõ câu hỏi của bạn tại đây"
                        autoComplete="off"
                        autoFocus
                        className="relative z-20 h-20 w-full bg-transparent px-2 text-center font-serif text-3xl font-bold leading-none text-white outline-none transition placeholder:text-white/34 sm:text-4xl md:text-5xl"
                      />
                      <div className="relative z-20 h-56 w-full max-w-[56rem] overflow-hidden">
                        <div className="absolute left-[6%] right-[6%] top-0 h-[3px] bg-gradient-to-r from-transparent via-[#d4af37] to-transparent blur-sm" />
                        <div className="absolute left-[6%] right-[6%] top-0 h-px bg-gradient-to-r from-transparent via-[#f8df85] to-transparent" />
                        <div className="absolute left-[36%] right-[36%] top-0 h-[6px] bg-gradient-to-r from-transparent via-white to-transparent blur-sm" />
                        <div className="absolute left-[36%] right-[36%] top-0 h-px bg-gradient-to-r from-transparent via-white to-transparent" />
                      </div>
                    </div>

                    <button
                      type="submit"
                      className="sr-only"
                    >
                      Xáo bài
                    </button>
                  </motion.form>
                </div>
              )}

              {step === 'selecting' && deckReady && (
                <motion.div
                  key="selection-hint"
                  initial={{ opacity: 0, y: 12 }}
                  animate={{ opacity: 1, y: 0 }}
                  exit={{ opacity: 0, y: 12 }}
                  className="absolute inset-x-4 top-4 z-20 mx-auto flex w-fit max-w-[calc(100%-32px)] items-center gap-3 rounded-full border border-white/12 bg-black/36 px-4 py-2 text-center text-xs font-semibold text-white/72 backdrop-blur-md"
                >
                  <span>{selectedIds.length < spread.targetCount ? 'Chọn theo trực giác' : 'Đã đủ lá'}</span>
                  <span className="h-1 w-1 rounded-full bg-[#d4af37]" />
                  <span>Chạm vào lá đã chọn để đổi ý</span>
                </motion.div>
              )}

              {readingPanelOpen && (
                <motion.aside
                  key="result"
                  initial={{ x: '100%' }}
                  animate={{ x: 0 }}
                  exit={{ x: '100%' }}
                  transition={{ duration: 0.82, ease: [0.16, 0.86, 0.18, 1] }}
                  className="fixed inset-y-0 right-0 z-50 h-[100svh] w-full overflow-y-auto overflow-x-hidden border-l border-white/[0.18] shadow-[-30px_0_90px_rgba(0,0,0,0.52)] backdrop-blur-[30px] md:w-[min(860px,58vw)] xl:w-[min(960px,52vw)] 2xl:w-[min(1040px,48vw)]"
                  style={{
                    backgroundColor: 'rgba(10,10,24,0.54)',
                    backgroundImage:
                      'linear-gradient(128deg, rgba(255,255,255,0.12) 0%, rgba(255,255,255,0.045) 26%, rgba(8,8,22,0.7) 100%), radial-gradient(circle at 0% 8%, rgba(212,175,55,0.18), transparent 34%), radial-gradient(circle at 100% 100%, rgba(120,231,255,0.12), transparent 38%)',
                    boxShadow:
                      '-32px 0 90px rgba(0,0,0,0.58), inset 1px 0 0 rgba(255,255,255,0.14), inset 0 0 42px rgba(255,255,255,0.055)',
                    WebkitBackdropFilter: 'blur(30px) saturate(1.22)',
                    backdropFilter: 'blur(30px) saturate(1.22)',
                  }}
                >
                  <div
                    className="sticky top-0 z-10 border-b border-white/10 px-5 py-5 backdrop-blur-2xl md:px-8 md:py-7"
                    style={{
                      backgroundImage:
                        'linear-gradient(180deg, rgba(8,8,22,0.98), rgba(8,8,22,0.9) 72%, rgba(8,8,22,0))',
                    }}
                  >
                    <div className="mx-auto flex max-w-[900px] items-start justify-between gap-5">
                      <div className="min-w-0">
                        <p className="font-mono text-[11px] uppercase tracking-[0.24em] text-[#d4af37]">Bản đọc đang mở</p>
                        <h2 className="mt-2 font-serif text-4xl leading-tight text-white md:text-5xl">{result?.free.spread_name ?? spread.desc}</h2>
                        <p className="mt-4 max-w-3xl text-sm italic leading-7 text-white/[0.62]">&ldquo;{result?.free.question ?? question}&rdquo;</p>
                      </div>
                      <button
                        type="button"
                        onClick={resetAll}
                        className="shrink-0 rounded-full border border-white/[0.14] bg-white/[0.07] px-4 py-2 text-xs font-bold text-white/70 transition hover:border-white/[0.28] hover:text-white"
                      >
                        Đóng
                      </button>
                    </div>
                  </div>

                  <div className="mx-auto max-w-[900px] px-5 pb-12 pt-7 md:px-8 md:pt-8">
                    <section>
                      <div className="flex items-center gap-3">
                        <h3 className="font-mono text-xs uppercase tracking-[0.2em] text-[#d4af37]">Luận giải</h3>
                        {isStreaming && (
                          <span className="rounded-full border border-[#d4af37]/25 bg-[#d4af37]/10 px-2.5 py-1 text-[10px] font-bold uppercase tracking-[0.12em] text-[#f8df85]">
                            Đang stream
                          </span>
                        )}
                      </div>
                      <div className="mt-4 rounded-lg border border-white/10 bg-white/[0.075] px-5 py-5 shadow-[inset_0_1px_0_rgba(255,255,255,0.08)] md:px-6 md:py-6">
                        <div className="flex items-start gap-4">
                          <div className="mt-1 grid h-8 w-8 shrink-0 place-items-center rounded-full border border-[#d4af37]/30 bg-[#d4af37]/[0.12] font-mono text-[10px] font-black text-[#f8df85]">
                            AI
                          </div>
                          <div className="min-w-0 flex-1">
                            <p className="font-mono text-[11px] uppercase tracking-[0.16em] text-white/[0.42]">Tarot reader</p>
                            <div className="mt-3 space-y-4 text-[15px] leading-8 text-white/[0.84] md:text-base md:leading-9">
                              {streamBlocks.length > 0 ? (
                                streamBlocks.map((block, index) => (
                                  <p key={`${index}-${block.slice(0, 18)}`} className="whitespace-pre-line">
                                    {block}
                                  </p>
                                ))
                              ) : (
                                <p className="text-white/[0.58]">Đang bắt đầu đọc năng lượng từ 5 lá bạn chọn...</p>
                              )}
                            </div>
                            <span ref={aiStreamEndRef} className="block h-px" />
                          </div>
                        </div>
                        {isStreaming && (
                          <span className="ml-12 mt-4 inline-block h-4 w-2 animate-pulse bg-[#d4af37]" />
                        )}
                      </div>
                    </section>

                    {displayCards.length > 0 && (
                      <section className="mt-8 border-t border-white/10 pt-6">
                        <h3 className="font-mono text-xs uppercase tracking-[0.2em] text-[#d4af37]">Các lá đã bốc</h3>
                        <div className="mt-4 grid grid-cols-1 gap-3 sm:grid-cols-2 xl:grid-cols-5">
                          {displayCards.map((card, index) => {
                            const active = activeCard?.id === card.id
                            return (
                              <button
                                key={card.id}
                                type="button"
                                onClick={() => setActiveCardId(card.id)}
                                className="min-w-0 rounded-lg border px-3 py-3 text-left transition"
                                style={{
                                  backgroundColor: active ? 'rgba(212,175,55,0.1)' : 'rgba(255,255,255,0.035)',
                                  borderColor: active ? 'rgba(212,175,55,0.72)' : 'rgba(255,255,255,0.1)',
                                  color: active ? '#f8df85' : 'rgba(255,255,255,0.68)',
                                }}
                              >
                                <span className="block font-mono text-[10px] uppercase leading-4 tracking-[0.12em]">{index + 1}. {card.position}</span>
                                <span className="mt-2 block truncate text-sm font-bold text-white">{card.name_vi}</span>
                              </button>
                            )
                          })}
                        </div>
                      </section>
                    )}

                    {activeCard && (
                      <section className="mt-8 rounded-lg border border-white/10 bg-white/[0.035] p-5 md:p-6">
                        <p className="font-mono text-xs uppercase tracking-[0.18em] text-[#78e7ff]">{activeCard.position}</p>
                        <h3 className="mt-2 font-serif text-4xl leading-tight text-white">{activeCard.name_vi}</h3>
                        <p className="mt-3 text-xs leading-5 text-white/[0.56]">{formatKeywords(activeCard.keywords_vi)}</p>
                        <p className="mt-5 text-[15px] leading-8 text-white/[0.76] md:text-base md:leading-8">
                          {activeCard.meaning_vi || activeCard.upright_vi || 'Lá này đang nhắc bạn quan sát lại câu hỏi từ một góc rõ ràng hơn trước khi hành động.'}
                        </p>
                      </section>
                    )}

                    <div className="mt-8 flex gap-3">
                      <button
                        type="button"
                        onClick={reshuffle}
                        className="rounded-full border border-white/12 bg-white/5 px-5 py-2.5 text-xs font-bold text-white/70 transition hover:border-white/[0.24] hover:text-white"
                      >
                        Trải lại câu này
                      </button>
                    </div>
                  </div>
                </motion.aside>
              )}
            </AnimatePresence>
          </div>
        </div>
      </section>

      {step === 'selecting' && deckReady && (
        <div className="fixed inset-x-0 bottom-0 z-40 px-4 pb-4 md:px-8 md:pb-6">
          <div className="mx-auto flex max-w-2xl items-center justify-between gap-3 rounded-full border border-white/12 bg-black/48 px-4 py-3 shadow-2xl backdrop-blur-xl">
            <div className="min-w-0">
              <p className="truncate text-xs text-white/55">{question}</p>
              <p className="mt-0.5 font-mono text-[11px] uppercase tracking-[0.16em] text-[#d4af37]">
                {spread.desc}
              </p>
            </div>
            <div className="shrink-0 rounded-full border border-[#d4af37]/28 bg-[#d4af37]/10 px-4 py-2 text-xs font-bold text-[#f8df85]">
              {canRead ? 'Đang mở bài...' : `${selectedIds.length}/${spread.targetCount} lá`}
            </div>
          </div>
        </div>
      )}
    </main>
  )
}
