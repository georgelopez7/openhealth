import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/(pages)/dashboard')({
  component: RouteComponent,
})

function RouteComponent() {
  return <div>Hello "/(pages)/dashboard"!</div>
}
