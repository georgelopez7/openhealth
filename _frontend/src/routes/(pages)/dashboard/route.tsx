import { createFileRoute } from "@tanstack/react-router";

import { PageLayout } from "#/components/(layouts)/page-layout/page-layout";

export const Route = createFileRoute("/(pages)/dashboard")({
  component: RouteComponent,
});

function RouteComponent() {
  return (
    <PageLayout>
      <div>Hello "/(pages)/dashboard"!</div>
    </PageLayout>
  );
}
