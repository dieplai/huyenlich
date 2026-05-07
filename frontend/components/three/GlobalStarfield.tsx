'use client'
import { useEffect, useRef } from 'react'
import * as THREE from 'three'

export function GlobalStarfield({ className }: { className?: string }) {
  const mountRef = useRef<HTMLDivElement>(null)

  useEffect(() => {
    const mount = mountRef.current
    if (!mount) return

    // Renderer
    const renderer = new THREE.WebGLRenderer({ alpha: true, antialias: false })
    renderer.setPixelRatio(Math.min(window.devicePixelRatio, 1.5))
    renderer.setSize(mount.clientWidth, mount.clientHeight)
    renderer.setClearColor(0x000000, 0)
    mount.appendChild(renderer.domElement)

    // Scene + camera
    const scene = new THREE.Scene()
    const camera = new THREE.PerspectiveCamera(75, mount.clientWidth / mount.clientHeight, 0.1, 200)
    camera.position.z = 10

    // Stars
    const starCount = 7000
    const starPos = new Float32Array(starCount * 3)
    const starColors = new Float32Array(starCount * 3)
    for (let i = 0; i < starCount; i++) {
      const theta = Math.random() * Math.PI * 2
      const phi = Math.acos(2 * Math.random() - 1)
      const r = 20 + Math.random() * 60
      starPos[i*3]   = r * Math.sin(phi) * Math.cos(theta)
      starPos[i*3+1] = r * Math.sin(phi) * Math.sin(theta) * 0.5
      starPos[i*3+2] = r * Math.cos(phi)
      const warm = Math.random()
      starColors[i*3]   = 0.6 + warm * 0.4
      starColors[i*3+1] = 0.5 + warm * 0.2
      starColors[i*3+2] = 0.4 + (1 - warm) * 0.4
    }
    const starGeo = new THREE.BufferGeometry()
    starGeo.setAttribute('position', new THREE.BufferAttribute(starPos, 3))
    starGeo.setAttribute('color', new THREE.BufferAttribute(starColors, 3))
    const starMat = new THREE.PointsMaterial({ size: 0.06, vertexColors: true, transparent: true, opacity: 0.85, sizeAttenuation: true })
    const stars = new THREE.Points(starGeo, starMat)
    scene.add(stars)

    // Nebula particles
    const nebCount = 2000
    const nebPos = new Float32Array(nebCount * 3)
    for (let i = 0; i < nebCount; i++) {
      nebPos[i*3]   = (Math.random() - 0.5) * 60
      nebPos[i*3+1] = (Math.random() - 0.5) * 20
      nebPos[i*3+2] = (Math.random() - 0.5) * 30 - 5
    }
    const nebGeo = new THREE.BufferGeometry()
    nebGeo.setAttribute('position', new THREE.BufferAttribute(nebPos, 3))
    const nebMat = new THREE.PointsMaterial({ size: 0.15, color: 0x5b3fa0, transparent: true, opacity: 0.18, sizeAttenuation: true })
    scene.add(new THREE.Points(nebGeo, nebMat))

    // Mouse parallax
    let mx = 0, my = 0
    const onMouse = (e: MouseEvent) => {
      mx = (e.clientX / window.innerWidth  - 0.5) * 2
      my = (e.clientY / window.innerHeight - 0.5) * 2
    }
    window.addEventListener('mousemove', onMouse)

    // Resize
    const onResize = () => {
      const w = mount.clientWidth, h = mount.clientHeight
      camera.aspect = w / h
      camera.updateProjectionMatrix()
      renderer.setSize(w, h)
    }
    window.addEventListener('resize', onResize)

    // Animate
    let raf: number
    const clock = new THREE.Clock()
    const animate = () => {
      raf = requestAnimationFrame(animate)
      const t = clock.getElapsedTime()
      stars.rotation.y = t * 0.018
      stars.rotation.x = t * 0.005
      camera.position.x += (mx * 0.8 - camera.position.x) * 0.02
      camera.position.y += (-my * 0.5 - camera.position.y) * 0.02
      camera.lookAt(scene.position)
      renderer.render(scene, camera)
    }
    animate()

    return () => {
      cancelAnimationFrame(raf)
      window.removeEventListener('mousemove', onMouse)
      window.removeEventListener('resize', onResize)
      renderer.dispose()
      starGeo.dispose(); starMat.dispose()
      nebGeo.dispose(); nebMat.dispose()
      mount.removeChild(renderer.domElement)
    }
  }, [])

  return <div ref={mountRef} className={`absolute inset-0 pointer-events-none ${className ?? ''}`} />
}
