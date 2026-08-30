import { createFileRoute } from "@tanstack/react-router";
import NotFoundLayout from "#/components/(layouts)/not-found-layout/not-found-layout";
import { PageLayout } from "#/components/(layouts)/page-layout/page-layout";

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
