// Polyfill React internals cho R3F 8.x chạy được với React 19
// React 19 chuyển ReactCurrentOwner sang ReactSharedInternals
import * as React from 'react'

type ReactInternals = {
  ReactCurrentOwner?: unknown
  ReactSharedInternals?: { ReactCurrentOwner?: unknown }
}

const internals = (React as unknown as { __SECRET_INTERNALS_DO_NOT_USE_OR_YOU_WILL_BE_FIRED?: ReactInternals })
  .__SECRET_INTERNALS_DO_NOT_USE_OR_YOU_WILL_BE_FIRED

if (internals && !internals.ReactCurrentOwner) {
  const shared = (React as unknown as { __SECRET_INTERNALS_DO_NOT_USE_OR_YOU_WILL_BE_FIRED: ReactInternals & { ReactSharedInternals?: { ReactCurrentOwner?: unknown } } })
    .__SECRET_INTERNALS_DO_NOT_USE_OR_YOU_WILL_BE_FIRED

  const owner = shared?.ReactSharedInternals?.ReactCurrentOwner
    ?? (React as unknown as { __CLIENT_INTERNALS_DO_NOT_USE_OR_WARN_USERS_THEY_CANNOT_UPGRADE?: { ReactCurrentOwner?: unknown } })
       .__CLIENT_INTERNALS_DO_NOT_USE_OR_WARN_USERS_THEY_CANNOT_UPGRADE?.ReactCurrentOwner
    ?? { current: null }

  if (shared) {
    shared.ReactCurrentOwner = owner
  }
}

export {}
