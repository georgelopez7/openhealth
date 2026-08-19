import { createFileRoute } from "@tanstack/react-router";

import { PageLayout } from "#/components/(layouts)/page-layout/page-layout";
import NotFoundLayout from "#/components/(layouts)/not-found-layout/not-found-layout";

export const Route = createFileRoute("/$")({
  component: NotFoundPage,
});

function NotFoundPage() {
  return (
    <PageLayout>
      <NotFoundLayout />
    </PageLayout>
  );
}
